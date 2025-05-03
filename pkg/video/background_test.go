package video

import (
	"testing"
)

// TestBackgroundTilemap tests accessing the background tilemap
func TestBackgroundTilemap(t *testing.T) {
	display := NewMockDisplay()
	gpu := NewGPU(display)
	
	// The background tilemap is located at 0x9800-0x9BFF (1KB)
	// Each byte in this area represents a tile number (0-255)
	
	// Write some tile numbers to the background tilemap
	gpu.Write(0x9800, 1) // Top-left corner tile
	gpu.Write(0x9801, 2) // Next tile to the right
	gpu.Write(0x9820, 3) // Tile on the next row (32 bytes per row)
	
	// Test reading from the background tilemap
	if gpu.Read(0x9800) != 1 {
		t.Errorf("Expected background tilemap at 0x9800 to be 1, got %d", gpu.Read(0x9800))
	}
	
	if gpu.Read(0x9801) != 2 {
		t.Errorf("Expected background tilemap at 0x9801 to be 2, got %d", gpu.Read(0x9801))
	}
	
	if gpu.Read(0x9820) != 3 {
		t.Errorf("Expected background tilemap at 0x9820 to be 3, got %d", gpu.Read(0x9820))
	}
	
	// Test getting a tile number from the background map
	tileNum := gpu.GetBackgroundTileNumber(0, 0) // Top-left corner
	if tileNum != 1 {
		t.Errorf("Expected background tile at (0,0) to be 1, got %d", tileNum)
	}
	
	tileNum = gpu.GetBackgroundTileNumber(1, 0) // One tile to the right
	if tileNum != 2 {
		t.Errorf("Expected background tile at (1,0) to be 2, got %d", tileNum)
	}
	
	tileNum = gpu.GetBackgroundTileNumber(0, 1) // One tile down
	if tileNum != 3 {
		t.Errorf("Expected background tile at (0,1) to be 3, got %d", tileNum)
	}
}

// TestRenderBackground tests rendering the background
func TestRenderBackground(t *testing.T) {
	display := NewMockDisplay()
	gpu := NewGPU(display)
	
	// Set up a simple background tilemap
	gpu.Write(0x9800, 0) // Use tile 0 for the top-left corner
	
	// Set up a simple tile pattern in VRAM for tile 0
	// This is the same pattern as in the tile tests
	tileData := []byte{
		0x3C, 0x3C, // Row 1: colors 0, 3, 3, 3, 3, 0, 0, 0
		0x42, 0x42, // Row 2: colors 0, 1, 0, 0, 0, 1, 0, 0
		0x81, 0xB9, // Row 3: colors 0, 2, 0, 0, 0, 2, 0, 0
		0x81, 0xA5, // Row 4: colors 0, 2, 0, 0, 0, 2, 0, 0
		0x81, 0xA5, // Row 5: colors 0, 2, 0, 0, 0, 2, 0, 0
		0x81, 0xBD, // Row 6: colors 0, 2, 0, 0, 0, 2, 0, 0
		0x42, 0x42, // Row 7: colors 0, 1, 0, 0, 0, 1, 0, 0
		0x3C, 0x3C, // Row 8: colors 0, 3, 3, 3, 3, 0, 0, 0
	}
	
	// Write tile data to VRAM
	for i, b := range tileData {
		gpu.Write(0x8000+uint16(i), b)
	}
	
	// Set the background palette
	gpu.Write(0xFF47, 0xE4) // 11 10 01 00 -> colors 3, 2, 1, 0
	
	// Render the background
	gpu.RenderBackground()
	
	// Check a few specific pixels on the display
	// The top-left corner of the screen should show the tile we set up
	expectedPixels := []struct {
		x, y  int
		color byte
	}{
		{0, 0, 0}, // Top-left corner (color 0)
		{1, 0, 3}, // Top row, second pixel (color 3)
		{5, 0, 0}, // Top row, sixth pixel (color 0)
		{1, 1, 1}, // Second row, second pixel (color 1)
		{5, 1, 1}, // Second row, sixth pixel (color 1)
		{1, 7, 3}, // Bottom row, second pixel (color 3)
	}
	
	for _, p := range expectedPixels {
		if display.GetPixel(p.x, p.y) != p.color {
			t.Errorf("Expected display pixel at (%d,%d) to be %d, got %d", 
				p.x, p.y, p.color, display.GetPixel(p.x, p.y))
		}
	}
}

// TestBackgroundScrolling tests the background scrolling functionality
func TestBackgroundScrolling(t *testing.T) {
	display := NewMockDisplay()
	gpu := NewGPU(display)
	
	// Set up a simple background tilemap
	gpu.Write(0x9800, 0) // Use tile 0 for the top-left corner
	gpu.Write(0x9801, 1) // Use tile 1 for the next tile to the right
	
	// Set up tile patterns in VRAM
	// Tile 0: All pixels are color 1
	for i := 0; i < 16; i++ {
		gpu.Write(0x8000+uint16(i), 0xFF)
	}
	
	// Tile 1: All pixels are color 2
	for i := 0; i < 16; i++ {
		gpu.Write(0x8010+uint16(i), 0xFF)
	}
	
	// Set the background palette
	gpu.Write(0xFF47, 0xE4) // 11 10 01 00 -> colors 3, 2, 1, 0
	
	// Test with no scrolling
	gpu.Write(0xFF42, 0) // SCY = 0
	gpu.Write(0xFF43, 0) // SCX = 0
	gpu.RenderBackground()
	
	// The top-left corner should be from tile 0 (color 1)
	if display.GetPixel(0, 0) != 1 {
		t.Errorf("Expected display pixel at (0,0) to be 1, got %d", display.GetPixel(0, 0))
	}
	
	// Test with horizontal scrolling (SCX = 8)
	gpu.Write(0xFF43, 8) // SCX = 8
	gpu.RenderBackground()
	
	// The top-left corner should now be from tile 1 (color 2)
	if display.GetPixel(0, 0) != 2 {
		t.Errorf("Expected display pixel at (0,0) to be 2, got %d", display.GetPixel(0, 0))
	}
}
