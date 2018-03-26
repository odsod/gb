package gb

func (gb *GameBoy) StepTimer(cycles uint) {
	gb.divCycles += uint16(cycles)
}

func (gb *GameBoy) DIV() uint8 {
	return uint8(gb.divCycles >> 8)
}
