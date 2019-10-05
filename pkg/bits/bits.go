package bits

// ConcatWord concatenates two bytes into a single word
func ConcatWord(a uint8, b uint8) uint16 {
	return (uint16(a) << 8) | uint16(b)
}

// SplitWord splits a word into two bytes
func SplitWord(data uint16) (uint8, uint8) {
	return uint8((data & 0xFF00) >> 8), uint8(data & 0xFF)
}

// BitOfByte determines if given bit of byte is set
func BitOfByte(data uint8, bit uint8) bool {
	return !(data & (1 << bit) == 0)
}

// BitOfWord determines if given bit of word is set
func BitOfWord(data uint16, bit uint8) bool {
	return !(data & (1 << bit) == 0)
}
