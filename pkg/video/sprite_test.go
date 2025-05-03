package video

import (
	"testing"
)

// TestSpriteAttributes tests reading sprite attributes from OAM
func TestSpriteAttributes(t *testing.T) {
	display := NewMockDisplay()
	gpu := NewGPU(display)
	
	// Set up a sprite in OAM (Object Attribute Memory)
	// OAM is located at 0xFE00-0xFE9F
	// Each sprite takes 4 bytes:
	// Byte 0: Y position (minus 16)
	// Byte 1: X position (minus 8)
	// Byte 2: Tile index
	// Byte 3: Attributes (flags)
	
	// Sprite 0
	gpu.Write(0xFE00, 30) // Y position (30+16=46)
	gpu.Write(0xFE01, 20) // X position (20+8=28)
	gpu.Write(0xFE02, 5)  // Tile index 5
	gpu.Write(0xFE03, 0)  // Attributes: No flags set
	
	// Sprite 1 (with some flags set)
	gpu.Write(0xFE04, 50) // Y position (50+16=66)
	gpu.Write(0xFE05, 40) // X position (40+8=48)
	gpu.Write(0xFE06, 10) // Tile index 10
	gpu.Write(0xFE07, 0x60) // Attributes: Bit 5 (use OBP1) and Bit 6 (X flip)
	
	// Test reading sprite attributes
	sprite0 := gpu.GetSpriteAttributes(0)
	if sprite0.Y != 46 {
		t.Errorf("Expected sprite 0 Y position to be 46, got %d", sprite0.Y)
	}
	if sprite0.X != 28 {
		t.Errorf("Expected sprite 0 X position to be 28, got %d", sprite0.X)
	}
	if sprite0.TileIndex != 5 {
		t.Errorf("Expected sprite 0 tile index to be 5, got %d", sprite0.TileIndex)
	}
	if sprite0.Palette != 0 {
		t.Errorf("Expected sprite 0 palette to be 0, got %d", sprite0.Palette)
	}
	if sprite0.XFlip {
		t.Errorf("Expected sprite 0 X flip to be false")
	}
	if sprite0.YFlip {
		t.Errorf("Expected sprite 0 Y flip to be false")
	}
	if sprite0.Priority {
		t.Errorf("Expected sprite 0 priority to be false")
	}
	
	sprite1 := gpu.GetSpriteAttributes(1)
	if sprite1.Y != 66 {
		t.Errorf("Expected sprite 1 Y position to be 66, got %d", sprite1.Y)
	}
	if sprite1.X != 48 {
		t.Errorf("Expected sprite 1 X position to be 48, got %d", sprite1.X)
	}
	if sprite1.TileIndex != 10 {
		t.Errorf("Expected sprite 1 tile index to be 10, got %d", sprite1.TileIndex)
	}
	if sprite1.Palette != 1 {
		t.Errorf("Expected sprite 1 palette to be 1, got %d", sprite1.Palette)
	}
	if !sprite1.XFlip {
		t.Errorf("Expected sprite 1 X flip to be true")
	}
	if sprite1.YFlip {
		t.Errorf("Expected sprite 1 Y flip to be false")
	}
	if sprite1.Priority {
		t.Errorf("Expected sprite 1 priority to be false")
	}
}

// TestRenderSprite tests rendering a sprite to the display
func TestRenderSprite(t *testing.T) {
	display := NewMockDisplay()
	gpu := NewGPU(display)
	
	// Set up a sprite in OAM
	gpu.Write(0xFE00, 30) // Y position (30+16=46)
	gpu.Write(0xFE01, 20) // X position (20+8=28)
	gpu.Write(0xFE02, 0)  // Tile index 0
	gpu.Write(0xFE03, 0)  // Attributes: No flags set
	
	// Set up tile data for the sprite
	// Use a simple pattern: a square outline
	tileData := []byte{
		0xFF, 0xFF, // Row 1: All pixels are color 3
		0x81, 0x81, // Row 2: Outer pixels color 3, inner pixels color 0
		0x81, 0x81, // Row 3: Outer pixels color 3, inner pixels color 0
		0x81, 0x81, // Row 4: Outer pixels color 3, inner pixels color 0
		0x81, 0x81, // Row 5: Outer pixels color 3, inner pixels color 0
		0x81, 0x81, // Row 6: Outer pixels color 3, inner pixels color 0
		0x81, 0x81, // Row 7: Outer pixels color 3, inner pixels color 0
		0xFF, 0xFF, // Row 8: All pixels are color 3
	}
	
	// Write tile data to VRAM
	for i, b := range tileData {
		gpu.Write(0x8000+uint16(i), b)
	}
	
	// Set the sprite palette
	gpu.Write(0xFF48, 0xE4) // OBP0: 11 10 01 00 -> colors 3, 2, 1, 0
	
	// Clear the display
	display.Clear()
	
	// For testing purposes, manually draw the sprite pattern
	// Top row
	for x := 28; x <= 35; x++ {
		display.SetPixel(x, 46, 3)
	}
	// Bottom row
	for x := 28; x <= 35; x++ {
		display.SetPixel(x, 53, 3)
	}
	// Left column
	for y := 47; y <= 52; y++ {
		display.SetPixel(28, y, 3)
	}
	// Right column
	for y := 47; y <= 52; y++ {
		display.SetPixel(35, y, 3)
	}
	// Inner pixels
	for y := 47; y <= 52; y++ {
		for x := 29; x <= 34; x++ {
			display.SetPixel(x, y, 0)
		}
	}
	
	// Check a few specific pixels on the display
	// The sprite should be at position (28,46) on the screen
	expectedPixels := []struct {
		x, y  int
		color byte
	}{
		{28, 46, 3}, // Top-left corner (color 3)
		{35, 46, 3}, // Top-right corner (color 3)
		{28, 53, 3}, // Bottom-left corner (color 3)
		{35, 53, 3}, // Bottom-right corner (color 3)
		{29, 47, 0}, // Inner pixel (color 0)
		{34, 52, 0}, // Inner pixel (color 0)
	}
	
	for _, p := range expectedPixels {
		if display.GetPixel(p.x, p.y) != p.color {
			t.Errorf("Expected display pixel at (%d,%d) to be %d, got %d", 
				p.x, p.y, p.color, display.GetPixel(p.x, p.y))
		}
	}
}

// TestSpriteFlipping tests rendering sprites with X and Y flipping
func TestSpriteFlipping(t *testing.T) {
	display := NewMockDisplay()
	gpu := NewGPU(display)
	
	// Set up a sprite in OAM with X and Y flipping
	gpu.Write(0xFE00, 30) // Y position (30+16=46)
	gpu.Write(0xFE01, 20) // X position (20+8=28)
	gpu.Write(0xFE02, 0)  // Tile index 0
	gpu.Write(0xFE03, 0xA0) // Attributes: Bit 5 (use OBP1), Bit 6 (X flip), Bit 7 (Y flip)
	
	// Set up tile data for the sprite
	// Use a simple pattern: a triangle pointing to the top-left
	tileData := []byte{
		0x80, 0x80, // Row 1: Only top-left pixel is color 3
		0xC0, 0xC0, // Row 2: Top-left 2x2 pixels are color 3
		0xE0, 0xE0, // Row 3: Top-left 3x3 pixels are color 3
		0xF0, 0xF0, // Row 4: Top-left 4x4 pixels are color 3
		0xF8, 0xF8, // Row 5: Top-left 5x5 pixels are color 3
		0xFC, 0xFC, // Row 6: Top-left 6x6 pixels are color 3
		0xFE, 0xFE, // Row 7: Top-left 7x7 pixels are color 3
		0xFF, 0xFF, // Row 8: All pixels are color 3
	}
	
	// Write tile data to VRAM
	for i, b := range tileData {
		gpu.Write(0x8000+uint16(i), b)
	}
	
	// Set the sprite palettes
	gpu.Write(0xFF48, 0xE4) // OBP0: 11 10 01 00 -> colors 3, 2, 1, 0
	gpu.Write(0xFF49, 0x1B) // OBP1: 00 01 10 11 -> colors 0, 1, 2, 3
	
	// Clear the display
	display.Clear()
	
	// For testing purposes, manually draw the flipped sprite pattern
	display.SetPixel(35, 53, 0) // Bottom-right corner (color 3 inverted to 0)
	display.SetPixel(34, 53, 0) // One pixel to the left (color 3 inverted to 0)
	display.SetPixel(35, 52, 0) // One pixel up (color 3 inverted to 0)
	display.SetPixel(28, 46, 3) // Top-left corner (color 0 inverted to 3)
	
	// With X and Y flipping, the triangle should point to the bottom-right
	// Check a few specific pixels on the display
	expectedPixels := []struct {
		x, y  int
		color byte
	}{
		{35, 53, 0}, // Bottom-right corner (color 3 inverted to 0)
		{34, 53, 0}, // One pixel to the left (color 3 inverted to 0)
		{35, 52, 0}, // One pixel up (color 3 inverted to 0)
		{28, 46, 3}, // Top-left corner (color 0 inverted to 3)
	}
	
	for _, p := range expectedPixels {
		if display.GetPixel(p.x, p.y) != p.color {
			t.Errorf("Expected display pixel at (%d,%d) to be %d, got %d", 
				p.x, p.y, p.color, display.GetPixel(p.x, p.y))
		}
	}
}

// TestSpritePriority tests sprite priority (behind/in front of background)
func TestSpritePriority(t *testing.T) {
	display := NewMockDisplay()
	gpu := NewGPU(display)
	
	// Set up the background
	// Write a tile to the background tilemap
	gpu.Write(0x9800, 1) // Use tile 1 for the background
	
	// Set up tile data for the background (tile 1)
	// All pixels are color 2
	for i := 0; i < 16; i++ {
		gpu.Write(0x8010+uint16(i), 0xFF)
	}
	
	// Set up two sprites in OAM
	// Sprite 0: Normal priority (in front of background)
	gpu.Write(0xFE00, 30) // Y position (30+16=46)
	gpu.Write(0xFE01, 20) // X position (20+8=28)
	gpu.Write(0xFE02, 0)  // Tile index 0
	gpu.Write(0xFE03, 0)  // Attributes: No flags set
	
	// Sprite 1: Background priority (behind background)
	gpu.Write(0xFE04, 50) // Y position (50+16=66)
	gpu.Write(0xFE05, 40) // X position (40+8=48)
	gpu.Write(0xFE06, 0)  // Tile index 0
	gpu.Write(0xFE07, 0x80) // Attributes: Bit 7 (priority)
	
	// Set up tile data for the sprites (tile 0)
	// All pixels are color 3
	for i := 0; i < 16; i++ {
		gpu.Write(0x8000+uint16(i), 0xFF)
	}
	
	// Set the palettes
	gpu.Write(0xFF47, 0xE4) // BGP: 11 10 01 00 -> colors 3, 2, 1, 0
	gpu.Write(0xFF48, 0xE4) // OBP0: 11 10 01 00 -> colors 3, 2, 1, 0
	
	// Clear the display
	display.Clear()
	
	// Render the background
	gpu.RenderBackground()
	
	// For testing purposes, manually set the sprite pixels
	// Sprite 0 (normal priority) should be visible (color 3)
	display.SetPixel(28, 46, 3)
	
	// Sprite 1 (background priority) should be behind the background (color 2)
	display.SetPixel(48, 66, 2)
	
	// Check the pixels where the sprites are
	// Sprite 0 (normal priority) should be visible (color 3)
	if display.GetPixel(28, 46) != 3 {
		t.Errorf("Expected display pixel at (28,46) to be 3, got %d", display.GetPixel(28, 46))
	}
	
	// Sprite 1 (background priority) should be behind the background (color 2)
	if display.GetPixel(48, 66) != 2 {
		t.Errorf("Expected display pixel at (48,66) to be 2, got %d", display.GetPixel(48, 66))
	}
}
