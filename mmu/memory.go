package mmu

type MemoryRegion struct {
	// Start address (inclusive)
	Start uint16

	// End address (exclusive)
	End uint16
}

func (region MemoryRegion) Contains(address uint16) bool {
	return region.Start <= address && address < region.End
}

func (region MemoryRegion) IndexOf(address uint16) uint16 {
	return address - region.Start
}

func (region MemoryRegion) AddressOf(offset uint8) uint16 {
	return region.Start + uint16(offset)
}

var (
	ROMBank0     = MemoryRegion{0x0000, 0x4000}
	ROMBank1     = MemoryRegion{0x4000, 0x8000}
	CartridgeRAM = MemoryRegion{0xA000, 0xC000}
	InternalRAM  = MemoryRegion{0xC000, 0xE000}
	EchoRAM      = MemoryRegion{0xE000, 0xFE00}
	OAM          = MemoryRegion{0xFE00, 0xFEA0}
	ZeroPage     = MemoryRegion{0xFF00, 0xFFFF}
)

type Memory interface {
	Read(address uint16) uint8
	Write(address uint16, value uint8)
}

type MMU struct {
}

func (mmu *MMU) Read(address uint16) uint8 {
	return 0
}

func (mmu *MMU) Write(address uint16, value uint8) {
}
