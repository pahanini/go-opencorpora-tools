package opencorpora

import (
	"encoding/gob"
	"errors"
	"github.com/pahanini/mafsa"
	"io"
	"os"
	"path/filepath"
)

// ErrWordNotFound returned if word not found in dict
var ErrWordNotFound = errors.New("word not found")

// Morph is a morphological dict based at opencorpora data
type Morph struct {
	tag   Tag
	metas []Meta
	tree  *mafsa.MinTree
}

// LoadMorph loads morph from file
func LoadMorph(fp string) (*Morph, error) {
	m := NewMorph()
	err := m.Load(fp)
	return m, err
}

// NewMorph creates new Morph instance
func NewMorph() *Morph {
	return &Morph{Tag{}, []Meta{}, nil}
}

// Tag returns word's tag or nil if word not found
func (m *Morph) Tag(word string) (Tag, error) {
	_, index := m.tree.IndexedTraverse([]rune(word))
	if index <= 1 {
		return Tag{}, ErrWordNotFound
	}
	return m.metas[index].Tag, nil
}

// Load reads and decodes file with specified filepath
func (m *Morph) Load(fp string) error {
	fp, err := filepath.Abs(fp)
	if err != nil {
		return err
	}
	f, err := os.Open(fp)
	if err != nil {
		return err
	}
	defer f.Close()
	return m.read(f)
}

// read reads and decodes MorphData from reader
func (m *Morph) read(r io.Reader) (err error) {
	var md MorphData
	decoder := gob.NewDecoder(r)
	if err := decoder.Decode(&md); err != nil {
		return err
	}
	return m.readMorphData(&md)
}

// readMorphData reads data from *MorphData struct and saves in Morph
func (m *Morph) readMorphData(md *MorphData) (err error) {
	m.tag = md.Tag
	m.metas = md.Metas
	m.tree, err = new(mafsa.Decoder).Decode(md.Tree)
	return
}
