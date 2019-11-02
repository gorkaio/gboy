package cpu_test

import (
	"testing"
)

func TestPushPutsRegisterContentToSPMemoryAddress(t *testing.T) {
	tests := []struct{
		opc opcode
		srcReg string
	}{
		{opcode{0xC5}, "BC"},
		{opcode{0xD5}, "DE"},
		{opcode{0xE5}, "HL"},
		{opcode{0xF5}, "AF"},
	}

	for _, test := range(tests) {
		testDescription := testDescription{
			"PUSH BC puts register content into SP memory address",
			test.opc,
			regMap{"SP": 0x1234, test.srcReg: 0x5670},
			regMap{"SP": 0x1232},
			memMap{},
			memMap{0x1232: 0x70, 0x1233: 0x56},
			16,
		}
		testCase := buildTestCase(testDescription)
		testCase.Run(t)
	}
}
