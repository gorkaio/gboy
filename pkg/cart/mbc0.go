package cart

func NewMBC0(data []byte) CartController {
	return &MBC0{
		memory: data,
	}
}

type MBC0 struct {
	memory []byte
}

func (r *MBC0) Read(addr uint16) byte {
	return r.memory[addr]
}

func (r *MBC0) Write(addr uint16, data byte) {}
