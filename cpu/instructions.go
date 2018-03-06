package cpu

import "fmt"

func (cpu *CPU) execute(opcode uint8) (cycles uint) {
	switch opcode {

	case 0x0: // NOP
		return 4

	case 0x7: // RLCA
		cpu.executeGeneratedCBInstruction(0x7)
		return 4

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

	case 0xf3: // DI
		cpu.DisableInterruptsAfterNextInstruction()
		return 4

	case 0xfb: // EI
		cpu.EnableInterruptsAfterNextInstruction()
		return 4

	default:
		return cpu.executeGeneratedInstruction(opcode)
	}
}
