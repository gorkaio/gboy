package cpu

import (
	"fmt"
)

const (
	unknownOpCode = "?"
	NOP           = 0x00
	LD_BC_D16     = 0x01
	INC_BC        = 0x03
	INC_B         = 0x04
	DEC_B         = 0x05
	LD_B_D8       = 0x06
	DEC_BC        = 0x0B
	INC_C         = 0x0C
	DEC_C         = 0x0D
	LD_C_D8       = 0x0E
	LD_DE_D16	  = 0x11
	INC_DE	      = 0x13
	INC_D	      = 0x14
	DEC_D	      = 0x15
	LD_D_D8       = 0x16
	DEC_DE	      = 0x1B
	INC_E	      = 0x1C
	DEC_E	      = 0x1D
	LD_E_D8	      = 0x1E
	JR_NZ_R8 	  = 0x20
	LD_HL_D16     = 0x21
	INC_HL        = 0x23
	INC_H         = 0x24
	DEC_H         = 0x25
	LD_H_D8       = 0x26
	DEC_HL        = 0x2B
	INC_L         = 0x2C
	DEC_L         = 0x2D
	LD_SP_D16     = 0x31
	LDD_HL_A      = 0x32
	INC_SP        = 0x33
	DEC_SP        = 0x3B
	INC_A         = 0x3C
	DEC_A         = 0x3D
	LD_A_D8       = 0x3E
	XOR_A         = 0xAF
	JMP           = 0xC3
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
	NOP: {
		mnemonic:   "NOP",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			return 4
		},
	},
	JMP: {
		mnemonic:   "JMP %#04x",
		argLengths: []int{lword},
		length:     3,
		handler: func(cpu *CPU, args ...int) int {
			cpu.PC.Set(uint16(args[0]))
			return 16
		},
	},
	XOR_A: {
		mnemonic:   "XOR A",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			cpu.A.Set(0)
			cpu.SetZ()
			cpu.ClearN()
			return 4
		},
	},
	LD_BC_D16: {
		mnemonic:   "LD BC, %#04x",
		argLengths: []int{lword},
		length:     3,
		handler: func(cpu *CPU, args ...int) int {
			d16 := uint16(args[0])
			cpu.BC.Set(d16)
			return 12
		},
	},
	LD_DE_D16: {
		mnemonic:   "LD DE, %#04x",
		argLengths: []int{lword},
		length:     3,
		handler: func(cpu *CPU, args ...int) int {
			d16 := uint16(args[0])
			cpu.DE.Set(d16)
			return 12
		},
	},
	LD_HL_D16: {
		mnemonic:   "LD HL, %#04x",
		argLengths: []int{lword},
		length:     3,
		handler: func(cpu *CPU, args ...int) int {
			d16 := uint16(args[0])
			cpu.HL.Set(d16)
			return 12
		},
	},
	LD_SP_D16: {
		mnemonic:   "LD SP, %#04x",
		argLengths: []int{lword},
		length:     3,
		handler: func(cpu *CPU, args ...int) int {
			d16 := uint16(args[0])
			cpu.SP.Set(d16)
			return 12
		},
	},
	LD_A_D8: {
		mnemonic:   "LD A, %#02x",
		argLengths: []int{lbyte},
		length:     2,
		handler: func(cpu *CPU, args ...int) int {
			d8 := uint8(args[0])
			cpu.A.Set(d8)
			return 8
		},
	},
	LD_B_D8: {
		mnemonic:   "LD B, %#02x",
		argLengths: []int{lbyte},
		length:     2,
		handler: func(cpu *CPU, args ...int) int {
			d8 := uint8(args[0])
			cpu.B.Set(d8)
			return 8
		},
	},
	LD_C_D8: {
		mnemonic:   "LD C, %#02x",
		argLengths: []int{lbyte},
		length:     2,
		handler: func(cpu *CPU, args ...int) int {
			d8 := uint8(args[0])
			cpu.C.Set(d8)
			return 8
		},
	},
	LD_D_D8: {
		mnemonic:   "LD D, %#02x",
		argLengths: []int{lbyte},
		length:     2,
		handler: func(cpu *CPU, args ...int) int {
			d8 := uint8(args[0])
			cpu.D.Set(d8)
			return 8
		},
	},
	LD_E_D8: {
		mnemonic:   "LD E, %#02x",
		argLengths: []int{lbyte},
		length:     2,
		handler: func(cpu *CPU, args ...int) int {
			d8 := uint8(args[0])
			cpu.E.Set(d8)
			return 8
		},
	},
	LD_H_D8: {
		mnemonic:   "LD H, %#02x",
		argLengths: []int{lbyte},
		length:     2,
		handler: func(cpu *CPU, args ...int) int {
			d8 := uint8(args[0])
			cpu.H.Set(d8)
			return 8
		},
	},
	LDD_HL_A: {
		mnemonic:   "LD (HL-), A",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			cpu.memoryWriteByte(cpu.HL.Get(), cpu.A.Get())
			cpu.HL.Dec()
			return 8
		},
	},
	INC_A: {
		mnemonic:   "INC A",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			cpu.A.Inc()
			cpu.ClearN()
			cpu.UpdateZ(cpu.A.Get())
			return 4
		},
	},
	INC_B: {
		mnemonic:   "INC B",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			cpu.B.Inc()
			cpu.ClearN()
			cpu.UpdateZ(cpu.B.Get())
			return 4
		},
	},
	INC_C: {
		mnemonic:   "INC C",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			cpu.C.Inc()
			cpu.ClearN()
			cpu.UpdateZ(cpu.C.Get())
			return 4
		},
	},
	INC_D: {
		mnemonic:   "INC D",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			cpu.D.Inc()
			cpu.ClearN()
			cpu.UpdateZ(cpu.D.Get())
			return 4
		},
	},
	INC_E: {
		mnemonic:   "INC E",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			cpu.E.Inc()
			cpu.ClearN()
			cpu.UpdateZ(cpu.E.Get())
			return 4
		},
	},
	INC_H: {
		mnemonic:   "INC H",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			cpu.H.Inc()
			cpu.ClearN()
			cpu.UpdateZ(cpu.H.Get())
			return 4
		},
	},
	INC_L: {
		mnemonic:   "INC L",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			cpu.L.Inc()
			cpu.ClearN()
			cpu.UpdateZ(cpu.L.Get())
			return 4
		},
	},
	DEC_A: {
		mnemonic:   "DEC A",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			cpu.A.Dec()
			cpu.SetN()
			cpu.UpdateZ(cpu.A.Get())
			return 4
		},
	},
	DEC_B: {
		mnemonic:   "DEC B",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			cpu.B.Dec()
			cpu.SetN()
			cpu.UpdateZ(cpu.B.Get())
			return 4
		},
	},
	DEC_C: {
		mnemonic:   "DEC C",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			cpu.C.Dec()
			cpu.SetN()
			cpu.UpdateZ(cpu.C.Get())
			return 4
		},
	},
	DEC_D: {
		mnemonic:   "DEC D",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			cpu.D.Dec()
			cpu.SetN()
			cpu.UpdateZ(cpu.D.Get())
			return 4
		},
	},
	DEC_E: {
		mnemonic:   "DEC E",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			cpu.E.Dec()
			cpu.SetN()
			cpu.UpdateZ(cpu.E.Get())
			return 4
		},
	},
	DEC_H: {
		mnemonic:   "DEC H",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			cpu.H.Dec()
			cpu.SetN()
			cpu.UpdateZ(cpu.H.Get())
			return 4
		},
	},
	DEC_L: {
		mnemonic:   "DEC L",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			cpu.L.Dec()
			cpu.SetN()
			cpu.UpdateZ(cpu.L.Get())
			return 4
		},
	},
	INC_BC: {
		mnemonic:   "INC BC",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			cpu.BC.Inc()
			return 8
		},
	},
	INC_DE: {
		mnemonic:   "INC DE",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			cpu.DE.Inc()
			return 8
		},
	},
	INC_HL: {
		mnemonic:   "INC HL",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			cpu.HL.Inc()
			return 8
		},
	},
	INC_SP: {
		mnemonic:   "INC SP",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			cpu.SP.Inc()
			return 8
		},
	},
	DEC_BC: {
		mnemonic:   "DEC BC",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			cpu.BC.Dec()
			return 8
		},
	},
	DEC_DE: {
		mnemonic:   "DEC DE",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			cpu.DE.Dec()
			return 8
		},
	},
	DEC_HL: {
		mnemonic:   "DEC HL",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			cpu.HL.Dec()
			return 8
		},
	},
	DEC_SP: {
		mnemonic:   "DEC SP",
		argLengths: []int{},
		length:     1,
		handler: func(cpu *CPU, args ...int) int {
			cpu.SP.Dec()
			return 8
		},
	},
	JR_NZ_R8: {
		mnemonic:   "JR NZ, %#02x",
		argLengths: []int{lbyte},
		length:     2,
		handler: func(cpu *CPU, args ...int) int {
			if !cpu.Z() {
				rel := int8(args[0]) // Signed relative address jump distance
				address := int(cpu.PC.Get()) + int(rel)
				cpu.PC.Set(uint16(address))
				return 12
			}
			return 8
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
			arg := flipWord(uint16((data & 0xFFFF0000) >> 16))
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

func flipWord(data uint16) uint16 {
	return ((data & 0xFF00) >> 8) | ((data & 0x00FF) << 8)
}
