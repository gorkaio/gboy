package video

// GPU represents the Game Boy's Graphics Processing Unit
type GPU struct {
	display Display
	vram    []byte // 8KB Video RAM (0x8000-0x9FFF)
	
	// Video registers
	lcdc byte // LCD Control (0xFF40)
	stat byte // LCDC Status (0xFF41)
	scy  byte // Scroll Y (0xFF42)
	scx  byte // Scroll X (0xFF43)
	ly   byte // LCD Y-Coordinate (0xFF44)
	lyc  byte // LY Compare (0xFF45)
	bgp  byte // BG Palette Data (0xFF47)
	obp0 byte // Object Palette 0 (0xFF48)
	obp1 byte // Object Palette 1 (0xFF49)
	wy   byte // Window Y Position (0xFF4A)
	wx   byte // Window X Position (0xFF4B)
	
	// Timing
	modeClock int // Clock for the current mode
	mode      byte // Current GPU mode
}

// NewGPU creates a new GPU instance
func NewGPU(display Display) *GPU {
	gpu := &GPU{
		display: display,
		vram:    make([]byte, 8192), // 8KB of VRAM
		lcdc:    0x91,               // Default value for LCDC
	}
	
	// Clear VRAM
	for i := range gpu.vram {
		gpu.vram[i] = 0
	}
	
	return gpu
}

// Read reads a byte from the GPU memory space
func (gpu *GPU) Read(address uint16) byte {
	// VRAM (0x8000-0x9FFF)
	if address >= 0x8000 && address <= 0x9FFF {
		return gpu.vram[address-0x8000]
	}
	
	// Video registers
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
	
	return 0xFF // Default return for unmapped addresses
}

// Write writes a byte to the GPU memory space
func (gpu *GPU) Write(address uint16, value byte) {
	// VRAM (0x8000-0x9FFF)
	if address >= 0x8000 && address <= 0x9FFF {
		gpu.vram[address-0x8000] = value
		return
	}
	
	// Video registers
	switch address {
	case 0xFF40: // LCDC
		gpu.lcdc = value
	case 0xFF41: // STAT
		// Only bits 3-6 are writable
		gpu.stat = (value & 0x78) | (gpu.stat & 0x87)
	case 0xFF42: // SCY
		gpu.scy = value
	case 0xFF43: // SCX
		gpu.scx = value
	case 0xFF44: // LY - Read-only
		// LY is read-only, writes are ignored
	case 0xFF45: // LYC
		gpu.lyc = value
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

// Update advances the GPU state by the specified number of cycles
func (gpu *GPU) Update(cycles int) {
	// This will be implemented in a later step
}

// IsLCDEnabled returns whether the LCD is enabled (LCDC bit 7)
func (gpu *GPU) IsLCDEnabled() bool {
	return gpu.lcdc&0x80 == 0x80
}
