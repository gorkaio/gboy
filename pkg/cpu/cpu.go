package cpu

import (
	"fmt"
	memory "github.com/gorkaio/gboy/pkg/memory"
	tablewriter "github.com/olekukonko/tablewriter"
	"os"
)

const (
	flagZ = 0x80
	flagN = 0x40
	flagH = 0x20
	flagC = 0x10
)

// CPU structure
type CPU struct {
	AF           WordRegister
	BC           WordRegister
	DE           WordRegister
	HL           WordRegister
	PC           WordRegister
	SP           WordRegister
	A            *ByteRegister
	F            *ByteRegister
	B            *ByteRegister
	C            *ByteRegister
	D            *ByteRegister
	E            *ByteRegister
	H            *ByteRegister
	L            *ByteRegister
	memory       memory.MemoryInterface
	debugEnabled bool
}

// New initialises a new Z80 cpu
func New(memory memory.MemoryInterface) *CPU {
	cpu := CPU{
		AF:     newWordRegister(),
		BC:     newWordRegister(),
		DE:     newWordRegister(),
		HL:     newWordRegister(),
		SP:     newWordRegister(),
		PC:     newWordRegister(),
		memory: memory,
	}
	cpu.PC.Set(0x100)
	cpu.A = cpu.AF.H()
	cpu.F = cpu.AF.L()
	cpu.B = cpu.BC.H()
	cpu.C = cpu.BC.L()
	cpu.D = cpu.DE.H()
	cpu.E = cpu.DE.L()
	cpu.H = cpu.HL.H()
	cpu.L = cpu.HL.L()
	cpu.debugEnabled = true
	return &cpu
}

// DebugEnable enables CPU debugging
func (cpu *CPU) DebugEnable() {
	cpu.debugEnabled = true
}

// DebugDisable disables CPU debugging
func (cpu *CPU) DebugDisable() {
	cpu.debugEnabled = false
}

// Step executes next instruction and returns cycles consumed
func (cpu *CPU) Step() (int, error) {
	op, err := cpu.opCodeAt(cpu.PC.Get())
	if cpu.debugEnabled {
		cpu.printStatus()
		fmt.Println(op.String())
	}

	if err != nil {
		return 0, err
	}

	cycles := op.handler(cpu, op.args...)
	if cpu.debugEnabled {
		cpu.printStatus()
		fmt.Printf("Cycles consumed: %d\n", cycles)
	}
	return cycles, nil
}

func (cpu *CPU) memoryReadWord(address uint16) uint16 {
	l := cpu.memory.Read(address)
	h := cpu.memory.Read(address + 1)
	return concatWord(h, l)
}

func (cpu *CPU) memoryReadDWord(address uint16) uint32 {
	a := uint32(cpu.memoryReadByte(address))
	b := uint32(cpu.memoryReadByte(address + 1))
	c := uint32(cpu.memoryReadByte(address + 2))
	d := uint32(cpu.memoryReadByte(address + 3))
	return (a<<24 | b<<16 | c<<8 | d)
}

func (cpu *CPU) memoryReadByte(address uint16) uint8 {
	return cpu.memory.Read(address)
}

func (cpu *CPU) memoryWriteByte(address uint16, data uint8) {
	cpu.memory.Write(address, data)
}

func (cpu *CPU) memoryWriteWord(address uint16, data uint16) {
	h, l := splitWord(data)
	cpu.memory.Write(address, l)
	cpu.memory.Write(address+1, h)
}

func (cpu *CPU) printStatus() {
	registerTable := tablewriter.NewWriter(os.Stdout)
	registerTable.SetHeader([]string{"Reg", "Value"})
	registerTable.Append([]string{"AF", fmt.Sprintf("%#04x", cpu.AF.Get())})
	registerTable.Append([]string{"BC", fmt.Sprintf("%#04x", cpu.BC.Get())})
	registerTable.Append([]string{"DE", fmt.Sprintf("%#04x", cpu.DE.Get())})
	registerTable.Append([]string{"HL", fmt.Sprintf("%#04x", cpu.HL.Get())})
	registerTable.Append([]string{"SP", fmt.Sprintf("%#04x", cpu.SP.Get())})
	registerTable.Append([]string{"PC", fmt.Sprintf("%#04x", cpu.PC.Get())})
	registerTable.Render()
	flagTable := tablewriter.NewWriter(os.Stdout)
	flagTable.SetHeader([]string{"Z", "N", "H", "C"})
	flagTable.Append([]string{
		fmt.Sprintf("%t", cpu.F.Get()&0x80 > 1),
		fmt.Sprintf("%t", cpu.F.Get()&0x40 > 1),
		fmt.Sprintf("%t", cpu.F.Get()&0x20 > 1),
		fmt.Sprintf("%t", cpu.F.Get()&0x10 > 1),
	})
	flagTable.Render()
}
