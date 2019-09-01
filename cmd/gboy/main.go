package main

import (
	"fmt"
	"github.com/gorkaio/gboy/pkg/gameboy"
	"github.com/gorkaio/gboy/pkg/cpu"
	"github.com/gorkaio/gboy/pkg/memory"
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

	m := memory.New()
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
