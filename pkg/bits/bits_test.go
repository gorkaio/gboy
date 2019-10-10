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

func TestDeterminesIfABitOfAByteIsSet(t *testing.T) {
	assert.True(t, bits.BitOfByte(0x08, 3))
	assert.False(t, bits.BitOfByte(0x12, 0))
}

func TestDeterminesIfABitOfAWordIsSet(t *testing.T) {
	assert.True(t, bits.BitOfWord(0x0101, 8))
	assert.False(t, bits.BitOfWord(0x0801, 1))
}

func TestFlipsBytesInAWord(t *testing.T) {
	assert.Equal(t, uint16(0xFECA), bits.FlipWord(0xCAFE))
}