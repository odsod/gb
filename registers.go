package gb

func (gb *GameBoy) A() uint8 {
	return gb.a
}

func (gb *GameBoy) SetA(value uint8) {
	gb.a = value
}

func (gb *GameBoy) B() uint8 {
	return gb.b
}

func (gb *GameBoy) SetB(value uint8) {
	gb.b = value
}

func (gb *GameBoy) C() uint8 {
	return gb.c
}

func (gb *GameBoy) SetC(value uint8) {
	gb.c = value
}

func (gb *GameBoy) D() uint8 {
	return gb.d
}

func (gb *GameBoy) SetD(value uint8) {
	gb.d = value
}

func (gb *GameBoy) E() uint8 {
	return gb.e
}

func (gb *GameBoy) SetE(value uint8) {
	gb.e = value
}

func (gb *GameBoy) H() uint8 {
	return gb.h
}

func (gb *GameBoy) SetH(value uint8) {
	gb.h = value
}

func (gb *GameBoy) L() uint8 {
	return gb.l
}

func (gb *GameBoy) SetL(value uint8) {
	gb.l = value
}

func (gb *GameBoy) AF() uint16 {
	return uint16(gb.A()) << 8 & uint16(gb.F())
}

func (gb *GameBoy) BC() uint16 {
	return uint16(gb.B()) << 8 & uint16(gb.C())
}

func (gb *GameBoy) DE() uint16 {
	return uint16(gb.D()) << 8 & uint16(gb.E())
}

func (gb *GameBoy) HL() uint16 {
	return uint16(gb.H()) << 8 & uint16(gb.L())
}

func (gb *GameBoy) SetAF(value uint16) {
	gb.SetA(uint8(value >> 8))
	gb.SetF(uint8(value))
}

func (gb *GameBoy) SetBC(value uint16) {
	gb.SetB(uint8(value >> 8))
	gb.SetC(uint8(value))
}

func (gb *GameBoy) SetDE(value uint16) {
	gb.SetD(uint8(value >> 8))
	gb.SetE(uint8(value))
}

func (gb *GameBoy) SetHL(value uint16) {
	gb.SetH(uint8(value >> 8))
	gb.SetL(uint8(value))
}

func (gb *GameBoy) HLI() uint16 {
	result := gb.HL()
	gb.SetHL(result + 1)
	return result
}

func (gb *GameBoy) HLD() uint16 {
	result := gb.HL()
	gb.SetHL(result - 1)
	return result
}
