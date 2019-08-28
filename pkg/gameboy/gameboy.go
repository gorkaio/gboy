package gameboy

import (
	"fmt"

	"github.com/gorkaio/gboy/pkg/cpu"
	"github.com/gorkaio/gboy/pkg/memory"
)

const cyclesPerScanline = 456
const cyclesPerFrame = 69905

// Gameboy struct
type Gameboy struct {
	cpu             cpu.CPUInterface
	mem             memory.MemoryInterface
	scanlineCounter int
	paused          bool
}

// New initialises a new Gameboy System
func New(mem memory.MemoryInterface, cpu cpu.CPUInterface) (*Gameboy, error) {
	gameboy := &Gameboy{
		mem:             mem,
		cpu:             cpu,
		scanlineCounter: cyclesPerScanline,
		paused:          false,
	}

	return gameboy, nil
}

func (gb *Gameboy) LoadCart(romfile string) error {
	err := gb.mem.Load(romfile)
	if err != nil {
		return err
	}
	return nil
}

func (gb *Gameboy) Eject() {
	gb.mem.Eject()
}

// Run runs the emulation
func (gb *Gameboy) Run() {
	for !gb.paused {
		gb.Update()
	}
}

// Update runs the system update cycle for a single frame
func (gb *Gameboy) Update() {
	cyclesConsumed := 0
	for cyclesConsumed < cyclesPerFrame {
		cycles, err := gb.cpu.Step()
		if err != nil {
			fmt.Println(err.Error())
		}
		gb.updateGraphics(cycles)
		cyclesConsumed += cycles
	}
	// render
}

func (gb *Gameboy) updateGraphics(cycles int) {
	gb.scanlineCounter -= cycles
	if gb.scanlineCounter <= 0 {
		currentScanline := gb.mem.Read(0xFF44) + 1
		if currentScanline > 153 {
			currentScanline = 0
		}
		gb.mem.Write(0xFF44, currentScanline)
		gb.scanlineCounter = cyclesPerScanline
	}
}
