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

func TestDetectsHalfCarryInByteAdds(t *testing.T) {
	assert.True(t, bits.HalfCarryAddByte(0x0F, 1))
	assert.False(t, bits.HalfCarryAddByte(0xF0, 2))
}

func TestDetectsHalfCarryInByteSubs(t *testing.T) {
	assert.True(t, bits.HalfCarrySubByte(0x10, 1))
	assert.False(t, bits.HalfCarrySubByte(0x04, 2))
}

func TestDetectsHalfCarryInWordAdds(t *testing.T) {
	assert.True(t, bits.HalfCarryAddWord(0x0FFF, 1))
	assert.False(t, bits.HalfCarryAddWord(0x00FF, 2))
}

func TestDetectsHalfCarryInWordSubs(t *testing.T) {
	assert.True(t, bits.HalfCarrySubWord(0x1000, 1))
	assert.False(t, bits.HalfCarrySubWord(0x0004, 2))
}

func TestDetectsCarryInByteAdds(t *testing.T) {
	assert.True(t, bits.CarryAddByte(0xFF, 2))
	assert.False(t, bits.CarryAddByte(0xF0, 5))
}

func TestDetectsCarryInWordAdds(t *testing.T) {
	assert.True(t, bits.CarryAddWord(0xFFFF, 2))
	assert.False(t, bits.CarryAddWord(0x0FF0, 5))
}

func TestDetectsCarryInByteSubs(t *testing.T) {
	assert.True(t, bits.CarrySubByte(0x01, 2))
	assert.False(t, bits.CarrySubByte(0xF0, 5))
}

func TestDetectsCarryInWordSubs(t *testing.T) {
	assert.True(t, bits.CarrySubWord(0x01, 2))
	assert.False(t, bits.CarrySubWord(0x0FF0, 5))
}

func TestAddAndSubAreTheSameThing(t *testing.T) {
	
}
