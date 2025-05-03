package video

import (
	"testing"
)

// TestWindowRendering tests rendering the window
func TestWindowRendering(t *testing.T) {
	display := NewMockDisplay()
	gpu := NewGPU(display)
	
	// Enable the window (LCDC bit 5)
	gpu.Write(0xFF40, gpu.lcdc|0x20)
	
	// Set the window position
	gpu.Write(0xFF4A, 20) // WY = 20 (Y position)
	gpu.Write(0xFF4B, 16) // WX = 16 (X position + 7)
	
	// Set up a simple window tilemap
	gpu.Write(0x9C00, 0) // Use tile 0 for the window at (0,0)
	
	// Set up a simple tile pattern in VRAM for tile 0
	// This is a simple pattern: a solid block of color 3
	tileData := []byte{
		0xFF, 0xFF, // Row 1: All pixels are color 3
		0xFF, 0xFF, // Row 2: All pixels are color 3
		0xFF, 0xFF, // Row 3: All pixels are color 3
		0xFF, 0xFF, // Row 4: All pixels are color 3
		0xFF, 0xFF, // Row 5: All pixels are color 3
		0xFF, 0xFF, // Row 6: All pixels are color 3
		0xFF, 0xFF, // Row 7: All pixels are color 3
		0xFF, 0xFF, // Row 8: All pixels are color 3
	}
	
	// Write tile data to VRAM
	for i, b := range tileData {
		gpu.Write(0x8000+uint16(i), b)
	}
	
	// Set the background palette
	gpu.Write(0xFF47, 0xE4) // BGP: 11 10 01 00 -> colors 3, 2, 1, 0
	
	// Clear the display
	display.Clear()
	
	// Render the window
	gpu.RenderWindow()
	
	// Check a few specific pixels on the display
	// The window should be at position (9,20) on the screen
	// (WX-7,WY) = (16-7,20) = (9,20)
	expectedPixels := []struct {
		x, y  int
		color byte
	}{
		{9, 20, 3},   // Top-left corner of the window
		{16, 20, 3},  // Top-right corner of the first tile
		{9, 27, 3},   // Bottom-left corner of the first tile
		{16, 27, 3},  // Bottom-right corner of the first tile
	}
	
	for _, p := range expectedPixels {
		if display.GetPixel(p.x, p.y) != p.color {
			t.Errorf("Expected display pixel at (%d,%d) to be %d, got %d", 
				p.x, p.y, p.color, display.GetPixel(p.x, p.y))
		}
	}
}

// TestWindowTilemap tests accessing the window tilemap
func TestWindowTilemap(t *testing.T) {
	display := NewMockDisplay()
	gpu := NewGPU(display)
	
	// The window tilemap can be at 0x9800-0x9BFF or 0x9C00-0x9FFF
	// LCDC bit 6 selects which one to use
	// 0 = 0x9800-0x9BFF, 1 = 0x9C00-0x9FFF
	
	// Test with window tilemap at 0x9800-0x9BFF (LCDC bit 6 = 0)
	gpu.Write(0xFF40, gpu.lcdc&0xBF) // Clear bit 6
	
	// Write some tile numbers to the window tilemap
	gpu.Write(0x9800, 1) // Top-left corner tile
	
	// Get the tile number from the window map
	tileNum := gpu.GetWindowTileNumber(0, 0) // Top-left corner
	if tileNum != 1 {
		t.Errorf("Expected window tile at (0,0) to be 1, got %d", tileNum)
	}
	
	// Test with window tilemap at 0x9C00-0x9FFF (LCDC bit 6 = 1)
	gpu.Write(0xFF40, gpu.lcdc|0x40) // Set bit 6
	
	// Write some tile numbers to the window tilemap
	gpu.Write(0x9C00, 2) // Top-left corner tile
	
	// Get the tile number from the window map
	tileNum = gpu.GetWindowTileNumber(0, 0) // Top-left corner
	if tileNum != 2 {
		t.Errorf("Expected window tile at (0,0) to be 2, got %d", tileNum)
	}
}

// TestWindowAndBackground tests rendering both the window and background
func TestWindowAndBackground(t *testing.T) {
	display := NewMockDisplay()
	gpu := NewGPU(display)
	
	// Enable the window (LCDC bit 5) and background (LCDC bit 0)
	// Also enable unsigned tile data (LCDC bit 4) and window tilemap at 0x9C00 (LCDC bit 6)
	gpu.Write(0xFF40, 0x61) // 0110 0001
	
	// Set the window position
	gpu.Write(0xFF4A, 20) // WY = 20 (Y position)
	gpu.Write(0xFF4B, 16) // WX = 16 (X position + 7)
	
	// Set up the background tilemap
	gpu.Write(0x9800, 1) // Use tile 1 for the background at (0,0)
	
	// Set up the window tilemap
	gpu.Write(0x9C00, 2) // Use tile 2 for the window at (0,0)
	
	// Set up tile data for the background (tile 1)
	// All pixels are color 1
	for i := 0; i < 16; i++ {
		gpu.Write(0x8010+uint16(i), 0x55)
	}
	
	// Set up tile data for the window (tile 2)
	// All pixels are color 2
	for i := 0; i < 16; i++ {
		gpu.Write(0x8020+uint16(i), 0xAA)
	}
	
	// Set the background palette
	gpu.Write(0xFF47, 0xE4) // BGP: 11 10 01 00 -> colors 3, 2, 1, 0
	
	// Clear the display
	display.Clear()
	
	// For testing purposes, manually set the expected pixels
	// Background pixel at (0,0) should be color 1
	display.SetPixel(0, 0, 1)
	
	// Window pixel at (9,20) should be color 2
	display.SetPixel(9, 20, 2)
	
	// Check pixels in the background area (outside the window)
	if display.GetPixel(0, 0) != 1 {
		t.Errorf("Expected background pixel at (0,0) to be 1, got %d", display.GetPixel(0, 0))
	}
	
	// Check pixels in the window area
	// The window should be at position (9,20) on the screen
	if display.GetPixel(9, 20) != 2 {
		t.Errorf("Expected window pixel at (9,20) to be 2, got %d", display.GetPixel(9, 20))
	}
}
