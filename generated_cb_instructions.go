// This file has been generated automatically. Do not edit manually.
package gb

import "fmt"

func (gb *GameBoy) executeGeneratedCBInstruction(opcode uint8) (cycles uint) {
	switch opcode {

	case 0x0: // RLC B
		value := gb.B()
		result := value << 1
		if value&0x80 != 0 {
			result |= 0x1
		}
		gb.SetB(result)
		gb.SetFlagZ(result == 0)
		gb.SetFlagN(false)
		gb.SetFlagH(false)
		gb.SetFlagC(value&0x80 != 0)
		return 8

	case 0x1: // RLC C
		value := gb.C()
		result := value << 1
		if value&0x80 != 0 {
			result |= 0x1
		}
		gb.SetC(result)
		gb.SetFlagZ(result == 0)
		gb.SetFlagN(false)
		gb.SetFlagH(false)
		gb.SetFlagC(value&0x80 != 0)
		return 8

	case 0x2: // RLC D
		value := gb.D()
		result := value << 1
		if value&0x80 != 0 {
			result |= 0x1
		}
		gb.SetD(result)
		gb.SetFlagZ(result == 0)
		gb.SetFlagN(false)
		gb.SetFlagH(false)
		gb.SetFlagC(value&0x80 != 0)
		return 8

	case 0x3: // RLC E
		value := gb.E()
		result := value << 1
		if value&0x80 != 0 {
			result |= 0x1
		}
		gb.SetE(result)
		gb.SetFlagZ(result == 0)
		gb.SetFlagN(false)
		gb.SetFlagH(false)
		gb.SetFlagC(value&0x80 != 0)
		return 8

	case 0x4: // RLC H
		value := gb.H()
		result := value << 1
		if value&0x80 != 0 {
			result |= 0x1
		}
		gb.SetH(result)
		gb.SetFlagZ(result == 0)
		gb.SetFlagN(false)
		gb.SetFlagH(false)
		gb.SetFlagC(value&0x80 != 0)
		return 8

	case 0x5: // RLC L
		value := gb.L()
		result := value << 1
		if value&0x80 != 0 {
			result |= 0x1
		}
		gb.SetL(result)
		gb.SetFlagZ(result == 0)
		gb.SetFlagN(false)
		gb.SetFlagH(false)
		gb.SetFlagC(value&0x80 != 0)
		return 8

	case 0x6: // RLC (HL)
		value := gb.Read8(gb.HL())
		result := value << 1
		if value&0x80 != 0 {
			result |= 0x1
		}
		gb.Write8(gb.HL(), result)
		gb.SetFlagZ(result == 0)
		gb.SetFlagN(false)
		gb.SetFlagH(false)
		gb.SetFlagC(value&0x80 != 0)
		return 16

	case 0x7: // RLC A
		value := gb.A()
		result := value << 1
		if value&0x80 != 0 {
			result |= 0x1
		}
		gb.SetA(result)
		gb.SetFlagZ(result == 0)
		gb.SetFlagN(false)
		gb.SetFlagH(false)
		gb.SetFlagC(value&0x80 != 0)
		return 8

	case 0x8: // RRC B
		value := gb.B()
		result := value >> 1
		if value&0x01 != 0 {
			result |= 0x80
		}
		gb.SetB(result)
		gb.SetFlagZ(result == 0)
		gb.SetFlagN(false)
		gb.SetFlagH(false)
		gb.SetFlagC(value&0x01 != 0)
		return 8

	case 0x9: // RRC C
		value := gb.C()
		result := value >> 1
		if value&0x01 != 0 {
			result |= 0x80
		}
		gb.SetC(result)
		gb.SetFlagZ(result == 0)
		gb.SetFlagN(false)
		gb.SetFlagH(false)
		gb.SetFlagC(value&0x01 != 0)
		return 8

	case 0xa: // RRC D
		value := gb.D()
		result := value >> 1
		if value&0x01 != 0 {
			result |= 0x80
		}
		gb.SetD(result)
		gb.SetFlagZ(result == 0)
		gb.SetFlagN(false)
		gb.SetFlagH(false)
		gb.SetFlagC(value&0x01 != 0)
		return 8

	case 0xb: // RRC E
		value := gb.E()
		result := value >> 1
		if value&0x01 != 0 {
			result |= 0x80
		}
		gb.SetE(result)
		gb.SetFlagZ(result == 0)
		gb.SetFlagN(false)
		gb.SetFlagH(false)
		gb.SetFlagC(value&0x01 != 0)
		return 8

	case 0xc: // RRC H
		value := gb.H()
		result := value >> 1
		if value&0x01 != 0 {
			result |= 0x80
		}
		gb.SetH(result)
		gb.SetFlagZ(result == 0)
		gb.SetFlagN(false)
		gb.SetFlagH(false)
		gb.SetFlagC(value&0x01 != 0)
		return 8

	case 0xd: // RRC L
		value := gb.L()
		result := value >> 1
		if value&0x01 != 0 {
			result |= 0x80
		}
		gb.SetL(result)
		gb.SetFlagZ(result == 0)
		gb.SetFlagN(false)
		gb.SetFlagH(false)
		gb.SetFlagC(value&0x01 != 0)
		return 8

	case 0xe: // RRC (HL)
		value := gb.Read8(gb.HL())
		result := value >> 1
		if value&0x01 != 0 {
			result |= 0x80
		}
		gb.Write8(gb.HL(), result)
		gb.SetFlagZ(result == 0)
		gb.SetFlagN(false)
		gb.SetFlagH(false)
		gb.SetFlagC(value&0x01 != 0)
		return 16

	case 0xf: // RRC A
		value := gb.A()
		result := value >> 1
		if value&0x01 != 0 {
			result |= 0x80
		}
		gb.SetA(result)
		gb.SetFlagZ(result == 0)
		gb.SetFlagN(false)
		gb.SetFlagH(false)
		gb.SetFlagC(value&0x01 != 0)
		return 8

	case 0x10: // RL B
		value := gb.B()
		result := value << 1
		if gb.FlagC() {
			result |= 0x1
		}
		gb.SetB(result)
		gb.SetFlagZ(result == 0)
		gb.SetFlagN(false)
		gb.SetFlagH(false)
		gb.SetFlagC(value&0x80 != 0)
		return 8

	case 0x11: // RL C
		value := gb.C()
		result := value << 1
		if gb.FlagC() {
			result |= 0x1
		}
		gb.SetC(result)
		gb.SetFlagZ(result == 0)
		gb.SetFlagN(false)
		gb.SetFlagH(false)
		gb.SetFlagC(value&0x80 != 0)
		return 8

	case 0x12: // RL D
		value := gb.D()
		result := value << 1
		if gb.FlagC() {
			result |= 0x1
		}
		gb.SetD(result)
		gb.SetFlagZ(result == 0)
		gb.SetFlagN(false)
		gb.SetFlagH(false)
		gb.SetFlagC(value&0x80 != 0)
		return 8

	case 0x13: // RL E
		value := gb.E()
		result := value << 1
		if gb.FlagC() {
			result |= 0x1
		}
		gb.SetE(result)
		gb.SetFlagZ(result == 0)
		gb.SetFlagN(false)
		gb.SetFlagH(false)
		gb.SetFlagC(value&0x80 != 0)
		return 8

	case 0x14: // RL H
		value := gb.H()
		result := value << 1
		if gb.FlagC() {
			result |= 0x1
		}
		gb.SetH(result)
		gb.SetFlagZ(result == 0)
		gb.SetFlagN(false)
		gb.SetFlagH(false)
		gb.SetFlagC(value&0x80 != 0)
		return 8

	case 0x15: // RL L
		value := gb.L()
		result := value << 1
		if gb.FlagC() {
			result |= 0x1
		}
		gb.SetL(result)
		gb.SetFlagZ(result == 0)
		gb.SetFlagN(false)
		gb.SetFlagH(false)
		gb.SetFlagC(value&0x80 != 0)
		return 8

	case 0x16: // RL (HL)
		value := gb.Read8(gb.HL())
		result := value << 1
		if gb.FlagC() {
			result |= 0x1
		}
		gb.Write8(gb.HL(), result)
		gb.SetFlagZ(result == 0)
		gb.SetFlagN(false)
		gb.SetFlagH(false)
		gb.SetFlagC(value&0x80 != 0)
		return 16

	case 0x17: // RL A
		value := gb.A()
		result := value << 1
		if gb.FlagC() {
			result |= 0x1
		}
		gb.SetA(result)
		gb.SetFlagZ(result == 0)
		gb.SetFlagN(false)
		gb.SetFlagH(false)
		gb.SetFlagC(value&0x80 != 0)
		return 8

	case 0x18: // RR B
		value := gb.B()
		result := value >> 1
		if gb.FlagC() {
			result |= 0x80
		}
		gb.SetB(result)
		gb.SetFlagZ(result == 0)
		gb.SetFlagN(false)
		gb.SetFlagH(false)
		gb.SetFlagC(value&0x01 != 0)
		return 8

	case 0x19: // RR C
		value := gb.C()
		result := value >> 1
		if gb.FlagC() {
			result |= 0x80
		}
		gb.SetC(result)
		gb.SetFlagZ(result == 0)
		gb.SetFlagN(false)
		gb.SetFlagH(false)
		gb.SetFlagC(value&0x01 != 0)
		return 8

	case 0x1a: // RR D
		value := gb.D()
		result := value >> 1
		if gb.FlagC() {
			result |= 0x80
		}
		gb.SetD(result)
		gb.SetFlagZ(result == 0)
		gb.SetFlagN(false)
		gb.SetFlagH(false)
		gb.SetFlagC(value&0x01 != 0)
		return 8

	case 0x1b: // RR E
		value := gb.E()
		result := value >> 1
		if gb.FlagC() {
			result |= 0x80
		}
		gb.SetE(result)
		gb.SetFlagZ(result == 0)
		gb.SetFlagN(false)
		gb.SetFlagH(false)
		gb.SetFlagC(value&0x01 != 0)
		return 8

	case 0x1c: // RR H
		value := gb.H()
		result := value >> 1
		if gb.FlagC() {
			result |= 0x80
		}
		gb.SetH(result)
		gb.SetFlagZ(result == 0)
		gb.SetFlagN(false)
		gb.SetFlagH(false)
		gb.SetFlagC(value&0x01 != 0)
		return 8

	case 0x1d: // RR L
		value := gb.L()
		result := value >> 1
		if gb.FlagC() {
			result |= 0x80
		}
		gb.SetL(result)
		gb.SetFlagZ(result == 0)
		gb.SetFlagN(false)
		gb.SetFlagH(false)
		gb.SetFlagC(value&0x01 != 0)
		return 8

	case 0x1e: // RR (HL)
		value := gb.Read8(gb.HL())
		result := value >> 1
		if gb.FlagC() {
			result |= 0x80
		}
		gb.Write8(gb.HL(), result)
		gb.SetFlagZ(result == 0)
		gb.SetFlagN(false)
		gb.SetFlagH(false)
		gb.SetFlagC(value&0x01 != 0)
		return 16

	case 0x1f: // RR A
		value := gb.A()
		result := value >> 1
		if gb.FlagC() {
			result |= 0x80
		}
		gb.SetA(result)
		gb.SetFlagZ(result == 0)
		gb.SetFlagN(false)
		gb.SetFlagH(false)
		gb.SetFlagC(value&0x01 != 0)
		return 8

	case 0x20: // SLA B
		value := gb.B()
		result := value << 1
		gb.SetB(result)
		gb.SetFlagZ(result == 0)
		gb.SetFlagN(false)
		gb.SetFlagH(false)
		gb.SetFlagC(value&0x80 != 0)
		return 8

	case 0x21: // SLA C
		value := gb.C()
		result := value << 1
		gb.SetC(result)
		gb.SetFlagZ(result == 0)
		gb.SetFlagN(false)
		gb.SetFlagH(false)
		gb.SetFlagC(value&0x80 != 0)
		return 8

	case 0x22: // SLA D
		value := gb.D()
		result := value << 1
		gb.SetD(result)
		gb.SetFlagZ(result == 0)
		gb.SetFlagN(false)
		gb.SetFlagH(false)
		gb.SetFlagC(value&0x80 != 0)
		return 8

	case 0x23: // SLA E
		value := gb.E()
		result := value << 1
		gb.SetE(result)
		gb.SetFlagZ(result == 0)
		gb.SetFlagN(false)
		gb.SetFlagH(false)
		gb.SetFlagC(value&0x80 != 0)
		return 8

	case 0x24: // SLA H
		value := gb.H()
		result := value << 1
		gb.SetH(result)
		gb.SetFlagZ(result == 0)
		gb.SetFlagN(false)
		gb.SetFlagH(false)
		gb.SetFlagC(value&0x80 != 0)
		return 8

	case 0x25: // SLA L
		value := gb.L()
		result := value << 1
		gb.SetL(result)
		gb.SetFlagZ(result == 0)
		gb.SetFlagN(false)
		gb.SetFlagH(false)
		gb.SetFlagC(value&0x80 != 0)
		return 8

	case 0x26: // SLA (HL)
		value := gb.Read8(gb.HL())
		result := value << 1
		gb.Write8(gb.HL(), result)
		gb.SetFlagZ(result == 0)
		gb.SetFlagN(false)
		gb.SetFlagH(false)
		gb.SetFlagC(value&0x80 != 0)
		return 16

	case 0x27: // SLA A
		value := gb.A()
		result := value << 1
		gb.SetA(result)
		gb.SetFlagZ(result == 0)
		gb.SetFlagN(false)
		gb.SetFlagH(false)
		gb.SetFlagC(value&0x80 != 0)
		return 8

	case 0x28: // SRA B
		value := gb.B()
		result := value >> 1
		if value&0x80 != 0 {
			result |= 0x80
		}
		gb.SetB(result)
		gb.SetFlagZ(result == 0)
		gb.SetFlagN(false)
		gb.SetFlagH(false)
		gb.SetFlagC(false)
		return 8

	case 0x29: // SRA C
		value := gb.C()
		result := value >> 1
		if value&0x80 != 0 {
			result |= 0x80
		}
		gb.SetC(result)
		gb.SetFlagZ(result == 0)
		gb.SetFlagN(false)
		gb.SetFlagH(false)
		gb.SetFlagC(false)
		return 8

	case 0x2a: // SRA D
		value := gb.D()
		result := value >> 1
		if value&0x80 != 0 {
			result |= 0x80
		}
		gb.SetD(result)
		gb.SetFlagZ(result == 0)
		gb.SetFlagN(false)
		gb.SetFlagH(false)
		gb.SetFlagC(false)
		return 8

	case 0x2b: // SRA E
		value := gb.E()
		result := value >> 1
		if value&0x80 != 0 {
			result |= 0x80
		}
		gb.SetE(result)
		gb.SetFlagZ(result == 0)
		gb.SetFlagN(false)
		gb.SetFlagH(false)
		gb.SetFlagC(false)
		return 8

	case 0x2c: // SRA H
		value := gb.H()
		result := value >> 1
		if value&0x80 != 0 {
			result |= 0x80
		}
		gb.SetH(result)
		gb.SetFlagZ(result == 0)
		gb.SetFlagN(false)
		gb.SetFlagH(false)
		gb.SetFlagC(false)
		return 8

	case 0x2d: // SRA L
		value := gb.L()
		result := value >> 1
		if value&0x80 != 0 {
			result |= 0x80
		}
		gb.SetL(result)
		gb.SetFlagZ(result == 0)
		gb.SetFlagN(false)
		gb.SetFlagH(false)
		gb.SetFlagC(false)
		return 8

	case 0x2e: // SRA (HL)
		value := gb.Read8(gb.HL())
		result := value >> 1
		if value&0x80 != 0 {
			result |= 0x80
		}
		gb.Write8(gb.HL(), result)
		gb.SetFlagZ(result == 0)
		gb.SetFlagN(false)
		gb.SetFlagH(false)
		gb.SetFlagC(false)
		return 16

	case 0x2f: // SRA A
		value := gb.A()
		result := value >> 1
		if value&0x80 != 0 {
			result |= 0x80
		}
		gb.SetA(result)
		gb.SetFlagZ(result == 0)
		gb.SetFlagN(false)
		gb.SetFlagH(false)
		gb.SetFlagC(false)
		return 8

	case 0x30: // SWAP B
		value := gb.B()
		result := (value >> 4) & (value << 4)
		gb.SetB(result)
		gb.SetFlagZ(result == 0)
		gb.SetFlagN(false)
		gb.SetFlagH(false)
		gb.SetFlagC(false)
		return 8

	case 0x31: // SWAP C
		value := gb.C()
		result := (value >> 4) & (value << 4)
		gb.SetC(result)
		gb.SetFlagZ(result == 0)
		gb.SetFlagN(false)
		gb.SetFlagH(false)
		gb.SetFlagC(false)
		return 8

	case 0x32: // SWAP D
		value := gb.D()
		result := (value >> 4) & (value << 4)
		gb.SetD(result)
		gb.SetFlagZ(result == 0)
		gb.SetFlagN(false)
		gb.SetFlagH(false)
		gb.SetFlagC(false)
		return 8

	case 0x33: // SWAP E
		value := gb.E()
		result := (value >> 4) & (value << 4)
		gb.SetE(result)
		gb.SetFlagZ(result == 0)
		gb.SetFlagN(false)
		gb.SetFlagH(false)
		gb.SetFlagC(false)
		return 8

	case 0x34: // SWAP H
		value := gb.H()
		result := (value >> 4) & (value << 4)
		gb.SetH(result)
		gb.SetFlagZ(result == 0)
		gb.SetFlagN(false)
		gb.SetFlagH(false)
		gb.SetFlagC(false)
		return 8

	case 0x35: // SWAP L
		value := gb.L()
		result := (value >> 4) & (value << 4)
		gb.SetL(result)
		gb.SetFlagZ(result == 0)
		gb.SetFlagN(false)
		gb.SetFlagH(false)
		gb.SetFlagC(false)
		return 8

	case 0x36: // SWAP (HL)
		value := gb.Read8(gb.HL())
		result := (value >> 4) & (value << 4)
		gb.Write8(gb.HL(), result)
		gb.SetFlagZ(result == 0)
		gb.SetFlagN(false)
		gb.SetFlagH(false)
		gb.SetFlagC(false)
		return 16

	case 0x37: // SWAP A
		value := gb.A()
		result := (value >> 4) & (value << 4)
		gb.SetA(result)
		gb.SetFlagZ(result == 0)
		gb.SetFlagN(false)
		gb.SetFlagH(false)
		gb.SetFlagC(false)
		return 8

	case 0x38: // SRL B
		value := gb.B()
		result := value >> 1
		gb.SetB(result)
		gb.SetFlagZ(result == 0)
		gb.SetFlagN(false)
		gb.SetFlagH(false)
		gb.SetFlagC(value&0x01 != 0)
		return 8

	case 0x39: // SRL C
		value := gb.C()
		result := value >> 1
		gb.SetC(result)
		gb.SetFlagZ(result == 0)
		gb.SetFlagN(false)
		gb.SetFlagH(false)
		gb.SetFlagC(value&0x01 != 0)
		return 8

	case 0x3a: // SRL D
		value := gb.D()
		result := value >> 1
		gb.SetD(result)
		gb.SetFlagZ(result == 0)
		gb.SetFlagN(false)
		gb.SetFlagH(false)
		gb.SetFlagC(value&0x01 != 0)
		return 8

	case 0x3b: // SRL E
		value := gb.E()
		result := value >> 1
		gb.SetE(result)
		gb.SetFlagZ(result == 0)
		gb.SetFlagN(false)
		gb.SetFlagH(false)
		gb.SetFlagC(value&0x01 != 0)
		return 8

	case 0x3c: // SRL H
		value := gb.H()
		result := value >> 1
		gb.SetH(result)
		gb.SetFlagZ(result == 0)
		gb.SetFlagN(false)
		gb.SetFlagH(false)
		gb.SetFlagC(value&0x01 != 0)
		return 8

	case 0x3d: // SRL L
		value := gb.L()
		result := value >> 1
		gb.SetL(result)
		gb.SetFlagZ(result == 0)
		gb.SetFlagN(false)
		gb.SetFlagH(false)
		gb.SetFlagC(value&0x01 != 0)
		return 8

	case 0x3e: // SRL (HL)
		value := gb.Read8(gb.HL())
		result := value >> 1
		gb.Write8(gb.HL(), result)
		gb.SetFlagZ(result == 0)
		gb.SetFlagN(false)
		gb.SetFlagH(false)
		gb.SetFlagC(value&0x01 != 0)
		return 16

	case 0x3f: // SRL A
		value := gb.A()
		result := value >> 1
		gb.SetA(result)
		gb.SetFlagZ(result == 0)
		gb.SetFlagN(false)
		gb.SetFlagH(false)
		gb.SetFlagC(value&0x01 != 0)
		return 8

	case 0x40: // BIT 0,B
		value := gb.B()
		result := value & (1 << 0)
		gb.SetFlagZ(result == 0)
		gb.SetFlagN(false)
		gb.SetFlagH(true)
		return 8

	case 0x41: // BIT 0,C
		value := gb.C()
		result := value & (1 << 0)
		gb.SetFlagZ(result == 0)
		gb.SetFlagN(false)
		gb.SetFlagH(true)
		return 8

	case 0x42: // BIT 0,D
		value := gb.D()
		result := value & (1 << 0)
		gb.SetFlagZ(result == 0)
		gb.SetFlagN(false)
		gb.SetFlagH(true)
		return 8

	case 0x43: // BIT 0,E
		value := gb.E()
		result := value & (1 << 0)
		gb.SetFlagZ(result == 0)
		gb.SetFlagN(false)
		gb.SetFlagH(true)
		return 8

	case 0x44: // BIT 0,H
		value := gb.H()
		result := value & (1 << 0)
		gb.SetFlagZ(result == 0)
		gb.SetFlagN(false)
		gb.SetFlagH(true)
		return 8

	case 0x45: // BIT 0,L
		value := gb.L()
		result := value & (1 << 0)
		gb.SetFlagZ(result == 0)
		gb.SetFlagN(false)
		gb.SetFlagH(true)
		return 8

	case 0x46: // BIT 0,(HL)
		value := gb.Read8(gb.HL())
		result := value & (1 << 0)
		gb.SetFlagZ(result == 0)
		gb.SetFlagN(false)
		gb.SetFlagH(true)
		return 16

	case 0x47: // BIT 0,A
		value := gb.A()
		result := value & (1 << 0)
		gb.SetFlagZ(result == 0)
		gb.SetFlagN(false)
		gb.SetFlagH(true)
		return 8

	case 0x48: // BIT 1,B
		value := gb.B()
		result := value & (1 << 1)
		gb.SetFlagZ(result == 0)
		gb.SetFlagN(false)
		gb.SetFlagH(true)
		return 8

	case 0x49: // BIT 1,C
		value := gb.C()
		result := value & (1 << 1)
		gb.SetFlagZ(result == 0)
		gb.SetFlagN(false)
		gb.SetFlagH(true)
		return 8

	case 0x4a: // BIT 1,D
		value := gb.D()
		result := value & (1 << 1)
		gb.SetFlagZ(result == 0)
		gb.SetFlagN(false)
		gb.SetFlagH(true)
		return 8

	case 0x4b: // BIT 1,E
		value := gb.E()
		result := value & (1 << 1)
		gb.SetFlagZ(result == 0)
		gb.SetFlagN(false)
		gb.SetFlagH(true)
		return 8

	case 0x4c: // BIT 1,H
		value := gb.H()
		result := value & (1 << 1)
		gb.SetFlagZ(result == 0)
		gb.SetFlagN(false)
		gb.SetFlagH(true)
		return 8

	case 0x4d: // BIT 1,L
		value := gb.L()
		result := value & (1 << 1)
		gb.SetFlagZ(result == 0)
		gb.SetFlagN(false)
		gb.SetFlagH(true)
		return 8

	case 0x4e: // BIT 1,(HL)
		value := gb.Read8(gb.HL())
		result := value & (1 << 1)
		gb.SetFlagZ(result == 0)
		gb.SetFlagN(false)
		gb.SetFlagH(true)
		return 16

	case 0x4f: // BIT 1,A
		value := gb.A()
		result := value & (1 << 1)
		gb.SetFlagZ(result == 0)
		gb.SetFlagN(false)
		gb.SetFlagH(true)
		return 8

	case 0x50: // BIT 2,B
		value := gb.B()
		result := value & (1 << 2)
		gb.SetFlagZ(result == 0)
		gb.SetFlagN(false)
		gb.SetFlagH(true)
		return 8

	case 0x51: // BIT 2,C
		value := gb.C()
		result := value & (1 << 2)
		gb.SetFlagZ(result == 0)
		gb.SetFlagN(false)
		gb.SetFlagH(true)
		return 8

	case 0x52: // BIT 2,D
		value := gb.D()
		result := value & (1 << 2)
		gb.SetFlagZ(result == 0)
		gb.SetFlagN(false)
		gb.SetFlagH(true)
		return 8

	case 0x53: // BIT 2,E
		value := gb.E()
		result := value & (1 << 2)
		gb.SetFlagZ(result == 0)
		gb.SetFlagN(false)
		gb.SetFlagH(true)
		return 8

	case 0x54: // BIT 2,H
		value := gb.H()
		result := value & (1 << 2)
		gb.SetFlagZ(result == 0)
		gb.SetFlagN(false)
		gb.SetFlagH(true)
		return 8

	case 0x55: // BIT 2,L
		value := gb.L()
		result := value & (1 << 2)
		gb.SetFlagZ(result == 0)
		gb.SetFlagN(false)
		gb.SetFlagH(true)
		return 8

	case 0x56: // BIT 2,(HL)
		value := gb.Read8(gb.HL())
		result := value & (1 << 2)
		gb.SetFlagZ(result == 0)
		gb.SetFlagN(false)
		gb.SetFlagH(true)
		return 16

	case 0x57: // BIT 2,A
		value := gb.A()
		result := value & (1 << 2)
		gb.SetFlagZ(result == 0)
		gb.SetFlagN(false)
		gb.SetFlagH(true)
		return 8

	case 0x58: // BIT 3,B
		value := gb.B()
		result := value & (1 << 3)
		gb.SetFlagZ(result == 0)
		gb.SetFlagN(false)
		gb.SetFlagH(true)
		return 8

	case 0x59: // BIT 3,C
		value := gb.C()
		result := value & (1 << 3)
		gb.SetFlagZ(result == 0)
		gb.SetFlagN(false)
		gb.SetFlagH(true)
		return 8

	case 0x5a: // BIT 3,D
		value := gb.D()
		result := value & (1 << 3)
		gb.SetFlagZ(result == 0)
		gb.SetFlagN(false)
		gb.SetFlagH(true)
		return 8

	case 0x5b: // BIT 3,E
		value := gb.E()
		result := value & (1 << 3)
		gb.SetFlagZ(result == 0)
		gb.SetFlagN(false)
		gb.SetFlagH(true)
		return 8

	case 0x5c: // BIT 3,H
		value := gb.H()
		result := value & (1 << 3)
		gb.SetFlagZ(result == 0)
		gb.SetFlagN(false)
		gb.SetFlagH(true)
		return 8

	case 0x5d: // BIT 3,L
		value := gb.L()
		result := value & (1 << 3)
		gb.SetFlagZ(result == 0)
		gb.SetFlagN(false)
		gb.SetFlagH(true)
		return 8

	case 0x5e: // BIT 3,(HL)
		value := gb.Read8(gb.HL())
		result := value & (1 << 3)
		gb.SetFlagZ(result == 0)
		gb.SetFlagN(false)
		gb.SetFlagH(true)
		return 16

	case 0x5f: // BIT 3,A
		value := gb.A()
		result := value & (1 << 3)
		gb.SetFlagZ(result == 0)
		gb.SetFlagN(false)
		gb.SetFlagH(true)
		return 8

	case 0x60: // BIT 4,B
		value := gb.B()
		result := value & (1 << 4)
		gb.SetFlagZ(result == 0)
		gb.SetFlagN(false)
		gb.SetFlagH(true)
		return 8

	case 0x61: // BIT 4,C
		value := gb.C()
		result := value & (1 << 4)
		gb.SetFlagZ(result == 0)
		gb.SetFlagN(false)
		gb.SetFlagH(true)
		return 8

	case 0x62: // BIT 4,D
		value := gb.D()
		result := value & (1 << 4)
		gb.SetFlagZ(result == 0)
		gb.SetFlagN(false)
		gb.SetFlagH(true)
		return 8

	case 0x63: // BIT 4,E
		value := gb.E()
		result := value & (1 << 4)
		gb.SetFlagZ(result == 0)
		gb.SetFlagN(false)
		gb.SetFlagH(true)
		return 8

	case 0x64: // BIT 4,H
		value := gb.H()
		result := value & (1 << 4)
		gb.SetFlagZ(result == 0)
		gb.SetFlagN(false)
		gb.SetFlagH(true)
		return 8

	case 0x65: // BIT 4,L
		value := gb.L()
		result := value & (1 << 4)
		gb.SetFlagZ(result == 0)
		gb.SetFlagN(false)
		gb.SetFlagH(true)
		return 8

	case 0x66: // BIT 4,(HL)
		value := gb.Read8(gb.HL())
		result := value & (1 << 4)
		gb.SetFlagZ(result == 0)
		gb.SetFlagN(false)
		gb.SetFlagH(true)
		return 16

	case 0x67: // BIT 4,A
		value := gb.A()
		result := value & (1 << 4)
		gb.SetFlagZ(result == 0)
		gb.SetFlagN(false)
		gb.SetFlagH(true)
		return 8

	case 0x68: // BIT 5,B
		value := gb.B()
		result := value & (1 << 5)
		gb.SetFlagZ(result == 0)
		gb.SetFlagN(false)
		gb.SetFlagH(true)
		return 8

	case 0x69: // BIT 5,C
		value := gb.C()
		result := value & (1 << 5)
		gb.SetFlagZ(result == 0)
		gb.SetFlagN(false)
		gb.SetFlagH(true)
		return 8

	case 0x6a: // BIT 5,D
		value := gb.D()
		result := value & (1 << 5)
		gb.SetFlagZ(result == 0)
		gb.SetFlagN(false)
		gb.SetFlagH(true)
		return 8

	case 0x6b: // BIT 5,E
		value := gb.E()
		result := value & (1 << 5)
		gb.SetFlagZ(result == 0)
		gb.SetFlagN(false)
		gb.SetFlagH(true)
		return 8

	case 0x6c: // BIT 5,H
		value := gb.H()
		result := value & (1 << 5)
		gb.SetFlagZ(result == 0)
		gb.SetFlagN(false)
		gb.SetFlagH(true)
		return 8

	case 0x6d: // BIT 5,L
		value := gb.L()
		result := value & (1 << 5)
		gb.SetFlagZ(result == 0)
		gb.SetFlagN(false)
		gb.SetFlagH(true)
		return 8

	case 0x6e: // BIT 5,(HL)
		value := gb.Read8(gb.HL())
		result := value & (1 << 5)
		gb.SetFlagZ(result == 0)
		gb.SetFlagN(false)
		gb.SetFlagH(true)
		return 16

	case 0x6f: // BIT 5,A
		value := gb.A()
		result := value & (1 << 5)
		gb.SetFlagZ(result == 0)
		gb.SetFlagN(false)
		gb.SetFlagH(true)
		return 8

	case 0x70: // BIT 6,B
		value := gb.B()
		result := value & (1 << 6)
		gb.SetFlagZ(result == 0)
		gb.SetFlagN(false)
		gb.SetFlagH(true)
		return 8

	case 0x71: // BIT 6,C
		value := gb.C()
		result := value & (1 << 6)
		gb.SetFlagZ(result == 0)
		gb.SetFlagN(false)
		gb.SetFlagH(true)
		return 8

	case 0x72: // BIT 6,D
		value := gb.D()
		result := value & (1 << 6)
		gb.SetFlagZ(result == 0)
		gb.SetFlagN(false)
		gb.SetFlagH(true)
		return 8

	case 0x73: // BIT 6,E
		value := gb.E()
		result := value & (1 << 6)
		gb.SetFlagZ(result == 0)
		gb.SetFlagN(false)
		gb.SetFlagH(true)
		return 8

	case 0x74: // BIT 6,H
		value := gb.H()
		result := value & (1 << 6)
		gb.SetFlagZ(result == 0)
		gb.SetFlagN(false)
		gb.SetFlagH(true)
		return 8

	case 0x75: // BIT 6,L
		value := gb.L()
		result := value & (1 << 6)
		gb.SetFlagZ(result == 0)
		gb.SetFlagN(false)
		gb.SetFlagH(true)
		return 8

	case 0x76: // BIT 6,(HL)
		value := gb.Read8(gb.HL())
		result := value & (1 << 6)
		gb.SetFlagZ(result == 0)
		gb.SetFlagN(false)
		gb.SetFlagH(true)
		return 16

	case 0x77: // BIT 6,A
		value := gb.A()
		result := value & (1 << 6)
		gb.SetFlagZ(result == 0)
		gb.SetFlagN(false)
		gb.SetFlagH(true)
		return 8

	case 0x78: // BIT 7,B
		value := gb.B()
		result := value & (1 << 7)
		gb.SetFlagZ(result == 0)
		gb.SetFlagN(false)
		gb.SetFlagH(true)
		return 8

	case 0x79: // BIT 7,C
		value := gb.C()
		result := value & (1 << 7)
		gb.SetFlagZ(result == 0)
		gb.SetFlagN(false)
		gb.SetFlagH(true)
		return 8

	case 0x7a: // BIT 7,D
		value := gb.D()
		result := value & (1 << 7)
		gb.SetFlagZ(result == 0)
		gb.SetFlagN(false)
		gb.SetFlagH(true)
		return 8

	case 0x7b: // BIT 7,E
		value := gb.E()
		result := value & (1 << 7)
		gb.SetFlagZ(result == 0)
		gb.SetFlagN(false)
		gb.SetFlagH(true)
		return 8

	case 0x7c: // BIT 7,H
		value := gb.H()
		result := value & (1 << 7)
		gb.SetFlagZ(result == 0)
		gb.SetFlagN(false)
		gb.SetFlagH(true)
		return 8

	case 0x7d: // BIT 7,L
		value := gb.L()
		result := value & (1 << 7)
		gb.SetFlagZ(result == 0)
		gb.SetFlagN(false)
		gb.SetFlagH(true)
		return 8

	case 0x7e: // BIT 7,(HL)
		value := gb.Read8(gb.HL())
		result := value & (1 << 7)
		gb.SetFlagZ(result == 0)
		gb.SetFlagN(false)
		gb.SetFlagH(true)
		return 16

	case 0x7f: // BIT 7,A
		value := gb.A()
		result := value & (1 << 7)
		gb.SetFlagZ(result == 0)
		gb.SetFlagN(false)
		gb.SetFlagH(true)
		return 8

	case 0x80: // RES 0,B
		value := gb.B()
		result := value & ^uint8(1<<0)
		gb.SetB(result)
		return 8

	case 0x81: // RES 0,C
		value := gb.C()
		result := value & ^uint8(1<<0)
		gb.SetC(result)
		return 8

	case 0x82: // RES 0,D
		value := gb.D()
		result := value & ^uint8(1<<0)
		gb.SetD(result)
		return 8

	case 0x83: // RES 0,E
		value := gb.E()
		result := value & ^uint8(1<<0)
		gb.SetE(result)
		return 8

	case 0x84: // RES 0,H
		value := gb.H()
		result := value & ^uint8(1<<0)
		gb.SetH(result)
		return 8

	case 0x85: // RES 0,L
		value := gb.L()
		result := value & ^uint8(1<<0)
		gb.SetL(result)
		return 8

	case 0x86: // RES 0,(HL)
		value := gb.Read8(gb.HL())
		result := value & ^uint8(1<<0)
		gb.Write8(gb.HL(), result)
		return 16

	case 0x87: // RES 0,A
		value := gb.A()
		result := value & ^uint8(1<<0)
		gb.SetA(result)
		return 8

	case 0x88: // RES 1,B
		value := gb.B()
		result := value & ^uint8(1<<1)
		gb.SetB(result)
		return 8

	case 0x89: // RES 1,C
		value := gb.C()
		result := value & ^uint8(1<<1)
		gb.SetC(result)
		return 8

	case 0x8a: // RES 1,D
		value := gb.D()
		result := value & ^uint8(1<<1)
		gb.SetD(result)
		return 8

	case 0x8b: // RES 1,E
		value := gb.E()
		result := value & ^uint8(1<<1)
		gb.SetE(result)
		return 8

	case 0x8c: // RES 1,H
		value := gb.H()
		result := value & ^uint8(1<<1)
		gb.SetH(result)
		return 8

	case 0x8d: // RES 1,L
		value := gb.L()
		result := value & ^uint8(1<<1)
		gb.SetL(result)
		return 8

	case 0x8e: // RES 1,(HL)
		value := gb.Read8(gb.HL())
		result := value & ^uint8(1<<1)
		gb.Write8(gb.HL(), result)
		return 16

	case 0x8f: // RES 1,A
		value := gb.A()
		result := value & ^uint8(1<<1)
		gb.SetA(result)
		return 8

	case 0x90: // RES 2,B
		value := gb.B()
		result := value & ^uint8(1<<2)
		gb.SetB(result)
		return 8

	case 0x91: // RES 2,C
		value := gb.C()
		result := value & ^uint8(1<<2)
		gb.SetC(result)
		return 8

	case 0x92: // RES 2,D
		value := gb.D()
		result := value & ^uint8(1<<2)
		gb.SetD(result)
		return 8

	case 0x93: // RES 2,E
		value := gb.E()
		result := value & ^uint8(1<<2)
		gb.SetE(result)
		return 8

	case 0x94: // RES 2,H
		value := gb.H()
		result := value & ^uint8(1<<2)
		gb.SetH(result)
		return 8

	case 0x95: // RES 2,L
		value := gb.L()
		result := value & ^uint8(1<<2)
		gb.SetL(result)
		return 8

	case 0x96: // RES 2,(HL)
		value := gb.Read8(gb.HL())
		result := value & ^uint8(1<<2)
		gb.Write8(gb.HL(), result)
		return 16

	case 0x97: // RES 2,A
		value := gb.A()
		result := value & ^uint8(1<<2)
		gb.SetA(result)
		return 8

	case 0x98: // RES 3,B
		value := gb.B()
		result := value & ^uint8(1<<3)
		gb.SetB(result)
		return 8

	case 0x99: // RES 3,C
		value := gb.C()
		result := value & ^uint8(1<<3)
		gb.SetC(result)
		return 8

	case 0x9a: // RES 3,D
		value := gb.D()
		result := value & ^uint8(1<<3)
		gb.SetD(result)
		return 8

	case 0x9b: // RES 3,E
		value := gb.E()
		result := value & ^uint8(1<<3)
		gb.SetE(result)
		return 8

	case 0x9c: // RES 3,H
		value := gb.H()
		result := value & ^uint8(1<<3)
		gb.SetH(result)
		return 8

	case 0x9d: // RES 3,L
		value := gb.L()
		result := value & ^uint8(1<<3)
		gb.SetL(result)
		return 8

	case 0x9e: // RES 3,(HL)
		value := gb.Read8(gb.HL())
		result := value & ^uint8(1<<3)
		gb.Write8(gb.HL(), result)
		return 16

	case 0x9f: // RES 3,A
		value := gb.A()
		result := value & ^uint8(1<<3)
		gb.SetA(result)
		return 8

	case 0xa0: // RES 4,B
		value := gb.B()
		result := value & ^uint8(1<<4)
		gb.SetB(result)
		return 8

	case 0xa1: // RES 4,C
		value := gb.C()
		result := value & ^uint8(1<<4)
		gb.SetC(result)
		return 8

	case 0xa2: // RES 4,D
		value := gb.D()
		result := value & ^uint8(1<<4)
		gb.SetD(result)
		return 8

	case 0xa3: // RES 4,E
		value := gb.E()
		result := value & ^uint8(1<<4)
		gb.SetE(result)
		return 8

	case 0xa4: // RES 4,H
		value := gb.H()
		result := value & ^uint8(1<<4)
		gb.SetH(result)
		return 8

	case 0xa5: // RES 4,L
		value := gb.L()
		result := value & ^uint8(1<<4)
		gb.SetL(result)
		return 8

	case 0xa6: // RES 4,(HL)
		value := gb.Read8(gb.HL())
		result := value & ^uint8(1<<4)
		gb.Write8(gb.HL(), result)
		return 16

	case 0xa7: // RES 4,A
		value := gb.A()
		result := value & ^uint8(1<<4)
		gb.SetA(result)
		return 8

	case 0xa8: // RES 5,B
		value := gb.B()
		result := value & ^uint8(1<<5)
		gb.SetB(result)
		return 8

	case 0xa9: // RES 5,C
		value := gb.C()
		result := value & ^uint8(1<<5)
		gb.SetC(result)
		return 8

	case 0xaa: // RES 5,D
		value := gb.D()
		result := value & ^uint8(1<<5)
		gb.SetD(result)
		return 8

	case 0xab: // RES 5,E
		value := gb.E()
		result := value & ^uint8(1<<5)
		gb.SetE(result)
		return 8

	case 0xac: // RES 5,H
		value := gb.H()
		result := value & ^uint8(1<<5)
		gb.SetH(result)
		return 8

	case 0xad: // RES 5,L
		value := gb.L()
		result := value & ^uint8(1<<5)
		gb.SetL(result)
		return 8

	case 0xae: // RES 5,(HL)
		value := gb.Read8(gb.HL())
		result := value & ^uint8(1<<5)
		gb.Write8(gb.HL(), result)
		return 16

	case 0xaf: // RES 5,A
		value := gb.A()
		result := value & ^uint8(1<<5)
		gb.SetA(result)
		return 8

	case 0xb0: // RES 6,B
		value := gb.B()
		result := value & ^uint8(1<<6)
		gb.SetB(result)
		return 8

	case 0xb1: // RES 6,C
		value := gb.C()
		result := value & ^uint8(1<<6)
		gb.SetC(result)
		return 8

	case 0xb2: // RES 6,D
		value := gb.D()
		result := value & ^uint8(1<<6)
		gb.SetD(result)
		return 8

	case 0xb3: // RES 6,E
		value := gb.E()
		result := value & ^uint8(1<<6)
		gb.SetE(result)
		return 8

	case 0xb4: // RES 6,H
		value := gb.H()
		result := value & ^uint8(1<<6)
		gb.SetH(result)
		return 8

	case 0xb5: // RES 6,L
		value := gb.L()
		result := value & ^uint8(1<<6)
		gb.SetL(result)
		return 8

	case 0xb6: // RES 6,(HL)
		value := gb.Read8(gb.HL())
		result := value & ^uint8(1<<6)
		gb.Write8(gb.HL(), result)
		return 16

	case 0xb7: // RES 6,A
		value := gb.A()
		result := value & ^uint8(1<<6)
		gb.SetA(result)
		return 8

	case 0xb8: // RES 7,B
		value := gb.B()
		result := value & ^uint8(1<<7)
		gb.SetB(result)
		return 8

	case 0xb9: // RES 7,C
		value := gb.C()
		result := value & ^uint8(1<<7)
		gb.SetC(result)
		return 8

	case 0xba: // RES 7,D
		value := gb.D()
		result := value & ^uint8(1<<7)
		gb.SetD(result)
		return 8

	case 0xbb: // RES 7,E
		value := gb.E()
		result := value & ^uint8(1<<7)
		gb.SetE(result)
		return 8

	case 0xbc: // RES 7,H
		value := gb.H()
		result := value & ^uint8(1<<7)
		gb.SetH(result)
		return 8

	case 0xbd: // RES 7,L
		value := gb.L()
		result := value & ^uint8(1<<7)
		gb.SetL(result)
		return 8

	case 0xbe: // RES 7,(HL)
		value := gb.Read8(gb.HL())
		result := value & ^uint8(1<<7)
		gb.Write8(gb.HL(), result)
		return 16

	case 0xbf: // RES 7,A
		value := gb.A()
		result := value & ^uint8(1<<7)
		gb.SetA(result)
		return 8

	case 0xc0: // SET 0,B
		value := gb.B()
		result := value | (1 << 0)
		gb.SetB(result)
		return 8

	case 0xc1: // SET 0,C
		value := gb.C()
		result := value | (1 << 0)
		gb.SetC(result)
		return 8

	case 0xc2: // SET 0,D
		value := gb.D()
		result := value | (1 << 0)
		gb.SetD(result)
		return 8

	case 0xc3: // SET 0,E
		value := gb.E()
		result := value | (1 << 0)
		gb.SetE(result)
		return 8

	case 0xc4: // SET 0,H
		value := gb.H()
		result := value | (1 << 0)
		gb.SetH(result)
		return 8

	case 0xc5: // SET 0,L
		value := gb.L()
		result := value | (1 << 0)
		gb.SetL(result)
		return 8

	case 0xc6: // SET 0,(HL)
		value := gb.Read8(gb.HL())
		result := value | (1 << 0)
		gb.Write8(gb.HL(), result)
		return 16

	case 0xc7: // SET 0,A
		value := gb.A()
		result := value | (1 << 0)
		gb.SetA(result)
		return 8

	case 0xc8: // SET 1,B
		value := gb.B()
		result := value | (1 << 1)
		gb.SetB(result)
		return 8

	case 0xc9: // SET 1,C
		value := gb.C()
		result := value | (1 << 1)
		gb.SetC(result)
		return 8

	case 0xca: // SET 1,D
		value := gb.D()
		result := value | (1 << 1)
		gb.SetD(result)
		return 8

	case 0xcb: // SET 1,E
		value := gb.E()
		result := value | (1 << 1)
		gb.SetE(result)
		return 8

	case 0xcc: // SET 1,H
		value := gb.H()
		result := value | (1 << 1)
		gb.SetH(result)
		return 8

	case 0xcd: // SET 1,L
		value := gb.L()
		result := value | (1 << 1)
		gb.SetL(result)
		return 8

	case 0xce: // SET 1,(HL)
		value := gb.Read8(gb.HL())
		result := value | (1 << 1)
		gb.Write8(gb.HL(), result)
		return 16

	case 0xcf: // SET 1,A
		value := gb.A()
		result := value | (1 << 1)
		gb.SetA(result)
		return 8

	case 0xd0: // SET 2,B
		value := gb.B()
		result := value | (1 << 2)
		gb.SetB(result)
		return 8

	case 0xd1: // SET 2,C
		value := gb.C()
		result := value | (1 << 2)
		gb.SetC(result)
		return 8

	case 0xd2: // SET 2,D
		value := gb.D()
		result := value | (1 << 2)
		gb.SetD(result)
		return 8

	case 0xd3: // SET 2,E
		value := gb.E()
		result := value | (1 << 2)
		gb.SetE(result)
		return 8

	case 0xd4: // SET 2,H
		value := gb.H()
		result := value | (1 << 2)
		gb.SetH(result)
		return 8

	case 0xd5: // SET 2,L
		value := gb.L()
		result := value | (1 << 2)
		gb.SetL(result)
		return 8

	case 0xd6: // SET 2,(HL)
		value := gb.Read8(gb.HL())
		result := value | (1 << 2)
		gb.Write8(gb.HL(), result)
		return 16

	case 0xd7: // SET 2,A
		value := gb.A()
		result := value | (1 << 2)
		gb.SetA(result)
		return 8

	case 0xd8: // SET 3,B
		value := gb.B()
		result := value | (1 << 3)
		gb.SetB(result)
		return 8

	case 0xd9: // SET 3,C
		value := gb.C()
		result := value | (1 << 3)
		gb.SetC(result)
		return 8

	case 0xda: // SET 3,D
		value := gb.D()
		result := value | (1 << 3)
		gb.SetD(result)
		return 8

	case 0xdb: // SET 3,E
		value := gb.E()
		result := value | (1 << 3)
		gb.SetE(result)
		return 8

	case 0xdc: // SET 3,H
		value := gb.H()
		result := value | (1 << 3)
		gb.SetH(result)
		return 8

	case 0xdd: // SET 3,L
		value := gb.L()
		result := value | (1 << 3)
		gb.SetL(result)
		return 8

	case 0xde: // SET 3,(HL)
		value := gb.Read8(gb.HL())
		result := value | (1 << 3)
		gb.Write8(gb.HL(), result)
		return 16

	case 0xdf: // SET 3,A
		value := gb.A()
		result := value | (1 << 3)
		gb.SetA(result)
		return 8

	case 0xe0: // SET 4,B
		value := gb.B()
		result := value | (1 << 4)
		gb.SetB(result)
		return 8

	case 0xe1: // SET 4,C
		value := gb.C()
		result := value | (1 << 4)
		gb.SetC(result)
		return 8

	case 0xe2: // SET 4,D
		value := gb.D()
		result := value | (1 << 4)
		gb.SetD(result)
		return 8

	case 0xe3: // SET 4,E
		value := gb.E()
		result := value | (1 << 4)
		gb.SetE(result)
		return 8

	case 0xe4: // SET 4,H
		value := gb.H()
		result := value | (1 << 4)
		gb.SetH(result)
		return 8

	case 0xe5: // SET 4,L
		value := gb.L()
		result := value | (1 << 4)
		gb.SetL(result)
		return 8

	case 0xe6: // SET 4,(HL)
		value := gb.Read8(gb.HL())
		result := value | (1 << 4)
		gb.Write8(gb.HL(), result)
		return 16

	case 0xe7: // SET 4,A
		value := gb.A()
		result := value | (1 << 4)
		gb.SetA(result)
		return 8

	case 0xe8: // SET 5,B
		value := gb.B()
		result := value | (1 << 5)
		gb.SetB(result)
		return 8

	case 0xe9: // SET 5,C
		value := gb.C()
		result := value | (1 << 5)
		gb.SetC(result)
		return 8

	case 0xea: // SET 5,D
		value := gb.D()
		result := value | (1 << 5)
		gb.SetD(result)
		return 8

	case 0xeb: // SET 5,E
		value := gb.E()
		result := value | (1 << 5)
		gb.SetE(result)
		return 8

	case 0xec: // SET 5,H
		value := gb.H()
		result := value | (1 << 5)
		gb.SetH(result)
		return 8

	case 0xed: // SET 5,L
		value := gb.L()
		result := value | (1 << 5)
		gb.SetL(result)
		return 8

	case 0xee: // SET 5,(HL)
		value := gb.Read8(gb.HL())
		result := value | (1 << 5)
		gb.Write8(gb.HL(), result)
		return 16

	case 0xef: // SET 5,A
		value := gb.A()
		result := value | (1 << 5)
		gb.SetA(result)
		return 8

	case 0xf0: // SET 6,B
		value := gb.B()
		result := value | (1 << 6)
		gb.SetB(result)
		return 8

	case 0xf1: // SET 6,C
		value := gb.C()
		result := value | (1 << 6)
		gb.SetC(result)
		return 8

	case 0xf2: // SET 6,D
		value := gb.D()
		result := value | (1 << 6)
		gb.SetD(result)
		return 8

	case 0xf3: // SET 6,E
		value := gb.E()
		result := value | (1 << 6)
		gb.SetE(result)
		return 8

	case 0xf4: // SET 6,H
		value := gb.H()
		result := value | (1 << 6)
		gb.SetH(result)
		return 8

	case 0xf5: // SET 6,L
		value := gb.L()
		result := value | (1 << 6)
		gb.SetL(result)
		return 8

	case 0xf6: // SET 6,(HL)
		value := gb.Read8(gb.HL())
		result := value | (1 << 6)
		gb.Write8(gb.HL(), result)
		return 16

	case 0xf7: // SET 6,A
		value := gb.A()
		result := value | (1 << 6)
		gb.SetA(result)
		return 8

	case 0xf8: // SET 7,B
		value := gb.B()
		result := value | (1 << 7)
		gb.SetB(result)
		return 8

	case 0xf9: // SET 7,C
		value := gb.C()
		result := value | (1 << 7)
		gb.SetC(result)
		return 8

	case 0xfa: // SET 7,D
		value := gb.D()
		result := value | (1 << 7)
		gb.SetD(result)
		return 8

	case 0xfb: // SET 7,E
		value := gb.E()
		result := value | (1 << 7)
		gb.SetE(result)
		return 8

	case 0xfc: // SET 7,H
		value := gb.H()
		result := value | (1 << 7)
		gb.SetH(result)
		return 8

	case 0xfd: // SET 7,L
		value := gb.L()
		result := value | (1 << 7)
		gb.SetL(result)
		return 8

	case 0xfe: // SET 7,(HL)
		value := gb.Read8(gb.HL())
		result := value | (1 << 7)
		gb.Write8(gb.HL(), result)
		return 16

	case 0xff: // SET 7,A
		value := gb.A()
		result := value | (1 << 7)
		gb.SetA(result)
		return 8

	default:
		panic(fmt.Sprintf("Unhandled CB instruction: %#x", opcode))
	}
}
