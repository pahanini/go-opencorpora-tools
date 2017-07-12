package opencorpora

import (
	"errors"
	"github.com/pahanini/mafsa"
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
	d := MorphData{}
	if err := d.Load(fp); err != nil {
		return nil, err
	}
	if err := m.readMorphData(&d); err != nil {
		return nil, err
	}
	return m, nil
}

// NewMorph creates new Morph instance
func NewMorph() *Morph {
	return &Morph{Tag{}, []Meta{}, nil}
}

// Tag returns word's tag or nil if word not found
func (m *Morph) Tag(word string) (Tag, error) {
	_, index := m.tree.IndexedTraverse([]rune(word))
	tag := Tag{}
	if index < 1 {
		return tag, ErrWordNotFound
	}
	for _, grammemeIndex := range m.metas[index-1].GrammemeIndexes {
		tag = append(tag, m.tag[grammemeIndex])
	}
	return tag, nil
}

// readMorphData reads data from *MorphData struct and saves in Morph
func (m *Morph) readMorphData(md *MorphData) (err error) {
	m.tag = md.Tag
	m.metas = md.Metas
	m.tree, err = new(mafsa.Decoder).Decode(md.Tree)
	return
}
