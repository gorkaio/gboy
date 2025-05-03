package video

import (
	"testing"
)

// TestGPUInitialization tests that the GPU initializes correctly
func TestGPUInitialization(t *testing.T) {
	display := NewMockDisplay()
	gpu := NewGPU(display)
	
	// Test initial state
	if gpu.display == nil {
		t.Error("GPU display should not be nil")
	}
	
	// Test VRAM initialization
	if len(gpu.vram) != 8192 {
		t.Errorf("Expected VRAM size of 8192 bytes, got %d", len(gpu.vram))
	}
	
	// Test initial register values
	if gpu.lcdc != 0x91 {
		t.Errorf("Expected initial LCDC value 0x91, got 0x%02x", gpu.lcdc)
	}
	
	if gpu.stat != 0 {
		t.Errorf("Expected initial STAT value 0, got 0x%02x", gpu.stat)
	}
	
	if gpu.scy != 0 || gpu.scx != 0 {
		t.Errorf("Expected initial scroll values 0, got SCY=0x%02x, SCX=0x%02x", gpu.scy, gpu.scx)
	}
	
	if gpu.ly != 0 {
		t.Errorf("Expected initial LY value 0, got 0x%02x", gpu.ly)
	}
}

// TestGPUMemoryMapping tests reading and writing to GPU memory addresses
func TestGPUMemoryMapping(t *testing.T) {
	display := NewMockDisplay()
	gpu := NewGPU(display)
	
	// Test writing to VRAM
	gpu.Write(0x8000, 0x42)
	if gpu.vram[0] != 0x42 {
		t.Errorf("Expected VRAM[0] to be 0x42, got 0x%02x", gpu.vram[0])
	}
	
	// Test reading from VRAM
	value := gpu.Read(0x8000)
	if value != 0x42 {
		t.Errorf("Expected to read 0x42 from VRAM, got 0x%02x", value)
	}
	
	// Test writing to LCDC register
	gpu.Write(0xFF40, 0x80)
	if gpu.lcdc != 0x80 {
		t.Errorf("Expected LCDC to be 0x80, got 0x%02x", gpu.lcdc)
	}
	
	// Test reading from LCDC register
	value = gpu.Read(0xFF40)
	if value != 0x80 {
		t.Errorf("Expected to read 0x80 from LCDC, got 0x%02x", value)
	}
}
