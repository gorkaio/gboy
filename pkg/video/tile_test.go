package video

import (
	"testing"
)

// TestTileDecoding tests the decoding of tile data from VRAM
func TestTileDecoding(t *testing.T) {
	display := NewMockDisplay()
	gpu := NewGPU(display)
	
	// Set up a simple tile pattern in VRAM
	// Each tile is 16 bytes (8x8 pixels, 2 bits per pixel)
	// First tile at 0x8000 (VRAM offset 0)
	tileData := []byte{
		0x3C, 0x3C, // Row 1: 00111100 00111100 -> 00000000 00111100 (colors 0, 3, 3, 3, 3, 0, 0, 0)
		0x42, 0x42, // Row 2: 01000010 01000010 -> 00000000 01000010 (colors 0, 1, 0, 0, 0, 1, 0, 0)
		0x81, 0xB9, // Row 3: 10000001 10111001 -> 00000000 10111001 (colors 0, 2, 0, 0, 0, 2, 0, 0)
		0x81, 0xA5, // Row 4: 10000001 10100101 -> 00000000 10100101 (colors 0, 2, 0, 0, 0, 2, 0, 0)
		0x81, 0xA5, // Row 5: 10000001 10100101 -> 00000000 10100101 (colors 0, 2, 0, 0, 0, 2, 0, 0)
		0x81, 0xBD, // Row 6: 10000001 10111101 -> 00000000 10111101 (colors 0, 2, 0, 0, 0, 2, 0, 0)
		0x42, 0x42, // Row 7: 01000010 01000010 -> 00000000 01000010 (colors 0, 1, 0, 0, 0, 1, 0, 0)
		0x3C, 0x3C, // Row 8: 00111100 00111100 -> 00000000 00111100 (colors 0, 3, 3, 3, 3, 0, 0, 0)
	}
	
	// Write tile data to VRAM
	for i, b := range tileData {
		gpu.Write(0x8000+uint16(i), b)
	}
	
	// Test decoding the first tile
	tile := gpu.GetTile(0)
	
	// Check a few specific pixels
	expectedPixels := []struct {
		x, y  int
		color byte
	}{
		{0, 0, 0}, // Top-left corner (color 0)
		{2, 0, 3}, // Top row, third pixel (color 3)
		{5, 0, 0}, // Top row, sixth pixel (color 0)
		{1, 1, 1}, // Second row, second pixel (color 1)
		{5, 1, 1}, // Second row, sixth pixel (color 1)
		{2, 7, 3}, // Bottom row, third pixel (color 3)
	}
	
	for _, p := range expectedPixels {
		if tile[p.y][p.x] != p.color {
			t.Errorf("Expected tile[%d][%d] to be %d, got %d", p.y, p.x, p.color, tile[p.y][p.x])
		}
	}
}

// TestRenderTile tests rendering a tile to the display
func TestRenderTile(t *testing.T) {
	display := NewMockDisplay()
	gpu := NewGPU(display)
	
	// Set up a simple tile pattern in VRAM (same as in TestTileDecoding)
	tileData := []byte{
		0x3C, 0x3C, // Row 1
		0x42, 0x42, // Row 2
		0x81, 0xB9, // Row 3
		0x81, 0xA5, // Row 4
		0x81, 0xA5, // Row 5
		0x81, 0xBD, // Row 6
		0x42, 0x42, // Row 7
		0x3C, 0x3C, // Row 8
	}
	
	// Write tile data to VRAM
	for i, b := range tileData {
		gpu.Write(0x8000+uint16(i), b)
	}
	
	// Render the tile to the display at position (10, 20)
	gpu.RenderTile(0, 10, 20)
	
	// Check a few specific pixels on the display
	expectedPixels := []struct {
		x, y  int
		color byte
	}{
		{10, 20, 0}, // Top-left corner (color 0)
		{12, 20, 3}, // Top row, third pixel (color 3)
		{15, 20, 0}, // Top row, sixth pixel (color 0)
		{11, 21, 1}, // Second row, second pixel (color 1)
		{15, 21, 1}, // Second row, sixth pixel (color 1)
		{12, 27, 3}, // Bottom row, third pixel (color 3)
	}
	
	for _, p := range expectedPixels {
		if display.GetPixel(p.x, p.y) != p.color {
			t.Errorf("Expected display pixel at (%d,%d) to be %d, got %d", 
				p.x, p.y, p.color, display.GetPixel(p.x, p.y))
		}
	}
}
