package main

import (
  "fmt"
  "os"
)

func main() {
  if len(os.Args) < 2 {
    fmt.Println("No ROM file specified!")
    os.Exit(1)
  }

  fmt.Println("GBoy!")
  fmt.Printf("Loading %q\n", os.Args[1])
}
