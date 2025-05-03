package video

import (
	"testing"
)

// TestGPUModes tests the GPU mode state machine
func TestGPUModes(t *testing.T) {
	display := NewMockDisplay()
	gpu := NewGPU(display)
	
	// Initial mode should be 0 (H-Blank)
	if gpu.GetMode() != 0 {
		t.Errorf("Expected initial GPU mode to be 0, got %d", gpu.GetMode())
	}
	
	// Test mode transitions
	// Mode 0 (H-Blank) -> Mode 2 (OAM Search)
	gpu.UpdateMode(80)
	if gpu.GetMode() != 2 {
		t.Errorf("Expected GPU mode to be 2 after 80 cycles, got %d", gpu.GetMode())
	}
	
	// Mode 2 (OAM Search) -> Mode 3 (Pixel Transfer)
	gpu.UpdateMode(80)
	if gpu.GetMode() != 3 {
		t.Errorf("Expected GPU mode to be 3 after 80 more cycles, got %d", gpu.GetMode())
	}
	
	// Mode 3 (Pixel Transfer) -> Mode 0 (H-Blank)
	gpu.UpdateMode(172)
	if gpu.GetMode() != 0 {
		t.Errorf("Expected GPU mode to be 0 after 172 more cycles, got %d", gpu.GetMode())
	}
	
	// Complete a full line (one full cycle of modes 2, 3, 0)
	// We've already done 80 + 80 + 172 = 332 cycles
	// A full line is 456 cycles, so we need 456 - 332 = 124 more cycles
	gpu.UpdateMode(124)
	
	// After a full line, we should be back in mode 2 (OAM Search)
	if gpu.GetMode() != 2 {
		t.Errorf("Expected GPU mode to be 2 after a full line, got %d", gpu.GetMode())
	}
	
	// Complete 143 more lines (144 visible lines total)
	for i := 0; i < 143; i++ {
		gpu.UpdateMode(456)
	}
	
	// After 144 lines, we should be in mode 1 (V-Blank)
	if gpu.GetMode() != 1 {
		t.Errorf("Expected GPU mode to be 1 after 144 lines, got %d", gpu.GetMode())
	}
	
	// V-Blank lasts for 10 lines (10 * 456 = 4560 cycles)
	// Let's do 9 lines
	for i := 0; i < 9; i++ {
		gpu.UpdateMode(456)
		if gpu.GetMode() != 1 {
			t.Errorf("Expected GPU mode to still be 1 during V-Blank, got %d", gpu.GetMode())
		}
	}
	
	// After the 10th line of V-Blank, we should go back to mode 2 (OAM Search)
	gpu.UpdateMode(456)
	if gpu.GetMode() != 2 {
		t.Errorf("Expected GPU mode to be 2 after V-Blank, got %d", gpu.GetMode())
	}
}

// TestLYRegister tests the LY register (current scanline)
func TestLYRegister(t *testing.T) {
	display := NewMockDisplay()
	gpu := NewGPU(display)
	
	// Initial LY should be 0
	if gpu.ly != 0 {
		t.Errorf("Expected initial LY to be 0, got %d", gpu.ly)
	}
	
	// Complete one line (456 cycles)
	gpu.UpdateMode(456)
	
	// LY should now be 1
	if gpu.ly != 1 {
		t.Errorf("Expected LY to be 1 after one line, got %d", gpu.ly)
	}
	
	// Complete 143 more lines (144 visible lines total)
	for i := 0; i < 143; i++ {
		gpu.UpdateMode(456)
	}
	
	// LY should now be 144 (first line of V-Blank)
	if gpu.ly != 144 {
		t.Errorf("Expected LY to be 144 after 144 lines, got %d", gpu.ly)
	}
	
	// Complete 9 more lines (9 of the 10 V-Blank lines)
	for i := 0; i < 9; i++ {
		gpu.UpdateMode(456)
	}
	
	// LY should now be 153 (last line of V-Blank)
	if gpu.ly != 153 {
		t.Errorf("Expected LY to be 153 after 153 lines, got %d", gpu.ly)
	}
	
	// Complete the last line of V-Blank
	gpu.UpdateMode(456)
	
	// LY should wrap back to 0
	if gpu.ly != 0 {
		t.Errorf("Expected LY to wrap back to 0, got %d", gpu.ly)
	}
}

// TestSTATRegister tests the STAT register (LCD status)
func TestSTATRegister(t *testing.T) {
	display := NewMockDisplay()
	gpu := NewGPU(display)
	
	// Initial STAT should have mode 0 (bits 0-1 = 00)
	if gpu.stat&0x03 != 0 {
		t.Errorf("Expected initial STAT mode bits to be 0, got %d", gpu.stat&0x03)
	}
	
	// Change to mode 2
	gpu.UpdateMode(80)
	
	// STAT should now have mode 2 (bits 0-1 = 10)
	if gpu.stat&0x03 != 2 {
		t.Errorf("Expected STAT mode bits to be 2, got %d", gpu.stat&0x03)
	}
	
	// Test LYC=LY flag (bit 2)
	// Set LYC to current LY value
	gpu.Write(0xFF45, gpu.ly) // LYC register
	
	// STAT bit 2 should be set (LYC=LY)
	if gpu.stat&0x04 != 0x04 {
		t.Errorf("Expected STAT bit 2 to be set when LYC=LY, got %02X", gpu.stat)
	}
	
	// Set LYC to a different value
	gpu.Write(0xFF45, gpu.ly+1)
	
	// STAT bit 2 should be clear (LYC≠LY)
	if gpu.stat&0x04 != 0 {
		t.Errorf("Expected STAT bit 2 to be clear when LYC≠LY, got %02X", gpu.stat)
	}
}
