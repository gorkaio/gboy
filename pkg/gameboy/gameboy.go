package gameboy

import (
	"fmt"
	"os"
	"github.com/gorkaio/gboy/pkg/cart"
	"github.com/gorkaio/gboy/pkg/memory"
	"github.com/gorkaio/gboy/pkg/video"
	"io/ioutil"
)

//go:generate mockgen -destination=mocks/memory_mock.go -package=gameboy_mock github.com/gorkaio/gboy/pkg/gameboy Memory
//go:generate mockgen -destination=mocks/cpu_mock.go -package=gameboy_mock github.com/gorkaio/gboy/pkg/gameboy CPU

const cyclesPerScanline = 456
const cyclesPerFrame = 69905

// Memory defines the interface for memory interaction
type Memory interface {
	Load(cart memory.Cart)
	Eject()
	Read(address uint16) uint8
	Write(address uint16, data uint8)
	GetGPU() *video.GPU
}

// CPU defines the interface for CPU interaction
type CPU interface {
	Step() (int, error)
}

// Gameboy struct
type Gameboy struct {
	cpu             CPU
	mem             Memory
	romfile         string
	scanlineCounter int
	paused          bool
}

// New initialises a new Gameboy System
func New(mem Memory, cpu CPU) (*Gameboy, error) {
	gameboy := &Gameboy{
		mem:             mem,
		cpu:             cpu,
		scanlineCounter: cyclesPerScanline,
		paused:          false,
	}

	return gameboy, nil
}

// LoadCart loads a cart using the loader
func (gb *Gameboy) LoadCart(romfile string) error {
	data, err := ioutil.ReadFile(romfile)
	if err != nil {
		return err
	}

	cart, err := cart.NewCart(data)
	if err != nil {
		return err
	}

	gb.mem.Load(cart)
	gb.romfile = romfile
	return nil
}

// Eject ejects a cart from memory
func (gb *Gameboy) Eject() {
	gb.romfile = ""
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
			// Format the error message in a parsable format
			fmt.Fprintf(os.Stderr, "ERROR: %s\n", err.Error())
			// Exit with error code 1
			os.Exit(1)
		}
		gb.updateGraphics(cycles)
		cyclesConsumed += cycles
	}
	
	// Get the GPU and update it
	gpu := gb.mem.GetGPU()
	gpu.Update(0) // Update with 0 cycles to trigger rendering
}

func (gb *Gameboy) updateGraphics(cycles int) {
	// Get the GPU and update it with the elapsed cycles
	gpu := gb.mem.GetGPU()
	gpu.Update(cycles)
	
	// Update the scanline counter for compatibility with existing code
	gb.scanlineCounter -= cycles
	if gb.scanlineCounter <= 0 {
		gb.scanlineCounter = cyclesPerScanline
	}
}
