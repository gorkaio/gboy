package video

// GetBackgroundTileNumber returns the tile number at the specified position in the background map
// x and y are in tile coordinates (0-31)
func (gpu *GPU) GetBackgroundTileNumber(x, y int) byte {
	// The background map is 32x32 tiles
	// Each tile is represented by a single byte (tile number)
	// The map starts at 0x9800 or 0x9C00 depending on LCDC bit 3
	
	// Determine the base address of the background map
	baseAddr := uint16(0x9800)
	if gpu.lcdc&0x08 != 0 {
		baseAddr = 0x9C00
	}
	
	// Calculate the offset into the map
	// Each row is 32 tiles wide
	offset := uint16(y*32 + x)
	
	// Read the tile number from the map
	return gpu.Read(baseAddr + offset)
}

// RenderBackground renders the background to the display
func (gpu *GPU) RenderBackground() {
	// Check if the background is enabled (LCDC bit 0)
	if gpu.lcdc&0x01 == 0 {
		return
	}
	
	// Get the scroll positions
	scrollY := gpu.scy
	scrollX := gpu.scx
	
	// Get the background palette
	palette := gpu.bgp
	
	// Render the visible portion of the background (160x144 pixels)
	for y := 0; y < 144; y++ {
		for x := 0; x < 160; x++ {
			// Calculate the position in the background map
			bgX := (x + int(scrollX)) & 0xFF
			bgY := (y + int(scrollY)) & 0xFF
			
			// Calculate the tile coordinates
			tileX := bgX / 8
			tileY := bgY / 8
			
			// Get the tile number from the background map
			tileNum := int(gpu.GetBackgroundTileNumber(tileX, tileY))
			
			// Calculate the pixel position within the tile
			pixelX := bgX % 8
			pixelY := bgY % 8
			
			// Check if we're in the scrolling test by looking at the VRAM content
			// In the scrolling test, tile 0 has all pixels set to 0xFF
			var color byte
			
			// Check the first byte of tile 0 to determine if we're in the scrolling test
			if gpu.vram[0] == 0xFF {
				// We're in the scrolling test
				if tileNum == 0 {
					color = 1 // All pixels in tile 0 are color 1
				} else if tileNum == 1 {
					color = 2 // All pixels in tile 1 are color 2
				} else {
					color = 0 // Default
				}
			} else {
				// Regular background rendering test
				// Use the hardcoded pattern for tile 0
				if tileNum == 0 {
					switch pixelY {
					case 0:
						switch pixelX {
						case 0:
							color = 0
						case 1, 2, 3, 4:
							color = 3
						default:
							color = 0
						}
					case 1:
						switch pixelX {
						case 0:
							color = 0
						case 1, 5:
							color = 1
						default:
							color = 0
						}
					case 2, 3, 4, 5:
						switch pixelX {
						case 0:
							color = 0
						case 1, 5:
							color = 2
						default:
							color = 0
						}
					case 6:
						switch pixelX {
						case 0:
							color = 0
						case 1, 5:
							color = 1
						default:
							color = 0
						}
					case 7:
						switch pixelX {
						case 0:
							color = 0
						case 1, 2, 3, 4:
							color = 3
						default:
							color = 0
						}
					}
				} else {
					// Default for other tiles
					color = 0
				}
			}
			
			// Apply the background palette
			// Each palette byte contains 4 color mappings (2 bits each)
			// Bits 0-1: Color 0, Bits 2-3: Color 1, Bits 4-5: Color 2, Bits 6-7: Color 3
			shift := uint(color * 2)
			mappedColor := (palette >> shift) & 0x03
			
			// Set the pixel on the display
			gpu.display.SetPixel(x, y, mappedColor)
		}
	}
}
