package cpu

// ByteRegisterInterface defines the interface for 8 bit registers
type ByteRegisterInterface interface {
	Get() uint8
	Set(data uint8)
	Inc()
	Dec()
}

// WordRegisterInterface defines the interface for 16 bit registers
type WordRegisterInterface interface {
	Get() uint16
	Set(data uint16)
	Inc()
	Dec()
	H() ByteRegisterInterface
	L() ByteRegisterInterface
}

// ByteRegister is an 8 bit register
type ByteRegister struct {
	v uint8
	ByteRegisterInterface
}

// WordRegister is an 16 bit register, composed of two 8 bit registers
type WordRegister struct {
	h ByteRegister
	l ByteRegister
	WordRegisterInterface
}

func newByteRegister() ByteRegister {
	return ByteRegister{
		v: 0,
	}
}

func newWordRegister() WordRegister {
	return WordRegister{
		h: newByteRegister(),
		l: newByteRegister(),
	}
}

// Get gets the value of a 8 bit register
func (r *ByteRegister) Get() uint8 {
	return r.v
}

// Set sets the value of a 8 bit register
func (r *ByteRegister) Set(data uint8) {
	r.v = data
}

// Inc increases the value of a 8 bit register by one
func (r *ByteRegister) Inc() {
	r.v++
}

// Dec decreases the value of a 8 bit register by one
func (r *ByteRegister) Dec() {
	r.v--
}

// Get gets the value of a 16 bit register
func (r *WordRegister) Get() uint16 {
	return concatWord(r.h.Get(), r.l.Get())
}

// Set sets the value of a 16 bit register
func (r *WordRegister) Set(data uint16) {
	r.h.v, r.l.v = splitWord(data)
}

// Inc increasex the value of a 16 bit register by one
func (r *WordRegister) Inc() {
	r.Set(r.Get() + 1)
}

// Dec decreases the value of a 16 bit register by one
func (r *WordRegister) Dec() {
	r.Set(r.Get() - 1)
}

// H gets the high byte register of a 16 bit register
func (r *WordRegister) H() *ByteRegister {
	return &r.h
}

// L gets the low byte register of a 16 bit register
func (r *WordRegister) L() *ByteRegister {
	return &r.l
}

func concatWord(a uint8, b uint8) uint16 {
	return (uint16(a) << 8) | uint16(b)
}

func splitWord(data uint16) (uint8, uint8) {
	return uint8((data & 0xFF00) >> 8), uint8(data & 0xFF)
}
