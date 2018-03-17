package gb

func (gb *GameBoy) SP() uint16 {
	return gb.sp
}

func (gb *GameBoy) SetSP(value uint16) {
	gb.sp = value
}

func (gb *GameBoy) Push8(value uint8) {
	panic("Not implemented")
}

func (gb *GameBoy) Push16(value uint16) {
	panic("Not implemented")
}

func (gb *GameBoy) Pop8() uint8 {
	panic("Not implemented")
}

func (gb *GameBoy) Pop16() uint16 {
	lo := gb.Pop8()
	hi := gb.Pop8()
	return (uint16(hi) << 8) & uint16(lo)
}
