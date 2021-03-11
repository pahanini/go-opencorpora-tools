package opencorpora

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/require"
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
					<grammeme parent="ANim">
						<name>anim</name>
						<alias>од</alias>
						<description>одушевлённое</description>
					</grammeme>
					<grammeme parent="NMbr">
						<name>sing</name>
						<alias>ед</alias>
						<description>единственное число</description>
					</grammeme>
					<grammeme parent="CAse">
						<name>nomn</name>
						<alias>им</alias>
						<description>именительный падеж</description>
					</grammeme>
					<grammeme parent="CAse">
						<name>gent</name>
						<alias>рд</alias>
						<description>родительный падеж</description>
					</grammeme>
					<grammeme parent="ms-f">
						<name>masc</name>
						<alias>мр</alias>
						<description>мужской род</description>
					</grammeme>
			</grammemes>
			<lemmata>
					<lemma id="1" rev="1">
							<l t="ёж">
								<g v="NOUN"/>
								<g v="anim"/>
								<g v="masc"/>
							</l>
							<f t="ёж">
								<g v="sing"/>
								<g v="nomn"/>
							</f>
							<f t="ежа">
								<g v="sing"/>
								<g v="gent"/>
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
	require.Len(t, tag[0], 5)
	require.Equal(t, "NOUN", tag[0][0].Name)
}
