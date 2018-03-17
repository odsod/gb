package gb

const (
	flagZBit = 1 << 7
	flagNBit = 1 << 6
	flagHBit = 1 << 5
	flagCBit = 1 << 4
)

func (gb *GameBoy) F() uint8 {
	var f uint8
	if gb.flagZ {
		f |= flagZBit
	}
	if gb.flagN {
		f |= flagNBit
	}
	if gb.flagH {
		f |= flagHBit
	}
	if gb.flagC {
		f |= flagCBit
	}
	return f
}

func (gb *GameBoy) SetF(f uint8) {
	gb.flagZ = f&flagZBit > 0
	gb.flagN = f&flagNBit > 0
	gb.flagH = f&flagHBit > 0
	gb.flagC = f&flagCBit > 0
}

func (gb *GameBoy) FlagZ() bool {
	return gb.flagZ
}

func (gb *GameBoy) FlagN() bool {
	return gb.flagN
}

func (gb *GameBoy) FlagH() bool {
	return gb.flagH
}

func (gb *GameBoy) FlagC() bool {
	return gb.flagC
}

func (gb *GameBoy) SetFlagZ(value bool) {
	gb.flagZ = value
}

func (gb *GameBoy) SetFlagN(value bool) {
	gb.flagN = value
}

func (gb *GameBoy) SetFlagH(value bool) {
	gb.flagH = value
}

func (gb *GameBoy) SetFlagC(value bool) {
	gb.flagC = value
}

// Carry gives the numeric value of the C flag
func (gb *GameBoy) Carry() uint8 {
	if gb.flagC {
		return 1
	}
	return 0
}
