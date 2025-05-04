package main

import (
	"fmt"
	"os"
	"runtime"

	"github.com/gorkaio/gboy/pkg/cpu"
	"github.com/gorkaio/gboy/pkg/gameboy"
	"github.com/gorkaio/gboy/pkg/memory"
	"github.com/gorkaio/gboy/pkg/video"
	"github.com/veandco/go-sdl2/sdl"
)

func main() {
	// Lock the main thread for SDL
	runtime.LockOSThread()

	if len(os.Args) < 2 {
		fmt.Println("No ROM file specified!")
		fmt.Println("Usage: ./gboy <rom_file>")
		os.Exit(1)
	}

	romfile := os.Args[1]
	fmt.Println("GBoy!")
	fmt.Printf("Loading %q...\n", romfile)
	
	// Initialize SDL subsystems
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		fmt.Printf("Error initializing SDL: %v\n", err)
		os.Exit(1)
	}
	defer sdl.Quit()
	
	// Initialize the SDL display with a scale factor of 3
	fmt.Println("Initializing SDL display...")
	display, err := video.NewSDLDisplay(3)
	if err != nil {
		fmt.Printf("Error creating SDL display: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("SDL display initialized successfully")
	defer display.Close()

	// Initialize memory with the display
	fmt.Println("Initializing memory...")
	m := memory.New(display)
	c := cpu.New(m)
	gb, err := gameboy.New(m, c)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	// Load the cart
	fmt.Println("Loading cart...")
	err = gb.LoadCart(romfile)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println("Cart loaded successfully")

	// Run the emulator
	fmt.Println("Starting emulation...")
	gb.Run()
}
