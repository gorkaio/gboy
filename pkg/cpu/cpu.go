package cpu

//go:generate mockgen -destination=mocks/memory_mock.go -package=cpu_mock github.com/gorkaio/gboy/pkg/cpu Memory

import (
	"fmt"
	"github.com/gorkaio/gboy/pkg/bits"
	"github.com/olekukonko/tablewriter"
	"os"
)

// Memory defines the interface for memory interaction
type Memory interface {
	Read(address uint16) byte
	Write(address uint16, data byte)
}

const (
	flagZ = byte(0x80)
	flagN = byte(0x40)
	flagH = byte(0x20)
	flagC = byte(0x10)
)

// CPU structure
type CPU struct {
	AF, BC, DE, HL, SP, PC *WordRegister
	A, F, B, C, D, E, H, L *ByteRegister
	memory                 Memory
	debugEnabled           bool
	imeFlag                bool
}

// State reflects the CPU status
type State struct {
	AF, BC, DE, HL, SP, PC uint16
	IME                    bool
}

// New initialises a new Z80 cpu
func New(memory Memory) *CPU {
	cpu := CPU{
		AF:           newMaskedWordRegister(0xFFF0),
		BC:           newWordRegister(),
		DE:           newWordRegister(),
		HL:           newWordRegister(),
		SP:           newWordRegister(),
		PC:           newWordRegister(),
		memory:       memory,
		debugEnabled: false,
		imeFlag:      false,
	}
	cpu.PC.Set(0x100)
	cpu.SP.Set(0xFFFE)
	cpu.A = cpu.AF.H()
	cpu.F = cpu.AF.L()
	cpu.B = cpu.BC.H()
	cpu.C = cpu.BC.L()
	cpu.D = cpu.DE.H()
	cpu.E = cpu.DE.L()
	cpu.H = cpu.HL.H()
	cpu.L = cpu.HL.L()
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
	if err != nil {
		return 0, err
	}

	if cpu.debugEnabled {
		cpu.printStatus()
		fmt.Println(op.String())
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
	return bits.ConcatWord(h, l)
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
	h, l := bits.SplitWord(data)
	cpu.memory.Write(address, l)
	cpu.memory.Write(address+1, h)
}

// SetFlagZ sets or clears the Zero Flag
func (cpu *CPU) SetFlagZ(value bool) {
	if value {
		cpu.F.Set(cpu.F.Get() | flagZ)
	} else {
		cpu.F.Set(cpu.F.Get() &^ flagZ)
	}
}

// SetFlagC sets or clears the Carry Flag
func (cpu *CPU) SetFlagC(value bool) {
	if value {
		cpu.F.Set(cpu.F.Get() | flagC)
	} else {
		cpu.F.Set(cpu.F.Get() &^ flagC)
	}
}

// SetFlagN sets or clears the Negative Flag
func (cpu *CPU) SetFlagN(value bool) {
	if value {
		cpu.F.Set(cpu.F.Get() | flagN)
	} else {
		cpu.F.Set(cpu.F.Get() &^ flagN)
	}
}

// SetFlagH sets or clears the Half-Carry Flag
func (cpu *CPU) SetFlagH(value bool) {
	if value {
		cpu.F.Set(cpu.F.Get() | flagH)
	} else {
		cpu.F.Set(cpu.F.Get() &^ flagH)
	}
}

// FlagZ returns the status of the Zero flag
func (cpu *CPU) FlagZ() bool {
	return cpu.F.Get()&flagZ == flagZ
}

// FlagC returns the status of the Carry flag
func (cpu *CPU) FlagC() bool {
	return cpu.F.Get()&flagC == flagC
}

// FlagN returns the status of the Negative flag
func (cpu *CPU) FlagN() bool {
	return cpu.F.Get()&flagN == flagN
}

// FlagH returns the status of the Half-Carry flag
func (cpu *CPU) FlagH() bool {
	return cpu.F.Get()&flagH == flagH
}

// DisableInterrupts clears the interrupt master enable flag
func (cpu *CPU) DisableInterrupts() {
	cpu.imeFlag = false
}

// EnableInterrupts sets the interrupt master enable flag
func (cpu *CPU) EnableInterrupts() {
	cpu.imeFlag = true
}

// InterruptsEnabled reads the status of the interrupt master enable flag
func (cpu *CPU) InterruptsEnabled() bool {
	return cpu.imeFlag
}

// Status returns the CPU register status
func (cpu *CPU) Status() State {
	return State{
		AF:  cpu.AF.Get(),
		BC:  cpu.BC.Get(),
		DE:  cpu.DE.Get(),
		HL:  cpu.HL.Get(),
		SP:  cpu.SP.Get(),
		PC:  cpu.PC.Get(),
		IME: cpu.imeFlag,
	}
}

// SetStatus sets CPU status
func (cpu *CPU) SetStatus(state State) {
	cpu.AF.Set(state.AF)
	cpu.BC.Set(state.BC)
	cpu.DE.Set(state.DE)
	cpu.HL.Set(state.HL)
	cpu.SP.Set(state.SP)
	cpu.PC.Set(state.PC)
	cpu.imeFlag = state.IME
}

func (cpu *CPU) printStatus() {
	status := cpu.Status()
	registerTable := tablewriter.NewWriter(os.Stdout)
	registerTable.SetHeader([]string{"Reg", "Value"})
	registerTable.Append([]string{"AF", fmt.Sprintf("%#04x", status.AF)})
	registerTable.Append([]string{"BC", fmt.Sprintf("%#04x", status.BC)})
	registerTable.Append([]string{"DE", fmt.Sprintf("%#04x", status.DE)})
	registerTable.Append([]string{"HL", fmt.Sprintf("%#04x", status.HL)})
	registerTable.Append([]string{"SP", fmt.Sprintf("%#04x", status.SP)})
	registerTable.Append([]string{"PC", fmt.Sprintf("%#04x", status.PC)})
	registerTable.Render()
}
