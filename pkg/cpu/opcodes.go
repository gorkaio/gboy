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
			cycles := cpu.nop()
			cpu.PC.Inc()
			return cycles
		},
	},
	0x01: {
		mnemonic:   "LD BC, %#04x",
		argLengths: []int{lword},
		length:     3,
		handler: func(cpu *CPU, args ...int) int {
			cycles := cpu.ldR16d16(cpu.BC, uint16(args[0]))
			cpu.PC.IncBy(3)
			return cycles
		},
	},
	0x02: {
		mnemonic:   "LD (BC), A",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			cycles := cpu.ldaR16R8(cpu.BC, cpu.A)
			cpu.PC.Inc()
			return cycles
		},
	},
	0x03: {
		mnemonic:   "INC BC",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			cycles := cpu.incR16(cpu.BC)
			cpu.PC.Inc()
			return cycles
		},
	},
	0x04: {
		mnemonic:   "INC B",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			cycles := cpu.incR8(cpu.B)
			cpu.PC.Inc()
			return cycles
		},
	},
	0x05: {
		mnemonic:   "DEC B",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			cycles := cpu.decR8(cpu.B)
			cpu.PC.Inc()
			return cycles
		},
	},
	0x06: {
		mnemonic:   "LD B, %#02x",
		argLengths: []int{lbyte},
		length:     2,
		handler: func(cpu *CPU, args ...int) int {
			cycles := cpu.ldR8d8(cpu.B, uint8(args[0]))
			cpu.PC.IncBy(2)
			return cycles
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
			cpu.PC.Inc()
			return 4
		},
	},
	0x08: {
		mnemonic:   "LD (%#04x), SP",
		argLengths: []int{lword},
		length:     3,
		handler: func(cpu *CPU, args ...int) int {
			cycles := cpu.lda16R16(uint16(args[0]), cpu.SP)
			cpu.PC.IncBy(3)
			return cycles
		},
	},
	0x09: {
		mnemonic:   "ADD HL, BC",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			cycles := cpu.addR16R16(cpu.HL, cpu.BC)
			cpu.PC.Inc()
			return cycles
		},
	},
	0x0A: {
		mnemonic:   "LD A, (BC)",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			cycles := cpu.ldR8aR16(cpu.A, cpu.BC)
			cpu.PC.Inc()
			return cycles
		},
	},
	0x0B: {
		mnemonic:   "DEC BC",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			cycles := cpu.decR16(cpu.BC)
			cpu.PC.Inc()
			return cycles
		},
	},
	0x0C: {
		mnemonic:   "INC C",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			cycles := cpu.incR8(cpu.C)
			cpu.PC.Inc()
			return cycles
		},
	},
	0x0D: {
		mnemonic:   "DEC C",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			cycles := cpu.decR8(cpu.C)
			cpu.PC.Inc()
			return cycles
		},
	},
	0x0E: {
		mnemonic:   "LD C, %#02x",
		argLengths: []int{lbyte},
		length:     2,
		handler: func(cpu *CPU, args ...int) int {
			cycles := cpu.ldR8d8(cpu.C, uint8(args[0]))
			cpu.PC.IncBy(2)
			return cycles
		},
	},
	/* TODO: 0x0F */
	/* TODO: 0x10 */
	0x11: {
		mnemonic:   "LD DE, %#04x",
		argLengths: []int{lword},
		length:     3,
		handler: func(cpu *CPU, args ...int) int {
			cycles := cpu.ldR16d16(cpu.DE, uint16(args[0]))
			cpu.PC.IncBy(3)
			return cycles
		},
	},
	0x12: {
		mnemonic:   "LD (DE), A",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			cycles := cpu.ldaR16R8(cpu.DE, cpu.A)
			cpu.PC.Inc()
			return cycles
		},
	},
	0x13: {
		mnemonic:   "INC DE",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			cycles := cpu.incR16(cpu.DE)
			cpu.PC.Inc()
			return cycles
		},
	},
	0x14: {
		mnemonic:   "INC D",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			cycles := cpu.incR8(cpu.D)
			cpu.PC.Inc()
			return cycles
		},
	},
	0x15: {
		mnemonic:   "DEC D",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			cycles := cpu.decR8(cpu.D)
			cpu.PC.Inc()
			return cycles
		},
	},
	0x16: {
		mnemonic:   "LD D, %#02x",
		argLengths: []int{lbyte},
		length:     2,
		handler: func(cpu *CPU, args ...int) int {
			cycles := cpu.ldR8d8(cpu.D, uint8(args[0]))
			cpu.PC.IncBy(2)
			return cycles
		},
	},
	/* TODO: 0x17 */
	0x18: {
		mnemonic:   "JR %#02x",
		argLengths: []int{lbyte},
		length:     2,
		handler: func(cpu *CPU, args ...int) int {
			return cpu.jr(int8(args[0]))
		},
	},
	0x19: {
		mnemonic:   "ADD HL, DE",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			cycles := cpu.addR16R16(cpu.HL, cpu.DE)
			cpu.PC.Inc()
			return cycles
		},
	},
	0x1A: {
		mnemonic:   "LD A, (DE)",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			cycles := cpu.ldR8aR16(cpu.A, cpu.DE)
			cpu.PC.Inc()
			return cycles
		},
	},
	0x1B: {
		mnemonic:   "DEC DE",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			cycles := cpu.decR16(cpu.DE)
			cpu.PC.Inc()
			return cycles
		},
	},
	0x1C: {
		mnemonic:   "INC E",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			cycles := cpu.incR8(cpu.E)
			cpu.PC.Inc()
			return cycles
		},
	},
	0x1D: {
		mnemonic:   "DEC E",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			cycles := cpu.decR8(cpu.E)
			cpu.PC.Inc()
			return cycles
		},
	},
	0x1E: {
		mnemonic:   "LD E, %#02x",
		argLengths: []int{lbyte},
		length:     2,
		handler: func(cpu *CPU, args ...int) int {
			cycles := cpu.ldR8d8(cpu.E, uint8(args[0]))
			cpu.PC.IncBy(2)
			return cycles
		},
	},
	/* TODO: 0x1F */
	0x20: {
		mnemonic:   "JR NZ, %#02x",
		argLengths: []int{lbyte},
		length:     2,
		handler: func(cpu *CPU, args ...int) int {
			if !cpu.FlagZ() {
				return cpu.jr(int8(args[0]))
			}
			cpu.PC.IncBy(2)
			return 8
		},
	},
	0x21: {
		mnemonic:   "LD HL, %#04x",
		argLengths: []int{lword},
		length:     3,
		handler: func(cpu *CPU, args ...int) int {
			cycles := cpu.ldR16d16(cpu.HL, uint16(args[0]))
			cpu.PC.IncBy(3)
			return cycles
		},
	},
	0x22: {
		mnemonic:   "LDI (HL), A",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			cycles := cpu.ldiaR16R8(cpu.HL, cpu.A)
			cpu.PC.Inc()
			return cycles
		},
	},
	0x23: {
		mnemonic:   "INC HL",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			cycles :=  cpu.incR16(cpu.HL)
			cpu.PC.Inc()
			return cycles
		},
	},
	0x24: {
		mnemonic:   "INC H",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			cycles := cpu.incR8(cpu.H)
			cpu.PC.Inc()
			return cycles
		},
	},
	0x25: {
		mnemonic:   "DEC H",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			cycles := cpu.decR8(cpu.H)
			cpu.PC.Inc()
			return cycles
		},
	},
	0x26: {
		mnemonic:   "LD H, %#02x",
		argLengths: []int{lbyte},
		length:     2,
		handler: func(cpu *CPU, args ...int) int {
			cycles := cpu.ldR8d8(cpu.H, uint8(args[0]))
			cpu.PC.IncBy(2)
			return cycles
		},
	},
	/* TODO: 0x27 */
	0x28: {
		mnemonic:   "JR Z, %#02x",
		argLengths: []int{lbyte},
		length:     2,
		handler: func(cpu *CPU, args ...int) int {
			if cpu.FlagZ() {
				return cpu.jr(int8(args[0]))
			}
			cpu.PC.IncBy(2)
			return 8
		},
	},
	0x29: {
		mnemonic:   "ADD HL, HL",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			cycles := cpu.addR16R16(cpu.HL, cpu.HL)
			cpu.PC.Inc()
			return cycles
		},
	},
	0x2A: {
		mnemonic:   "LDI A, (HL)",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			cycles := cpu.ldiR8aR16(cpu.A, cpu.HL)
			cpu.PC.Inc()
			return cycles
		},
	},
	0x2B: {
		mnemonic:   "DEC HL",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			cycles := cpu.decR16(cpu.HL)
			cpu.PC.Inc()
			return cycles
		},
	},
	0x2C: {
		mnemonic:   "INC L",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			cycles := cpu.incR8(cpu.L)
			cpu.PC.Inc()
			return cycles
		},
	},
	0x2D: {
		mnemonic:   "DEC L",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			cycles := cpu.decR8(cpu.L)
			cpu.PC.Inc()
			return cycles
		},
	},
	0x2E: {
		mnemonic:   "LD L, %#02x",
		argLengths: []int{lbyte},
		length:     2,
		handler: func(cpu *CPU, args ...int) int {
			cycles := cpu.ldR8d8(cpu.L, uint8(args[0]))
			cpu.PC.IncBy(2)
			return cycles
		},
	},
	/* TODO: 0x2F */
	0x30: {
		mnemonic:   "JR NC, %#02x",
		argLengths: []int{lbyte},
		length:     2,
		handler: func(cpu *CPU, args ...int) int {
			if !cpu.FlagC() {
				return cpu.jr(int8(args[0]))
			}
			cpu.PC.IncBy(2)
			return 8
		},
	},
	0x31: {
		mnemonic:   "LD SP, %#04x",
		argLengths: []int{lword},
		length:     3,
		handler: func(cpu *CPU, args ...int) int {
			cycles := cpu.ldR16d16(cpu.SP, uint16(args[0]))
			cpu.PC.IncBy(3)
			return cycles
		},
	},
	0x32: {
		mnemonic:   "LDD (HL), A",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			cycles := cpu.lddaR16R8(cpu.HL, cpu.A)
			cpu.PC.Inc()
			return cycles
		},
	},
	0x33: {
		mnemonic:   "INC SP",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			cycles := cpu.incR16(cpu.SP)
			cpu.PC.Inc()
			return cycles
		},
	},
	0x34: {
		mnemonic:   "INC (HL)",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			a := cpu.HL.Get()
			d8 := cpu.memoryReadByte(cpu.HL.Get())
			result := d8 + 1
			halfCarry := bits.HalfCarryAddByte(d8, 1)
			cpu.memoryWriteByte(a, result)
			cpu.SetFlagN(false)
			cpu.SetFlagZ(result == 0)
			cpu.SetFlagH(halfCarry)
			cpu.PC.Inc()
			return 12
		},
	},
	0x35: {
		mnemonic:   "DEC (HL)",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			a := cpu.HL.Get()
			d8 := cpu.memoryReadByte(cpu.HL.Get())
			result := d8 - 1
			halfCarry := bits.HalfCarrySubByte(d8, 1)
			cpu.memoryWriteByte(a, result)
			cpu.SetFlagN(true)
			cpu.SetFlagZ(result == 0)
			cpu.SetFlagH(halfCarry)
			cpu.PC.Inc()
			return 12
		},
	},
	0x36: {
		mnemonic:   "LD (HL), %#02x",
		argLengths: []int{lbyte},
		length:     2,
		handler: func(cpu *CPU, args ...int) int {
			d8 := byte(args[0])
			cycles := cpu.ldaR16d8(cpu.HL, d8)
			cpu.PC.IncBy(2)
			return cycles
		},
	},
	/* TODO: 0x37 */
	0x38: {
		mnemonic:   "JR C, %#02x",
		argLengths: []int{lbyte},
		length:     2,
		handler: func(cpu *CPU, args ...int) int {
			if cpu.FlagC() {
				return cpu.jr(int8(args[0]))
			}
			cpu.PC.IncBy(2)
			return 8
		},
	},
	0x39: {
		mnemonic:   "ADD HL, SP",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			cycles := cpu.addR16R16(cpu.HL, cpu.SP)
			cpu.PC.Inc()
			return cycles
		},
	},
	0x3A: {
		mnemonic:   "LDD A, (HL)",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			cycles := cpu.lddR8aR16(cpu.A, cpu.HL)
			cpu.PC.Inc()
			return cycles
		},
	},
	0x3B: {
		mnemonic:   "DEC SP",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			cycles := cpu.decR16(cpu.SP)
			cpu.PC.Inc()
			return cycles
		},
	},
	0x3C: {
		mnemonic:   "INC A",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			cycles := cpu.incR8(cpu.A)
			cpu.PC.Inc()
			return cycles
		},
	},
	0x3D: {
		mnemonic:   "DEC A",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			cycles := cpu.decR8(cpu.A)
			cpu.PC.Inc()
			return cycles
		},
	},
	0x3E: {
		mnemonic:   "LD A, %#02x",
		argLengths: []int{lbyte},
		length:     2,
		handler: func(cpu *CPU, args ...int) int {
			d8 := byte(args[0])
			cycles := cpu.ldR8d8(cpu.A, d8)
			cpu.PC.IncBy(2)
			return cycles
		},
	},
	0x40: {
		mnemonic:   "LD B, B",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			cycles := cpu.ldR8R8(cpu.B, cpu.B)
			cpu.PC.Inc()
			return cycles
		},
	},
	0x41: {
		mnemonic:   "LD B, C",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			cycles := cpu.ldR8R8(cpu.B, cpu.C)
			cpu.PC.Inc()
			return cycles
		},
	},
	0x42: {
		mnemonic:   "LD B, D",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			cycles := cpu.ldR8R8(cpu.B, cpu.D)
			cpu.PC.Inc()
			return cycles
		},
	},
	0x43: {
		mnemonic:   "LD B, E",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			cycles := cpu.ldR8R8(cpu.B, cpu.E)
			cpu.PC.Inc()
			return cycles
		},
	},
	0x44: {
		mnemonic:   "LD B, H",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			cycles := cpu.ldR8R8(cpu.B, cpu.H)
			cpu.PC.Inc()
			return cycles
		},
	},
	0x45: {
		mnemonic:   "LD B, L",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			cycles := cpu.ldR8R8(cpu.B, cpu.L)
			cpu.PC.Inc()
			return cycles
		},
	},
	0x46: {
		mnemonic:   "LD B, (HL)",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			cycles := cpu.ldR8aR16(cpu.B, cpu.HL)
			cpu.PC.Inc()
			return cycles
		},
	},
	0x47: {
		mnemonic:   "LD B, A",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			cycles := cpu.ldR8R8(cpu.B, cpu.A)
			cpu.PC.Inc()
			return cycles
		},
	},
	0x48: {
		mnemonic:   "LD C, B",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			cycles := cpu.ldR8R8(cpu.C, cpu.B)
			cpu.PC.Inc()
			return cycles
		},
	},
	0x49: {
		mnemonic:   "LD C, C",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			cycles := cpu.ldR8R8(cpu.C, cpu.C)
			cpu.PC.Inc()
			return cycles
		},
	},
	0x4A: {
		mnemonic:   "LD C, D",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			cycles := cpu.ldR8R8(cpu.C, cpu.D)
			cpu.PC.Inc()
			return cycles
		},
	},
	0x4B: {
		mnemonic:   "LD C, E",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			cycles := cpu.ldR8R8(cpu.C, cpu.E)
			cpu.PC.Inc()
			return cycles
		},
	},
	0x4C: {
		mnemonic:   "LD C, H",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			cycles := cpu.ldR8R8(cpu.C, cpu.H)
			cpu.PC.Inc()
			return cycles
		},
	},
	0x4D: {
		mnemonic:   "LD C, L",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			cycles := cpu.ldR8R8(cpu.C, cpu.L)
			cpu.PC.Inc()
			return cycles
		},
	},
	0x4E: {
		mnemonic:   "LD C, (HL)",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			cycles := cpu.ldR8aR16(cpu.C, cpu.HL)
			cpu.PC.Inc()
			return cycles
		},
	},
	0x4F: {
		mnemonic:   "LD C, A",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			cycles := cpu.ldR8R8(cpu.C, cpu.A)
			cpu.PC.Inc()
			return cycles
		},
	},
	0x50: {
		mnemonic:   "LD D, B",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			cycles := cpu.ldR8R8(cpu.D, cpu.B)
			cpu.PC.Inc()
			return cycles
		},
	},
	0x51: {
		mnemonic:   "LD D, C",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			cycles := cpu.ldR8R8(cpu.D, cpu.C)
			cpu.PC.Inc()
			return cycles
		},
	},
	0x52: {
		mnemonic:   "LD D, D",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			cycles := cpu.ldR8R8(cpu.D, cpu.D)
			cpu.PC.Inc()
			return cycles
		},
	},
	0x53: {
		mnemonic:   "LD D, E",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			cycles := cpu.ldR8R8(cpu.D, cpu.E)
			cpu.PC.Inc()
			return cycles
		},
	},
	0x54: {
		mnemonic:   "LD D, H",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			cycles := cpu.ldR8R8(cpu.D, cpu.H)
			cpu.PC.Inc()
			return cycles
		},
	},
	0x55: {
		mnemonic:   "LD D, L",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			cycles := cpu.ldR8R8(cpu.D, cpu.L)
			cpu.PC.Inc()
			return cycles
		},
	},
	0x56: {
		mnemonic:   "LD D, (HL)",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			cycles := cpu.ldR8aR16(cpu.D, cpu.HL)
			cpu.PC.Inc()
			return cycles
		},
	},
	0x57: {
		mnemonic:   "LD D, A",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			cycles := cpu.ldR8R8(cpu.D, cpu.A)
			cpu.PC.Inc()
			return cycles
		},
	},
	0x58: {
		mnemonic:   "LD E, B",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			cycles := cpu.ldR8R8(cpu.E, cpu.B)
			cpu.PC.Inc()
			return cycles
		},
	},
	0x59: {
		mnemonic:   "LD E, C",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			cycles := cpu.ldR8R8(cpu.E, cpu.C)
			cpu.PC.Inc()
			return cycles
		},
	},
	0x5A: {
		mnemonic:   "LD E, D",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			cycles := cpu.ldR8R8(cpu.E, cpu.D)
			cpu.PC.Inc()
			return cycles
		},
	},
	0x5B: {
		mnemonic:   "LD E, E",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			cycles := cpu.ldR8R8(cpu.E, cpu.E)
			cpu.PC.Inc()
			return cycles
		},
	},
	0x5C: {
		mnemonic:   "LD E, H",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			cycles := cpu.ldR8R8(cpu.E, cpu.H)
			cpu.PC.Inc()
			return cycles
		},
	},
	0x5D: {
		mnemonic:   "LD E, L",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			cycles := cpu.ldR8R8(cpu.E, cpu.L)
			cpu.PC.Inc()
			return cycles
		},
	},
	0x5E: {
		mnemonic:   "LD E, (HL)",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			cycles := cpu.ldR8aR16(cpu.E, cpu.HL)
			cpu.PC.Inc()
			return cycles
		},
	},
	0x5F: {
		mnemonic:   "LD E, A",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			cycles := cpu.ldR8R8(cpu.E, cpu.A)
			cpu.PC.Inc()
			return cycles
		},
	},
	0x60: {
		mnemonic:   "LD H, B",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			cycles := cpu.ldR8R8(cpu.H, cpu.B)
			cpu.PC.Inc()
			return cycles
		},
	},
	0x61: {
		mnemonic:   "LD H, C",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			cycles := cpu.ldR8R8(cpu.H, cpu.C)
			cpu.PC.Inc()
			return cycles
		},
	},
	0x62: {
		mnemonic:   "LD H, D",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			cycles := cpu.ldR8R8(cpu.H, cpu.D)
			cpu.PC.Inc()
			return cycles
		},
	},
	0x63: {
		mnemonic:   "LD H, E",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			cycles := cpu.ldR8R8(cpu.H, cpu.E)
			cpu.PC.Inc()
			return cycles
		},
	},
	0x64: {
		mnemonic:   "LD H, H",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			cycles := cpu.ldR8R8(cpu.H, cpu.H)
			cpu.PC.Inc()
			return cycles
		},
	},
	0x65: {
		mnemonic:   "LD H, L",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			cycles := cpu.ldR8R8(cpu.H, cpu.L)
			cpu.PC.Inc()
			return cycles
		},
	},
	0x66: {
		mnemonic:   "LD H, (HL)",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			cycles := cpu.ldR8aR16(cpu.H, cpu.HL)
			cpu.PC.Inc()
			return cycles
		},
	},
	0x67: {
		mnemonic:   "LD H, A",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			cycles := cpu.ldR8R8(cpu.H, cpu.A)
			cpu.PC.Inc()
			return cycles
		},
	},
	0x68: {
		mnemonic:   "LD L, B",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			cycles := cpu.ldR8R8(cpu.L, cpu.B)
			cpu.PC.Inc()
			return cycles
		},
	},
	0x69: {
		mnemonic:   "LD L, C",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			cycles := cpu.ldR8R8(cpu.L, cpu.C)
			cpu.PC.Inc()
			return cycles
		},
	},
	0x6A: {
		mnemonic:   "LD L, D",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			cycles := cpu.ldR8R8(cpu.L, cpu.D)
			cpu.PC.Inc()
			return cycles
		},
	},
	0x6B: {
		mnemonic:   "LD L, E",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			cycles := cpu.ldR8R8(cpu.L, cpu.E)
			cpu.PC.Inc()
			return cycles
		},
	},
	0x6C: {
		mnemonic:   "LD L, H",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			cycles := cpu.ldR8R8(cpu.L, cpu.H)
			cpu.PC.Inc()
			return cycles
		},
	},
	0x6D: {
		mnemonic:   "LD L, L",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			cycles := cpu.ldR8R8(cpu.L, cpu.L)
			cpu.PC.Inc()
			return cycles
		},
	},
	0x6E: {
		mnemonic:   "LD L, (HL)",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			cycles := cpu.ldR8aR16(cpu.L, cpu.HL)
			cpu.PC.Inc()
			return cycles
		},
	},
	0x6F: {
		mnemonic:   "LD L, A",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			cycles := cpu.ldR8R8(cpu.L, cpu.A)
			cpu.PC.Inc()
			return cycles
		},
	},
	0x70: {
		mnemonic:   "LD (HL), B",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			cycles := cpu.ldaR16R8(cpu.HL, cpu.B)
			cpu.PC.Inc()
			return cycles
		},
	},
	0x71: {
		mnemonic:   "LD (HL), C",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			cycles := cpu.ldaR16R8(cpu.HL, cpu.C)
			cpu.PC.Inc()
			return cycles
		},
	},
	0x72: {
		mnemonic:   "LD (HL), D",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			cycles := cpu.ldaR16R8(cpu.HL, cpu.D)
			cpu.PC.Inc()
			return cycles
		},
	},
	0x73: {
		mnemonic:   "LD (HL), E",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			cycles := cpu.ldaR16R8(cpu.HL, cpu.E)
			cpu.PC.Inc()
			return cycles
		},
	},
	0x74: {
		mnemonic:   "LD (HL), H",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			cycles := cpu.ldaR16R8(cpu.HL, cpu.H)
			cpu.PC.Inc()
			return cycles
		},
	},
	0x75: {
		mnemonic:   "LD (HL), L",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			cycles := cpu.ldaR16R8(cpu.HL, cpu.L)
			cpu.PC.Inc()
			return cycles
		},
	},
	/* TODO: 0x76 */
	0x77: {
		mnemonic:   "LD (HL), A",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			cycles := cpu.ldaR16R8(cpu.HL, cpu.A)
			cpu.PC.Inc()
			return cycles
		},
	},
	0x78: {
		mnemonic:   "LD A, B",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			cycles := cpu.ldR8R8(cpu.A, cpu.B)
			cpu.PC.Inc()
			return cycles
		},
	},
	0x79: {
		mnemonic:   "LD A, C",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			cycles := cpu.ldR8R8(cpu.A, cpu.C)
			cpu.PC.Inc()
			return cycles
		},
	},
	0x7A: {
		mnemonic:   "LD A, D",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			cycles := cpu.ldR8R8(cpu.A, cpu.D)
			cpu.PC.Inc()
			return cycles
		},
	},
	0x7B: {
		mnemonic:   "LD A, E",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			cycles := cpu.ldR8R8(cpu.A, cpu.E)
			cpu.PC.Inc()
			return cycles
		},
	},
	0x7C: {
		mnemonic:   "LD A, H",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			cycles := cpu.ldR8R8(cpu.A, cpu.H)
			cpu.PC.Inc()
			return cycles
		},
	},
	0x7D: {
		mnemonic:   "LD A, L",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			cycles := cpu.ldR8R8(cpu.A, cpu.L)
			cpu.PC.Inc()
			return cycles
		},
	},
	0x7E: {
		mnemonic:   "LD A, (HL)",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			cycles := cpu.ldR8aR16(cpu.A, cpu.HL)
			cpu.PC.Inc()
			return cycles
		},
	},
	0x7F: {
		mnemonic:   "LD A, A",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			cycles := cpu.ldR8R8(cpu.A, cpu.A)
			cpu.PC.Inc()
			return cycles
		},
	},
	0x80: {
		mnemonic:   "ADD A, B",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			cycles := cpu.addR8R8(cpu.A, cpu.B)
			cpu.PC.Inc()
			return cycles
		},
	},
	0x81: {
		mnemonic:   "ADD A, C",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			cycles := cpu.addR8R8(cpu.A, cpu.C)
			cpu.PC.Inc()
			return cycles
		},
	},
	0x82: {
		mnemonic:   "ADD A, D",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			cycles := cpu.addR8R8(cpu.A, cpu.D)
			cpu.PC.Inc()
			return cycles
		},
	},
	0x83: {
		mnemonic:   "ADD A, E",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			cycles := cpu.addR8R8(cpu.A, cpu.E)
			cpu.PC.Inc()
			return cycles
		},
	},
	0x84: {
		mnemonic:   "ADD A, H",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			cycles := cpu.addR8R8(cpu.A, cpu.H)
			cpu.PC.Inc()
			return cycles
		},
	},
	0x85: {
		mnemonic:   "ADD A, L",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			cycles := cpu.addR8R8(cpu.A, cpu.L)
			cpu.PC.Inc()
			return cycles
		},
	},
	0x86: {
		mnemonic:   "ADD A, (HL)",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			cycles := cpu.addR8aR16(cpu.A, cpu.HL)
			cpu.PC.Inc()
			return cycles
		},
	},
	0x87: {
		mnemonic:   "ADD A, A",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			cycles := cpu.addR8R8(cpu.A, cpu.A)
			cpu.PC.Inc()
			return cycles
		},
	},
	0x88: {
		mnemonic:   "ADC A, B",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			cycles := cpu.adcR8R8(cpu.A, cpu.B)
			cpu.PC.Inc()
			return cycles
		},
	},
	0x89: {
		mnemonic:   "ADC A, C",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			cycles := cpu.adcR8R8(cpu.A, cpu.C)
			cpu.PC.Inc()
			return cycles
		},
	},
	0x8A: {
		mnemonic:   "ADC A, D",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			cycles := cpu.adcR8R8(cpu.A, cpu.D)
			cpu.PC.Inc()
			return cycles
		},
	},
	0x8B: {
		mnemonic:   "ADC A, E",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			cycles := cpu.adcR8R8(cpu.A, cpu.E)
			cpu.PC.Inc()
			return cycles
		},
	},
	0x8C: {
		mnemonic:   "ADC A, H",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			cycles := cpu.adcR8R8(cpu.A, cpu.H)
			cpu.PC.Inc()
			return cycles
		},
	},
	0x8D: {
		mnemonic:   "ADC A, L",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			cycles := cpu.adcR8R8(cpu.A, cpu.L)
			cpu.PC.Inc()
			return cycles
		},
	},
	0x8E: {
		mnemonic:   "ADC A, (HL)",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			cycles := cpu.adcR8aR16(cpu.A, cpu.HL)
			cpu.PC.Inc()
			return cycles
		},
	},
	0x8F: {
		mnemonic:   "ADC A, A",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			cycles := cpu.adcR8R8(cpu.A, cpu.A)
			cpu.PC.Inc()
			return cycles
		},
	},
	0x90: {
		mnemonic:   "SUB B",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			cycles := cpu.subR8(cpu.B)
			cpu.PC.Inc()
			return cycles
		},
	},
	0x91: {
		mnemonic:   "SUB C",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			cycles := cpu.subR8(cpu.C)
			cpu.PC.Inc()
			return cycles
		},
	},
	0x92: {
		mnemonic:   "SUB D",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			cycles := cpu.subR8(cpu.D)
			cpu.PC.Inc()
			return cycles
		},
	},
	0x93: {
		mnemonic:   "SUB E",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			cycles := cpu.subR8(cpu.E)
			cpu.PC.Inc()
			return cycles
		},
	},
	0x94: {
		mnemonic:   "SUB H",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			cycles := cpu.subR8(cpu.H)
			cpu.PC.Inc()
			return cycles
		},
	},
	0x95: {
		mnemonic:   "SUB L",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			cycles := cpu.subR8(cpu.L)
			cpu.PC.Inc()
			return cycles
		},
	},
	0x96: {
		mnemonic:   "SUB (HL)",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			cycles := cpu.subaR16(cpu.HL)
			cpu.PC.Inc()
			return cycles
		},
	},
	0x97: {
		mnemonic:   "SUB A",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			cycles := cpu.subR8(cpu.A)
			cpu.PC.Inc()
			return cycles
		},
	},
	0x98: {
		mnemonic:   "SBC B",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			cycles := cpu.sbcR8(cpu.B)
			cpu.PC.Inc()
			return cycles
		},
	},
	0x99: {
		mnemonic:   "SBC C",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			cycles := cpu.sbcR8(cpu.C)
			cpu.PC.Inc()
			return cycles
		},
	},
	0x9A: {
		mnemonic:   "SBC D",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			cycles := cpu.sbcR8(cpu.D)
			cpu.PC.Inc()
			return cycles
		},
	},
	0x9B: {
		mnemonic:   "SBC E",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			cycles := cpu.sbcR8(cpu.E)
			cpu.PC.Inc()
			return cycles
		},
	},
	0x9C: {
		mnemonic:   "SBC H",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			cycles := cpu.sbcR8(cpu.H)
			cpu.PC.Inc()
			return cycles
		},
	},
	0x9D: {
		mnemonic:   "SBC L",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			cycles := cpu.sbcR8(cpu.L)
			cpu.PC.Inc()
			return cycles
		},
	},
	0x9E: {
		mnemonic:   "SBC (HL)",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			cycles := cpu.sbcR8aR16(cpu.A, cpu.HL)
			cpu.PC.Inc()
			return cycles
		},
	},
	0x9F: {
		mnemonic:   "SBC A",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			cycles := cpu.sbcR8(cpu.A)
			cpu.PC.Inc()
			return cycles
		},
	},
	0xA0: {
		mnemonic:   "AND B",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			cycles := cpu.andR8(cpu.B)
			cpu.PC.Inc()
			return cycles
		},
	},
	0xA1: {
		mnemonic:   "AND C",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			cycles := cpu.andR8(cpu.C)
			cpu.PC.Inc()
			return cycles
		},
	},
	0xA2: {
		mnemonic:   "AND D",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			cycles := cpu.andR8(cpu.D)
			cpu.PC.Inc()
			return cycles
		},
	},
	0xA3: {
		mnemonic:   "AND E",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			cycles := cpu.andR8(cpu.E)
			cpu.PC.Inc()
			return cycles
		},
	},
	0xA4: {
		mnemonic:   "AND H",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			cycles := cpu.andR8(cpu.H)
			cpu.PC.Inc()
			return cycles
		},
	},
	0xA5: {
		mnemonic:   "AND L",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			cycles := cpu.andR8(cpu.L)
			cpu.PC.Inc()
			return cycles
		},
	},
	0xA6: {
		mnemonic:   "AND (HL)",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			cycles := cpu.andaR16(cpu.HL)
			cpu.PC.Inc()
			return cycles
		},
	},
	0xA7: {
		mnemonic:   "AND A",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			cycles := cpu.andR8(cpu.A)
			cpu.PC.Inc()
			return cycles
		},
	},
	0xA8: {
		mnemonic:   "XOR B",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			cycles := cpu.xorR8(cpu.B)
			cpu.PC.Inc()
			return cycles
		},
	},
	0xA9: {
		mnemonic:   "XOR C",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			cycles := cpu.xorR8(cpu.C)
			cpu.PC.Inc()
			return cycles
		},
	},
	0xAA: {
		mnemonic:   "XOR D",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			cycles := cpu.xorR8(cpu.D)
			cpu.PC.Inc()
			return cycles
		},
	},
	0xAB: {
		mnemonic:   "XOR E",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			cycles := cpu.xorR8(cpu.E)
			cpu.PC.Inc()
			return cycles
		},
	},
	0xAC: {
		mnemonic:   "XOR H",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			cycles := cpu.xorR8(cpu.H)
			cpu.PC.Inc()
			return cycles
		},
	},
	0xAD: {
		mnemonic:   "XOR L",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			cycles := cpu.xorR8(cpu.L)
			cpu.PC.Inc()
			return cycles
		},
	},
	0xAE: {
		mnemonic:   "XOR (HL)",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			cycles := cpu.xoraR16(cpu.HL)
			cpu.PC.Inc()
			return cycles
		},
	},
	0xAF: {
		mnemonic:   "XOR A",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			cycles := cpu.xorR8(cpu.A)
			cpu.PC.Inc()
			return cycles
		},
	},
	0xB0: {
		mnemonic:   "OR B",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			cycles := cpu.orR8(cpu.B)
			cpu.PC.Inc()
			return cycles
		},
	},
	0xB1: {
		mnemonic:   "OR C",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			cycles := cpu.orR8(cpu.C)
			cpu.PC.Inc()
			return cycles
		},
	},
	0xB2: {
		mnemonic:   "OR D",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			cycles := cpu.orR8(cpu.D)
			cpu.PC.Inc()
			return cycles
		},
	},
	0xB3: {
		mnemonic:   "OR E",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			cycles := cpu.orR8(cpu.E)
			cpu.PC.Inc()
			return cycles
		},
	},
	0xB4: {
		mnemonic:   "OR H",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			cycles := cpu.orR8(cpu.H)
			cpu.PC.Inc()
			return cycles
		},
	},
	0xB5: {
		mnemonic:   "OR L",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			cycles := cpu.orR8(cpu.L)
			cpu.PC.Inc()
			return cycles
		},
	},
	0xB6: {
		mnemonic:   "OR (HL)",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			cycles := cpu.oraR16(cpu.HL)
			cpu.PC.Inc()
			return cycles
		},
	},
	0xB7: {
		mnemonic:   "OR A",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			cycles := cpu.orR8(cpu.A)
			cpu.PC.Inc()
			return cycles
		},
	},
	0xB8: {
		mnemonic:   "CP B",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			cycles := cpu.cpR8(cpu.B)
			cpu.PC.Inc()
			return cycles
		},
	},
	0xB9: {
		mnemonic:   "CP C",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			cycles := cpu.cpR8(cpu.C)
			cpu.PC.Inc()
			return cycles
		},
	},
	0xBA: {
		mnemonic:   "CP D",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			cycles := cpu.cpR8(cpu.D)
			cpu.PC.Inc()
			return cycles
		},
	},
	0xBB: {
		mnemonic:   "CP E",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			cycles := cpu.cpR8(cpu.E)
			cpu.PC.Inc()
			return cycles
		},
	},
	0xBC: {
		mnemonic:   "CP H",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			cycles := cpu.cpR8(cpu.H)
			cpu.PC.Inc()
			return cycles
		},
	},
	0xBD: {
		mnemonic:   "CP L",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			cycles := cpu.cpR8(cpu.L)
			cpu.PC.Inc()
			return cycles
		},
	},
	0xBE: {
		mnemonic:   "CP (HL)",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			cycles := cpu.cpA8(cpu.L)
			cpu.PC.Inc()
			return cycles
		},
	},
	0xBF: {
		mnemonic:   "CP A",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			cycles := cpu.cpR8(cpu.A)
			cpu.PC.Inc()
			return cycles
		},
	},
	0xC0: {
		mnemonic:   "RETNZ",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			if (cpu.FlagZ()) {
				cpu.PC.Inc()
				return 8
			}
			addr := cpu.memoryReadWord(cpu.SP.Get())
			cpu.PC.Set(addr)
			cpu.SP.IncBy(2)
			return 20
		},
	},
	0xC1: {
		mnemonic:   "POP BC",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			cycles := cpu.popR16(cpu.BC)
			cpu.PC.Inc()
			return cycles
		},
	},
	0xC2: {
		mnemonic:   "JR NZ, %#04x",
		argLengths: []int{lword},
		length:     3,
		handler: func(cpu *CPU, args ...int) int {
			if !cpu.FlagZ() {
				return cpu.jmp(uint16(args[0]))
			}
			cpu.PC.IncBy(3)
			return 12
		},
	},
	0xC3: {
		mnemonic:   "JMP %#04x",
		argLengths: []int{lword},
		length:     3,
		handler: func(cpu *CPU, args ...int) int {
			return cpu.jmp(uint16(args[0]))
		},
	},
	0xC4: {
		mnemonic: "CALL NZ, %#04x",
		argLengths: []int{lword},
		length: 3,
		handler: func(cpu *CPU, args ...int) int {
			if cpu.FlagZ() {
				cpu.PC.IncBy(3)
				return 12
			}
			cpu.pushR16(cpu.PC)
			cpu.PC.Set(uint16(args[0]))
			return 24
		},
	},
	0xC5: {
		mnemonic:   "PUSH BC",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			cycles := cpu.pushR16(cpu.BC)
			cpu.PC.Inc()
			return cycles
		},
	},
	0xC6: {
		mnemonic:   "ADD %#02x",
		argLengths: []int{lbyte},
		length:     2,
		handler: func(cpu *CPU, args ...int) int {
			cycles := cpu.addR8d8(cpu.A, byte(args[0]))
			cpu.PC.IncBy(2)
			return cycles
		},
	},
	0xC7: {
		mnemonic:   "RST 00H",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			return cpu.rst(0x0000)
		},
	},
	0xC8: {
		mnemonic:   "RETZ",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			if (!cpu.FlagZ()) {
				cpu.PC.Inc()
				return 8
			}
			addr := cpu.memoryReadWord(cpu.SP.Get())
			cpu.PC.Set(addr)
			cpu.SP.IncBy(2)
			return 20
		},
	},
	0xC9: {
		mnemonic:   "RET",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			addr := cpu.memoryReadWord(cpu.SP.Get())
			cpu.PC.Set(addr)
			cpu.SP.IncBy(2)
			return 16
		},
	},
	0xCA: {
		mnemonic:   "JP Z, %#04x",
		argLengths: []int{lword},
		length:     3,
		handler: func(cpu *CPU, args ...int) int {
			if cpu.FlagZ() {
				return cpu.jmp(uint16(args[0]))
			}
			cpu.PC.IncBy(3)
			return 12
		},
	},
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
		mnemonic:   "ADC %#02x",
		argLengths: []int{lbyte},
		length:     2,
		handler: func(cpu *CPU, args ...int) int {
			cycles := cpu.adcR8d8(cpu.A, byte(args[0]))
			cpu.PC.IncBy(2)
			return cycles
		},
	},
	0xCF: {
		mnemonic:   "RST 08H",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			return cpu.rst(0x0008)
		},
	},
	0xD0: {
		mnemonic:   "RETNC",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			if (cpu.FlagC()) {
				cpu.PC.Inc()
				return 8
			}
			addr := cpu.memoryReadWord(cpu.SP.Get())
			cpu.PC.Set(addr)
			cpu.SP.IncBy(2)
			return 20
		},
	},
	0xD1: {
		mnemonic:   "POP DE",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			cycles := cpu.popR16(cpu.DE)
			cpu.PC.Inc()
			return cycles
		},
	},
	0xD2: {
		mnemonic:   "JP NC, %#04x",
		argLengths: []int{lword},
		length:     3,
		handler: func(cpu *CPU, args ...int) int {
			if !cpu.FlagC() {
				return cpu.jmp(uint16(args[0]))
			}
			cpu.PC.IncBy(3)
			return 12
		},
	},
	/* TODO: 0xD3 */
	/* TODO: 0xD4 */
	0xD5: {
		mnemonic:   "PUSH DE",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			cycles := cpu.pushR16(cpu.DE)
			cpu.PC.Inc()
			return cycles
		},
	},
	0xD6: {
		mnemonic:   "SUB %#02x",
		argLengths: []int{lbyte},
		length:     2,
		handler: func(cpu *CPU, args ...int) int {
			cycles := cpu.subd8(byte(args[0]))
			cpu.PC.IncBy(2)
			return cycles
		},
	},
	0xD7: {
		mnemonic:   "RST 10H",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			return cpu.rst(0x0010)
		},
	},
	0xD8: {
		mnemonic:   "RETC",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			if (!cpu.FlagC()) {
				cpu.PC.Inc()
				return 8
			}
			addr := cpu.memoryReadWord(cpu.SP.Get())
			cpu.PC.Set(addr)
			cpu.SP.IncBy(2)
			return 20
		},
	},
	0xD9: {
		mnemonic:   "RETI",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			addr := cpu.memoryReadWord(cpu.SP.Get())
			cpu.PC.Set(addr)
			cpu.SP.IncBy(2)
			cpu.EnableInterrupts()
			return 16
		},
	},
	0xDA: {
		mnemonic:   "JP C, %#04x",
		argLengths: []int{lword},
		length:     3,
		handler: func(cpu *CPU, args ...int) int {
			if cpu.FlagC() {
				return cpu.jmp(uint16(args[0]))
			}
			cpu.PC.IncBy(3)
			return 12
		},
	},
	/* TODO: 0xDB */
	/* TODO: 0xDC */
	/* TODO: 0xDD */
	0xDE: {
		mnemonic:   "SBC %#02x",
		argLengths: []int{lbyte},
		length:     2,
		handler: func(cpu *CPU, args ...int) int {
			cycles := cpu.sbcD8(byte(args[0]))
			cpu.PC.IncBy(2)
			return cycles
		},
	},
	0xDF: {
		mnemonic:   "RST 18H",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			return cpu.rst(0x0018)
		},
	},
	0xE0: {
		mnemonic:   "LDH (%#02x), A",
		argLengths: []int{lbyte},
		length:     2,
		handler: func(cpu *CPU, args ...int) int {
			cycles := cpu.ldha8R8(byte(args[0]), cpu.A)
			cpu.PC.IncBy(2)
			return cycles
		},
	},
	0xE1: {
		mnemonic:   "POP HL",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			cycles := cpu.popR16(cpu.HL)
			cpu.PC.Inc()
			return cycles
		},
	},
	0xE2: {
		mnemonic:   "LD (C), A",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			cycles := cpu.ldhR8R8(cpu.C, cpu.A)
			cpu.PC.Inc()
			return cycles
		},
	},
	/* TODO: 0xE3 */
	/* TODO: 0xE4 */
	0xE5: {
		mnemonic:   "PUSH HL",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			cycles := cpu.pushR16(cpu.HL)
			cpu.PC.Inc()
			return cycles
		},
	},
	0xE6: {
		mnemonic:   "AND %#02x",
		argLengths: []int{lbyte},
		length:     2,
		handler: func(cpu *CPU, args ...int) int {
			cycles := cpu.andD8(byte(args[0]))
			cpu.PC.IncBy(2)
			return cycles
		},
	},
	0xE7: {
		mnemonic:   "RST 20H",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			return cpu.rst(0x0020)
		},
	},
	0xE8: {
		mnemonic:   "ADD SP, %#02x",
		argLengths: []int{lbyte},
		length:     2,
		handler: func(cpu *CPU, args ...int) int {
			cycles := cpu.addSP(int8(args[0]))
			cpu.PC.IncBy(2)
			return cycles
		},
	},
	0xE9: {
		mnemonic:   "JP (HL)",
		argLengths: []int{lword},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			return cpu.jmpR16(cpu.HL)
		},
	},
	0xEA: {
		mnemonic:   "LD (%#04x), A",
		argLengths: []int{lword},
		length:     3,
		handler: func(cpu *CPU, args ...int) int {
			cycles := cpu.lda16R8(uint16(args[0]), cpu.A)
			cpu.PC.IncBy(3)
			return cycles
		},
	},
	/* TODO: 0xEB */
	/* TODO: 0xEC */
	/* TODO: 0xED */
	0xEE: {
		mnemonic:   "XOR %#02x",
		argLengths: []int{lbyte},
		length:     2,
		handler: func(cpu *CPU, args ...int) int {
			cycles := cpu.xorD8(byte(args[0]))
			cpu.PC.IncBy(2)
			return cycles
		},
	},
	0xEF: {
		mnemonic:   "RST 28H",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			return cpu.rst(0x0028)
		},
	},
	0xF0: {
		mnemonic:   "LDH A, (%#02x)",
		argLengths: []int{lbyte},
		length:     2,
		handler: func(cpu *CPU, args ...int) int {
			cycles := cpu.ldhR8a8(cpu.A, byte(args[0]))
			cpu.PC.IncBy(2)
			return cycles
		},
	},
	0xF1: {
		mnemonic:   "POP AF",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			cycles := cpu.popR16(cpu.AF)
			cpu.PC.Inc()
			return cycles
		},
	},
	0xF2: {
		mnemonic:   "LD A, (C)",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			cycles := cpu.ldR8hR8(cpu.A, cpu.C)
			cpu.PC.Inc()
			return cycles
		},
	},
	0xF3: {
		mnemonic:   "DI",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			cpu.DisableInterrupts()
			cpu.PC.Inc()
			return 4
		},
	},
	/* TODO: 0xF4 */
	0xF5: {
		mnemonic:   "PUSH AF",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			cycles := cpu.pushR16(cpu.AF)
			cpu.PC.Inc()
			return cycles
		},
	},
	0xF6: {
		mnemonic:   "OR %#02",
		argLengths: []int{lbyte},
		length:     2,
		handler: func(cpu *CPU, args ...int) int {
			cycles := cpu.orD8(byte(args[0]))
			cpu.PC.IncBy(2)
			return cycles
		},
	},
	0xF7: {
		mnemonic:   "RST 30H",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			return cpu.rst(0x0030)
		},
	},
	0xF8: {
		mnemonic:   "LDHL SP, %#02x",
		argLengths: []int{lbyte},
		length:     2,
		handler: func(cpu *CPU, args ...int) int {
			cycles := cpu.ldR16R16a8(cpu.HL, cpu.SP, int8(args[0]))
			cpu.PC.IncBy(2)
			return cycles
		},
	},
	0xF9: {
		mnemonic:   "LD SP, HL",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			cycles := cpu.ldR16R16(cpu.SP, cpu.HL)
			cpu.PC.Inc()
			return cycles
		},
	},
	0xFA: {
		mnemonic:   "LD A, (%#04x)",
		argLengths: []int{lword},
		length:     3,
		handler: func(cpu *CPU, args ...int) int {
			cycles := cpu.ldR8a16(cpu.A, uint16(args[0]))
			cpu.PC.IncBy(3)
			return cycles
		},
	},
	0xFB: {
		mnemonic:   "EI",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			cpu.EnableInterrupts()
			cpu.PC.Inc()
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
			cycles := cpu.cpD8(byte(args[0]))
			cpu.PC.IncBy(2)
			return cycles
		},
	},
	0xFF: {
		mnemonic:   "RST 38H",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			return cpu.rst(0x0038)
		},
	},
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

func (cpu *CPU) jmp(a16 uint16) int {
	cpu.jump(a16)
	return 16
}

func (cpu *CPU) jr(r8 int8) int {
	var a16 uint16
	if r8 < 0 {
		a16 = cpu.PC.Get() - uint16(^r8 - 1)
	} else {
		a16 = cpu.PC.Get() + uint16(r8)
	}
	cpu.jump(a16)
	return 12
}

func (cpu *CPU) jmpR16(r *WordRegister) int {
	cpu.jump(r.Get())
	return 4
}

func (cpu *CPU) ldR8aR16(r1 *ByteRegister, r2 *WordRegister) int {
	r1.Set(cpu.memoryReadByte(r2.Get()))
	return 8
}

func (cpu *CPU) lda16R16(a16 uint16, r *WordRegister) int {
	h, l := bits.SplitWord(r.Get())
	cpu.memoryWriteByte(a16, l)
	cpu.memoryWriteByte(a16+1, h)
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
	cpu.push(r.Get())
	return 16
}

func (cpu *CPU) popR16(r *WordRegister) int {
	r.Set(cpu.pop())
	return 12
}

func (cpu *CPU) ldR16R16a8(r1, r2 *WordRegister, r8 int8) int {
	var carry, halfCarry bool
	var a16 uint16
	if r8 < 0 {
		rel := uint16(^r8 - 1)
		a16 = uint16(r2.Get() - rel)
		halfCarry = bits.HalfCarrySubWord(r2.Get(), rel)
		carry = bits.CarrySubWord(r2.Get(), rel)
	} else {
		rel := uint16(r8)
		a16 = uint16(r2.Get() + rel)
		halfCarry = bits.HalfCarryAddWord(r2.Get(), rel)
		carry = bits.CarryAddWord(r2.Get(), rel)
	}
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
	result, flags := cpu.subByte(cpu.A.Get(), r.Get(), false)
	cpu.A.Set(result)
	cpu.F.Set(flags)
	return 4
}

func (cpu *CPU) sbcR8(r *ByteRegister) int {
	result, flags := cpu.subByte(cpu.A.Get(), r.Get(), cpu.FlagC())
	cpu.A.Set(result)
	cpu.F.Set(flags)
	return 4
}

func (cpu *CPU) sbcD8(d8 byte) int {
	result, flags := cpu.subByte(cpu.A.Get(), d8, cpu.FlagC())
	cpu.A.Set(result)
	cpu.F.Set(flags)
	return 8
}

func (cpu *CPU) andR8(r *ByteRegister) int {
	result, flags := cpu.and(cpu.A.Get(), r.Get())
	cpu.A.Set(result)
	cpu.F.Set(flags)
	return 4
}

func (cpu *CPU) andD8(d8 byte) int {
	result, flags := cpu.and(cpu.A.Get(), d8)
	cpu.A.Set(result)
	cpu.F.Set(flags)
	return 8
}

func (cpu *CPU) andaR16(r *WordRegister) int {
	result, flags := cpu.and(cpu.A.Get(), cpu.memoryReadByte(r.Get()))
	cpu.A.Set(result)
	cpu.F.Set(flags)
	return 8
}

func (cpu *CPU) orR8(r *ByteRegister) int {
	result, flags := cpu.or(cpu.A.Get(), r.Get())
	cpu.A.Set(result)
	cpu.F.Set(flags)
	return 4
}

func (cpu *CPU) orD8(d8 byte) int {
	result, flags := cpu.or(cpu.A.Get(), d8)
	cpu.A.Set(result)
	cpu.F.Set(flags)
	return 8
}

func (cpu *CPU) oraR16(r *WordRegister) int {
	result, flags := cpu.or(cpu.A.Get(), cpu.memoryReadByte(r.Get()))
	cpu.A.Set(result)
	cpu.F.Set(flags)
	return 8
}

func (cpu *CPU) xorR8(r *ByteRegister) int {
	result, flags := cpu.xor(cpu.A.Get(), r.Get())
	cpu.A.Set(result)
	cpu.F.Set(flags)
	return 4
}

func (cpu *CPU) xorD8(d8 byte) int {
	result, flags := cpu.xor(cpu.A.Get(), d8)
	cpu.A.Set(result)
	cpu.F.Set(flags)
	return 8
}

func (cpu *CPU) xoraR16(r *WordRegister) int {
	result, flags := cpu.xor(cpu.A.Get(), cpu.memoryReadByte(r.Get()))
	cpu.A.Set(result)
	cpu.F.Set(flags)
	return 8
}

func (cpu *CPU) cpR8(r *ByteRegister) int {
	flags := cpu.cmp(cpu.A.Get(), r.Get())
	cpu.F.Set(flags)
	return 4
}

func (cpu *CPU) cpD8(d8 byte) int {
	flags := cpu.cmp(cpu.A.Get(), d8)
	cpu.F.Set(flags)
	return 8
}

func (cpu *CPU) cpA8(r *ByteRegister) int {
	flags := cpu.cmp(cpu.A.Get(), cpu.memoryReadByte(cpu.HL.Get()))
	cpu.F.Set(flags)
	return 8
}

func (cpu *CPU) subd8(d8 byte) int {
	result, flags := cpu.subByte(cpu.A.Get(), d8, false)
	cpu.A.Set(result)
	cpu.F.Set(flags)
	return 8
}

func (cpu *CPU) subaR16(r *WordRegister) int {
	d8 := cpu.memoryReadByte(r.Get())
	result, flags := cpu.subByte(cpu.A.Get(), d8, false)
	cpu.A.Set(result)
	cpu.F.Set(flags)
	return 8
}

func (cpu *CPU) addR8R8(r1, r2 *ByteRegister) int {
	result, flags := cpu.addByte(r1.Get(), r2.Get(), false)
	r1.Set(result)
	cpu.F.Set(flags)
	return 4
}

func (cpu *CPU) adcR8R8(r1, r2 *ByteRegister) int {
	result, flags := cpu.addByte(r1.Get(), r2.Get(), cpu.FlagC())
	r1.Set(result)
	cpu.F.Set(flags)
	return 4
}

func (cpu *CPU) adcR8aR16(r1 *ByteRegister, r2 *WordRegister) int {
	d8 := cpu.memoryReadByte(r2.Get())
	result, flags := cpu.addByte(r1.Get(), d8, cpu.FlagC())
	r1.Set(result)
	cpu.F.Set(flags)
	return 8
}

func (cpu *CPU) sbcR8aR16(r1 *ByteRegister, r2 *WordRegister) int {
	result, flags := cpu.subByte(r1.Get(), cpu.memoryReadByte(r2.Get()), cpu.FlagC())
	r1.Set(result)
	cpu.F.Set(flags)
	return 8
}

func (cpu *CPU) addR16R16(r1, r2 *WordRegister) int {
	result, flags := cpu.addWord(r1.Get(), r2.Get(), false)
	r1.Set(result)
	cpu.F.Set(flags)
	return 8
}

func (cpu *CPU) addR8d8(r1 *ByteRegister, d8 byte) int {
	result, flags := cpu.addByte(r1.Get(), d8, false)
	r1.Set(result)
	cpu.F.Set(flags)
	return 8
}

func (cpu *CPU) adcR8d8(r1 *ByteRegister, d8 byte) int {
	result, flags := cpu.addByte(r1.Get(), d8, cpu.FlagC())
	r1.Set(result)
	cpu.F.Set(flags)
	return 8
}

func (cpu *CPU) addSP(r8 int8) int {
	var result uint16
	var flags byte
	if r8 < 0 {
		result, flags = cpu.subWord(cpu.SP.Get(), uint16(^r8 - 1), false)
	} else {
		result, flags = cpu.addWord(cpu.SP.Get(), uint16(r8), false)
	}
	cpu.SP.Set(result)
	cpu.F.Set(flags & ^flagZ & ^flagN)
	return 16
}

func (cpu *CPU) addR8aR16(r1 *ByteRegister, r2 *WordRegister) int {
	d8 := cpu.memoryReadByte(r2.Get())
	result, flags := cpu.addByte(r1.Get(), d8, false)
	r1.Set(result)
	cpu.F.Set(flags)
	return 8
}

func (cpu *CPU) incR8(r *ByteRegister) int {
	result, flags := cpu.addByte(r.Get(), 1, false)
	r.Set(result)
	cpu.F.Set((flags & ^flagC) | (cpu.F.Get() & flagC))
	return 4
}

func (cpu *CPU) incR16(r *WordRegister) int {
	result, flags := cpu.addWord(r.Get(), 1, false)
	r.Set(result)
	cpu.F.Set((flags & ^flagC) | (cpu.F.Get() & flagC))
	return 8
}

func (cpu *CPU) decR8(r *ByteRegister) int {
	result, flags := cpu.subByte(r.Get(), 1, false)
	r.Set(result)
	cpu.F.Set((flags & ^flagC) | (cpu.F.Get() & flagC))
	return 4
}

func (cpu *CPU) decR16(r *WordRegister) int {
	result, flags := cpu.subWord(r.Get(), 1, false)
	r.Set(result)
	cpu.F.Set((flags & ^flagC) | (cpu.F.Get() & flagC))
	return 8
}

func (cpu *CPU) rst(addr uint16) int {
	cpu.push(cpu.PC.Get())
	cpu.jump(addr)
	return 16
}

// Instructions

func (cpu *CPU) and(op byte, value byte) (byte, byte) {
	result := op & value
	flags := buildFlags(result == 0, false, true, false)
	return result, flags
}

func (cpu *CPU) or(op byte, value byte) (byte, byte) {
	result := op | value
	flags := buildFlags(result == 0, false, false, false)
	return result, flags
}

func (cpu *CPU) xor(op byte, value byte) (byte, byte) {
	result := op ^ value
	flags := buildFlags(result == 0, false, false, false)
	return result, flags
}

func (cpu *CPU) cmp(op byte, value byte) byte {
	diff := op - value
	return buildFlags(diff == 0, true, bits.HalfCarrySubByte(op, value), bits.CarrySubByte(op, value))
}

func (cpu *CPU) addByte(op byte, value byte, carryBit bool) (byte, byte) {
	if carryBit {
		value++
	}
	result := op + value
	flags := buildFlags(result == 0, false, bits.HalfCarryAddByte(op, value), bits.CarryAddByte(op, value))
	return result, flags
}

func (cpu *CPU) addWord(op uint16, value uint16, carryBit bool) (uint16, byte) {
	if carryBit {
		value++
	}
	result := op + value
	flags := buildFlags(result == 0, false, bits.HalfCarryAddWord(op, value), bits.CarryAddWord(op, value))
	return result, flags
}

func (cpu *CPU) subByte(op byte, value byte, carryBit bool) (byte, byte) {
	if carryBit {
		value++
	}
	result := op - value
	flags := buildFlags(result == 0, true, bits.HalfCarrySubByte(op, value), bits.CarrySubByte(op, value))
	return result, flags
}

func (cpu *CPU) subWord(op uint16, value uint16, carryBit bool) (uint16, byte) {
	if carryBit {
		value++
	}
	result := op - value
	flags := buildFlags(result == 0, true, bits.HalfCarrySubWord(op, value), bits.CarrySubWord(op, value))
	return result, flags
}

func (cpu *CPU) jump(a16 uint16) {
	cpu.PC.Set(a16)
}

func (cpu *CPU) push(v uint16) {
	cpu.SP.DecBy(2)
	cpu.memoryWriteWord(cpu.SP.Get(), v)
}

func (cpu *CPU) pop() uint16 {
	value := cpu.memoryReadWord(cpu.SP.Get())
	cpu.SP.IncBy(2)
	return value
}

func buildFlags(Z, N, H, C bool) byte {
	var flags byte
	if (Z) {
		flags |= flagZ
	}
	if (N) {
		flags |= flagN
	}
	if (H) {
		flags |= flagH
	}
	if (C) {
		flags |= flagC
	}
	return flags
}