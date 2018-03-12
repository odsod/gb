package cpu

import "fmt"

func (cpu *CPU) execute(opcode uint8) (cycles uint) {
	switch opcode {

	case 0x0: // NOP
		return 4

	case 0x7: // RLCA
		cpu.executeGeneratedCBInstruction(0x7)
		return 4

	case 0x8: // LD (a16),SP
		cpu.Write16(cpu.Fetch16(), cpu.SP())
		return 20

	case 0xf: // RRCA
		cpu.executeGeneratedCBInstruction(0xf)
		return 4

	case 0x10: // STOP 0
		operand := cpu.Fetch8()
		if operand != 0x00 {
			panic(fmt.Sprintf("Expected 0x00 after STOP (0x10) but was %#x", operand))
		}
		// TODO: Not implemented
		return 4

	case 0x17: // RLA
		cpu.executeGeneratedCBInstruction(0x17)
		return 4

	case 0x1f: // RRA
		cpu.executeGeneratedCBInstruction(0x1f)
		return 4

	case 0x27: // DAA
		// As seen on http://forums.nesdev.com/viewtopic.php?f=20&t=15944
		a := cpu.A()
		c := false
		if !cpu.FlagN() {
			// After an addition:
			// Adjust if (half-)carry occurred or if result is out of bounds
			if cpu.FlagC() || a > 0x99 {
				a += 0x60
				c = true
			}
			if cpu.FlagH() || (a&0xF > 0x9) {
				a += 0x6
			}
		} else {
			// After a subtraction:
			// Only adjust if (half-)carry occurred
			if cpu.FlagC() {
				a -= 0x60
			}
			if cpu.FlagH() {
				a -= 0x6
			}
		}
		cpu.SetFlagZ(a == 0)
		cpu.SetFlagH(false)
		cpu.SetFlagC(c)
		cpu.SetA(a)
		return 4

	case 0x2f: // CPL
		cpu.SetA(^cpu.A())
		cpu.SetFlagN(true)
		cpu.SetFlagH(true)
		return 4

	case 0x37: // SCF
		// Set carry flag.
		cpu.SetFlagN(false)
		cpu.SetFlagH(false)
		cpu.SetFlagC(true)
		return 4

	case 0x3f: // CCF
		// Complement carry flag.
		cpu.SetFlagN(false)
		cpu.SetFlagH(false)
		cpu.SetFlagC(!cpu.FlagC())
		return 4

	case 0x76: // HALT
		// TODO: Not implemented
		return 4

	case 0xcb: // PREFIX CB
		return 4 + cpu.executeGeneratedCBInstruction(cpu.Fetch8())

	case 0xd9: // RETI
		cpu.EnableInterruptsAfterNextInstruction()
		cpu.SetPC(cpu.Pop16())
		return 16

	case 0xe8: // ADD SP,r8
		operand1 := int32(cpu.SP())
		operand2 := int32(int8(cpu.Fetch8()))
		result := uint16(operand1 + operand2)
		cpu.SetSP(result)
		cpu.SetFlagZ(false)
		cpu.SetFlagN(false)
		cpu.SetFlagH((operand1&0xF)+(operand2&0xF) > 0xF)
		cpu.SetFlagC((operand1&0xFF)+(operand2&0xFF) > 0xFF)
		return 16

	case 0xf3: // DI
		cpu.DisableInterruptsAfterNextInstruction()
		return 4

	case 0xf8: // LD HL,SP+r8
		operand1 := int32(cpu.SP())
		operand2 := int32(int8(cpu.Fetch8()))
		result := uint16(operand1 + operand2)
		cpu.SetHL(result)
		cpu.SetFlagZ(false)
		cpu.SetFlagN(false)
		cpu.SetFlagH((operand1&0xF)+(operand2&0xF) > 0xF)
		cpu.SetFlagC((operand1&0xFF)+(operand2&0xFF) > 0xFF)
		return 12

	case 0xfb: // EI
		cpu.EnableInterruptsAfterNextInstruction()
		return 4

	default:
		return cpu.executeGeneratedInstruction(opcode)
	}
}
