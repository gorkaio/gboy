package cpu

import (
	"fmt"

	"github.com/gorkaio/gboy/pkg/bits"
)

const (
	lbyte = 1
	lword = 2
)

type opHandler func(cpu *CPU, args ...int) int

type op struct {
	mnemonic string
	args     []int
	length   uint16
	handler  opHandler
}

func (op *op) String() string {
	args := make([]interface{}, len(op.args))
	for i, arg := range op.args {
		args[i] = arg
	}

	return fmt.Sprintf(op.mnemonic, args...)
}

type opDefinition struct {
	mnemonic   string
	length     uint16
	argLengths []int
	handler    opHandler
}

var opDefinitions = map[uint8]opDefinition{
	0x00: {
		mnemonic:   "NOP",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			return cpu.nop()
		},
	},
	0x01: {
		mnemonic:   "LD BC, %#04x",
		argLengths: []int{lword},
		length:     3,
		handler: func(cpu *CPU, args ...int) int {
			return cpu.ldR16d16(cpu.BC, uint16(args[0]))
		},
	},
	0x02: {
		mnemonic:   "LD (BC), A",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			return cpu.ldaR16R8(cpu.BC, cpu.A)
		},
	},
	0x03: {
		mnemonic:   "INC BC",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			return cpu.incR16(cpu.BC)
		},
	},
	0x04: {
		mnemonic:   "INC B",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			return cpu.incR8(cpu.B)
		},
	},
	0x05: {
		mnemonic:   "DEC B",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			return cpu.decR8(cpu.B)
		},
	},
	0x06: {
		mnemonic:   "LD B, %#02x",
		argLengths: []int{lbyte},
		length:     2,
		handler: func(cpu *CPU, args ...int) int {
			return cpu.ldR8d8(cpu.B, uint8(args[0]))
		},
	},
	0x07: {
		mnemonic:   "RLCA",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			cpu.F.Set(0)
			cpu.SetFlagC(bits.BitOfByte(cpu.A.Get(), 7))
			cpu.A.Set(cpu.A.Get() << 1)
			return 4
		},
	},
	0x08: {
		mnemonic:   "LD (%#04x), SP",
		argLengths: []int{lword},
		length:     3,
		handler: func(cpu *CPU, args ...int) int {
			return cpu.lda16R16(uint16(args[0]), cpu.SP)
		},
	},
	0x09: {
		mnemonic: "ADD HL, BC",
		argLengths: []int{},
		length: 1,
		handler: func(cpu *CPU, args ...int) int {
			return cpu.addR16R16(cpu.HL, cpu.BC)
		},
	},
	0x0A: {
		mnemonic:   "LD A, (BC)",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			return cpu.ldR8aR16(cpu.A, cpu.BC)
		},
	},
	0x0B: {
		mnemonic:   "DEC BC",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			return cpu.decR16(cpu.BC)
		},
	},
	0x0C: {
		mnemonic:   "INC C",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			return cpu.incR8(cpu.C)
		},
	},
	0x0D: {
		mnemonic:   "DEC C",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			return cpu.decR8(cpu.C)
		},
	},
	0x0E: {
		mnemonic:   "LD C, %#02x",
		argLengths: []int{lbyte},
		length:     2,
		handler: func(cpu *CPU, args ...int) int {
			return cpu.ldR8d8(cpu.C, uint8(args[0]))
		},
	},
	/* TODO: 0x0F */
	/* TODO: 0x10 */
	0x11: {
		mnemonic:   "LD DE, %#04x",
		argLengths: []int{lword},
		length:     3,
		handler: func(cpu *CPU, args ...int) int {
			return cpu.ldR16d16(cpu.DE, uint16(args[0]))
		},
	},
	0x12: {
		mnemonic:   "LD (DE), A",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			return cpu.ldaR16R8(cpu.DE, cpu.A)
		},
	},
	0x13: {
		mnemonic:   "INC DE",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			return cpu.incR16(cpu.DE)
		},
	},
	0x14: {
		mnemonic:   "INC D",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			return cpu.incR8(cpu.D)
		},
	},
	0x15: {
		mnemonic:   "DEC D",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			return cpu.decR8(cpu.D)
		},
	},
	0x16: {
		mnemonic:   "LD D, %#02x",
		argLengths: []int{lbyte},
		length:     2,
		handler: func(cpu *CPU, args ...int) int {
			return cpu.ldR8d8(cpu.D, uint8(args[0]))
		},
	},
	/* TODO: 0x17 */
	/* TODO: 0x18 */
	0x19: {
		mnemonic: "ADD HL, DE",
		argLengths: []int{},
		length: 1,
		handler: func(cpu *CPU, args ...int) int {
			return cpu.addR16R16(cpu.HL, cpu.DE)
		},
	},
	0x1A: {
		mnemonic:   "LD A, (DE)",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			return cpu.ldR8aR16(cpu.A, cpu.DE)
		},
	},
	0x1B: {
		mnemonic:   "DEC DE",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			return cpu.decR16(cpu.DE)
		},
	},
	0x1C: {
		mnemonic:   "INC E",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			return cpu.incR8(cpu.E)
		},
	},
	0x1D: {
		mnemonic:   "DEC E",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			return cpu.decR8(cpu.E)
		},
	},
	0x1E: {
		mnemonic:   "LD E, %#02x",
		argLengths: []int{lbyte},
		length:     2,
		handler: func(cpu *CPU, args ...int) int {
			return cpu.ldR8d8(cpu.E, uint8(args[0]))
		},
	},
	/* TODO: 0x1F */
	0x20: {
		mnemonic:   "JR NZ, %#02x",
		argLengths: []int{lbyte},
		length:     2,
		handler: func(cpu *CPU, args ...int) int {
			if !cpu.FlagZ() {
				rel := int8(args[0]) // Signed relative address jump distance
				address := int(cpu.PC.Get()) + int(rel)
				cpu.PC.Set(uint16(address))
				return 12
			}
			return 8
		},
	},
	0x21: {
		mnemonic:   "LD HL, %#04x",
		argLengths: []int{lword},
		length:     3,
		handler: func(cpu *CPU, args ...int) int {
			return cpu.ldR16d16(cpu.HL, uint16(args[0]))
		},
	},
	0x22: {
		mnemonic:   "LDI (HL), A",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			return cpu.ldiaR16R8(cpu.HL, cpu.A)
		},
	},
	0x23: {
		mnemonic:   "INC HL",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			return cpu.incR16(cpu.HL)
		},
	},
	0x24: {
		mnemonic:   "INC H",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			return cpu.incR8(cpu.H)
		},
	},
	0x25: {
		mnemonic:   "DEC H",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			return cpu.decR8(cpu.H)
		},
	},
	0x26: {
		mnemonic:   "LD H, %#02x",
		argLengths: []int{lbyte},
		length:     2,
		handler: func(cpu *CPU, args ...int) int {
			return cpu.ldR8d8(cpu.H, uint8(args[0]))
		},
	},
	/* TODO: 0x27 */
	/* TODO: 0x28 */
	0x29: {
		mnemonic:   "ADD HL, HL",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			return cpu.addR16R16(cpu.HL, cpu.HL)
		},
	},
	0x2A: {
		mnemonic:   "LDI A, (HL)",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			return cpu.ldiR8aR16(cpu.A, cpu.HL)
		},
	},
	0x2B: {
		mnemonic:   "DEC HL",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			return cpu.decR16(cpu.HL)
		},
	},
	0x2C: {
		mnemonic:   "INC L",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			return cpu.incR8(cpu.L)
		},
	},
	0x2D: {
		mnemonic:   "DEC L",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			return cpu.decR8(cpu.L)
		},
	},
	0x2E: {
		mnemonic:   "LD L, %#02x",
		argLengths: []int{lbyte},
		length:     2,
		handler: func(cpu *CPU, args ...int) int {
			return cpu.ldR8d8(cpu.L, uint8(args[0]))
		},
	},
	/* TODO: 0x2F */
	/* TODO: 0x30 */
	0x31: {
		mnemonic:   "LD SP, %#04x",
		argLengths: []int{lword},
		length:     3,
		handler: func(cpu *CPU, args ...int) int {
			return cpu.ldR16d16(cpu.SP, uint16(args[0]))
		},
	},
	0x32: {
		mnemonic:   "LDD (HL), A",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			return cpu.lddaR16R8(cpu.HL, cpu.A)
		},
	},
	0x33: {
		mnemonic:   "INC SP",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			return cpu.incR16(cpu.SP)
		},
	},
	0x34: {
		mnemonic: "INC (HL)",
		argLengths: []int{},
		length: 1,
		handler: func(cpu *CPU, args ...int) int {
			a := cpu.HL.Get()
			d8 := cpu.memoryReadByte(cpu.HL.Get())
			halfCarry := bits.HalfCarryAddByte(d8, 1)
			cpu.memoryWriteByte(a, d8 + 1)
			cpu.SetFlagN(false)
			cpu.SetFlagZ(d8 == 0xFF)
			cpu.SetFlagH(halfCarry)
			return 12
		},
	},
	0x35: {
		mnemonic: "DEC (HL)",
		argLengths: []int{},
		length: 1,
		handler: func(cpu *CPU, args ...int) int {
			a := cpu.HL.Get()
			d8 := cpu.memoryReadByte(cpu.HL.Get())
			halfCarry := bits.HalfCarrySubByte(d8, 1)
			cpu.memoryWriteByte(a, d8 - 1)
			cpu.SetFlagN(true)
			cpu.SetFlagZ(d8 == 1)
			cpu.SetFlagH(halfCarry)
			return 12
		},
	},
	0x36: {
		mnemonic:   "LD (HL), %#02x",
		argLengths: []int{lbyte},
		length:     2,
		handler: func(cpu *CPU, args ...int) int {
			d8 := byte(args[0])
			return cpu.ldaR16d8(cpu.HL, d8)
		},
	},
	/* TODO: 0x37 */
	/* TODO: 0x38 */
	0x39: {
		mnemonic: "ADD HL, SP",
		argLengths: []int{},
		length: 1,
		handler: func(cpu *CPU, args ...int) int {
			return cpu.addR16R16(cpu.HL, cpu.SP)
		},
	},
	0x3A: {
		mnemonic:   "LDD A, (HL)",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			return cpu.lddR8aR16(cpu.A, cpu.HL)
		},
	},
	0x3B: {
		mnemonic:   "DEC SP",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			return cpu.decR16(cpu.SP)
		},
	},
	0x3C: {
		mnemonic:   "INC A",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			return cpu.incR8(cpu.A)
		},
	},
	0x3D: {
		mnemonic:   "DEC A",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			return cpu.decR8(cpu.A)
		},
	},
	0x3E: {
		mnemonic:   "LD A, %#02x",
		argLengths: []int{lbyte},
		length:     2,
		handler: func(cpu *CPU, args ...int) int {
			d8 := byte(args[0])
			return cpu.ldR8d8(cpu.A, d8)
		},
	},
	0x40: {
		mnemonic:   "LD B, B",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			return cpu.ldR8R8(cpu.B, cpu.B)
		},
	},
	0x41: {
		mnemonic:   "LD B, C",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			return cpu.ldR8R8(cpu.B, cpu.C)
		},
	},
	0x42: {
		mnemonic:   "LD B, D",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			return cpu.ldR8R8(cpu.B, cpu.D)
		},
	},
	0x43: {
		mnemonic:   "LD B, E",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			return cpu.ldR8R8(cpu.B, cpu.E)
		},
	},
	0x44: {
		mnemonic:   "LD B, H",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			return cpu.ldR8R8(cpu.B, cpu.H)
		},
	},
	0x45: {
		mnemonic:   "LD B, L",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			return cpu.ldR8R8(cpu.B, cpu.L)
		},
	},
	0x46: {
		mnemonic:   "LD B, (HL)",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			return cpu.ldR8aR16(cpu.B, cpu.HL)
		},
	},
	0x47: {
		mnemonic:   "LD B, A",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			return cpu.ldR8R8(cpu.B, cpu.A)
		},
	},
	0x48: {
		mnemonic:   "LD C, B",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			return cpu.ldR8R8(cpu.C, cpu.B)
		},
	},
	0x49: {
		mnemonic:   "LD C, C",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			return cpu.ldR8R8(cpu.C, cpu.C)
		},
	},
	0x4A: {
		mnemonic:   "LD C, D",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			return cpu.ldR8R8(cpu.C, cpu.D)
		},
	},
	0x4B: {
		mnemonic:   "LD C, E",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			return cpu.ldR8R8(cpu.C, cpu.E)
		},
	},
	0x4C: {
		mnemonic:   "LD C, H",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			return cpu.ldR8R8(cpu.C, cpu.H)
		},
	},
	0x4D: {
		mnemonic:   "LD C, L",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			return cpu.ldR8R8(cpu.C, cpu.L)
		},
	},
	0x4E: {
		mnemonic:   "LD C, (HL)",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			return cpu.ldR8aR16(cpu.C, cpu.HL)
		},
	},
	0x4F: {
		mnemonic:   "LD C, A",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			return cpu.ldR8R8(cpu.C, cpu.A)
		},
	},
	0x50: {
		mnemonic:   "LD D, B",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			return cpu.ldR8R8(cpu.D, cpu.B)
		},
	},
	0x51: {
		mnemonic:   "LD D, C",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			return cpu.ldR8R8(cpu.D, cpu.C)
		},
	},
	0x52: {
		mnemonic:   "LD D, D",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			return cpu.ldR8R8(cpu.D, cpu.D)
		},
	},
	0x53: {
		mnemonic:   "LD D, E",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			return cpu.ldR8R8(cpu.D, cpu.E)
		},
	},
	0x54: {
		mnemonic:   "LD D, H",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			return cpu.ldR8R8(cpu.D, cpu.H)
		},
	},
	0x55: {
		mnemonic:   "LD D, L",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			return cpu.ldR8R8(cpu.D, cpu.L)
		},
	},
	0x56: {
		mnemonic:   "LD D, (HL)",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			return cpu.ldR8aR16(cpu.D, cpu.HL)
		},
	},
	0x57: {
		mnemonic:   "LD D, A",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			return cpu.ldR8R8(cpu.D, cpu.A)
		},
	},
	0x58: {
		mnemonic:   "LD E, B",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			return cpu.ldR8R8(cpu.E, cpu.B)
		},
	},
	0x59: {
		mnemonic:   "LD E, C",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			return cpu.ldR8R8(cpu.E, cpu.C)
		},
	},
	0x5A: {
		mnemonic:   "LD E, D",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			return cpu.ldR8R8(cpu.E, cpu.D)
		},
	},
	0x5B: {
		mnemonic:   "LD E, E",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			return cpu.ldR8R8(cpu.E, cpu.E)
		},
	},
	0x5C: {
		mnemonic:   "LD E, H",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			return cpu.ldR8R8(cpu.E, cpu.H)
		},
	},
	0x5D: {
		mnemonic:   "LD E, L",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			return cpu.ldR8R8(cpu.E, cpu.L)
		},
	},
	0x5E: {
		mnemonic:   "LD E, (HL)",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			return cpu.ldR8aR16(cpu.E, cpu.HL)
		},
	},
	0x5F: {
		mnemonic:   "LD E, A",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			return cpu.ldR8R8(cpu.E, cpu.A)
		},
	},
	0x60: {
		mnemonic:   "LD H, B",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			return cpu.ldR8R8(cpu.H, cpu.B)
		},
	},
	0x61: {
		mnemonic:   "LD H, C",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			return cpu.ldR8R8(cpu.H, cpu.C)
		},
	},
	0x62: {
		mnemonic:   "LD H, D",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			return cpu.ldR8R8(cpu.H, cpu.D)
		},
	},
	0x63: {
		mnemonic:   "LD H, E",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			return cpu.ldR8R8(cpu.H, cpu.E)
		},
	},
	0x64: {
		mnemonic:   "LD H, H",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			return cpu.ldR8R8(cpu.H, cpu.H)
		},
	},
	0x65: {
		mnemonic:   "LD H, L",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			return cpu.ldR8R8(cpu.H, cpu.L)
		},
	},
	0x66: {
		mnemonic:   "LD H, (HL)",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			return cpu.ldR8aR16(cpu.H, cpu.HL)
		},
	},
	0x67: {
		mnemonic:   "LD H, A",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			return cpu.ldR8R8(cpu.H, cpu.A)
		},
	},
	0x68: {
		mnemonic:   "LD L, B",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			return cpu.ldR8R8(cpu.L, cpu.B)
		},
	},
	0x69: {
		mnemonic:   "LD L, C",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			return cpu.ldR8R8(cpu.L, cpu.C)
		},
	},
	0x6A: {
		mnemonic:   "LD L, D",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			return cpu.ldR8R8(cpu.L, cpu.D)
		},
	},
	0x6B: {
		mnemonic:   "LD L, E",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			return cpu.ldR8R8(cpu.L, cpu.E)
		},
	},
	0x6C: {
		mnemonic:   "LD L, H",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			return cpu.ldR8R8(cpu.L, cpu.H)
		},
	},
	0x6D: {
		mnemonic:   "LD L, L",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			return cpu.ldR8R8(cpu.L, cpu.L)
		},
	},
	0x6E: {
		mnemonic:   "LD L, (HL)",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			return cpu.ldR8aR16(cpu.L, cpu.HL)
		},
	},
	0x6F: {
		mnemonic:   "LD L, A",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			return cpu.ldR8R8(cpu.L, cpu.A)
		},
	},
	0x70: {
		mnemonic:   "LD (HL), B",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			return cpu.ldaR16R8(cpu.HL, cpu.B)
		},
	},
	0x71: {
		mnemonic:   "LD (HL), C",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			return cpu.ldaR16R8(cpu.HL, cpu.C)
		},
	},
	0x72: {
		mnemonic:   "LD (HL), D",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			return cpu.ldaR16R8(cpu.HL, cpu.D)
		},
	},
	0x73: {
		mnemonic:   "LD (HL), E",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			return cpu.ldaR16R8(cpu.HL, cpu.E)
		},
	},
	0x74: {
		mnemonic:   "LD (HL), H",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			return cpu.ldaR16R8(cpu.HL, cpu.H)
		},
	},
	0x75: {
		mnemonic:   "LD (HL), L",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			return cpu.ldaR16R8(cpu.HL, cpu.L)
		},
	},
	/* TODO: 0x76 */
	0x77: {
		mnemonic:   "LD (HL), A",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			return cpu.ldaR16R8(cpu.HL, cpu.A)
		},
	},
	0x78: {
		mnemonic:   "LD A, B",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			return cpu.ldR8R8(cpu.A, cpu.B)
		},
	},
	0x79: {
		mnemonic:   "LD A, C",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			return cpu.ldR8R8(cpu.A, cpu.C)
		},
	},
	0x7A: {
		mnemonic:   "LD A, D",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			return cpu.ldR8R8(cpu.A, cpu.D)
		},
	},
	0x7B: {
		mnemonic:   "LD A, E",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			return cpu.ldR8R8(cpu.A, cpu.E)
		},
	},
	0x7C: {
		mnemonic:   "LD A, H",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			return cpu.ldR8R8(cpu.A, cpu.H)
		},
	},
	0x7D: {
		mnemonic:   "LD A, L",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			return cpu.ldR8R8(cpu.A, cpu.L)
		},
	},
	0x7E: {
		mnemonic:   "LD A, (HL)",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			return cpu.ldR8aR16(cpu.A, cpu.HL)
		},
	},
	0x7F: {
		mnemonic:   "LD A, A",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			return cpu.ldR8R8(cpu.A, cpu.A)
		},
	},
	0x80: {
		mnemonic:   "ADD A, B",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			return cpu.addR8R8(cpu.A, cpu.B)
		},
	},
	0x81: {
		mnemonic:   "ADD A, C",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			return cpu.addR8R8(cpu.A, cpu.C)
		},
	},
	0x82: {
		mnemonic:   "ADD A, D",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			return cpu.addR8R8(cpu.A, cpu.D)
		},
	},
	0x83: {
		mnemonic:   "ADD A, E",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			return cpu.addR8R8(cpu.A, cpu.E)
		},
	},
	0x84: {
		mnemonic:   "ADD A, H",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			return cpu.addR8R8(cpu.A, cpu.H)
		},
	},
	0x85: {
		mnemonic:   "ADD A, L",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			return cpu.addR8R8(cpu.A, cpu.L)
		},
	},
	0x86: {
		mnemonic:   "ADD A, (HL)",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			return cpu.addR8aR16(cpu.A, cpu.HL)
		},
	},
	0x87: {
		mnemonic:   "ADD A, A",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			return cpu.addR8R8(cpu.A, cpu.A)
		},
	},
	/* TODO: 0x88 */
	/* TODO: 0x89 */
	/* TODO: 0x8A */
	/* TODO: 0x8B */
	/* TODO: 0x8C */
	/* TODO: 0x8D */
	/* TODO: 0x8E */
	/* TODO: 0x8F */
	0x90: {
		mnemonic:   "SUB B",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			return cpu.subR8(cpu.B)
		},
	},
	0x91: {
		mnemonic:   "SUB C",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			return cpu.subR8(cpu.B)
		},
	},
	0x92: {
		mnemonic:   "SUB D",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			return cpu.subR8(cpu.B)
		},
	},
	0x93: {
		mnemonic:   "SUB E",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			return cpu.subR8(cpu.B)
		},
	},
	0x94: {
		mnemonic:   "SUB H",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			return cpu.subR8(cpu.B)
		},
	},
	0x95: {
		mnemonic:   "SUB L",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			return cpu.subR8(cpu.B)
		},
	},
	0x96: {
		mnemonic:   "SUB (HL)",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			return cpu.subaR16(cpu.HL)
		},
	},
	0x97: {
		mnemonic:   "SUB A",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			return cpu.subR8(cpu.A)
		},
	},
	/* TODO: 0x98 */
	/* TODO: 0x99 */
	/* TODO: 0x9A */
	/* TODO: 0x9B */
	/* TODO: 0x9C */
	/* TODO: 0x9D */
	/* TODO: 0x9E */
	/* TODO: 0x9F */
	/* TODO: 0xA0 */
	/* TODO: 0xA1 */
	/* TODO: 0xA2 */
	/* TODO: 0xA3 */
	/* TODO: 0xA4 */
	/* TODO: 0xA5 */
	/* TODO: 0xA6 */
	/* TODO: 0xA7 */
	/* TODO: 0xA8 */
	/* TODO: 0xA9 */
	/* TODO: 0xAA */
	/* TODO: 0xAB */
	/* TODO: 0xAC */
	/* TODO: 0xAD */
	/* TODO: 0xAE */
	0xAF: {
		mnemonic:   "XOR A",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			cpu.A.Set(0)
			cpu.SetFlagZ(true)
			cpu.SetFlagN(false)
			return 4
		},
	},
	/* TODO: 0xB0 */
	/* TODO: 0xB1 */
	/* TODO: 0xB2 */
	/* TODO: 0xB3 */
	/* TODO: 0xB4 */
	/* TODO: 0xB5 */
	/* TODO: 0xB6 */
	/* TODO: 0xB7 */
	/* TODO: 0xB8 */
	/* TODO: 0xB9 */
	/* TODO: 0xBA */
	/* TODO: 0xBB */
	/* TODO: 0xBC */
	/* TODO: 0xBD */
	/* TODO: 0xBE */
	/* TODO: 0xBF */
	/* TODO: 0xC0 */
	0xC1: {
		mnemonic:   "POP BC",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			return cpu.popR16(cpu.BC)
		},
	},
	/* TODO: 0xC2 */
	0xC3: {
		mnemonic:   "JMP %#04x",
		argLengths: []int{lword},
		length:     3,
		handler: func(cpu *CPU, args ...int) int {
			cpu.PC.Set(uint16(args[0]))
			return 16
		},
	},
	/* TODO: 0xC4 */
	0xC5: {
		mnemonic:   "PUSH BC",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			return cpu.pushR16(cpu.BC)
		},
	},
	/* TODO: 0xC6 */
	/* TODO: 0xC7 */
	/* TODO: 0xC8 */
	/* TODO: 0xC9 */
	/* TODO: 0xCA */
	/* TODO: 0xCB */
	/* TODO: 0xCC */
	0xCD: {
		mnemonic:   "CALL %#04x",
		argLengths: []int{lword},
		length:     3,
		handler: func(cpu *CPU, args ...int) int {
			address := uint16(args[0])
			cpu.SP.Dec()
			cpu.memoryWriteWord(cpu.SP.Get(), cpu.PC.Get())
			cpu.PC.Set(address)
			return 24
		},
	},
	0xCE: {
		mnemonic:   "ADD A, %#02x",
		argLengths: []int{lbyte},
		length:     2,
		handler: func(cpu *CPU, args ...int) int {
			return cpu.addR8d8(cpu.A, byte(args[0]))
		},
	},
	/* TODO: 0xCF */
	/* TODO: 0xD0 */
	0xD1: {
		mnemonic:   "POP DE",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			return cpu.popR16(cpu.DE)
		},
	},
	/* TODO: 0xD2 */
	/* TODO: 0xD3 */
	/* TODO: 0xD4 */
	0xD5: {
		mnemonic:   "PUSH DE",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			return cpu.pushR16(cpu.DE)
		},
	},
	0xD6: {
		mnemonic:   "SUB %#02x",
		argLengths: []int{lbyte},
		length:     2,
		handler: func(cpu *CPU, args ...int) int {
			return cpu.subd8(byte(args[0]))
		},
	},
	/* TODO: 0xD7 */
	/* TODO: 0xD8 */
	/* TODO: 0xD9 */
	/* TODO: 0xDA */
	/* TODO: 0xDB */
	/* TODO: 0xDC */
	/* TODO: 0xDD */
	/* TODO: 0xDE */
	/* TODO: 0xDF */
	0xE0: {
		mnemonic:   "LDH (%#02x), A",
		argLengths: []int{lbyte},
		length:     2,
		handler: func(cpu *CPU, args ...int) int {
			return cpu.ldha8R8(byte(args[0]), cpu.A)
		},
	},
	0xE1: {
		mnemonic:   "POP HL",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			return cpu.popR16(cpu.HL)
		},
	},
	0xE2: {
		mnemonic:   "LD (C), A",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			return cpu.ldhR8R8(cpu.C, cpu.A)
		},
	},
	/* TODO: 0xE3 */
	/* TODO: 0xE4 */
	0xE5: {
		mnemonic:   "PUSH HL",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			return cpu.pushR16(cpu.HL)
		},
	},
	/* TODO: 0xE6 */
	/* TODO: 0xE7 */
	0xE8: {
		mnemonic: "ADD SP, %#02x",
		argLengths: []int{lbyte},
		length: 2,
		handler: func(cpu *CPU, args ...int) int {
			return cpu.addSP(int8(args[0]))
		},
	},
	/* TODO: 0xE9 */
	0xEA: {
		mnemonic:   "LD (%#04x), A",
		argLengths: []int{lword},
		length:     3,
		handler: func(cpu *CPU, args ...int) int {
			return cpu.lda16R8(uint16(args[0]), cpu.A)
		},
	},
	/* TODO: 0xEB */
	/* TODO: 0xEC */
	/* TODO: 0xED */
	/* TODO: 0xEE */
	/* TODO: 0xEF */
	0xF0: {
		mnemonic:   "LDH A, (%#02x)",
		argLengths: []int{lbyte},
		length:     2,
		handler: func(cpu *CPU, args ...int) int {
			return cpu.ldhR8a8(cpu.A, byte(args[0]))
		},
	},
	0xF1: {
		mnemonic:   "POP AF",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			return cpu.popR16(cpu.AF)
		},
	},
	0xF2: {
		mnemonic:   "LD A, (C)",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			return cpu.ldR8hR8(cpu.A, cpu.C)
		},
	},
	0xF3: {
		mnemonic:   "DI",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			cpu.DisableInterrupts()
			return 4
		},
	},
	/* TODO: 0xF4 */
	0xF5: {
		mnemonic:   "PUSH AF",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			return cpu.pushR16(cpu.AF)
		},
	},
	/* TODO: 0xF6 */
	/* TODO: 0xF7 */
	0xF8: {
		mnemonic:   "LDHL SP, %#02x",
		argLengths: []int{lbyte},
		length:     2,
		handler: func(cpu *CPU, args ...int) int {
			return cpu.ldR16R16a8(cpu.HL, cpu.SP, int8(args[0]))
		},
	},
	0xF9: {
		mnemonic:   "LD SP, HL",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			return cpu.ldR16R16(cpu.SP, cpu.HL)
		},
	},
	0xFA: {
		mnemonic:   "LD A, (%#04x)",
		argLengths: []int{lword},
		length:     3,
		handler: func(cpu *CPU, args ...int) int {
			return cpu.ldR8a16(cpu.A, uint16(args[0]))
		},
	},
	0xFB: {
		mnemonic:   "EI",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			cpu.EnableInterrupts()
			return 4
		},
	},
	/* TODO: 0xFC */
	/* TODO: 0xFD */
	0xFE: {
		mnemonic:   "CP %#02x",
		argLengths: []int{lbyte},
		length:     2,
		handler: func(cpu *CPU, args ...int) int {
			d8 := uint8(args[0])
			cpu.SetFlagZ(cpu.A.Get() == d8)
			cpu.SetFlagC(cpu.A.Get() < d8)
			cpu.SetFlagN(true)
			return 8
		},
	},
	/* TODO: 0xFF */
}

func opCodeFrom(data uint32) (op, error) {
	opCode := byte((data & 0xFF000000) >> 24)
	opDefinition, f := opDefinitions[opCode]
	if f != true {
		return op{}, fmt.Errorf("Unknown opcode %#02x", opCode)
	}

	args := []int{}
	data = data << 8
	for _, s := range opDefinition.argLengths {
		switch s {
		case lbyte:
			arg := (data & 0xFF000000) >> 24
			args = append(args, int(arg))
			data = data << 8
		case lword:
			arg := bits.FlipWord(uint16((data & 0xFFFF0000) >> 16))
			args = append(args, int(arg))
			data = data << 16
		default:
			panic("Unknown argument type")
		}
	}

	op := op{
		mnemonic: opDefinition.mnemonic,
		args:     args,
		length:   opDefinition.length,
		handler:  opDefinition.handler,
	}
	return op, nil
}

func (cpu *CPU) opCodeAt(address uint16) (op, error) {
	data := cpu.memoryReadDWord(address)
	return opCodeFrom(data)
}

func (cpu *CPU) nop() int {
	return 4
}

func (cpu *CPU) ldR8aR16(r1 *ByteRegister, r2 *WordRegister) int {
	r1.Set(cpu.memoryReadByte(r2.Get()))
	return 8
}

func (cpu *CPU) lda16R16(a16 uint16, r *WordRegister) int {
	h, l := bits.SplitWord(r.Get())
	cpu.memoryWriteByte(a16, l)
	cpu.memoryWriteByte(a16 + 1, h)
	return 20
}

func (cpu *CPU) ldR16d16(r *WordRegister, d16 uint16) int {
	r.Set(d16)
	return 12
}

func (cpu *CPU) ldaR16R8(r1 *WordRegister, r2 *ByteRegister) int {
	cpu.memoryWriteByte(r1.Get(), r2.Get())
	return 8
}

func (cpu *CPU) ldR8hR8(r1, r2 *ByteRegister) int {
	a16 := 0xFF00 + uint16(r2.Get())
	d8 := cpu.memoryReadByte(a16)
	r1.Set(d8)
	return 8
}

func (cpu *CPU) ldhR8R8(r1, r2 *ByteRegister) int {
	a16 := 0xFF00 + uint16(r1.Get())
	d8 := r2.Get()
	cpu.memoryWriteByte(a16, d8)
	return 8
}

func (cpu *CPU) ldR16R16(r1, r2 *WordRegister) int {
	r1.Set(r2.Get())
	return 8
}

func (cpu *CPU) pushR16(r *WordRegister) int {
	cpu.SP.DecBy(2)
	cpu.memoryWriteWord(cpu.SP.Get(), r.Get())
	return 16
}

func (cpu *CPU) popR16(r *WordRegister) int {
	a16 := cpu.memoryReadWord(cpu.SP.Get())
	r.Set(a16)
	cpu.SP.IncBy(2)
	return 12
}

func (cpu *CPU) ldR16R16a8(r1, r2 *WordRegister, r8 int8) int {
	var carry, halfCarry bool
	if (r8 < 0) {
		halfCarry = bits.HalfCarrySubWord(r2.Get(), -uint16(r8))
		carry = bits.CarrySubWord(r2.Get(), -uint16(r8))
	} else {
		halfCarry = bits.HalfCarryAddWord(r2.Get(), uint16(r8))
		carry = bits.CarryAddWord(r2.Get(), uint16(r8))
	}
	a16 := uint16(int(r2.Get()) + int(r8))
	l := cpu.memoryReadByte(a16)
	h := cpu.memoryReadByte(a16 + 1)
	d16 := bits.ConcatWord(h, l)
	r1.Set(d16)
	cpu.SetFlagZ(false)
	cpu.SetFlagN(false)
	cpu.SetFlagH(halfCarry)
	cpu.SetFlagC(carry)
	return 12
}

func (cpu *CPU) ldha8R8(a8 byte, r2 *ByteRegister) int {
	a16 := 0xFF00 + uint16(a8)
	d8 := r2.Get()
	cpu.memoryWriteByte(a16, d8)
	return 12
}

func (cpu *CPU) ldhR8a8(r *ByteRegister, a8 byte) int {
	a16 := 0xFF00 + uint16(a8)
	d8 := cpu.memoryReadByte(a16)
	r.Set(d8)
	return 12
}

func (cpu *CPU) lda16R8(a16 uint16, r *ByteRegister) int {
	cpu.memoryWriteByte(a16, r.Get())
	return 16
}

func (cpu *CPU) ldR8a16(r *ByteRegister, a16 uint16) int {
	d8 := cpu.memoryReadByte(a16)
	r.Set(d8)
	return 16
}

func (cpu *CPU) ldaR16d8(r1 *WordRegister, d8 byte) int {
	cpu.memoryWriteByte(r1.Get(), d8)
	return 12
}

func (cpu *CPU) lddaR16R8(r1 *WordRegister, r2 *ByteRegister) int {
	cpu.memoryWriteByte(r1.Get(), r2.Get())
	r1.Dec()
	return 8
}

func (cpu *CPU) lddR8aR16(r1 *ByteRegister, r2 *WordRegister) int {
	d8 := cpu.memoryReadByte(r2.Get())
	r1.Set(d8)
	r2.Dec()
	return 8
}

func (cpu *CPU) ldiR8aR16(r1 *ByteRegister, r2 *WordRegister) int {
	d8 := cpu.memoryReadByte(r2.Get())
	r1.Set(d8)
	r2.Inc()
	return 8
}

func (cpu *CPU) ldiaR16R8(r1 *WordRegister, r2 *ByteRegister) int {
	cpu.memoryWriteByte(r1.Get(), r2.Get())
	r1.Inc()
	return 8
}

func (cpu *CPU) ldR8d8(r *ByteRegister, d8 byte) int {
	r.Set(d8)
	return 8
}

func (cpu *CPU) ldR8R8(r1, r2 *ByteRegister) int {
	r1.Set(r2.Get())
	return 4
}

func (cpu *CPU) subR8(r *ByteRegister) int {
	d8 := cpu.A.Get() - r.Get()
	cpu.SetFlagZ(d8 == 0)
	cpu.SetFlagN(true)
	cpu.SetFlagC(d8 >= cpu.A.Get()) // TODO: Is this correct?
	cpu.SetFlagH(false)             // TODO: Fix this
	cpu.A.Set(d8)
	return 4
}

func (cpu *CPU) subd8(v8 byte) int {
	d8 := cpu.A.Get() - v8
	cpu.SetFlagZ(d8 == 0)
	cpu.SetFlagN(true)
	cpu.SetFlagC(d8 >= cpu.A.Get()) // TODO: Is this correct?
	cpu.SetFlagH(false)             // TODO: Fix this
	cpu.A.Set(d8)
	return 8
}

func (cpu *CPU) subaR16(r *WordRegister) int {
	d8 := cpu.A.Get() - cpu.memoryReadByte(r.Get())
	cpu.SetFlagZ(d8 == 0)
	cpu.SetFlagN(true)
	cpu.SetFlagC(d8 >= cpu.A.Get()) // TODO: Is this correct?
	cpu.SetFlagH(false)             // TODO: Fix this
	cpu.A.Set(d8)
	return 8
}

func (cpu *CPU) addR8R8(r1, r2 *ByteRegister) int {
	halfCarry := bits.HalfCarryAddByte(r1.Get(), r2.Get())
	carry := bits.CarryAddByte(r1.Get(), r2.Get())
	r1.Set(r1.Get() + r2.Get())
	cpu.SetFlagZ(r1.Get() == 0)
	cpu.SetFlagN(false)
	cpu.SetFlagC(carry)
	cpu.SetFlagH(halfCarry)
	return 4
}

func (cpu *CPU) addR16R16(r1, r2 *WordRegister) int {
	halfCarry := bits.HalfCarryAddWord(r1.Get(), r2.Get())
	carry := bits.CarryAddWord(r1.Get(), r2.Get())
	r1.Set(r1.Get() + r2.Get())
	cpu.SetFlagZ(r1.Get() == 0)
	cpu.SetFlagN(false)
	cpu.SetFlagC(carry)
	cpu.SetFlagH(halfCarry)
	return 8
}

func (cpu *CPU) addR8d8(r1 *ByteRegister, v8 byte) int {
	halfCarry := bits.HalfCarryAddByte(r1.Get(), v8)
	carry := bits.CarryAddByte(r1.Get(), v8)
	r1.Set(r1.Get() + v8)
	cpu.SetFlagZ(r1.Get() == 0)
	cpu.SetFlagN(false)
	cpu.SetFlagC(carry)
	cpu.SetFlagH(halfCarry)
	return 8
}

func (cpu *CPU) addSP(r8 int8) int {
	var carry, halfCarry bool
	if (r8 < 0) {
		halfCarry = bits.HalfCarrySubWord(cpu.SP.Get(), -uint16(r8))
		carry = bits.CarrySubWord(cpu.SP.Get(), -uint16(r8))
	} else {
		halfCarry = bits.HalfCarryAddWord(cpu.SP.Get(), uint16(r8))
		carry = bits.CarryAddWord(cpu.SP.Get(), uint16(r8))
	}
	cpu.SP.Set(uint16(int(cpu.SP.Get()) + int(r8)))
	cpu.SetFlagN(false)
	cpu.SetFlagZ(false)
	cpu.SetFlagH(halfCarry)
	cpu.SetFlagC(carry)
	return 16
}

func (cpu *CPU) addR8aR16(r1 *ByteRegister, r2 *WordRegister) int {
	d8 := cpu.memoryReadByte(r2.Get())
	halfCarry := bits.HalfCarryAddByte(r1.Get(), d8)
	carry := bits.CarryAddByte(r1.Get(), d8)
	r1.Set(r1.Get() + d8)
	cpu.SetFlagZ(r1.Get() == 0)
	cpu.SetFlagN(false)
	cpu.SetFlagC(carry)
	cpu.SetFlagH(halfCarry)
	return 8
}

func (cpu *CPU) incR8(r *ByteRegister) int {
	halfCarry := bits.HalfCarryAddByte(r.Get(), 1)
	r.Inc()
	cpu.SetFlagN(false)
	cpu.SetFlagZ(r.Get() == 0)
	cpu.SetFlagH(halfCarry)
	return 4
}

func (cpu *CPU) incR16(r *WordRegister) int {
	halfCarry := bits.HalfCarryAddWord(r.Get(), 1)
	r.Inc()
	cpu.SetFlagN(false)
	cpu.SetFlagZ(r.Get() == 0)
	cpu.SetFlagH(halfCarry)
	return 8
}

func (cpu *CPU) decR8(r *ByteRegister) int {
	halfCarry := bits.HalfCarrySubByte(r.Get(), 1)
	r.Dec()
	cpu.SetFlagN(true)
	cpu.SetFlagZ(r.Get() == 0)
	cpu.SetFlagH(halfCarry)
	return 4
}

func (cpu *CPU) decR16(r *WordRegister) int {
	halfCarry := bits.HalfCarrySubByte(r.H().Get(), 1) && bits.HalfCarrySubByte(r.L().Get(), 1)
	r.Dec()
	cpu.SetFlagN(true)
	cpu.SetFlagZ(r.Get() == 0)
	cpu.SetFlagH(halfCarry)
	return 8
}
