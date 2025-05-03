package main

import (
	"fmt"
	"github.com/gorkaio/gboy/pkg/gameboy"
	"github.com/gorkaio/gboy/pkg/cpu"
	"github.com/gorkaio/gboy/pkg/memory"
	"github.com/gorkaio/gboy/pkg/video"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("No ROM file specified!")
		os.Exit(1)
	}

	romfile := os.Args[1]
	fmt.Println("GBoy!")
	fmt.Printf("Loading %q...\n", romfile)
	
	// Initialize the SDL display with a scale factor of 3
	display, err := video.NewSDLDisplay(3)
	if err != nil {
		fmt.Printf("Error creating SDL display: %v\n", err)
		os.Exit(1)
	}
	defer display.Close()

	// Initialize memory with the display
	m := memory.New(display)
	c := cpu.New(m)
	gb, err := gameboy.New(m, c)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	err = gb.LoadCart(romfile)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	gb.Run()
}
