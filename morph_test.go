package opencorpora

import (
	"bytes"
	"github.com/stretchr/testify/require"
	"testing"
)

const dict = `
	<?xml version="1.0" encoding="utf-8" standalone="yes"?>
	<dictionary version="0.92" revision="394634">
	    <grammemes>
	        <grammeme parent="">
	            <name>POST</name>
	            <alias>ЧР</alias>
	            <description>часть речи</description>
	        </grammeme>
					<grammeme parent="POST">
							<name>NOUN</name>
							<alias>СУЩ</alias>
							<description>имя существительное</description>
					</grammeme>
			</grammemes>
			<lemmata>
					<lemma id="1" rev="1">
							<l t="ёж">
									<g v="POST"/>
									<g v="NOUN"/>
							</l>
							<f t="ёж">
									<g v="NOUN"/>
							</f>
							<f t="ежа">
									<g v="POST"/>
									<g v="NOUN"/>
							</f>
					</lemma>
			</lemmata>
	</dictionary>
	`

func TestMorth(t *testing.T) {
	d := &MorphData{}
	buf := bytes.NewBufferString(dict)
	err := d.ImportFromReader(buf)
	require.NoError(t, err)

	m := NewMorph()
	err = m.readMorphData(d)
	require.NoError(t, err)

	tag, err := m.Tag("ежа")
	require.NoError(t, err)
	require.Len(t, tag, 2)
	require.Equal(t, "POST", tag[0].Name)
}
