package gameboy

import (
	"fmt"
	"github.com/gorkaio/gboy/pkg/cart"
	"github.com/gorkaio/gboy/pkg/memory"
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
