package opencorpora

import (
	"compress/bzip2"
	"encoding/gob"
	"encoding/xml"
	"github.com/pahanini/mafsa"
	"io"
	"net/http"
	"os"
	"sort"
)

const dictURL = "http://opencorpora.org/files/export/dict/dict.opcorpora.xml.bz2"

// MorphData is a struct to encode and save
type MorphData struct {
	Tree  []byte
	Tag   Tag
	Metas []Meta
}

// ImportFromWeb imports data from opencorpora site
func (d *MorphData) ImportFromWeb() (err error) {
	resp, err := http.Get(dictURL)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	bz2 := bzip2.NewReader(resp.Body)
	return d.ImportFromReader(bz2)
}

// ImportFromXMLFile reads XML file  and saves in MorphData
func (d *MorphData) ImportFromXMLFile(fp string) (err error) {
	f, err := os.Open(fp)
	if err != nil {
		return err
	}
	return d.ImportFromReader(f)
}

// ImportFromReader reads XML data from reader and saves in MorphData
func (d *MorphData) ImportFromReader(r io.Reader) (err error) {
	decoder := xml.NewDecoder(r)
	d.Tag = Tag{}
	d.Metas = []Meta{}

	// ww temporary keeps all dict before add
	// it to d.buildTree, tm associates tag names with Grammemes
	ww := wordForms{}
	tm := tagMap{}

	for {
		// exit loop if nothing to decode
		t, _ := decoder.Token()
		if t == nil {
			break
		}

		switch se := t.(type) {
		case xml.StartElement:
			// place pointer to grammeme into d.Tag
			// and temporary tagMap
			if se.Name.Local == "grammeme" {
				var g Grammeme
				if err = decoder.DecodeElement(&g, &se); err != nil {
					return err
				}
				d.Tag = append(d.Tag, g)
				tm[g.Name] = len(d.Tag) - 1
			}
			// generate all wordForms from lemma
			if se.Name.Local == "lemma" {
				var l Lemma
				if err = decoder.DecodeElement(&l, &se); err != nil {
					return err
				}
				ww = append(ww, newWordForm(l.Main, tm))
				for _, f := range l.Forms {
					ww = append(ww, newWordForm(f, tm))
				}
			}
		}
	}

	bt := mafsa.New()
	sort.Sort(ww)
	for _, w := range ww {
		if err = bt.Insert(w.word); err != nil {
			return err
		}
		d.Metas = append(d.Metas, w.meta)
	}
	bt.Finish()
	d.Tree, err = bt.MarshalBinary()
	return err
}

// Save saves MorphData to file
func (d *MorphData) Save(fp string) error {
	f, err := os.Create(fp)
	if err != nil {
		return err
	}
	defer f.Close()
	return d.writeMorphData(f)
}

// Load loads MorphData from file
func (d *MorphData) Load(fp string) error {
	f, err := os.Open(fp)
	if err != nil {
		return err
	}
	defer f.Close()
	return d.readMorphData(f)
}

// writeMorphData creates, encodes and writes MorphData to a io.Writer
func (d *MorphData) writeMorphData(w io.Writer) (err error) {
	g := gob.NewEncoder(w)
	return g.Encode(d)
}

// readMorphData reads, decodes and saves MorphData from a io.Reader
func (d *MorphData) readMorphData(r io.Reader) (err error) {
	g := gob.NewDecoder(r)
	return g.Decode(d)
}

// --- internal data structs ----

type tagMap map[string]int

type wordForm struct {
	word string
	meta Meta
}

func newWordForm(f Form, tm tagMap) wordForm {
	wf := wordForm{
		word: f.Value,
		meta: Meta{},
	}
	for _, n := range f.GrammemeNames {
		if grammemeIndex, ok := tm[n.Value]; ok {
			wf.meta.GrammemeIndexes = append(wf.meta.GrammemeIndexes, int8(grammemeIndex))
		}
	}
	return wf
}

type wordForms []wordForm

func (w wordForms) Len() int {
	return len(w)
}

func (w wordForms) Less(i, j int) bool {
	return w[i].word < w[j].word
}

func (w wordForms) Swap(i, j int) {
	w[i], w[j] = w[j], w[i]
}
