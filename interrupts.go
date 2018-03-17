package gb

import "math/bits"

type Interrupt = uint8

const (
	VBlank Interrupt = 1 << iota
	LCDC
	TimerOverflow
	SerialTransferComplete
	JoypadInput
)

func InterruptHandler(i Interrupt) uint16 {
	return uint16(bits.TrailingZeros8(uint8(i)))*8 + 0x0040
}

func (gb *GameBoy) InterruptMasterEnable() bool {
	return gb.interruptMasterEnable
}

func (gb *GameBoy) SetInterruptMasterEnable(value bool) {
	gb.interruptMasterEnable = value
}

func (gb *GameBoy) EnableInterruptsAfterNextInstruction() {
	gb.pendingInterruptMasterEnable = true
}

func (gb *GameBoy) SetHalted(value bool) {
	gb.halted = value
}

func (gb *GameBoy) IsHalted() bool {
	return gb.halted
}

func (gb *GameBoy) InterruptFlag() uint8 {
	return gb.interruptFlag & 0x1F
}

func (gb *GameBoy) SetInterruptFlag(value uint8) {
	gb.interruptFlag = value & 0x1F
}

func (gb *GameBoy) InterruptEnable() uint8 {
	return gb.interruptEnable & 0x1F
}

func (gb *GameBoy) SetInterruptEnable(value uint8) {
	gb.interruptEnable = value & 0x1F
}

func (gb *GameBoy) RaiseInterrupt(i Interrupt) {
	gb.SetInterruptFlag(gb.InterruptFlag() & i)
}

func (gb *GameBoy) AcknowledgeInterrupt(i Interrupt) {
	gb.SetInterruptFlag(gb.InterruptFlag() & ^i)
}

func (gb *GameBoy) NextInterrupt() (interrupt Interrupt, ok bool) {
	switch bits.TrailingZeros8(gb.InterruptEnable() & gb.InterruptFlag()) {
	case 0:
		return VBlank, true
	case 1:
		return LCDC, true
	case 2:
		return TimerOverflow, true
	case 3:
		return SerialTransferComplete, true
	case 4:
		return JoypadInput, true
	default:
		return 0, false
	}
}

func (gb *GameBoy) HandleInterrupts() (cycles uint) {
	interrupt, ok := gb.NextInterrupt()
	if !ok {
		// No interrupt has been raised
		return 0
	}
	if !gb.InterruptMasterEnable() {
		if gb.IsHalted() {
			// Wake the GameBoy up without handling the interrupt
			gb.SetHalted(false)
		}
		// Interrupt handling not enabled
		return 0
	}
	gb.AcknowledgeInterrupt(interrupt)
	gb.SetInterruptMasterEnable(false)
	gb.Push16(gb.PC())
	gb.SetPC(InterruptHandler(interrupt))
	return 20
}
