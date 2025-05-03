package video

import (
	"testing"
)

// TestVBlankInterrupt tests that a V-Blank interrupt is requested when entering V-Blank mode
func TestVBlankInterrupt(t *testing.T) {
	display := NewMockDisplay()
	gpu := NewGPU(display)
	
	// Create a mock interrupt controller
	mockInterrupt := &MockInterrupt{}
	gpu.SetInterruptController(mockInterrupt)
	
	// Enable the LCD
	gpu.Write(0xFF40, gpu.lcdc|0x80)
	
	// Simulate completing all visible scanlines (0-143) and entering V-Blank
	for i := 0; i < 144; i++ {
		gpu.ly = byte(i)
	}
	
	// Trigger V-Blank by updating the mode
	gpu.UpdateMode(456) // Complete a full line
	
	// Check that we're in V-Blank mode
	if gpu.GetMode() != 1 {
		t.Errorf("Expected GPU mode to be 1 (V-Blank), got %d", gpu.GetMode())
	}
	
	// Check that a V-Blank interrupt was requested
	if !mockInterrupt.vblankRequested {
		t.Errorf("Expected V-Blank interrupt to be requested")
	}
}

// TestLCDCStatusInterrupts tests that LCD Status interrupts are requested for various conditions
func TestLCDCStatusInterrupts(t *testing.T) {
	display := NewMockDisplay()
	gpu := NewGPU(display)
	
	// Create a mock interrupt controller
	mockInterrupt := &MockInterrupt{}
	gpu.SetInterruptController(mockInterrupt)
	
	// Enable the LCD
	gpu.Write(0xFF40, gpu.lcdc|0x80)
	
	// Test cases for STAT interrupts
	testCases := []struct {
		name           string
		statValue      byte // Value to write to STAT register
		updateAction   func() // Action to trigger the interrupt
		shouldTrigger  bool // Whether an interrupt should be triggered
	}{
		{
			name:      "H-Blank Interrupt",
			statValue: 0x08, // Enable H-Blank interrupt (bit 3)
			updateAction: func() {
				gpu.mode = 3 // Set mode to Pixel Transfer
				gpu.UpdateMode(172) // Complete pixel transfer, enter H-Blank
			},
			shouldTrigger: true,
		},
		{
			name:      "V-Blank Interrupt",
			statValue: 0x10, // Enable V-Blank interrupt (bit 4)
			updateAction: func() {
				gpu.ly = 143 // Last visible scanline
				gpu.UpdateMode(456) // Complete a full line, enter V-Blank
			},
			shouldTrigger: true,
		},
		{
			name:      "OAM Interrupt",
			statValue: 0x20, // Enable OAM interrupt (bit 5)
			updateAction: func() {
				gpu.mode = 0 // Set mode to H-Blank
				gpu.UpdateMode(80) // Complete H-Blank, enter OAM Search
			},
			shouldTrigger: true,
		},
		{
			name:      "LYC=LY Interrupt",
			statValue: 0x40, // Enable LYC=LY interrupt (bit 6)
			updateAction: func() {
				gpu.lyc = 50 // Set LYC to 50
				gpu.ly = 49 // Set LY to 49
				gpu.UpdateMode(456) // Complete a full line, LY becomes 50
			},
			shouldTrigger: true,
		},
	}
	
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Reset the mock interrupt
			mockInterrupt.lcdstatRequested = false
			
			// Set up the STAT register
			gpu.Write(0xFF41, tc.statValue)
			
			// Perform the action to trigger the interrupt
			tc.updateAction()
			
			// Check if the interrupt was triggered
			if tc.shouldTrigger && !mockInterrupt.lcdstatRequested {
				t.Errorf("Expected LCD STAT interrupt to be requested for %s", tc.name)
			} else if !tc.shouldTrigger && mockInterrupt.lcdstatRequested {
				t.Errorf("Did not expect LCD STAT interrupt to be requested for %s", tc.name)
			}
		})
	}
}

// MockInterrupt is a mock implementation of an interrupt controller
type MockInterrupt struct {
	vblankRequested  bool
	lcdstatRequested bool
}

// RequestVBlankInterrupt requests a V-Blank interrupt
func (m *MockInterrupt) RequestVBlankInterrupt() {
	m.vblankRequested = true
}

// RequestLCDSTATInterrupt requests an LCD STAT interrupt
func (m *MockInterrupt) RequestLCDSTATInterrupt() {
	m.lcdstatRequested = true
}
