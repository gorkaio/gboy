package bits

// ConcatWord concatenates two bytes into a single word
func ConcatWord(a, b uint8) uint16 {
	return (uint16(a) << 8) | uint16(b)
}

// SplitWord splits a word into two bytes
func SplitWord(data uint16) (uint8, uint8) {
	return uint8((data & 0xFF00) >> 8), uint8(data & 0xFF)
}

// BitOfByte determines if given bit of byte is set
func BitOfByte(data, bit uint8) bool {
	return !(data&(1<<bit) == 0)
}

// BitOfWord determines if given bit of word is set
func BitOfWord(data uint16, bit uint8) bool {
	return !(data&(1<<bit) == 0)
}

// HalfCarryAddByte determines half carry for byte adding
func HalfCarryAddByte(data, value byte) bool {
	return ((data&0xF) + (value&0xf)) & 0x10 == 0x10
}

// HalfCarrySubByte determines half carry for byte substraction
func HalfCarrySubByte(data, value byte) bool {
	return ((data&0xF) - (value&0xf)) & 0x10 == 0x10
}

// FlipWord interchanges the High and Low bytes of a word
func FlipWord(data uint16) uint16 {
	return ((data & 0xFF00) >> 8) | ((data & 0x00FF) << 8)
}
