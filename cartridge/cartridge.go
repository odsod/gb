package cartridge

import (
	"fmt"

	"github.com/odsod/gb/mmu"
)

// ROM offsets
const (
	RST00                            = 0x0000
	RST08                            = 0x0008
	RST10                            = 0x0010
	RST18                            = 0x0018
	RST20                            = 0x0020
	RST28                            = 0x0028
	RST30                            = 0x0030
	RST38                            = 0x0038
	VBlankInterrupt                  = 0x0040
	LCDCStatusInterrupt              = 0x0048
	TimerOverflowInterrupt           = 0x0050
	SerialTransferCompletionInterupt = 0x0058
	InputPortInterrputs              = 0x0060
	CodeExecutionStart               = 0x0100
	NintendoLogoStart                = 0x0104
	NintendoLogoEnd                  = 0x0133
	GameTitleStart                   = 0x0134
	GameTitleEnd                     = 0x0142
	GameBoyColorFlag                 = 0x0143
	LicenceCodeHighNibble            = 0x0144
	LicenceCodeLowNibble             = 0x0145
	SuperGameBoyFlag                 = 0x0146
	CartridgeType                    = 0x0147
	ROMSize                          = 0x0148
	RAMSize                          = 0x0149
	DestinationCode                  = 0x014A
	LicenseeCode                     = 0x014B
	MaskROMVersionNumber             = 0x014C
	ComplementCheck                  = 0x014D
	ChecksumHighByte                 = 0x014E
	ChecksumLowByte                  = 0x014F
)

const (
	// ROMBankSize is 16KB per bank
	ROMBankSize = 0x4000

	// RAMBankSize is 8KB per bank
	RAMBankSize = 0x2000
)

// ROMBank is a bank of read-only memory
type ROMBank [ROMBankSize]byte

// RAMBank is a bank of read/write memory
type RAMBank [RAMBankSize]byte

// Cartridge models a game cartridge
type Cartridge interface {
	mmu.Memory

	// LoadROM initializes the cartridge ROM
	LoadROM(rom []byte)

	// LoadRAM initializes the cartridge RAM
	LoadRAM(ram []byte)

	// HasRAM is true if this cartridge has onboard RAM
	HasRAM() bool

	// HasBattery is true if this cartridge has an onboard battery
	HasBattery() bool
}

// FromROM creates a Cartridge for the provided ROM
func FromROM(rom []byte) Cartridge {
	var cartridge Cartridge
	switch rom[CartridgeType] {
	case 0x00:
		cartridge = &basic{}
	case 0x01:
		cartridge = &mbc1{}
	case 0x02:
		cartridge = &mbc1{ram: true}
	case 0x03:
		cartridge = &mbc1{ram: true, battery: true}
	case 0x05:
		cartridge = &mbc2{}
	case 0x06:
		cartridge = &mbc2{battery: true}
	case 0x08:
		cartridge = &basic{ram: true}
	case 0x09:
		cartridge = &basic{ram: true, battery: true}
	default:
		panic(fmt.Sprintf("Unsupported cartridge type: %#x", rom[CartridgeType]))
	}
	cartridge.LoadROM(rom)
	return cartridge
}
