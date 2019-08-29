package bits_test

import (
	"github.com/gorkaio/gboy/pkg/bits"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConcatenatesBytesIntoWord(t *testing.T) {
	h, l := byte(0xCA), byte(0xFE)
	expected := uint16(0xCAFE)
	actual := bits.ConcatWord(h, l)
	assert.Equal(t, expected, actual)
}

func TestSplitsWordsIntoBytes(t *testing.T) {
	expectedH, expectedL := byte(0xCA), byte(0xFE)
	actualH, actualL := bits.SplitWord(0xCAFE)
	assert.Equal(t, expectedH, actualH)
	assert.Equal(t, expectedL, actualL)
}
