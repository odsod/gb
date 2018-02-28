package cpu

type Registers struct {
	// A, B, C, D, E, H and L are 8-bit general purpose registers
	A, B, C, D, E, H, L byte

	// F is the flags register
	F Flags
}

func (regs Registers) AF() uint16 {
	return uint16(regs.A) << 8 & uint16(regs.F.Read())
}

func (regs Registers) BC() uint16 {
	return uint16(regs.B) << 8 & uint16(regs.C)
}

func (regs Registers) DE() uint16 {
	return uint16(regs.D) << 8 & uint16(regs.E)
}

func (regs Registers) HL() uint16 {
	return uint16(regs.H) << 8 & uint16(regs.L)
}
