package cart

type mbc0 struct {
	memory []byte
	MemoryBankController
}

// NewMBC0 creates a new memory bank controller of type 0
func NewMBC0(data []byte) MemoryBankController {
	return &mbc0{
		memory: data,
	}
}

func (r *mbc0) Read(addr uint16) byte {
	return r.memory[addr]
}

func (r *mbc0) Write(addr uint16, data byte) {}
