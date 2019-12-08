package cpu

import "github.com/gorkaio/gboy/pkg/bits"

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

func (cpu *CPU) call(a16 uint16) {
	cpu.push(cpu.PC.Get())
	cpu.jump(a16)
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