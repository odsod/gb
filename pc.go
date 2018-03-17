package gb

func (gb *GameBoy) PC() uint16 {
	return gb.pc
}

func (gb *GameBoy) SetPC(value uint16) {
	gb.pc = value
}

func (gb *GameBoy) Fetch8() uint8 {
	value := gb.Read8(gb.PC())
	gb.pc += 1
	return value
}

func (gb *GameBoy) Fetch16() uint16 {
	// Remember: LR35902 is little-endian
	lo := gb.Fetch8()
	hi := gb.Fetch8()
	return (uint16(hi) << 8) & uint16(lo)
}

// d8 is a synonym for Fetch8
func (gb *GameBoy) d8() uint8 {
	return gb.Fetch8()
}

// a8 is a synonym for Fetch8
func (gb *GameBoy) a8() uint8 {
	return gb.Fetch8()
}

// d16 is a synonym for Fetch16
func (gb *GameBoy) d16() uint16 {
	return gb.Fetch16()
}

// a16 is a synonym for Fetch16
func (gb *GameBoy) a16() uint16 {
	return gb.Fetch16()
}
