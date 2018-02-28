package cartridge

import "github.com/odsod/gb/memory"

type mbc2 struct {
	battery bool
	ram     bool
}

func (cartridge *mbc2) LoadROM(rom []byte) {
	panic("Not implemented")
}

func (cartridge *mbc2) LoadRAM(ram []byte) {
	panic("Not implemented")
}

func (cartridge *mbc2) HasBattery() bool {
	panic("Not implemented")
}

func (cartridge *mbc2) HasRAM() bool {
	panic("Not implemented")
}

func (cartridge *mbc2) Read(address memory.Address) (byte, error) {
	panic("Not implemented")
}

func (cartridge *mbc2) Write(address memory.Address, value byte) error {
	panic("Not implemented")
}
