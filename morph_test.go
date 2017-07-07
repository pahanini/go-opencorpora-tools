package opencorpora

import (
	"github.com/stretchr/testify/require"
	"os"
	"testing"
)

func TestMorphReaderToChans(t *testing.T) {
	f, err := os.Open("testdata/test.xml")
	require.NoError(t, err)
	defer f.Close()

	gCh := make(chan Grammeme, 1000)
	lCh := make(chan Lemma, 1000)

	err = MorphReaderToChans(f, gCh, lCh)

	require.NoError(t, err)
	require.Len(t, gCh, 67)
	require.Len(t, lCh, 21)
}
