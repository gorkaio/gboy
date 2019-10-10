package cpu_test

import (
	mocks "github.com/gorkaio/gboy/pkg/cpu/mocks"
)

const (
	FlagZ = 0x80
	FlagN = 0x40
	FlagH = 0x20
	FlagC = 0x10
)

type regMap map[string]int
type memMap map[uint16]uint8
type opcode []byte
type testDescription struct {
	description      string
	opcode           opcode
	regsGiven        regMap
	regsExpected     regMap
	memReadExpected  memMap
	memWriteExpected memMap
	cycles           int
}

func readInstruction(mem *mocks.MockMemory, address uint16, instruction []byte) {
	for i, b := range instruction {
		mem.EXPECT().Read(address + uint16(i)).Return(b)
	}
}

func instruction(bytes ...byte) []byte {
	result := make([]byte, 4)
	for i, byte := range bytes {
		result[i] = byte
	}
	return result
}

func buildTestCase(test testDescription) testCase {
	testCase := NewTestCase(instruction(test.opcode...))
	for k, v := range test.regsGiven {
		testCase.WithRegister(k, v)
	}

	for addr, dat := range test.memReadExpected {
		testCase.ExpectMemoryRead(addr, dat)
	}

	for addr, dat := range test.memWriteExpected {
		testCase.ExpectMemoryWrite(addr, dat)
	}

	for k, v := range test.regsExpected {
		testCase.ExpectRegister(k, v)
	}

	// Check PC has incremented same bytes than opcode length unless it's set in initial state or expected state
	if testCase.expectedState.regs["PC"] == testCase.initialState.regs["PC"] && testCase.initialState.regs["PC"] == 0 {
		testCase.ExpectRegister("PC", int(testCase.initialState.regs["PC"])+len(test.opcode))
	}
	testCase.ExpectCycles(test.cycles)
	testCase.WithDescription(test.description)

	return testCase
}
