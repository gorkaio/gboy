package cpu_test

import (
	"testing"
)

func TestRst(t *testing.T) {
	tests := []struct{
		opc opcode
		value int
	}{
		{opcode{0xC7}, 0x0000},
		{opcode{0xD7}, 0x0010},
		{opcode{0xE7}, 0x0020},
		{opcode{0xF7}, 0x0030},
		{opcode{0xCF}, 0x0008},
		{opcode{0xDF}, 0x0018},
		{opcode{0xEF}, 0x0028},
		{opcode{0xFF}, 0x0038},
	}

	for _, test := range(tests) {
		testDescription := testDescription{
			"RST puts PC into SP and sets PC",
			test.opc,
			regMap{"SP": 0x1234, "PC": 0x5678},
			regMap{"SP": 0x1232, "PC": test.value},
			memMap{},
			memMap{0x1233: 0x56, 0x1232: 0x78},
			16,
		}
		testCase := buildTestCase(testDescription)
		testCase.Run(t)
	}
}