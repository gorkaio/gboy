package cpu

import (
	"fmt"
)

const (
	unknownOpCode = "?"
	NOP           = 0x00
	JMP           = 0xC3
	XOR_A         = 0xAF
	LD_HL_D16     = 0x21
	LD_C_D8       = 0x0E
	LD_B_D8       = 0x06
	LDD_HL_A      = 0x32
	DEC_B         = 0x05
	JR_NZ_R8 	  = 0x20
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
			cpu.F.Set(cpu.F.Get() | flagZ)
			return 4
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
