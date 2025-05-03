package video

// GPU represents the Game Boy's Graphics Processing Unit
type GPU struct {
	display Display
	vram    [0x2000]byte // 8KB of video RAM (0x8000-0x9FFF)
	oam     [0xA0]byte   // Object Attribute Memory (0xFE00-0xFE9F)
	lcdc    byte        // LCD Control register (FF40)
	stat    byte        // LCD Status register (FF41)
	scy     byte        // Scroll Y register (FF42)
	scx     byte        // Scroll X register (FF43)
	ly      byte        // Current scanline (FF44)
	lyc     byte        // Scanline compare (FF45)
	bgp     byte        // Background palette (FF47)
	obp0    byte        // Object palette 0 (FF48)
	obp1    byte        // Object palette 1 (FF49)
	wy      byte        // Window Y position (FF4A)
	wx      byte        // Window X position minus 7 (FF4B)
	
	// Mode state machine
	mode          byte // Current GPU mode (0-3)
	modeCycles    int  // Cycles spent in current mode
	scanlineCycles int  // Cycles spent on current scanline
}

// NewGPU creates a new GPU
func NewGPU(display Display) *GPU {
	return &GPU{
		display: display,
		mode:    0, // Start in H-Blank mode
		vram:    [0x2000]byte{}, // 8KB of VRAM
		lcdc:    0x91,               // Default value for LCDC
	}
}

// Read reads a byte from the GPU memory
func (gpu *GPU) Read(address uint16) byte {
	if address >= 0x8000 && address <= 0x9FFF {
		// VRAM
		return gpu.vram[address-0x8000]
	} else if address >= 0xFE00 && address <= 0xFE9F {
		// OAM (Object Attribute Memory)
		return gpu.oam[address-0xFE00]
	} else if address >= 0xFF40 && address <= 0xFF4B {
		// GPU registers
		switch address {
		case 0xFF40: // LCDC
			return gpu.lcdc
		case 0xFF41: // STAT
			return gpu.stat
		case 0xFF42: // SCY
			return gpu.scy
		case 0xFF43: // SCX
			return gpu.scx
		case 0xFF44: // LY
			return gpu.ly
		case 0xFF45: // LYC
			return gpu.lyc
		case 0xFF47: // BGP
			return gpu.bgp
		case 0xFF48: // OBP0
			return gpu.obp0
		case 0xFF49: // OBP1
			return gpu.obp1
		case 0xFF4A: // WY
			return gpu.wy
		case 0xFF4B: // WX
			return gpu.wx
		}
	}
	return 0xFF // Default return for unmapped addresses
}

// Write writes a byte to the GPU memory
func (gpu *GPU) Write(address uint16, value byte) {
	if address >= 0x8000 && address <= 0x9FFF {
		// VRAM
		gpu.vram[address-0x8000] = value
	} else if address >= 0xFE00 && address <= 0xFE9F {
		// OAM (Object Attribute Memory)
		gpu.oam[address-0xFE00] = value
	} else if address >= 0xFF40 && address <= 0xFF4B {
		// GPU registers
		switch address {
		case 0xFF40: // LCDC
			gpu.lcdc = value
		case 0xFF41: // STAT
			// Only bits 3-6 of STAT are writable
			// Bits 0-2 are read-only
			gpu.stat = (value & 0xF8) | (gpu.stat & 0x07)
		case 0xFF42: // SCY
			gpu.scy = value
		case 0xFF43: // SCX
			gpu.scx = value
		case 0xFF44: // LY (read-only)
			// LY is read-only, writes are ignored
		case 0xFF45: // LYC
			gpu.lyc = value
			// Update the LYC=LY flag in STAT
			gpu.updateLYCFlag()
		case 0xFF47: // BGP
			gpu.bgp = value
		case 0xFF48: // OBP0
			gpu.obp0 = value
		case 0xFF49: // OBP1
			gpu.obp1 = value
		case 0xFF4A: // WY
			gpu.wy = value
		case 0xFF4B: // WX
			gpu.wx = value
		}
	}
}

// Update advances the GPU state by the specified number of cycles
func (gpu *GPU) Update(cycles int) {
	// If LCD is disabled, don't update
	if !gpu.IsLCDEnabled() {
		return
	}
	
	// Update the GPU mode
	gpu.UpdateMode(cycles)
	
	// Render the screen at the end of V-Blank (when starting a new frame)
	if gpu.ly == 0 && gpu.mode == 2 {
		// Clear the display
		gpu.display.Clear()
		
		// Render the background
		if gpu.lcdc&0x01 != 0 { // Check if background is enabled (LCDC bit 0)
			gpu.RenderBackground()
		}
		
		// Render the window
		if gpu.lcdc&0x20 != 0 { // Check if window is enabled (LCDC bit 5)
			gpu.RenderWindow()
		}
		
		// Render the sprites
		if gpu.lcdc&0x02 != 0 { // Check if sprites are enabled (LCDC bit 1)
			gpu.RenderSprites()
		}
		
		// Refresh the display
		gpu.display.Refresh()
	}
}

// IsLCDEnabled returns whether the LCD is enabled (LCDC bit 7)
func (gpu *GPU) IsLCDEnabled() bool {
	return gpu.lcdc&0x80 == 0x80
}
