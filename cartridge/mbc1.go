package cartridge

import "github.com/odsod/gb/memory"

type mbc1 struct {
	battery bool
	ram     bool
}

func (cartridge *mbc1) LoadROM(rom []byte) {
	panic("Not implemented")
}

func (cartridge *mbc1) LoadRAM(ram []byte) {
	panic("Not implemented")
}

func (cartridge *mbc1) HasBattery() bool {
	panic("Not implemented")
}

func (cartridge *mbc1) HasRAM() bool {
	panic("Not implemented")
}

func (cartridge *mbc1) Read(address memory.Address) (byte, error) {
	panic("Not implemented")
}

func (cartridge *mbc1) Write(address memory.Address, value byte) error {
	panic("Not implemented")
}
