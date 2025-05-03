package video

import (
	"testing"
)

// TestDisplayInterface verifies that our mock display implements the Display interface
func TestDisplayInterface(t *testing.T) {
	var _ Display = &MockDisplay{}
}

// TestMockDisplay tests the basic functionality of the mock display
func TestMockDisplay(t *testing.T) {
	display := NewMockDisplay()
	
	// Test initial state
	if display.Width() != 160 || display.Height() != 144 {
		t.Errorf("Expected dimensions 160x144, got %dx%d", display.Width(), display.Height())
	}
	
	// Test setting and getting a pixel
	display.SetPixel(10, 20, 2)
	if display.GetPixel(10, 20) != 2 {
		t.Errorf("Expected pixel value 2, got %d", display.GetPixel(10, 20))
	}
	
	// Test clearing the display
	display.Clear()
	if display.GetPixel(10, 20) != 0 {
		t.Errorf("Expected pixel value 0 after clear, got %d", display.GetPixel(10, 20))
	}
}
