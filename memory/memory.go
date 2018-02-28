package memory

import "fmt"

type Address uint16

type Region struct {
	StartInclusive Address
	EndExclusive   Address
}

func (a Address) In(r Region) bool {
	return r.StartInclusive <= a && a < r.EndExclusive
}

func (a Address) IndexIn(r Region) uint {
	return uint(a) - uint(r.StartInclusive)
}

var (
	ROMBank0     = Region{0x0000, 0x4000}
	ROMBank1     = Region{0x4000, 0x8000}
	CartridgeRAM = Region{0xA000, 0xC000}
	InternalRAM  = Region{0xC000, 0xE000}
	EchoRAM      = Region{0xE000, 0xFE00}
	OAM          = Region{0xFE00, 0xFEA0}
)

type Mapper interface {
	Read(Address) (byte, error)
	Write(Address, byte) error
}

type BadAddressError struct {
	Address Address
}

func (e BadAddressError) Error() string {
	return fmt.Sprintf("Bad address: %#x", e.Address)
}
