package video

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

// SDLDisplay is a concrete implementation of the Display interface using SDL2
type SDLDisplay struct {
	window   *sdl.Window
	renderer *sdl.Renderer
	surface  *sdl.Surface
	width    int
	height   int
	scale    int // Scaling factor for the display
	pixels   []uint32 // Pixel buffer (one uint32 per pixel in RGBA format)
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

	// Create a surface
	surface, err := sdl.CreateRGBSurface(0, int32(width), int32(height), 32, 0xFF000000, 0x00FF0000, 0x0000FF00, 0x000000FF)
	if err != nil {
		renderer.Destroy()
		window.Destroy()
		sdl.Quit()
		return nil, fmt.Errorf("failed to create surface: %v", err)
	}

	// Create a pixel buffer
	pixels := make([]uint32, width*height)

	return &SDLDisplay{
		window:   window,
		renderer: renderer,
		surface:  surface,
		width:    width,
		height:   height,
		scale:    scale,
		pixels:   pixels,
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
		offset := y*d.width + x

		// Set the RGBA value (0xRRGGBBAA)
		d.pixels[offset] = uint32(r)<<24 | uint32(g)<<16 | uint32(b)<<8 | 0xFF
	}
}

// Refresh updates the display with the current pixel data
func (d *SDLDisplay) Refresh() error {
	// Lock the surface
	err := d.surface.Lock()
	if err != nil {
		return fmt.Errorf("failed to lock surface: %v", err)
	}

	// Get the surface pixels
	pixels := d.surface.Pixels()

	// Copy our pixel data to the surface
	for i := 0; i < len(d.pixels); i++ {
		// Convert from RGBA to ABGR (SDL's format)
		pixel := d.pixels[i]
		r := byte(pixel >> 24)
		g := byte(pixel >> 16)
		b := byte(pixel >> 8)
		a := byte(pixel)
		
		// Write to surface pixels (4 bytes per pixel)
		offset := i * 4
		pixels[offset] = b   // B
		pixels[offset+1] = g // G
		pixels[offset+2] = r // R
		pixels[offset+3] = a // A
	}

	// Unlock the surface
	d.surface.Unlock()

	// Create a texture from the surface
	texture, err := d.renderer.CreateTextureFromSurface(d.surface)
	if err != nil {
		return fmt.Errorf("failed to create texture from surface: %v", err)
	}
	defer texture.Destroy()

	// Clear the renderer
	if err := d.renderer.Clear(); err != nil {
		return fmt.Errorf("failed to clear renderer: %v", err)
	}

	// Copy the texture to the renderer, scaling it to the window size
	if err := d.renderer.Copy(texture, nil, nil); err != nil {
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
	// Set all pixels to white (0xFFFFFFFF)
	for i := range d.pixels {
		d.pixels[i] = 0xFFFFFFFF
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
	if d.surface != nil {
		d.surface.Free()
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
