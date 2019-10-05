package cpu_test

import (
	"testing"
	"github.com/gorkaio/gboy/pkg/cpu"
)

func TestNOP(t *testing.T) {
	for _, test := range testCasesForNOP() {
		testInstruction(t, test)
	}
}

func testCasesForNOP() []testCase {
	return []testCase{
		{
			description: "NOP does not change registers or memory",
			instruction: []byte{0x00, 0x00, 0x00, 0x00},
			initialState: cpu.State{
				AF: 0,
				BC: 0,
				DE: 0,
				HL: 0,
				SP: 0,
				PC: PCStartAddress,
				IME: false,
			},
			expectedState: cpu.State{
				AF: 0,
				BC: 0,
				DE: 0,
				HL: 0,
				SP: 0,
				PC: PCStartAddress + 1,
				IME: false,
			},
			expectedReads: []memoryAccess{},
			expectedWrites: []memoryAccess{},
			expectedCycles: 4,
		},
	}
}