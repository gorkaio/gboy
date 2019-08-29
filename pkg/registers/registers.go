package registers

import (
	"github.com/gorkaio/gboy/pkg/bits"
)

type RegisterInterface interface{}

// ByteRegisterInterface defines the interface for 8 bit registers
type ByteRegisterInterface interface {
	Get() uint8
	Set(data uint8)
	Inc()
	IncBy(q uint8)
	Dec()
	DecBy(q uint8)
	GetMask() uint8
	SetMask(uint8)
	RegisterInterface
}

// WordRegisterInterface defines the interface for 16 bit registers
type WordRegisterInterface interface {
	Get() uint16
	Set(data uint16)
	Inc()
	IncBy(q uint16)
	Dec()
	DecBy(q uint16)
	GetMask() uint16
	SetMask(uint16)
	H() ByteRegisterInterface
	L() ByteRegisterInterface
	RegisterInterface
}

// ByteRegister is an 8 bit register
type ByteRegister struct {
	value uint8
	mask  byte
	ByteRegisterInterface
}

// WordRegister is an 16 bit register, composed of two 8 bit registers
type WordRegister struct {
	highByte *ByteRegister
	lowByte  *ByteRegister
	WordRegisterInterface
}

// NewByteRegister creates a new 8 bit register
func NewByteRegister() *ByteRegister {
	return &ByteRegister{
		value: 0,
		mask:  0xFF,
	}
}

// NewMaskedByteRegister creates a new 8 bit register with some bits unavailable
func NewMaskedByteRegister(mask byte) *ByteRegister {
	return &ByteRegister{
		value: 0,
		mask:  mask,
	}
}

// NewWordRegister creates a new 16 bit register
func NewWordRegister() *WordRegister {
	return &WordRegister{
		highByte: NewByteRegister(),
		lowByte:  NewByteRegister(),
	}
}

// NewMaskedWordRegister creates a new 16 bit register with some bits unavailable
func NewMaskedWordRegister(mask uint16) *WordRegister {
	maskH, maskL := bits.SplitWord(mask)
	return &WordRegister{
		highByte: NewMaskedByteRegister(maskH),
		lowByte:  NewMaskedByteRegister(maskL),
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
func (r *WordRegister) H() ByteRegisterInterface {
	return r.highByte
}

// L gets the low byte register of a 16 bit register
func (r *WordRegister) L() ByteRegisterInterface {
	return r.lowByte
}
