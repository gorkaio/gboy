package video

// Display defines the interface for a Game Boy display
type Display interface {
	// SetPixel sets the color of a pixel at the specified coordinates
	// Color values: 0 = White, 1 = Light Gray, 2 = Dark Gray, 3 = Black
	SetPixel(x, y int, color byte)
	
	// Refresh updates the display with the current pixel data
	Refresh() error
	
	// Clear sets all pixels to color 0 (white)
	Clear()
	
	// Width returns the width of the display in pixels
	Width() int
	
	// Height returns the height of the display in pixels
	Height() int
}

// MockDisplay is a simple in-memory implementation of the Display interface for testing
type MockDisplay struct {
	pixels [][]byte
	width  int
	height int
}

// NewMockDisplay creates a new MockDisplay with Game Boy LCD dimensions (160x144)
func NewMockDisplay() *MockDisplay {
	width, height := 160, 144
	pixels := make([][]byte, height)
	for i := range pixels {
		pixels[i] = make([]byte, width)
	}
	
	return &MockDisplay{
		pixels: pixels,
		width:  width,
		height: height,
	}
}

// SetPixel sets a pixel color at the specified coordinates
func (d *MockDisplay) SetPixel(x, y int, color byte) {
	if x >= 0 && x < d.width && y >= 0 && y < d.height {
		d.pixels[y][x] = color & 0x03 // Ensure color is in range 0-3
	}
}

// GetPixel returns the color of a pixel at the specified coordinates
// This is not part of the Display interface but useful for testing
func (d *MockDisplay) GetPixel(x, y int) byte {
	if x >= 0 && x < d.width && y >= 0 && y < d.height {
		return d.pixels[y][x]
	}
	return 0
}

// Refresh does nothing in the mock implementation
func (d *MockDisplay) Refresh() error {
	// No-op for mock display
	return nil
}

// Clear sets all pixels to white (0)
func (d *MockDisplay) Clear() {
	for y := range d.pixels {
		for x := range d.pixels[y] {
			d.pixels[y][x] = 0
		}
	}
}

// Width returns the display width
func (d *MockDisplay) Width() int {
	return d.width
}

// Height returns the display height
func (d *MockDisplay) Height() int {
	return d.height
}
