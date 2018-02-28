package cartridge

import (
	"errors"
	"fmt"

	"github.com/odsod/gb/memory"
)

// basic cartridges have up to 2 ROM banks and a single RAM bank
type basic struct {
	romBank0 ROMBank
	romBank1 ROMBank
	ramBank  RAMBank
	battery  bool
	ram      bool
}

func (cartridge *basic) LoadROM(rom []byte) {
	if len(rom) > 2*ROMBankSize {
		panic(fmt.Sprintf("Can't load ROM of size %dB on to a basic cartridge", len(rom)))
	}
	if len(rom) > ROMBankSize {
		copy(cartridge.romBank1[:], rom[ROMBankSize:])
	}
	copy(cartridge.romBank0[:], rom[0:])
}

func (cartridge *basic) LoadRAM(ram []byte) {
	if len(ram) > len(cartridge.ramBank) {
		panic(fmt.Sprintf("Can't load RAM of size %dB into a single bank", len(ram)))
	}
	copy(cartridge.ramBank[:], ram)
}

func (cartridge *basic) HasBattery() bool {
	return cartridge.battery
}

func (cartridge *basic) HasRAM() bool {
	return cartridge.ram
}

func (cartridge *basic) Read(address memory.Address) (byte, error) {
	switch {
	case address.In(memory.ROMBank0):
		return cartridge.romBank0[address.IndexIn(memory.ROMBank0)], nil
	case address.In(memory.ROMBank1):
		return cartridge.romBank1[address.IndexIn(memory.ROMBank1)], nil
	case address.In(memory.CartridgeRAM):
		return cartridge.ramBank[address.IndexIn(memory.CartridgeRAM)], nil
	default:
		return 0, memory.BadAddressError{Address: address}
	}
}

func (cartridge *basic) Write(address memory.Address, value byte) error {
	if !address.In(memory.CartridgeRAM) {
		return memory.BadAddressError{Address: address}
	}
	if !cartridge.ram {
		return errors.New("Tried to write to non-existent cartridge RAM")
	}
	cartridge.ramBank[address.IndexIn(memory.CartridgeRAM)] = value
	return nil
}
