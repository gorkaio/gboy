package cpu_test

import (
	"fmt"
	"testing"
)

func TestLoadImmediateFor8bitRegisters(t *testing.T) {
	registerOpcodes := map[string]opcode{
		"A": opcode{0x3E, 0x12},
		"B": opcode{0x06, 0x12}, "C": opcode{0x0E, 0x12},
		"D": opcode{0x16, 0x12}, "E": opcode{0x1E, 0x12},
		"H": opcode{0x26, 0x12}, "L": opcode{0x2E, 0x12},
	}

	for r8, opc := range registerOpcodes {
		testDescription := testDescription{
			fmt.Sprintf("'LD %s, %#02x' loads %#02x value into %s", r8, 0x12, 0x12, r8),
			opc,
			regMap{},
			regMap{r8: 0x12},
			memMap{},
			memMap{},
			8,
		}
		testCase := buildTestCase(testDescription)
		testCase.Run(t)
	}
}

func TestLoadImmediateFor16bitRegisters(t *testing.T) {
	registerOpcodes := map[string]opcode{
		"BC": opcode{0x01, 0x34, 0x12}, "DE": opcode{0x11, 0x34, 0x12},
		"HL": opcode{0x21, 0x34, 0x12}, "SP": opcode{0x31, 0x34, 0x12},
	}

	for r16, opc := range registerOpcodes {
		testDescription := testDescription{
			fmt.Sprintf("'LD %s, %#04x' loads %#04x value into %s", r16, 0x1234, 0x1234, r16),
			opc,
			regMap{},
			regMap{r16: 0x1234},
			memMap{},
			memMap{},
			12,
		}
		testCase := buildTestCase(testDescription)
		testCase.Run(t)
	}
}

func TestLoadIndirectFromRegister(t *testing.T) {
	tests := []struct {
		opc      opcode
		src, dst string
		data     uint8
		address  uint16
		cycles   int
	}{
		{opc: opcode{0x02}, dst: "BC", src: "A", data: 0x56, address: 0x1234, cycles: 8},
		{opc: opcode{0x12}, dst: "DE", src: "A", data: 0x56, address: 0x1234, cycles: 8},
		{opc: opcode{0x70}, dst: "HL", src: "B", data: 0x56, address: 0x1234, cycles: 8},
		{opc: opcode{0x71}, dst: "HL", src: "C", data: 0x56, address: 0x1234, cycles: 8},
		{opc: opcode{0x72}, dst: "HL", src: "D", data: 0x56, address: 0x1234, cycles: 8},
		{opc: opcode{0x73}, dst: "HL", src: "E", data: 0x56, address: 0x1234, cycles: 8},
		{opc: opcode{0x74}, dst: "HL", src: "H", data: 0x12, address: 0x1234, cycles: 8},
		{opc: opcode{0x75}, dst: "HL", src: "L", data: 0x34, address: 0x1234, cycles: 8},
		{opc: opcode{0x77}, dst: "HL", src: "A", data: 0x56, address: 0x1234, cycles: 8},
	}

	for _, test := range tests {
		testDescription := testDescription{
			fmt.Sprintf("'LD (%s), %s' loads %s into memory address %s", test.dst, test.src, test.src, test.dst),
			test.opc,
			regMap{test.src: int(test.data), test.dst: int(test.address)},
			regMap{},
			memMap{},
			memMap{test.address: test.data},
			test.cycles,
		}
		testCase := buildTestCase(testDescription)
		testCase.Run(t)
	}
}

func TestLoadIndirectToRegister(t *testing.T) {
	tests := []struct {
		opc      opcode
		src, dst string
		data     uint8
		address  uint16
		cycles   int
	}{
		{opc: opcode{0x0A}, dst: "A", src: "BC", data: 0x56, address: 0x1234, cycles: 8},
		{opc: opcode{0x1A}, dst: "A", src: "DE", data: 0x56, address: 0x1234, cycles: 8},
		{opc: opcode{0x7E}, dst: "A", src: "HL", data: 0x56, address: 0x1234, cycles: 8},
		{opc: opcode{0x46}, dst: "B", src: "HL", data: 0x56, address: 0x1234, cycles: 8},
		{opc: opcode{0x4E}, dst: "C", src: "HL", data: 0x56, address: 0x1234, cycles: 8},
		{opc: opcode{0x56}, dst: "D", src: "HL", data: 0x56, address: 0x1234, cycles: 8},
		{opc: opcode{0x5E}, dst: "E", src: "HL", data: 0x56, address: 0x1234, cycles: 8},
		{opc: opcode{0x66}, dst: "H", src: "HL", data: 0x56, address: 0x1234, cycles: 8},
		{opc: opcode{0x6E}, dst: "L", src: "HL", data: 0x56, address: 0x1234, cycles: 8},
	}

	for _, test := range tests {
		testDescription := testDescription{
			fmt.Sprintf("'LD %s, (%s)' loads memory content from address %s into %s", test.dst, test.src, test.src, test.dst),
			test.opc,
			regMap{test.src: int(test.address)},
			regMap{test.dst: int(test.data)},
			memMap{test.address: test.data},
			memMap{},
			test.cycles,
		}
		testCase := buildTestCase(testDescription)
		testCase.Run(t)
	}
}

func TestLoadIndirectOther(t *testing.T) {
	testDescriptions := []testDescription{
		{
			"'LD (HL), %#02x' loads value %#02x into memory address HL",
			opcode{0x36, 0x56},
			regMap{"HL": 0x1234},
			regMap{"HL": 0x1234},
			memMap{},
			memMap{0x1234: 0x56},
			12,
		},
		{
			"'LDI (HL), A' loads A into memory address HL and increments HL",
			opcode{0x22},
			regMap{"A": 0x56, "HL": 0x1234},
			regMap{"HL": 0x1235},
			memMap{},
			memMap{0x1234: 0x56},
			8,
		},
		{
			"'LDD (HL), A' loads A into memory address HL and decrements HL",
			opcode{0x32},
			regMap{"A": 0x56, "HL": 0x1234},
			regMap{"A": 0x56, "HL": 0x1233},
			memMap{},
			memMap{0x1234: 0x56},
			8,
		},
		{
			"'LDI A, (HL)' loads content from memory address HL into A and increments HL",
			opcode{0x2A},
			regMap{"HL": 0x1234},
			regMap{"A": 0x56, "HL": 0x1235},
			memMap{0x1234: 0x56},
			memMap{},
			8,
		},
		{
			"'LDD A, (HL)' loads content from memory address HL into A and decrements HL",
			opcode{0x3A},
			regMap{"HL": 0x1234},
			regMap{"A": 0x56, "HL": 0x1233},
			memMap{0x1234: 0x56},
			memMap{},
			8,
		},
		{
			"'LD (C), A' loads A into memory address (0xFF00 + C)",
			opcode{0xE2},
			regMap{"A": 0x56, "C": 0x12},
			regMap{},
			memMap{},
			memMap{0xFF12: 0x56},
			8,
		},
		{
			"'LD A, (C)' loads memory address content from (0xFF00 + C) into A",
			opcode{0xF2},
			regMap{"A": 0x56, "C": 0x12},
			regMap{},
			memMap{0xFF12: 0x56},
			memMap{},
			8,
		},
	}

	for _, testDescription := range testDescriptions {
		testCase := buildTestCase(testDescription)
		testCase.Run(t)
	}
}

func TestLoadRegisterToRegister(t *testing.T) {
	tests := []struct {
		opc      opcode
		src, dst string
		cycles   int
	}{
		{opc: opcode{0x40}, dst: "B", src: "B", cycles: 4},
		{opc: opcode{0x41}, dst: "B", src: "C", cycles: 4},
		{opc: opcode{0x42}, dst: "B", src: "D", cycles: 4},
		{opc: opcode{0x43}, dst: "B", src: "E", cycles: 4},
		{opc: opcode{0x44}, dst: "B", src: "H", cycles: 4},
		{opc: opcode{0x45}, dst: "B", src: "L", cycles: 4},
		{opc: opcode{0x47}, dst: "B", src: "A", cycles: 4},

		{opc: opcode{0x48}, dst: "C", src: "B", cycles: 4},
		{opc: opcode{0x49}, dst: "C", src: "C", cycles: 4},
		{opc: opcode{0x4A}, dst: "C", src: "D", cycles: 4},
		{opc: opcode{0x4B}, dst: "C", src: "E", cycles: 4},
		{opc: opcode{0x4C}, dst: "C", src: "H", cycles: 4},
		{opc: opcode{0x4D}, dst: "C", src: "L", cycles: 4},
		{opc: opcode{0x4F}, dst: "C", src: "A", cycles: 4},

		{opc: opcode{0x50}, dst: "D", src: "B", cycles: 4},
		{opc: opcode{0x51}, dst: "D", src: "C", cycles: 4},
		{opc: opcode{0x52}, dst: "D", src: "D", cycles: 4},
		{opc: opcode{0x53}, dst: "D", src: "E", cycles: 4},
		{opc: opcode{0x54}, dst: "D", src: "H", cycles: 4},
		{opc: opcode{0x55}, dst: "D", src: "L", cycles: 4},
		{opc: opcode{0x57}, dst: "D", src: "A", cycles: 4},

		{opc: opcode{0x58}, dst: "E", src: "B", cycles: 4},
		{opc: opcode{0x59}, dst: "E", src: "C", cycles: 4},
		{opc: opcode{0x5A}, dst: "E", src: "D", cycles: 4},
		{opc: opcode{0x5B}, dst: "E", src: "E", cycles: 4},
		{opc: opcode{0x5C}, dst: "E", src: "H", cycles: 4},
		{opc: opcode{0x5D}, dst: "E", src: "L", cycles: 4},
		{opc: opcode{0x5F}, dst: "E", src: "A", cycles: 4},

		{opc: opcode{0x60}, dst: "H", src: "B", cycles: 4},
		{opc: opcode{0x61}, dst: "H", src: "C", cycles: 4},
		{opc: opcode{0x62}, dst: "H", src: "D", cycles: 4},
		{opc: opcode{0x63}, dst: "H", src: "E", cycles: 4},
		{opc: opcode{0x64}, dst: "H", src: "H", cycles: 4},
		{opc: opcode{0x65}, dst: "H", src: "L", cycles: 4},
		{opc: opcode{0x67}, dst: "H", src: "A", cycles: 4},

		{opc: opcode{0x68}, dst: "L", src: "B", cycles: 4},
		{opc: opcode{0x69}, dst: "L", src: "C", cycles: 4},
		{opc: opcode{0x6A}, dst: "L", src: "D", cycles: 4},
		{opc: opcode{0x6B}, dst: "L", src: "E", cycles: 4},
		{opc: opcode{0x6C}, dst: "L", src: "H", cycles: 4},
		{opc: opcode{0x6D}, dst: "L", src: "L", cycles: 4},
		{opc: opcode{0x6F}, dst: "L", src: "A", cycles: 4},

		{opc: opcode{0x78}, dst: "A", src: "B", cycles: 4},
		{opc: opcode{0x79}, dst: "A", src: "C", cycles: 4},
		{opc: opcode{0x7A}, dst: "A", src: "D", cycles: 4},
		{opc: opcode{0x7B}, dst: "A", src: "E", cycles: 4},
		{opc: opcode{0x7C}, dst: "A", src: "H", cycles: 4},
		{opc: opcode{0x7D}, dst: "A", src: "L", cycles: 4},
		{opc: opcode{0x7F}, dst: "A", src: "A", cycles: 4},
	}

	for _, test := range tests {
		testDescription := testDescription{
			fmt.Sprintf("'LD %s, %s' loads %s into %s", test.dst, test.src, test.src, test.dst),
			test.opc,
			regMap{test.src: 0x12},
			regMap{test.dst: 0x12},
			memMap{},
			memMap{},
			test.cycles,
		}
		testCase := buildTestCase(testDescription)
		testCase.Run(t)
	}
}
