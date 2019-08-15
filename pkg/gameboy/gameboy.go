package gameboy

import (
	"fmt"

	cart "github.com/gorkaio/gboy/pkg/cart"
	cpu "github.com/gorkaio/gboy/pkg/cpu"
	memory "github.com/gorkaio/gboy/pkg/memory"
)

// Gameboy struct
type Gameboy struct {
	cpu     *cpu.CPU
	romfile string
}

// New initialises a new Gameboy System
func New(romfile string) (*Gameboy, error) {
	cart, err := cart.New()
	if err != nil {
		return nil, err
	}

	err = cart.Load(romfile)
	if (err != nil) {
		return nil, err
	}
	
	mem, err := memory.New(cart)
	if err != nil {
		return nil, err
	}

	gameboy := &Gameboy{
		cpu:     cpu.New(mem),
		romfile: romfile,
	}
	return gameboy, nil
}

// Run executes game
func (gb *Gameboy) Run() {
	for {
		_, err := gb.cpu.Step()
		if err != nil {
			fmt.Println(err.Error())
		}
	}
}
