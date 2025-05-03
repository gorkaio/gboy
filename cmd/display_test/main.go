package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorkaio/gboy/pkg/video"
)

func main() {
	// Create a new SDL display with a scale factor of 3
	display, err := video.NewSDLDisplay(3)
	if err != nil {
		fmt.Printf("Error creating SDL display: %v\n", err)
		os.Exit(1)
	}
	defer display.Close()

	// Set up signal handling for clean shutdown
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sigChan
		fmt.Println("Shutting down...")
		display.Close()
		os.Exit(0)
	}()

	// Draw a test pattern
	drawTestPattern(display)

	// Main loop
	for {
		// Refresh the display
		if err := display.Refresh(); err != nil {
			fmt.Printf("Error refreshing display: %v\n", err)
			break
		}

		// Sleep to limit the refresh rate
		time.Sleep(16 * time.Millisecond) // ~60 FPS
	}
}

// drawTestPattern draws a test pattern on the display
func drawTestPattern(display video.Display) {
	// Clear the display
	display.Clear()

	// Draw a border
	for x := 0; x < display.Width(); x++ {
		display.SetPixel(x, 0, 3)                       // Top border
		display.SetPixel(x, display.Height()-1, 3)      // Bottom border
	}
	for y := 0; y < display.Height(); y++ {
		display.SetPixel(0, y, 3)                       // Left border
		display.SetPixel(display.Width()-1, y, 3)       // Right border
	}

	// Draw a checkerboard pattern in the center
	for y := 20; y < 124; y++ {
		for x := 20; x < 140; x++ {
			if (x/8 + y/8) % 2 == 0 {
				display.SetPixel(x, y, 2) // Dark gray
			} else {
				display.SetPixel(x, y, 1) // Light gray
			}
		}
	}

	// Draw the Game Boy logo
	drawGameBoyLogo(display, 40, 60)
}

// drawGameBoyLogo draws a simplified Game Boy logo
func drawGameBoyLogo(display video.Display, x, y int) {
	// Draw "GAME BOY" text (simplified)
	// G
	for i := 0; i < 8; i++ {
		display.SetPixel(x+i, y, 3)
		display.SetPixel(x, y+i, 3)
	}
	display.SetPixel(x+7, y+1, 3)
	display.SetPixel(x+7, y+2, 3)
	for i := 4; i < 8; i++ {
		display.SetPixel(x+i, y+4, 3)
	}
	display.SetPixel(x+7, y+5, 3)
	display.SetPixel(x+7, y+6, 3)
	display.SetPixel(x+7, y+7, 3)
	display.SetPixel(x+6, y+7, 3)
	display.SetPixel(x+5, y+7, 3)
	display.SetPixel(x+4, y+7, 3)

	// A
	for i := 0; i < 8; i++ {
		display.SetPixel(x+10, y+i, 3)
		display.SetPixel(x+17, y+i, 3)
	}
	for i := 11; i < 17; i++ {
		display.SetPixel(x+i, y, 3)
		display.SetPixel(x+i, y+4, 3)
	}

	// M
	for i := 0; i < 8; i++ {
		display.SetPixel(x+20, y+i, 3)
		display.SetPixel(x+27, y+i, 3)
	}
	display.SetPixel(x+21, y+1, 3)
	display.SetPixel(x+22, y+2, 3)
	display.SetPixel(x+23, y+3, 3)
	display.SetPixel(x+24, y+3, 3)
	display.SetPixel(x+25, y+2, 3)
	display.SetPixel(x+26, y+1, 3)

	// E
	for i := 0; i < 8; i++ {
		display.SetPixel(x+30, y+i, 3)
	}
	for i := 31; i < 37; i++ {
		display.SetPixel(x+i, y, 3)
		display.SetPixel(x+i, y+7, 3)
	}
	for i := 31; i < 35; i++ {
		display.SetPixel(x+i, y+4, 3)
	}

	// B
	for i := 0; i < 8; i++ {
		display.SetPixel(x+40, y+i, 3)
	}
	for i := 0; i < 6; i += 5 {
		display.SetPixel(x+41, y+i, 3)
		display.SetPixel(x+42, y+i, 3)
		display.SetPixel(x+43, y+i, 3)
		display.SetPixel(x+44, y+i, 3)
		display.SetPixel(x+45, y+i+1, 3)
		display.SetPixel(x+45, y+i+2, 3)
		display.SetPixel(x+44, y+i+3, 3)
		display.SetPixel(x+43, y+i+3, 3)
		display.SetPixel(x+42, y+i+3, 3)
		display.SetPixel(x+41, y+i+3, 3)
	}

	// O
	for i := 0; i < 8; i++ {
		display.SetPixel(x+50, y+i, 3)
		display.SetPixel(x+57, y+i, 3)
	}
	for i := 51; i < 57; i++ {
		display.SetPixel(x+i, y, 3)
		display.SetPixel(x+i, y+7, 3)
	}

	// Y
	for i := 0; i < 3; i++ {
		display.SetPixel(x+60, y+i, 3)
		display.SetPixel(x+67, y+i, 3)
	}
	display.SetPixel(x+61, y+3, 3)
	display.SetPixel(x+62, y+3, 3)
	display.SetPixel(x+65, y+3, 3)
	display.SetPixel(x+66, y+3, 3)
	for i := 4; i < 8; i++ {
		display.SetPixel(x+63, y+i, 3)
		display.SetPixel(x+64, y+i, 3)
	}
}
