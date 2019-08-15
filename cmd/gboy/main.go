package main

import (
	"fmt"
  "os"
  "github.com/gorkaio/gboy/pkg/gameboy"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("No ROM file specified!")
		os.Exit(1)
	}

  romfile := os.Args[1]
	fmt.Println("GBoy!")
  fmt.Printf("Loading %q...\n", romfile)
  
  gameboy, err := gameboy.New(romfile)
  if (err != nil) {
    fmt.Println(err.Error())
    os.Exit(1)
  }

  gameboy.Run()
}
