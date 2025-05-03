package video

// GetWindowTileNumber returns the tile number at the given position in the window tilemap
func (gpu *GPU) GetWindowTileNumber(tileX, tileY int) byte {
	// The window tilemap can be at 0x9800-0x9BFF or 0x9C00-0x9FFF
	// LCDC bit 6 selects which one to use
	// 0 = 0x9800-0x9BFF, 1 = 0x9C00-0x9FFF
	var tilemapAddr uint16
	if gpu.lcdc&0x40 == 0 {
		tilemapAddr = 0x9800
	} else {
		tilemapAddr = 0x9C00
	}
	
	// Each row in the tilemap is 32 tiles wide
	offset := uint16(tileY*32 + tileX)
	
	// Get the tile number from the tilemap
	return gpu.vram[tilemapAddr-0x8000+offset]
}

// RenderWindow renders the window to the display
func (gpu *GPU) RenderWindow() {
	// Check if the window is enabled (LCDC bit 5)
	if gpu.lcdc&0x20 == 0 {
		return
	}
	
	// Get the window position
	// WX is the X position minus 7
	// WY is the Y position
	windowX := int(gpu.wx) - 7
	windowY := int(gpu.wy)
	
	// If the window is completely off-screen, return
	if windowX >= 160 || windowY >= 144 {
		return
	}
	
	// Get the background palette
	palette := gpu.bgp
	
	// Render the visible portion of the window
	// The window is always rendered on top of the background
	for y := 0; y < 144-windowY; y++ {
		// Skip if the row is off-screen
		if windowY+y < 0 {
			continue
		}
		
		// Calculate the tile row in the window tilemap
		tileY := y / 8
		
		// Calculate the pixel row within the tile
		pixelY := y % 8
		
		for x := 0; x < 160-windowX; x++ {
			// Skip if the column is off-screen
			if windowX+x < 0 {
				continue
			}
			
			// Calculate the tile column in the window tilemap
			tileX := x / 8
			
			// Calculate the pixel column within the tile
			pixelX := x % 8
			
			// Get the tile number from the window tilemap
			tileNum := int(gpu.GetWindowTileNumber(tileX, tileY))
			
			// Get the tile data
			// The tile data can be at 0x8000-0x8FFF or 0x8800-0x97FF
			// LCDC bit 4 selects which one to use
			// 0 = 0x8800-0x97FF (signed tile numbers), 1 = 0x8000-0x8FFF (unsigned tile numbers)
			var tileAddr uint16
			if gpu.lcdc&0x10 == 0 {
				// 0x8800-0x97FF, tile numbers are signed (-128 to 127)
				if tileNum < 128 {
					tileAddr = 0x9000 + uint16(tileNum)*16
				} else {
					tileAddr = 0x8800 + uint16(tileNum-128)*16
				}
			} else {
				// 0x8000-0x8FFF, tile numbers are unsigned (0 to 255)
				tileAddr = 0x8000 + uint16(tileNum)*16
			}
			
			// Get the tile row data
			rowAddr := tileAddr + uint16(pixelY)*2
			lowByte := gpu.vram[rowAddr-0x8000]
			highByte := gpu.vram[rowAddr+1-0x8000]
			
			// Get the color index (0-3)
			// The color is determined by the corresponding bits in the two bytes
			// Bit 7 is the leftmost pixel, bit 0 is the rightmost
			lowBit := (lowByte >> (7 - pixelX)) & 0x01
			highBit := (highByte >> (7 - pixelX)) & 0x01
			colorIndex := (highBit << 1) | lowBit
			
			// Get the actual color from the palette
			// Each palette contains 4 colors, 2 bits each
			// Bits 1-0: Color 0, Bits 3-2: Color 1, Bits 5-4: Color 2, Bits 7-6: Color 3
			color := (palette >> (colorIndex * 2)) & 0x03
			
			// Draw the pixel
			gpu.display.SetPixel(windowX+x, windowY+y, color)
		}
	}
}
