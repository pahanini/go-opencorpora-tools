package opencorpora

// Grammeme represents opencorpora grammeme
type Grammeme struct {
	ID          string `xml:"name"`
	ParentID    string `xml:"parent,attr"`
	Alias       string `xml:"alias"`
	Description string `xml:"description"`
}

// GrammemeID represents ID of grammeme in lemmas
type GrammemeID struct {
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
	Value       string       `xml:"t,attr"`
	GrammemeIDs []GrammemeID `xml:"g"`
}
