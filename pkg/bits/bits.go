package bits

// ConcatWord concatenates two bytes into a single word
func ConcatWord(a uint8, b uint8) uint16 {
	return (uint16(a) << 8) | uint16(b)
}

// SplitWord splits a word into two bytes
func SplitWord(data uint16) (uint8, uint8) {
	return uint8((data & 0xFF00) >> 8), uint8(data & 0xFF)
}
