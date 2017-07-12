package opencorpora

import (
	"testing"
)

var m *Morph

func init() {
	m, _ = LoadMorph("morph.dict")
}

func BenchmarkMe(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if _, err := m.Tag("морфология"); err != nil {
			b.Fatal(err)
		}
	}
}
