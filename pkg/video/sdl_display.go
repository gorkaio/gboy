package video

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

// SDLDisplay is a concrete implementation of the Display interface using SDL2
type SDLDisplay struct {
	window   *sdl.Window
	renderer *sdl.Renderer
	texture  *sdl.Texture
	pixels   []byte
	width    int
	height   int
	scale    int // Scaling factor for the display
}

// NewSDLDisplay creates a new SDLDisplay with Game Boy LCD dimensions (160x144)
// and the specified scaling factor
func NewSDLDisplay(scale int) (*SDLDisplay, error) {
	// Initialize SDL
	if err := sdl.Init(sdl.INIT_VIDEO); err != nil {
		return nil, fmt.Errorf("failed to initialize SDL: %v", err)
	}

	// Game Boy screen dimensions
	width, height := 160, 144
	
	// Create a window
	window, err := sdl.CreateWindow(
		"Game Boy Emulator",
		sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		int32(width*scale), int32(height*scale),
		sdl.WINDOW_SHOWN,
	)
	if err != nil {
		sdl.Quit()
		return nil, fmt.Errorf("failed to create window: %v", err)
	}

	// Create a renderer
	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		window.Destroy()
		sdl.Quit()
		return nil, fmt.Errorf("failed to create renderer: %v", err)
	}

	// Create a texture
	texture, err := renderer.CreateTexture(
		sdl.PIXELFORMAT_ABGR8888,
		sdl.TEXTUREACCESS_STREAMING,
		int32(width), int32(height),
	)
	if err != nil {
		renderer.Destroy()
		window.Destroy()
		sdl.Quit()
		return nil, fmt.Errorf("failed to create texture: %v", err)
	}

	// Create a pixel buffer (4 bytes per pixel: RGBA)
	pixels := make([]byte, width*height*4)

	return &SDLDisplay{
		window:   window,
		renderer: renderer,
		texture:  texture,
		pixels:   pixels,
		width:    width,
		height:   height,
		scale:    scale,
	}, nil
}

// SetPixel sets the color of a pixel at the specified coordinates
// Color values: 0 = White, 1 = Light Gray, 2 = Dark Gray, 3 = Black
func (d *SDLDisplay) SetPixel(x, y int, color byte) {
	if x >= 0 && x < d.width && y >= 0 && y < d.height {
		// Convert Game Boy color (0-3) to RGBA
		var r, g, b byte
		switch color & 0x03 {
		case 0: // White
			r, g, b = 255, 255, 255
		case 1: // Light Gray
			r, g, b = 192, 192, 192
		case 2: // Dark Gray
			r, g, b = 96, 96, 96
		case 3: // Black
			r, g, b = 0, 0, 0
		}

		// Calculate the offset in the pixel buffer
		offset := (y*d.width + x) * 4

		// Set the RGBA values
		d.pixels[offset] = r     // R
		d.pixels[offset+1] = g   // G
		d.pixels[offset+2] = b   // B
		d.pixels[offset+3] = 255 // A (fully opaque)
	}
}

// Refresh updates the display with the current pixel data
func (d *SDLDisplay) Refresh() error {
	// Update the texture with the pixel data
	if err := d.texture.Update(nil, d.pixels, d.width*4); err != nil {
		return fmt.Errorf("failed to update texture: %v", err)
	}

	// Clear the renderer
	if err := d.renderer.Clear(); err != nil {
		return fmt.Errorf("failed to clear renderer: %v", err)
	}

	// Copy the texture to the renderer, scaling it to the window size
	if err := d.renderer.Copy(d.texture, nil, nil); err != nil {
		return fmt.Errorf("failed to copy texture to renderer: %v", err)
	}

	// Present the renderer
	d.renderer.Present()

	// Process events to keep the window responsive
	for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
		switch event.(type) {
		case *sdl.QuitEvent:
			// Handle quit event
			return nil
		}
	}

	return nil
}

// Clear sets all pixels to color 0 (white)
func (d *SDLDisplay) Clear() {
	// Set all pixels to white (255, 255, 255, 255)
	for i := 0; i < len(d.pixels); i += 4 {
		d.pixels[i] = 255     // R
		d.pixels[i+1] = 255   // G
		d.pixels[i+2] = 255   // B
		d.pixels[i+3] = 255   // A
	}
}

// Width returns the width of the display in pixels
func (d *SDLDisplay) Width() int {
	return d.width
}

// Height returns the height of the display in pixels
func (d *SDLDisplay) Height() int {
	return d.height
}

// Close cleans up the SDL resources
func (d *SDLDisplay) Close() error {
	if d.texture != nil {
		d.texture.Destroy()
	}
	if d.renderer != nil {
		d.renderer.Destroy()
	}
	if d.window != nil {
		d.window.Destroy()
	}
	sdl.Quit()
	return nil
}
