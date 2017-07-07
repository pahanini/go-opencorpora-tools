package opencorpora

import (
	"encoding/xml"
	"io"
)

// MorphReaderToChans decodes reader and sends data into chans
func MorphReaderToChans(r io.Reader, gCh chan<- Grammeme, lCh chan<- Lemma) error {
	decoder := xml.NewDecoder(r)
	for {
		t, _ := decoder.Token()
		if t == nil {
			break
		}
		switch se := t.(type) {
		case xml.StartElement:
			if se.Name.Local == "grammeme" {
				var g Grammeme
				if err := decoder.DecodeElement(&g, &se); err != nil {
					return err
				}
				gCh <- g
			}
			if se.Name.Local == "lemma" {
				var l Lemma
				if err := decoder.DecodeElement(&l, &se); err != nil {
					return err
				}
				lCh <- l
			}
		}
	}
	return nil
}
