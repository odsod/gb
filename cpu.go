package gb

func (gb *GameBoy) StepCPU() (cycles uint) {
	interruptCycles := gb.HandleInterrupts()
	if interruptCycles > 0 {
		return interruptCycles
	}
	return gb.execute(gb.Fetch8())
}

func (gb *GameBoy) execute(opcode uint8) (cycles uint) {
	switch opcode {

	case 0x0: // NOP
		return 4

	case 0x7: // RLCA
		gb.executeGeneratedCBInstruction(0x7)
		return 4

	case 0x8: // LD (a16),SP
		gb.Write16(gb.Fetch16(), gb.SP())
		return 20

	case 0xf: // RRCA
		gb.executeGeneratedCBInstruction(0xf)
		return 4

	case 0x10: // STOP 0
		panic("STOP not implemented")

	case 0x17: // RLA
		gb.executeGeneratedCBInstruction(0x17)
		return 4

	case 0x1f: // RRA
		gb.executeGeneratedCBInstruction(0x1f)
		return 4

	case 0x27: // DAA
		// As seen on http://forums.nesdev.com/viewtopic.php?f=20&t=15944
		a := gb.A()
		c := false
		if !gb.FlagN() {
			// After an addition:
			// Adjust if (half-)carry occurred or if result is out of bounds
			if gb.FlagC() || a > 0x99 {
				a += 0x60
				c = true
			}
			if gb.FlagH() || (a&0xF > 0x9) {
				a += 0x6
			}
		} else {
			// After a subtraction:
			// Only adjust if (half-)carry occurred
			if gb.FlagC() {
				a -= 0x60
			}
			if gb.FlagH() {
				a -= 0x6
			}
		}
		gb.SetFlagZ(a == 0)
		gb.SetFlagH(false)
		gb.SetFlagC(c)
		gb.SetA(a)
		return 4

	case 0x2f: // CPL
		gb.SetA(^gb.A())
		gb.SetFlagN(true)
		gb.SetFlagH(true)
		return 4

	case 0x37: // SCF
		// Set carry flag.
		gb.SetFlagN(false)
		gb.SetFlagH(false)
		gb.SetFlagC(true)
		return 4

	case 0x3f: // CCF
		// Complement carry flag.
		gb.SetFlagN(false)
		gb.SetFlagH(false)
		gb.SetFlagC(!gb.FlagC())
		return 4

	case 0x76: // HALT
		// TODO: Not implemented
		return 4

	case 0xcb: // PREFIX CB
		return 4 + gb.executeGeneratedCBInstruction(gb.Fetch8())

	case 0xd9: // RETI
		gb.EnableInterruptsAfterNextInstruction()
		return gb.executeGeneratedInstruction(0xc9) // RET

	case 0xe8: // ADD SP,r8
		operand1 := int32(gb.SP())
		operand2 := int32(int8(gb.Fetch8()))
		result := uint16(operand1 + operand2)
		gb.SetSP(result)
		gb.SetFlagZ(false)
		gb.SetFlagN(false)
		gb.SetFlagH((operand1&0xF)+(operand2&0xF) > 0xF)
		gb.SetFlagC((operand1&0xFF)+(operand2&0xFF) > 0xFF)
		return 16

	case 0xf3: // DI
		gb.SetInterruptMasterEnable(false)
		return 4

	case 0xf8: // LD HL,SP+r8
		operand1 := int32(gb.SP())
		operand2 := int32(int8(gb.Fetch8()))
		result := uint16(operand1 + operand2)
		gb.SetHL(result)
		gb.SetFlagZ(false)
		gb.SetFlagN(false)
		gb.SetFlagH((operand1&0xF)+(operand2&0xF) > 0xF)
		gb.SetFlagC((operand1&0xFF)+(operand2&0xFF) > 0xFF)
		return 12

	case 0xfb: // EI
		gb.EnableInterruptsAfterNextInstruction()
		return 4

	default:
		return gb.executeGeneratedInstruction(opcode)
	}
}
