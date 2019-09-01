package cpu

import (
	"github.com/gorkaio/gboy/pkg/bits"
)

// ByteRegister is an 8 bit register
type ByteRegister struct {
	value uint8
	mask  byte
}

// WordRegister is an 16 bit register, composed of two 8 bit registers
type WordRegister struct {
	highByte *ByteRegister
	lowByte  *ByteRegister
}

// NewByteRegister creates a new 8 bit register
func newByteRegister() *ByteRegister {
	return &ByteRegister{
		value: 0,
		mask:  0xFF,
	}
}

// NewMaskedByteRegister creates a new 8 bit register with some bits unavailable
func newMaskedByteRegister(mask byte) *ByteRegister {
	return &ByteRegister{
		value: 0,
		mask:  mask,
	}
}

// NewWordRegister creates a new 16 bit register
func newWordRegister() *WordRegister {
	return &WordRegister{
		highByte: newByteRegister(),
		lowByte:  newByteRegister(),
	}
}

// NewMaskedWordRegister creates a new 16 bit register with some bits unavailable
func newMaskedWordRegister(mask uint16) *WordRegister {
	maskH, maskL := bits.SplitWord(mask)
	return &WordRegister{
		highByte: newMaskedByteRegister(maskH),
		lowByte:  newMaskedByteRegister(maskL),
	}
}

// Get gets the value of a 8 bit register
func (r *ByteRegister) Get() uint8 {
	return r.value & r.mask
}

// Set sets the value of a 8 bit register
func (r *ByteRegister) Set(data uint8) {
	r.value = data & r.mask
}

// Inc increases the value of a 8 bit register by one
// TODO: Masked registers might have strange behaviours with inc/dec
func (r *ByteRegister) Inc() {
	r.IncBy(1)
}

// IncBy increases the value of a 8 bit register by q
// TODO: Masked registers might have strange behaviours with inc/dec
func (r *ByteRegister) IncBy(q uint8) {
	r.Set(r.value + q)
}

// Dec decreases the value of a 8 bit register by one
// TODO: Masked registers might have strange behaviours with inc/dec
func (r *ByteRegister) Dec() {
	r.DecBy(1)
}

// DecBy decreases the value of a 8 bit register by q
// TODO: Masked registers might have strange behaviours with inc/dec
func (r *ByteRegister) DecBy(q uint8) {
	r.Set(r.value - q)
}

// GetMask returns the mask applied to this register
func (r *ByteRegister) GetMask() uint8 {
	return r.mask
}

// SetMask sets a mask to this register
func (r *ByteRegister) SetMask(mask uint8) {
	r.mask = mask
}

// Get gets the value of a 16 bit register
func (r *WordRegister) Get() uint16 {
	return bits.ConcatWord(r.highByte.Get(), r.lowByte.Get())
}

// Set sets the value of a 16 bit register
func (r *WordRegister) Set(data uint16) {
	highByte, lowByte := bits.SplitWord(data)
	r.highByte.Set(highByte)
	r.lowByte.Set(lowByte)
}

// Inc increases the value of a 16 bit register by one
// TODO: Masked registers might have strange behaviours with inc/dec
func (r *WordRegister) Inc() {
	r.IncBy(1)
}

// IncBy increases the value of a 16 bit register by q
// TODO: Masked registers might have strange behaviours with inc/dec
func (r *WordRegister) IncBy(q uint16) {
	r.Set(r.Get() + q)
}

// Dec decreases the value of a 16 bit register by one
// TODO: Masked registers might have strange behaviours with inc/dec
func (r *WordRegister) Dec() {
	r.DecBy(1)
}

// DecBy decreases the value of a 16 bit register by q
// TODO: Masked registers might have strange behaviours with inc/dec
func (r *WordRegister) DecBy(q uint16) {
	r.Set(r.Get() - q)
}

// GetMask returns the mask applied to this register
func (r *WordRegister) GetMask() uint16 {
	return bits.ConcatWord(r.highByte.mask, r.lowByte.mask)
}

// SetMask sets a mask to this register
func (r *WordRegister) SetMask(mask uint16) {
	maskH, maskL := bits.SplitWord(mask)
	r.highByte.SetMask(maskH)
	r.lowByte.SetMask(maskL)
}

// H gets the high byte register of a 16 bit register
func (r *WordRegister) H() *ByteRegister {
	return r.highByte
}

// L gets the low byte register of a 16 bit register
func (r *WordRegister) L() *ByteRegister {
	return r.lowByte
}
