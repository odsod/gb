// This file has been generated automatically. Do not edit manually.
package gb

import "fmt"

func (gb *GameBoy) executeGeneratedInstruction(opcode uint8) (cycles uint) {
	switch opcode {

	case 0x0: // Unhandled: NOP
		panic("Unhandled instruction: NOP (0x0)")

	case 0x1: // LD BC,d16
		value := gb.d16()
		gb.SetBC(value)
		return 12

	case 0x2: // LD (BC),A
		value := gb.A()
		gb.Write8(gb.BC(), value)
		return 8

	case 0x3: // INC BC
		operand1 := gb.BC()
		const operand2 = 1
		result := operand1 + operand2
		gb.SetBC(result)
		return 8

	case 0x4: // INC B
		operand1 := gb.B()
		const operand2 = 1
		result := operand1 + operand2
		gb.SetB(result)
		gb.SetFlagZ(result == 0)
		gb.SetFlagN(false)
		gb.SetFlagH((operand1&0x0F)+(operand2&0x0F) > 0x0F)
		return 4

	case 0x5: // DEC B
		operand1 := gb.B()
		const operand2 = 1
		result := operand1 - operand2
		gb.SetB(result)
		gb.SetFlagZ(result == 0)
		gb.SetFlagN(true)
		gb.SetFlagH(int16(operand1&0x0F)-int16(operand2&0x0F) < 0)
		return 4

	case 0x6: // LD B,d8
		value := gb.d8()
		gb.SetB(value)
		return 8

	case 0x7: // Unhandled: RLCA
		panic("Unhandled instruction: RLCA (0x7)")

	case 0x8: // Unhandled: LD (a16),SP
		panic("Unhandled instruction: LD (a16),SP (0x8)")

	case 0x9: // ADD HL,BC
		operand1 := gb.HL()
		operand2 := gb.BC()
		result := operand1 + operand2
		gb.SetHL(result)
		gb.SetFlagN(false)
		gb.SetFlagH(uint32(operand1&0xFFF)+uint32(operand2&0xFFF) > 0xFFF)
		gb.SetFlagC(uint32(operand1)+uint32(operand2) > 0xFFFF)
		return 8

	case 0xa: // LD A,(BC)
		value := gb.Read8(gb.BC())
		gb.SetA(value)
		return 8

	case 0xb: // DEC BC
		operand1 := gb.BC()
		const operand2 = 1
		result := operand1 - operand2
		gb.SetBC(result)
		return 8

	case 0xc: // INC C
		operand1 := gb.C()
		const operand2 = 1
		result := operand1 + operand2
		gb.SetC(result)
		gb.SetFlagZ(result == 0)
		gb.SetFlagN(false)
		gb.SetFlagH((operand1&0x0F)+(operand2&0x0F) > 0x0F)
		return 4

	case 0xd: // DEC C
		operand1 := gb.C()
		const operand2 = 1
		result := operand1 - operand2
		gb.SetC(result)
		gb.SetFlagZ(result == 0)
		gb.SetFlagN(true)
		gb.SetFlagH(int16(operand1&0x0F)-int16(operand2&0x0F) < 0)
		return 4

	case 0xe: // LD C,d8
		value := gb.d8()
		gb.SetC(value)
		return 8

	case 0xf: // Unhandled: RRCA
		panic("Unhandled instruction: RRCA (0xf)")

	case 0x10: // Unhandled: STOP 0
		panic("Unhandled instruction: STOP 0 (0x10)")

	case 0x11: // LD DE,d16
		value := gb.d16()
		gb.SetDE(value)
		return 12

	case 0x12: // LD (DE),A
		value := gb.A()
		gb.Write8(gb.DE(), value)
		return 8

	case 0x13: // INC DE
		operand1 := gb.DE()
		const operand2 = 1
		result := operand1 + operand2
		gb.SetDE(result)
		return 8

	case 0x14: // INC D
		operand1 := gb.D()
		const operand2 = 1
		result := operand1 + operand2
		gb.SetD(result)
		gb.SetFlagZ(result == 0)
		gb.SetFlagN(false)
		gb.SetFlagH((operand1&0x0F)+(operand2&0x0F) > 0x0F)
		return 4

	case 0x15: // DEC D
		operand1 := gb.D()
		const operand2 = 1
		result := operand1 - operand2
		gb.SetD(result)
		gb.SetFlagZ(result == 0)
		gb.SetFlagN(true)
		gb.SetFlagH(int16(operand1&0x0F)-int16(operand2&0x0F) < 0)
		return 4

	case 0x16: // LD D,d8
		value := gb.d8()
		gb.SetD(value)
		return 8

	case 0x17: // Unhandled: RLA
		panic("Unhandled instruction: RLA (0x17)")

	case 0x18: // JR r8
		address := uint16(int32(gb.PC()) + int32(gb.Fetch8()))
		gb.SetPC(address)
		return 12

	case 0x19: // ADD HL,DE
		operand1 := gb.HL()
		operand2 := gb.DE()
		result := operand1 + operand2
		gb.SetHL(result)
		gb.SetFlagN(false)
		gb.SetFlagH(uint32(operand1&0xFFF)+uint32(operand2&0xFFF) > 0xFFF)
		gb.SetFlagC(uint32(operand1)+uint32(operand2) > 0xFFFF)
		return 8

	case 0x1a: // LD A,(DE)
		value := gb.Read8(gb.DE())
		gb.SetA(value)
		return 8

	case 0x1b: // DEC DE
		operand1 := gb.DE()
		const operand2 = 1
		result := operand1 - operand2
		gb.SetDE(result)
		return 8

	case 0x1c: // INC E
		operand1 := gb.E()
		const operand2 = 1
		result := operand1 + operand2
		gb.SetE(result)
		gb.SetFlagZ(result == 0)
		gb.SetFlagN(false)
		gb.SetFlagH((operand1&0x0F)+(operand2&0x0F) > 0x0F)
		return 4

	case 0x1d: // DEC E
		operand1 := gb.E()
		const operand2 = 1
		result := operand1 - operand2
		gb.SetE(result)
		gb.SetFlagZ(result == 0)
		gb.SetFlagN(true)
		gb.SetFlagH(int16(operand1&0x0F)-int16(operand2&0x0F) < 0)
		return 4

	case 0x1e: // LD E,d8
		value := gb.d8()
		gb.SetE(value)
		return 8

	case 0x1f: // Unhandled: RRA
		panic("Unhandled instruction: RRA (0x1f)")

	case 0x20: // JR NZ,r8
		if !gb.FlagZ() {
			address := uint16(int32(gb.PC()) + int32(gb.Fetch8()))
			gb.SetPC(address)
			return 12
		}
		return 8

	case 0x21: // LD HL,d16
		value := gb.d16()
		gb.SetHL(value)
		return 12

	case 0x22: // LD (HL+),A
		value := gb.A()
		gb.Write8(gb.HLI(), value)
		return 8

	case 0x23: // INC HL
		operand1 := gb.HL()
		const operand2 = 1
		result := operand1 + operand2
		gb.SetHL(result)
		return 8

	case 0x24: // INC H
		operand1 := gb.H()
		const operand2 = 1
		result := operand1 + operand2
		gb.SetH(result)
		gb.SetFlagZ(result == 0)
		gb.SetFlagN(false)
		gb.SetFlagH((operand1&0x0F)+(operand2&0x0F) > 0x0F)
		return 4

	case 0x25: // DEC H
		operand1 := gb.H()
		const operand2 = 1
		result := operand1 - operand2
		gb.SetH(result)
		gb.SetFlagZ(result == 0)
		gb.SetFlagN(true)
		gb.SetFlagH(int16(operand1&0x0F)-int16(operand2&0x0F) < 0)
		return 4

	case 0x26: // LD H,d8
		value := gb.d8()
		gb.SetH(value)
		return 8

	case 0x27: // Unhandled: DAA
		panic("Unhandled instruction: DAA (0x27)")

	case 0x28: // JR Z,r8
		if gb.FlagZ() {
			address := uint16(int32(gb.PC()) + int32(gb.Fetch8()))
			gb.SetPC(address)
			return 12
		}
		return 8

	case 0x29: // ADD HL,HL
		operand1 := gb.HL()
		operand2 := gb.HL()
		result := operand1 + operand2
		gb.SetHL(result)
		gb.SetFlagN(false)
		gb.SetFlagH(uint32(operand1&0xFFF)+uint32(operand2&0xFFF) > 0xFFF)
		gb.SetFlagC(uint32(operand1)+uint32(operand2) > 0xFFFF)
		return 8

	case 0x2a: // LD A,(HL+)
		value := gb.Read8(gb.HLI())
		gb.SetA(value)
		return 8

	case 0x2b: // DEC HL
		operand1 := gb.HL()
		const operand2 = 1
		result := operand1 - operand2
		gb.SetHL(result)
		return 8

	case 0x2c: // INC L
		operand1 := gb.L()
		const operand2 = 1
		result := operand1 + operand2
		gb.SetL(result)
		gb.SetFlagZ(result == 0)
		gb.SetFlagN(false)
		gb.SetFlagH((operand1&0x0F)+(operand2&0x0F) > 0x0F)
		return 4

	case 0x2d: // DEC L
		operand1 := gb.L()
		const operand2 = 1
		result := operand1 - operand2
		gb.SetL(result)
		gb.SetFlagZ(result == 0)
		gb.SetFlagN(true)
		gb.SetFlagH(int16(operand1&0x0F)-int16(operand2&0x0F) < 0)
		return 4

	case 0x2e: // LD L,d8
		value := gb.d8()
		gb.SetL(value)
		return 8

	case 0x2f: // Unhandled: CPL
		panic("Unhandled instruction: CPL (0x2f)")

	case 0x30: // JR NC,r8
		if !gb.FlagC() {
			address := uint16(int32(gb.PC()) + int32(gb.Fetch8()))
			gb.SetPC(address)
			return 12
		}
		return 8

	case 0x31: // LD SP,d16
		value := gb.d16()
		gb.SetSP(value)
		return 12

	case 0x32: // LD (HL-),A
		value := gb.A()
		gb.Write8(gb.HLD(), value)
		return 8

	case 0x33: // INC SP
		operand1 := gb.SP()
		const operand2 = 1
		result := operand1 + operand2
		gb.SetSP(result)
		return 8

	case 0x34: // INC (HL)
		operand1 := gb.Read8(gb.HL())
		const operand2 = 1
		result := operand1 + operand2
		gb.Write8(gb.HL(), result)
		gb.SetFlagZ(result == 0)
		gb.SetFlagN(false)
		gb.SetFlagH((operand1&0x0F)+(operand2&0x0F) > 0x0F)
		return 12

	case 0x35: // DEC (HL)
		operand1 := gb.Read8(gb.HL())
		const operand2 = 1
		result := operand1 - operand2
		gb.Write8(gb.HL(), result)
		gb.SetFlagZ(result == 0)
		gb.SetFlagN(true)
		gb.SetFlagH(int16(operand1&0x0F)-int16(operand2&0x0F) < 0)
		return 12

	case 0x36: // LD (HL),d8
		value := gb.d8()
		gb.Write8(gb.HL(), value)
		return 12

	case 0x37: // Unhandled: SCF
		panic("Unhandled instruction: SCF (0x37)")

	case 0x38: // JR C,r8
		if gb.FlagC() {
			address := uint16(int32(gb.PC()) + int32(gb.Fetch8()))
			gb.SetPC(address)
			return 12
		}
		return 8

	case 0x39: // ADD HL,SP
		operand1 := gb.HL()
		operand2 := gb.SP()
		result := operand1 + operand2
		gb.SetHL(result)
		gb.SetFlagN(false)
		gb.SetFlagH(uint32(operand1&0xFFF)+uint32(operand2&0xFFF) > 0xFFF)
		gb.SetFlagC(uint32(operand1)+uint32(operand2) > 0xFFFF)
		return 8

	case 0x3a: // LD A,(HL-)
		value := gb.Read8(gb.HLD())
		gb.SetA(value)
		return 8

	case 0x3b: // DEC SP
		operand1 := gb.SP()
		const operand2 = 1
		result := operand1 - operand2
		gb.SetSP(result)
		return 8

	case 0x3c: // INC A
		operand1 := gb.A()
		const operand2 = 1
		result := operand1 + operand2
		gb.SetA(result)
		gb.SetFlagZ(result == 0)
		gb.SetFlagN(false)
		gb.SetFlagH((operand1&0x0F)+(operand2&0x0F) > 0x0F)
		return 4

	case 0x3d: // DEC A
		operand1 := gb.A()
		const operand2 = 1
		result := operand1 - operand2
		gb.SetA(result)
		gb.SetFlagZ(result == 0)
		gb.SetFlagN(true)
		gb.SetFlagH(int16(operand1&0x0F)-int16(operand2&0x0F) < 0)
		return 4

	case 0x3e: // LD A,d8
		value := gb.d8()
		gb.SetA(value)
		return 8

	case 0x3f: // Unhandled: CCF
		panic("Unhandled instruction: CCF (0x3f)")

	case 0x40: // LD B,B
		value := gb.B()
		gb.SetB(value)
		return 4

	case 0x41: // LD B,C
		value := gb.C()
		gb.SetB(value)
		return 4

	case 0x42: // LD B,D
		value := gb.D()
		gb.SetB(value)
		return 4

	case 0x43: // LD B,E
		value := gb.E()
		gb.SetB(value)
		return 4

	case 0x44: // LD B,H
		value := gb.H()
		gb.SetB(value)
		return 4

	case 0x45: // LD B,L
		value := gb.L()
		gb.SetB(value)
		return 4

	case 0x46: // LD B,(HL)
		value := gb.Read8(gb.HL())
		gb.SetB(value)
		return 8

	case 0x47: // LD B,A
		value := gb.A()
		gb.SetB(value)
		return 4

	case 0x48: // LD C,B
		value := gb.B()
		gb.SetC(value)
		return 4

	case 0x49: // LD C,C
		value := gb.C()
		gb.SetC(value)
		return 4

	case 0x4a: // LD C,D
		value := gb.D()
		gb.SetC(value)
		return 4

	case 0x4b: // LD C,E
		value := gb.E()
		gb.SetC(value)
		return 4

	case 0x4c: // LD C,H
		value := gb.H()
		gb.SetC(value)
		return 4

	case 0x4d: // LD C,L
		value := gb.L()
		gb.SetC(value)
		return 4

	case 0x4e: // LD C,(HL)
		value := gb.Read8(gb.HL())
		gb.SetC(value)
		return 8

	case 0x4f: // LD C,A
		value := gb.A()
		gb.SetC(value)
		return 4

	case 0x50: // LD D,B
		value := gb.B()
		gb.SetD(value)
		return 4

	case 0x51: // LD D,C
		value := gb.C()
		gb.SetD(value)
		return 4

	case 0x52: // LD D,D
		value := gb.D()
		gb.SetD(value)
		return 4

	case 0x53: // LD D,E
		value := gb.E()
		gb.SetD(value)
		return 4

	case 0x54: // LD D,H
		value := gb.H()
		gb.SetD(value)
		return 4

	case 0x55: // LD D,L
		value := gb.L()
		gb.SetD(value)
		return 4

	case 0x56: // LD D,(HL)
		value := gb.Read8(gb.HL())
		gb.SetD(value)
		return 8

	case 0x57: // LD D,A
		value := gb.A()
		gb.SetD(value)
		return 4

	case 0x58: // LD E,B
		value := gb.B()
		gb.SetE(value)
		return 4

	case 0x59: // LD E,C
		value := gb.C()
		gb.SetE(value)
		return 4

	case 0x5a: // LD E,D
		value := gb.D()
		gb.SetE(value)
		return 4

	case 0x5b: // LD E,E
		value := gb.E()
		gb.SetE(value)
		return 4

	case 0x5c: // LD E,H
		value := gb.H()
		gb.SetE(value)
		return 4

	case 0x5d: // LD E,L
		value := gb.L()
		gb.SetE(value)
		return 4

	case 0x5e: // LD E,(HL)
		value := gb.Read8(gb.HL())
		gb.SetE(value)
		return 8

	case 0x5f: // LD E,A
		value := gb.A()
		gb.SetE(value)
		return 4

	case 0x60: // LD H,B
		value := gb.B()
		gb.SetH(value)
		return 4

	case 0x61: // LD H,C
		value := gb.C()
		gb.SetH(value)
		return 4

	case 0x62: // LD H,D
		value := gb.D()
		gb.SetH(value)
		return 4

	case 0x63: // LD H,E
		value := gb.E()
		gb.SetH(value)
		return 4

	case 0x64: // LD H,H
		value := gb.H()
		gb.SetH(value)
		return 4

	case 0x65: // LD H,L
		value := gb.L()
		gb.SetH(value)
		return 4

	case 0x66: // LD H,(HL)
		value := gb.Read8(gb.HL())
		gb.SetH(value)
		return 8

	case 0x67: // LD H,A
		value := gb.A()
		gb.SetH(value)
		return 4

	case 0x68: // LD L,B
		value := gb.B()
		gb.SetL(value)
		return 4

	case 0x69: // LD L,C
		value := gb.C()
		gb.SetL(value)
		return 4

	case 0x6a: // LD L,D
		value := gb.D()
		gb.SetL(value)
		return 4

	case 0x6b: // LD L,E
		value := gb.E()
		gb.SetL(value)
		return 4

	case 0x6c: // LD L,H
		value := gb.H()
		gb.SetL(value)
		return 4

	case 0x6d: // LD L,L
		value := gb.L()
		gb.SetL(value)
		return 4

	case 0x6e: // LD L,(HL)
		value := gb.Read8(gb.HL())
		gb.SetL(value)
		return 8

	case 0x6f: // LD L,A
		value := gb.A()
		gb.SetL(value)
		return 4

	case 0x70: // LD (HL),B
		value := gb.B()
		gb.Write8(gb.HL(), value)
		return 8

	case 0x71: // LD (HL),C
		value := gb.C()
		gb.Write8(gb.HL(), value)
		return 8

	case 0x72: // LD (HL),D
		value := gb.D()
		gb.Write8(gb.HL(), value)
		return 8

	case 0x73: // LD (HL),E
		value := gb.E()
		gb.Write8(gb.HL(), value)
		return 8

	case 0x74: // LD (HL),H
		value := gb.H()
		gb.Write8(gb.HL(), value)
		return 8

	case 0x75: // LD (HL),L
		value := gb.L()
		gb.Write8(gb.HL(), value)
		return 8

	case 0x76: // Unhandled: HALT
		panic("Unhandled instruction: HALT (0x76)")

	case 0x77: // LD (HL),A
		value := gb.A()
		gb.Write8(gb.HL(), value)
		return 8

	case 0x78: // LD A,B
		value := gb.B()
		gb.SetA(value)
		return 4

	case 0x79: // LD A,C
		value := gb.C()
		gb.SetA(value)
		return 4

	case 0x7a: // LD A,D
		value := gb.D()
		gb.SetA(value)
		return 4

	case 0x7b: // LD A,E
		value := gb.E()
		gb.SetA(value)
		return 4

	case 0x7c: // LD A,H
		value := gb.H()
		gb.SetA(value)
		return 4

	case 0x7d: // LD A,L
		value := gb.L()
		gb.SetA(value)
		return 4

	case 0x7e: // LD A,(HL)
		value := gb.Read8(gb.HL())
		gb.SetA(value)
		return 8

	case 0x7f: // LD A,A
		value := gb.A()
		gb.SetA(value)
		return 4

	case 0x80: // ADD A,B
		operand1 := gb.A()
		operand2 := gb.B()
		result := operand1 + operand2
		gb.SetA(result)
		gb.SetFlagZ(result == 0)
		gb.SetFlagN(false)
		gb.SetFlagH((operand1&0x0F)+(operand2&0x0F) > 0x0F)
		gb.SetFlagC(uint16(operand1)+uint16(operand2) > 0xFF)
		return 4

	case 0x81: // ADD A,C
		operand1 := gb.A()
		operand2 := gb.C()
		result := operand1 + operand2
		gb.SetA(result)
		gb.SetFlagZ(result == 0)
		gb.SetFlagN(false)
		gb.SetFlagH((operand1&0x0F)+(operand2&0x0F) > 0x0F)
		gb.SetFlagC(uint16(operand1)+uint16(operand2) > 0xFF)
		return 4

	case 0x82: // ADD A,D
		operand1 := gb.A()
		operand2 := gb.D()
		result := operand1 + operand2
		gb.SetA(result)
		gb.SetFlagZ(result == 0)
		gb.SetFlagN(false)
		gb.SetFlagH((operand1&0x0F)+(operand2&0x0F) > 0x0F)
		gb.SetFlagC(uint16(operand1)+uint16(operand2) > 0xFF)
		return 4

	case 0x83: // ADD A,E
		operand1 := gb.A()
		operand2 := gb.E()
		result := operand1 + operand2
		gb.SetA(result)
		gb.SetFlagZ(result == 0)
		gb.SetFlagN(false)
		gb.SetFlagH((operand1&0x0F)+(operand2&0x0F) > 0x0F)
		gb.SetFlagC(uint16(operand1)+uint16(operand2) > 0xFF)
		return 4

	case 0x84: // ADD A,H
		operand1 := gb.A()
		operand2 := gb.H()
		result := operand1 + operand2
		gb.SetA(result)
		gb.SetFlagZ(result == 0)
		gb.SetFlagN(false)
		gb.SetFlagH((operand1&0x0F)+(operand2&0x0F) > 0x0F)
		gb.SetFlagC(uint16(operand1)+uint16(operand2) > 0xFF)
		return 4

	case 0x85: // ADD A,L
		operand1 := gb.A()
		operand2 := gb.L()
		result := operand1 + operand2
		gb.SetA(result)
		gb.SetFlagZ(result == 0)
		gb.SetFlagN(false)
		gb.SetFlagH((operand1&0x0F)+(operand2&0x0F) > 0x0F)
		gb.SetFlagC(uint16(operand1)+uint16(operand2) > 0xFF)
		return 4

	case 0x86: // ADD A,(HL)
		operand1 := gb.A()
		operand2 := gb.Read8(gb.HL())
		result := operand1 + operand2
		gb.SetA(result)
		gb.SetFlagZ(result == 0)
		gb.SetFlagN(false)
		gb.SetFlagH((operand1&0x0F)+(operand2&0x0F) > 0x0F)
		gb.SetFlagC(uint16(operand1)+uint16(operand2) > 0xFF)
		return 8

	case 0x87: // ADD A,A
		operand1 := gb.A()
		operand2 := gb.A()
		result := operand1 + operand2
		gb.SetA(result)
		gb.SetFlagZ(result == 0)
		gb.SetFlagN(false)
		gb.SetFlagH((operand1&0x0F)+(operand2&0x0F) > 0x0F)
		gb.SetFlagC(uint16(operand1)+uint16(operand2) > 0xFF)
		return 4

	case 0x88: // ADC A,B
		operand1 := gb.A()
		operand2 := gb.B()
		result := operand1 + operand2 + gb.Carry()
		gb.SetA(result)
		gb.SetFlagZ(result == 0)
		gb.SetFlagN(false)
		gb.SetFlagH((operand1&0x0F)+(operand2&0x0F)+gb.Carry() > 0x0F)
		gb.SetFlagC(uint16(operand1)+uint16(operand2)+uint16(gb.Carry()) > 0xFF)
		return 4

	case 0x89: // ADC A,C
		operand1 := gb.A()
		operand2 := gb.C()
		result := operand1 + operand2 + gb.Carry()
		gb.SetA(result)
		gb.SetFlagZ(result == 0)
		gb.SetFlagN(false)
		gb.SetFlagH((operand1&0x0F)+(operand2&0x0F)+gb.Carry() > 0x0F)
		gb.SetFlagC(uint16(operand1)+uint16(operand2)+uint16(gb.Carry()) > 0xFF)
		return 4

	case 0x8a: // ADC A,D
		operand1 := gb.A()
		operand2 := gb.D()
		result := operand1 + operand2 + gb.Carry()
		gb.SetA(result)
		gb.SetFlagZ(result == 0)
		gb.SetFlagN(false)
		gb.SetFlagH((operand1&0x0F)+(operand2&0x0F)+gb.Carry() > 0x0F)
		gb.SetFlagC(uint16(operand1)+uint16(operand2)+uint16(gb.Carry()) > 0xFF)
		return 4

	case 0x8b: // ADC A,E
		operand1 := gb.A()
		operand2 := gb.E()
		result := operand1 + operand2 + gb.Carry()
		gb.SetA(result)
		gb.SetFlagZ(result == 0)
		gb.SetFlagN(false)
		gb.SetFlagH((operand1&0x0F)+(operand2&0x0F)+gb.Carry() > 0x0F)
		gb.SetFlagC(uint16(operand1)+uint16(operand2)+uint16(gb.Carry()) > 0xFF)
		return 4

	case 0x8c: // ADC A,H
		operand1 := gb.A()
		operand2 := gb.H()
		result := operand1 + operand2 + gb.Carry()
		gb.SetA(result)
		gb.SetFlagZ(result == 0)
		gb.SetFlagN(false)
		gb.SetFlagH((operand1&0x0F)+(operand2&0x0F)+gb.Carry() > 0x0F)
		gb.SetFlagC(uint16(operand1)+uint16(operand2)+uint16(gb.Carry()) > 0xFF)
		return 4

	case 0x8d: // ADC A,L
		operand1 := gb.A()
		operand2 := gb.L()
		result := operand1 + operand2 + gb.Carry()
		gb.SetA(result)
		gb.SetFlagZ(result == 0)
		gb.SetFlagN(false)
		gb.SetFlagH((operand1&0x0F)+(operand2&0x0F)+gb.Carry() > 0x0F)
		gb.SetFlagC(uint16(operand1)+uint16(operand2)+uint16(gb.Carry()) > 0xFF)
		return 4

	case 0x8e: // ADC A,(HL)
		operand1 := gb.A()
		operand2 := gb.Read8(gb.HL())
		result := operand1 + operand2 + gb.Carry()
		gb.SetA(result)
		gb.SetFlagZ(result == 0)
		gb.SetFlagN(false)
		gb.SetFlagH((operand1&0x0F)+(operand2&0x0F)+gb.Carry() > 0x0F)
		gb.SetFlagC(uint16(operand1)+uint16(operand2)+uint16(gb.Carry()) > 0xFF)
		return 8

	case 0x8f: // ADC A,A
		operand1 := gb.A()
		operand2 := gb.A()
		result := operand1 + operand2 + gb.Carry()
		gb.SetA(result)
		gb.SetFlagZ(result == 0)
		gb.SetFlagN(false)
		gb.SetFlagH((operand1&0x0F)+(operand2&0x0F)+gb.Carry() > 0x0F)
		gb.SetFlagC(uint16(operand1)+uint16(operand2)+uint16(gb.Carry()) > 0xFF)
		return 4

	case 0x90: // SUB B
		operand1 := gb.A()
		operand2 := gb.B()
		result := operand1 - operand2
		gb.SetA(result)
		gb.SetFlagZ(result == 0)
		gb.SetFlagN(true)
		gb.SetFlagH(int16(operand1&0x0F)-int16(operand2&0x0F) < 0)
		gb.SetFlagC(int16(operand1)-int16(operand2) < 0)
		return 4

	case 0x91: // SUB C
		operand1 := gb.A()
		operand2 := gb.C()
		result := operand1 - operand2
		gb.SetA(result)
		gb.SetFlagZ(result == 0)
		gb.SetFlagN(true)
		gb.SetFlagH(int16(operand1&0x0F)-int16(operand2&0x0F) < 0)
		gb.SetFlagC(int16(operand1)-int16(operand2) < 0)
		return 4

	case 0x92: // SUB D
		operand1 := gb.A()
		operand2 := gb.D()
		result := operand1 - operand2
		gb.SetA(result)
		gb.SetFlagZ(result == 0)
		gb.SetFlagN(true)
		gb.SetFlagH(int16(operand1&0x0F)-int16(operand2&0x0F) < 0)
		gb.SetFlagC(int16(operand1)-int16(operand2) < 0)
		return 4

	case 0x93: // SUB E
		operand1 := gb.A()
		operand2 := gb.E()
		result := operand1 - operand2
		gb.SetA(result)
		gb.SetFlagZ(result == 0)
		gb.SetFlagN(true)
		gb.SetFlagH(int16(operand1&0x0F)-int16(operand2&0x0F) < 0)
		gb.SetFlagC(int16(operand1)-int16(operand2) < 0)
		return 4

	case 0x94: // SUB H
		operand1 := gb.A()
		operand2 := gb.H()
		result := operand1 - operand2
		gb.SetA(result)
		gb.SetFlagZ(result == 0)
		gb.SetFlagN(true)
		gb.SetFlagH(int16(operand1&0x0F)-int16(operand2&0x0F) < 0)
		gb.SetFlagC(int16(operand1)-int16(operand2) < 0)
		return 4

	case 0x95: // SUB L
		operand1 := gb.A()
		operand2 := gb.L()
		result := operand1 - operand2
		gb.SetA(result)
		gb.SetFlagZ(result == 0)
		gb.SetFlagN(true)
		gb.SetFlagH(int16(operand1&0x0F)-int16(operand2&0x0F) < 0)
		gb.SetFlagC(int16(operand1)-int16(operand2) < 0)
		return 4

	case 0x96: // SUB (HL)
		operand1 := gb.A()
		operand2 := gb.Read8(gb.HL())
		result := operand1 - operand2
		gb.SetA(result)
		gb.SetFlagZ(result == 0)
		gb.SetFlagN(true)
		gb.SetFlagH(int16(operand1&0x0F)-int16(operand2&0x0F) < 0)
		gb.SetFlagC(int16(operand1)-int16(operand2) < 0)
		return 8

	case 0x97: // SUB A
		operand1 := gb.A()
		operand2 := gb.A()
		result := operand1 - operand2
		gb.SetA(result)
		gb.SetFlagZ(result == 0)
		gb.SetFlagN(true)
		gb.SetFlagH(int16(operand1&0x0F)-int16(operand2&0x0F) < 0)
		gb.SetFlagC(int16(operand1)-int16(operand2) < 0)
		return 4

	case 0x98: // SBC A,B
		operand1 := gb.A()
		operand2 := gb.B()
		result := operand1 - operand2 - gb.Carry()
		gb.SetA(result)
		gb.SetFlagZ(result == 0)
		gb.SetFlagN(true)
		gb.SetFlagH(int16(operand1&0x0F)-int16(operand2&0x0F)-int16(gb.Carry()) < 0)
		gb.SetFlagC(int16(operand1)-int16(operand2)-int16(gb.Carry()) < 0)
		return 4

	case 0x99: // SBC A,C
		operand1 := gb.A()
		operand2 := gb.C()
		result := operand1 - operand2 - gb.Carry()
		gb.SetA(result)
		gb.SetFlagZ(result == 0)
		gb.SetFlagN(true)
		gb.SetFlagH(int16(operand1&0x0F)-int16(operand2&0x0F)-int16(gb.Carry()) < 0)
		gb.SetFlagC(int16(operand1)-int16(operand2)-int16(gb.Carry()) < 0)
		return 4

	case 0x9a: // SBC A,D
		operand1 := gb.A()
		operand2 := gb.D()
		result := operand1 - operand2 - gb.Carry()
		gb.SetA(result)
		gb.SetFlagZ(result == 0)
		gb.SetFlagN(true)
		gb.SetFlagH(int16(operand1&0x0F)-int16(operand2&0x0F)-int16(gb.Carry()) < 0)
		gb.SetFlagC(int16(operand1)-int16(operand2)-int16(gb.Carry()) < 0)
		return 4

	case 0x9b: // SBC A,E
		operand1 := gb.A()
		operand2 := gb.E()
		result := operand1 - operand2 - gb.Carry()
		gb.SetA(result)
		gb.SetFlagZ(result == 0)
		gb.SetFlagN(true)
		gb.SetFlagH(int16(operand1&0x0F)-int16(operand2&0x0F)-int16(gb.Carry()) < 0)
		gb.SetFlagC(int16(operand1)-int16(operand2)-int16(gb.Carry()) < 0)
		return 4

	case 0x9c: // SBC A,H
		operand1 := gb.A()
		operand2 := gb.H()
		result := operand1 - operand2 - gb.Carry()
		gb.SetA(result)
		gb.SetFlagZ(result == 0)
		gb.SetFlagN(true)
		gb.SetFlagH(int16(operand1&0x0F)-int16(operand2&0x0F)-int16(gb.Carry()) < 0)
		gb.SetFlagC(int16(operand1)-int16(operand2)-int16(gb.Carry()) < 0)
		return 4

	case 0x9d: // SBC A,L
		operand1 := gb.A()
		operand2 := gb.L()
		result := operand1 - operand2 - gb.Carry()
		gb.SetA(result)
		gb.SetFlagZ(result == 0)
		gb.SetFlagN(true)
		gb.SetFlagH(int16(operand1&0x0F)-int16(operand2&0x0F)-int16(gb.Carry()) < 0)
		gb.SetFlagC(int16(operand1)-int16(operand2)-int16(gb.Carry()) < 0)
		return 4

	case 0x9e: // SBC A,(HL)
		operand1 := gb.A()
		operand2 := gb.Read8(gb.HL())
		result := operand1 - operand2 - gb.Carry()
		gb.SetA(result)
		gb.SetFlagZ(result == 0)
		gb.SetFlagN(true)
		gb.SetFlagH(int16(operand1&0x0F)-int16(operand2&0x0F)-int16(gb.Carry()) < 0)
		gb.SetFlagC(int16(operand1)-int16(operand2)-int16(gb.Carry()) < 0)
		return 8

	case 0x9f: // SBC A,A
		operand1 := gb.A()
		operand2 := gb.A()
		result := operand1 - operand2 - gb.Carry()
		gb.SetA(result)
		gb.SetFlagZ(result == 0)
		gb.SetFlagN(true)
		gb.SetFlagH(int16(operand1&0x0F)-int16(operand2&0x0F)-int16(gb.Carry()) < 0)
		gb.SetFlagC(int16(operand1)-int16(operand2)-int16(gb.Carry()) < 0)
		return 4

	case 0xa0: // AND B
		operand1 := gb.A()
		operand2 := gb.B()
		result := operand1 & operand2
		gb.SetA(result)
		gb.SetFlagZ(result == 0)
		gb.SetFlagN(false)
		gb.SetFlagH(true)
		gb.SetFlagC(false)
		return 4

	case 0xa1: // AND C
		operand1 := gb.A()
		operand2 := gb.C()
		result := operand1 & operand2
		gb.SetA(result)
		gb.SetFlagZ(result == 0)
		gb.SetFlagN(false)
		gb.SetFlagH(true)
		gb.SetFlagC(false)
		return 4

	case 0xa2: // AND D
		operand1 := gb.A()
		operand2 := gb.D()
		result := operand1 & operand2
		gb.SetA(result)
		gb.SetFlagZ(result == 0)
		gb.SetFlagN(false)
		gb.SetFlagH(true)
		gb.SetFlagC(false)
		return 4

	case 0xa3: // AND E
		operand1 := gb.A()
		operand2 := gb.E()
		result := operand1 & operand2
		gb.SetA(result)
		gb.SetFlagZ(result == 0)
		gb.SetFlagN(false)
		gb.SetFlagH(true)
		gb.SetFlagC(false)
		return 4

	case 0xa4: // AND H
		operand1 := gb.A()
		operand2 := gb.H()
		result := operand1 & operand2
		gb.SetA(result)
		gb.SetFlagZ(result == 0)
		gb.SetFlagN(false)
		gb.SetFlagH(true)
		gb.SetFlagC(false)
		return 4

	case 0xa5: // AND L
		operand1 := gb.A()
		operand2 := gb.L()
		result := operand1 & operand2
		gb.SetA(result)
		gb.SetFlagZ(result == 0)
		gb.SetFlagN(false)
		gb.SetFlagH(true)
		gb.SetFlagC(false)
		return 4

	case 0xa6: // AND (HL)
		operand1 := gb.A()
		operand2 := gb.Read8(gb.HL())
		result := operand1 & operand2
		gb.SetA(result)
		gb.SetFlagZ(result == 0)
		gb.SetFlagN(false)
		gb.SetFlagH(true)
		gb.SetFlagC(false)
		return 8

	case 0xa7: // AND A
		operand1 := gb.A()
		operand2 := gb.A()
		result := operand1 & operand2
		gb.SetA(result)
		gb.SetFlagZ(result == 0)
		gb.SetFlagN(false)
		gb.SetFlagH(true)
		gb.SetFlagC(false)
		return 4

	case 0xa8: // XOR B
		operand1 := gb.A()
		operand2 := gb.B()
		result := operand1 ^ operand2
		gb.SetA(result)
		gb.SetFlagZ(result == 0)
		gb.SetFlagN(false)
		gb.SetFlagH(false)
		gb.SetFlagC(false)
		return 4

	case 0xa9: // XOR C
		operand1 := gb.A()
		operand2 := gb.C()
		result := operand1 ^ operand2
		gb.SetA(result)
		gb.SetFlagZ(result == 0)
		gb.SetFlagN(false)
		gb.SetFlagH(false)
		gb.SetFlagC(false)
		return 4

	case 0xaa: // XOR D
		operand1 := gb.A()
		operand2 := gb.D()
		result := operand1 ^ operand2
		gb.SetA(result)
		gb.SetFlagZ(result == 0)
		gb.SetFlagN(false)
		gb.SetFlagH(false)
		gb.SetFlagC(false)
		return 4

	case 0xab: // XOR E
		operand1 := gb.A()
		operand2 := gb.E()
		result := operand1 ^ operand2
		gb.SetA(result)
		gb.SetFlagZ(result == 0)
		gb.SetFlagN(false)
		gb.SetFlagH(false)
		gb.SetFlagC(false)
		return 4

	case 0xac: // XOR H
		operand1 := gb.A()
		operand2 := gb.H()
		result := operand1 ^ operand2
		gb.SetA(result)
		gb.SetFlagZ(result == 0)
		gb.SetFlagN(false)
		gb.SetFlagH(false)
		gb.SetFlagC(false)
		return 4

	case 0xad: // XOR L
		operand1 := gb.A()
		operand2 := gb.L()
		result := operand1 ^ operand2
		gb.SetA(result)
		gb.SetFlagZ(result == 0)
		gb.SetFlagN(false)
		gb.SetFlagH(false)
		gb.SetFlagC(false)
		return 4

	case 0xae: // XOR (HL)
		operand1 := gb.A()
		operand2 := gb.Read8(gb.HL())
		result := operand1 ^ operand2
		gb.SetA(result)
		gb.SetFlagZ(result == 0)
		gb.SetFlagN(false)
		gb.SetFlagH(false)
		gb.SetFlagC(false)
		return 8

	case 0xaf: // XOR A
		operand1 := gb.A()
		operand2 := gb.A()
		result := operand1 ^ operand2
		gb.SetA(result)
		gb.SetFlagZ(result == 0)
		gb.SetFlagN(false)
		gb.SetFlagH(false)
		gb.SetFlagC(false)
		return 4

	case 0xb0: // OR B
		operand1 := gb.A()
		operand2 := gb.B()
		result := operand1 | operand2
		gb.SetA(result)
		gb.SetFlagZ(result == 0)
		gb.SetFlagN(false)
		gb.SetFlagH(false)
		gb.SetFlagC(false)
		return 4

	case 0xb1: // OR C
		operand1 := gb.A()
		operand2 := gb.C()
		result := operand1 | operand2
		gb.SetA(result)
		gb.SetFlagZ(result == 0)
		gb.SetFlagN(false)
		gb.SetFlagH(false)
		gb.SetFlagC(false)
		return 4

	case 0xb2: // OR D
		operand1 := gb.A()
		operand2 := gb.D()
		result := operand1 | operand2
		gb.SetA(result)
		gb.SetFlagZ(result == 0)
		gb.SetFlagN(false)
		gb.SetFlagH(false)
		gb.SetFlagC(false)
		return 4

	case 0xb3: // OR E
		operand1 := gb.A()
		operand2 := gb.E()
		result := operand1 | operand2
		gb.SetA(result)
		gb.SetFlagZ(result == 0)
		gb.SetFlagN(false)
		gb.SetFlagH(false)
		gb.SetFlagC(false)
		return 4

	case 0xb4: // OR H
		operand1 := gb.A()
		operand2 := gb.H()
		result := operand1 | operand2
		gb.SetA(result)
		gb.SetFlagZ(result == 0)
		gb.SetFlagN(false)
		gb.SetFlagH(false)
		gb.SetFlagC(false)
		return 4

	case 0xb5: // OR L
		operand1 := gb.A()
		operand2 := gb.L()
		result := operand1 | operand2
		gb.SetA(result)
		gb.SetFlagZ(result == 0)
		gb.SetFlagN(false)
		gb.SetFlagH(false)
		gb.SetFlagC(false)
		return 4

	case 0xb6: // OR (HL)
		operand1 := gb.A()
		operand2 := gb.Read8(gb.HL())
		result := operand1 | operand2
		gb.SetA(result)
		gb.SetFlagZ(result == 0)
		gb.SetFlagN(false)
		gb.SetFlagH(false)
		gb.SetFlagC(false)
		return 8

	case 0xb7: // OR A
		operand1 := gb.A()
		operand2 := gb.A()
		result := operand1 | operand2
		gb.SetA(result)
		gb.SetFlagZ(result == 0)
		gb.SetFlagN(false)
		gb.SetFlagH(false)
		gb.SetFlagC(false)
		return 4

	case 0xb8: // CP B
		operand1 := gb.A()
		operand2 := gb.B()
		result := operand1 - operand2
		gb.SetFlagZ(result == 0)
		gb.SetFlagN(true)
		gb.SetFlagH(int16(operand1&0x0F)-int16(operand2&0x0F) < 0)
		gb.SetFlagC(int16(operand1)-int16(operand2) < 0)
		return 4

	case 0xb9: // CP C
		operand1 := gb.A()
		operand2 := gb.C()
		result := operand1 - operand2
		gb.SetFlagZ(result == 0)
		gb.SetFlagN(true)
		gb.SetFlagH(int16(operand1&0x0F)-int16(operand2&0x0F) < 0)
		gb.SetFlagC(int16(operand1)-int16(operand2) < 0)
		return 4

	case 0xba: // CP D
		operand1 := gb.A()
		operand2 := gb.D()
		result := operand1 - operand2
		gb.SetFlagZ(result == 0)
		gb.SetFlagN(true)
		gb.SetFlagH(int16(operand1&0x0F)-int16(operand2&0x0F) < 0)
		gb.SetFlagC(int16(operand1)-int16(operand2) < 0)
		return 4

	case 0xbb: // CP E
		operand1 := gb.A()
		operand2 := gb.E()
		result := operand1 - operand2
		gb.SetFlagZ(result == 0)
		gb.SetFlagN(true)
		gb.SetFlagH(int16(operand1&0x0F)-int16(operand2&0x0F) < 0)
		gb.SetFlagC(int16(operand1)-int16(operand2) < 0)
		return 4

	case 0xbc: // CP H
		operand1 := gb.A()
		operand2 := gb.H()
		result := operand1 - operand2
		gb.SetFlagZ(result == 0)
		gb.SetFlagN(true)
		gb.SetFlagH(int16(operand1&0x0F)-int16(operand2&0x0F) < 0)
		gb.SetFlagC(int16(operand1)-int16(operand2) < 0)
		return 4

	case 0xbd: // CP L
		operand1 := gb.A()
		operand2 := gb.L()
		result := operand1 - operand2
		gb.SetFlagZ(result == 0)
		gb.SetFlagN(true)
		gb.SetFlagH(int16(operand1&0x0F)-int16(operand2&0x0F) < 0)
		gb.SetFlagC(int16(operand1)-int16(operand2) < 0)
		return 4

	case 0xbe: // CP (HL)
		operand1 := gb.A()
		operand2 := gb.Read8(gb.HL())
		result := operand1 - operand2
		gb.SetFlagZ(result == 0)
		gb.SetFlagN(true)
		gb.SetFlagH(int16(operand1&0x0F)-int16(operand2&0x0F) < 0)
		gb.SetFlagC(int16(operand1)-int16(operand2) < 0)
		return 8

	case 0xbf: // CP A
		operand1 := gb.A()
		operand2 := gb.A()
		result := operand1 - operand2
		gb.SetFlagZ(result == 0)
		gb.SetFlagN(true)
		gb.SetFlagH(int16(operand1&0x0F)-int16(operand2&0x0F) < 0)
		gb.SetFlagC(int16(operand1)-int16(operand2) < 0)
		return 4

	case 0xc0: // RET NZ
		if !gb.FlagZ() {
			address := gb.Pop16()
			gb.SetPC(address)
			return 20
		}
		return 8

	case 0xc1: // POP BC
		gb.SetBC(gb.Pop16())
		return 12

	case 0xc2: // JP NZ,a16
		if !gb.FlagZ() {
			address := gb.Fetch16()
			gb.SetPC(address)
			return 16
		}
		return 12

	case 0xc3: // JP a16
		address := gb.Fetch16()
		gb.SetPC(address)
		return 16

	case 0xc4: // CALL NZ,a16
		if !gb.FlagZ() {
			address := gb.Fetch16()
			gb.Push16(gb.PC())
			gb.SetPC(address)
			return 24
		}
		return 12

	case 0xc5: // PUSH BC
		gb.Push16(gb.BC())
		return 16

	case 0xc6: // ADD A,d8
		operand1 := gb.A()
		operand2 := gb.d8()
		result := operand1 + operand2
		gb.SetA(result)
		gb.SetFlagZ(result == 0)
		gb.SetFlagN(false)
		gb.SetFlagH((operand1&0x0F)+(operand2&0x0F) > 0x0F)
		gb.SetFlagC(uint16(operand1)+uint16(operand2) > 0xFF)
		return 8

	case 0xc7: // RST 00H
		const address = 0x00
		gb.Push16(gb.PC())
		gb.SetPC(address)
		return 16

	case 0xc8: // RET Z
		if gb.FlagZ() {
			address := gb.Pop16()
			gb.SetPC(address)
			return 20
		}
		return 8

	case 0xc9: // RET
		address := gb.Pop16()
		gb.SetPC(address)
		return 16

	case 0xca: // JP Z,a16
		if gb.FlagZ() {
			address := gb.Fetch16()
			gb.SetPC(address)
			return 16
		}
		return 12

	case 0xcb: // Unhandled: PREFIX CB
		panic("Unhandled instruction: PREFIX CB (0xcb)")

	case 0xcc: // CALL Z,a16
		if gb.FlagZ() {
			address := gb.Fetch16()
			gb.Push16(gb.PC())
			gb.SetPC(address)
			return 24
		}
		return 12

	case 0xcd: // CALL a16
		address := gb.Fetch16()
		gb.Push16(gb.PC())
		gb.SetPC(address)
		return 24

	case 0xce: // ADC A,d8
		operand1 := gb.A()
		operand2 := gb.d8()
		result := operand1 + operand2 + gb.Carry()
		gb.SetA(result)
		gb.SetFlagZ(result == 0)
		gb.SetFlagN(false)
		gb.SetFlagH((operand1&0x0F)+(operand2&0x0F)+gb.Carry() > 0x0F)
		gb.SetFlagC(uint16(operand1)+uint16(operand2)+uint16(gb.Carry()) > 0xFF)
		return 8

	case 0xcf: // RST 08H
		const address = 0x08
		gb.Push16(gb.PC())
		gb.SetPC(address)
		return 16

	case 0xd0: // RET NC
		if !gb.FlagC() {
			address := gb.Pop16()
			gb.SetPC(address)
			return 20
		}
		return 8

	case 0xd1: // POP DE
		gb.SetDE(gb.Pop16())
		return 12

	case 0xd2: // JP NC,a16
		if !gb.FlagC() {
			address := gb.Fetch16()
			gb.SetPC(address)
			return 16
		}
		return 12

	case 0xd3: // Not used
		panic("Unsupported opcode by Sharp LR35902: 0xd3")

	case 0xd4: // CALL NC,a16
		if !gb.FlagC() {
			address := gb.Fetch16()
			gb.Push16(gb.PC())
			gb.SetPC(address)
			return 24
		}
		return 12

	case 0xd5: // PUSH DE
		gb.Push16(gb.DE())
		return 16

	case 0xd6: // SUB d8
		operand1 := gb.A()
		operand2 := gb.d8()
		result := operand1 - operand2
		gb.SetA(result)
		gb.SetFlagZ(result == 0)
		gb.SetFlagN(true)
		gb.SetFlagH(int16(operand1&0x0F)-int16(operand2&0x0F) < 0)
		gb.SetFlagC(int16(operand1)-int16(operand2) < 0)
		return 8

	case 0xd7: // RST 10H
		const address = 0x10
		gb.Push16(gb.PC())
		gb.SetPC(address)
		return 16

	case 0xd8: // RET C
		if gb.FlagC() {
			address := gb.Pop16()
			gb.SetPC(address)
			return 20
		}
		return 8

	case 0xd9: // Unhandled: RETI
		panic("Unhandled instruction: RETI (0xd9)")

	case 0xda: // JP C,a16
		if gb.FlagC() {
			address := gb.Fetch16()
			gb.SetPC(address)
			return 16
		}
		return 12

	case 0xdb: // Not used
		panic("Unsupported opcode by Sharp LR35902: 0xdb")

	case 0xdc: // CALL C,a16
		if gb.FlagC() {
			address := gb.Fetch16()
			gb.Push16(gb.PC())
			gb.SetPC(address)
			return 24
		}
		return 12

	case 0xdd: // Not used
		panic("Unsupported opcode by Sharp LR35902: 0xdd")

	case 0xde: // SBC A,d8
		operand1 := gb.A()
		operand2 := gb.d8()
		result := operand1 - operand2 - gb.Carry()
		gb.SetA(result)
		gb.SetFlagZ(result == 0)
		gb.SetFlagN(true)
		gb.SetFlagH(int16(operand1&0x0F)-int16(operand2&0x0F)-int16(gb.Carry()) < 0)
		gb.SetFlagC(int16(operand1)-int16(operand2)-int16(gb.Carry()) < 0)
		return 8

	case 0xdf: // RST 18H
		const address = 0x18
		gb.Push16(gb.PC())
		gb.SetPC(address)
		return 16

	case 0xe0: // LDH (a8),A
		value := gb.A()
		gb.Write8(ZeroPage.AddressOf(gb.a8()), value)
		return 12

	case 0xe1: // POP HL
		gb.SetHL(gb.Pop16())
		return 12

	case 0xe2: // LD (C),A
		value := gb.A()
		gb.Write8(ZeroPage.AddressOf(gb.C()), value)
		return 8

	case 0xe3: // Not used
		panic("Unsupported opcode by Sharp LR35902: 0xe3")

	case 0xe4: // Not used
		panic("Unsupported opcode by Sharp LR35902: 0xe4")

	case 0xe5: // PUSH HL
		gb.Push16(gb.HL())
		return 16

	case 0xe6: // AND d8
		operand1 := gb.A()
		operand2 := gb.d8()
		result := operand1 & operand2
		gb.SetA(result)
		gb.SetFlagZ(result == 0)
		gb.SetFlagN(false)
		gb.SetFlagH(true)
		gb.SetFlagC(false)
		return 8

	case 0xe7: // RST 20H
		const address = 0x20
		gb.Push16(gb.PC())
		gb.SetPC(address)
		return 16

	case 0xe8: // Unhandled: ADD SP,r8
		panic("Unhandled instruction: ADD SP,r8 (0xe8)")

	case 0xe9: // JP (HL)
		address := gb.HL()
		gb.SetPC(address)
		return 4

	case 0xea: // LD (a16),A
		value := gb.A()
		gb.Write8(gb.a16(), value)
		return 16

	case 0xeb: // Not used
		panic("Unsupported opcode by Sharp LR35902: 0xeb")

	case 0xec: // Not used
		panic("Unsupported opcode by Sharp LR35902: 0xec")

	case 0xed: // Not used
		panic("Unsupported opcode by Sharp LR35902: 0xed")

	case 0xee: // XOR d8
		operand1 := gb.A()
		operand2 := gb.d8()
		result := operand1 ^ operand2
		gb.SetA(result)
		gb.SetFlagZ(result == 0)
		gb.SetFlagN(false)
		gb.SetFlagH(false)
		gb.SetFlagC(false)
		return 8

	case 0xef: // RST 28H
		const address = 0x28
		gb.Push16(gb.PC())
		gb.SetPC(address)
		return 16

	case 0xf0: // LDH A,(a8)
		value := gb.Read8(ZeroPage.AddressOf(gb.a8()))
		gb.SetA(value)
		return 12

	case 0xf1: // POP AF
		gb.SetAF(gb.Pop16())
		return 12

	case 0xf2: // LD A,(C)
		value := gb.Read8(ZeroPage.AddressOf(gb.C()))
		gb.SetA(value)
		return 8

	case 0xf3: // Unhandled: DI
		panic("Unhandled instruction: DI (0xf3)")

	case 0xf4: // Not used
		panic("Unsupported opcode by Sharp LR35902: 0xf4")

	case 0xf5: // PUSH AF
		gb.Push16(gb.AF())
		return 16

	case 0xf6: // OR d8
		operand1 := gb.A()
		operand2 := gb.d8()
		result := operand1 | operand2
		gb.SetA(result)
		gb.SetFlagZ(result == 0)
		gb.SetFlagN(false)
		gb.SetFlagH(false)
		gb.SetFlagC(false)
		return 8

	case 0xf7: // RST 30H
		const address = 0x30
		gb.Push16(gb.PC())
		gb.SetPC(address)
		return 16

	case 0xf8: // Unhandled: LD HL,SP+r8
		panic("Unhandled instruction: LD HL,SP+r8 (0xf8)")

	case 0xf9: // LD SP,HL
		value := gb.HL()
		gb.SetSP(value)
		return 8

	case 0xfa: // LD A,(a16)
		value := gb.Read8(gb.a16())
		gb.SetA(value)
		return 16

	case 0xfb: // Unhandled: EI
		panic("Unhandled instruction: EI (0xfb)")

	case 0xfc: // Not used
		panic("Unsupported opcode by Sharp LR35902: 0xfc")

	case 0xfd: // Not used
		panic("Unsupported opcode by Sharp LR35902: 0xfd")

	case 0xfe: // CP d8
		operand1 := gb.A()
		operand2 := gb.d8()
		result := operand1 - operand2
		gb.SetFlagZ(result == 0)
		gb.SetFlagN(true)
		gb.SetFlagH(int16(operand1&0x0F)-int16(operand2&0x0F) < 0)
		gb.SetFlagC(int16(operand1)-int16(operand2) < 0)
		return 8

	case 0xff: // RST 38H
		const address = 0x38
		gb.Push16(gb.PC())
		gb.SetPC(address)
		return 16

	default:
		panic(fmt.Sprintf("Opcode unhandled by generated instructions: %#x", opcode))
	}
}
