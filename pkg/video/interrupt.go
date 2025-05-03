package video

// InterruptController defines the interface for a component that can handle interrupt requests
type InterruptController interface {
	// RequestVBlankInterrupt requests a V-Blank interrupt (bit 0 of IF register)
	RequestVBlankInterrupt()
	
	// RequestLCDSTATInterrupt requests an LCD STAT interrupt (bit 1 of IF register)
	RequestLCDSTATInterrupt()
}

// SetInterruptController sets the interrupt controller for the GPU
func (gpu *GPU) SetInterruptController(controller InterruptController) {
	gpu.interruptController = controller
}

// checkAndRequestInterrupts checks if any interrupts should be requested based on the current GPU state
func (gpu *GPU) checkAndRequestInterrupts(oldMode byte) {
	// If no interrupt controller is set, return
	if gpu.interruptController == nil {
		return
	}
	
	// Check for V-Blank interrupt
	// V-Blank interrupt is requested when entering V-Blank mode (mode 1)
	if gpu.mode == 1 && oldMode != 1 {
		gpu.interruptController.RequestVBlankInterrupt()
	}
	
	// Check for LCD STAT interrupts
	// These are controlled by bits 3-6 of the STAT register
	
	// H-Blank interrupt (STAT bit 3)
	if gpu.mode == 0 && oldMode != 0 && gpu.stat&0x08 != 0 {
		gpu.interruptController.RequestLCDSTATInterrupt()
	}
	
	// V-Blank interrupt (STAT bit 4)
	if gpu.mode == 1 && oldMode != 1 && gpu.stat&0x10 != 0 {
		gpu.interruptController.RequestLCDSTATInterrupt()
	}
	
	// OAM interrupt (STAT bit 5)
	if gpu.mode == 2 && oldMode != 2 && gpu.stat&0x20 != 0 {
		gpu.interruptController.RequestLCDSTATInterrupt()
	}
	
	// LYC=LY interrupt (STAT bit 6)
	// This is triggered when LY becomes equal to LYC and the LYC=LY flag (STAT bit 2) is set
	if gpu.stat&0x04 != 0 && gpu.stat&0x40 != 0 {
		gpu.interruptController.RequestLCDSTATInterrupt()
	}
}
