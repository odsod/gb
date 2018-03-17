package gb

type GameBoy struct {
	pc                  uint16
	sp                  uint16
	a, b, c, d, e, h, l uint8

	flagZ, flagN, flagH, flagC bool

	halted                       bool
	interruptMasterEnable        bool
	pendingInterruptMasterEnable bool
	interruptEnable              uint8
	interruptFlag                uint8
}

func (gb *GameBoy) Step() (cycles uint) {
	cycles = gb.StepCPU()
	gb.StepPPU(cycles)
	return cycles
}
