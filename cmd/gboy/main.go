package main

import (
	"fmt"
  "os"
  "github.com/gorkaio/gboy/pkg/cart"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("No ROM file specified!")
		os.Exit(1)
	}

	fmt.Println("GBoy!")
  fmt.Printf("Loading %q...\n", os.Args[1])
  cart, err := cart.LoadFromFile(os.Args[1])
  if (err != nil) {
    fmt.Println(err)
    os.Exit(1)
  }
  fmt.Printf("  title: %q\n", cart.Title())
  fmt.Printf("  type: %s\n", cart.Type())
}
