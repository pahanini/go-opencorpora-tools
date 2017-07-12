package opencorpora

// Grammeme represents opencorpora grammeme
type Grammeme struct {
	Name        string `xml:"name"`
	Description string `xml:"description"`
}

// Tag is a slice of grammeme
type Tag []Grammeme

// Meta keeps WF metadata
type Meta struct {
	GrammemeIndexes []int8
}

// --- parser types are only require for xml parse ---

// GrammemeName represents name of grammeme in lemmas
type GrammemeName struct {
	Value string `xml:"v,attr"`
}

// Lemma represents lemma
type Lemma struct {
	ID    int    `xml:"id,attr"`
	Rev   int    `xml:"rev,attr"`
	Main  Form   `xml:"l"`
	Forms []Form `xml:"f"`
}

// Form represents lemma form
type Form struct {
	Value         string         `xml:"t,attr"`
	GrammemeNames []GrammemeName `xml:"g"`
}
