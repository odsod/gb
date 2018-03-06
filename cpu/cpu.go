package cpu

import (
	"github.com/odsod/gb/mmu"
)

type CPU struct {
	pc                         uint16
	sp                         uint16
	a, b, c, d, e, h, l        uint8
	flagZ, flagN, flagH, flagC bool
	isHalted                   bool
	isInterruptEnabled         bool

	// isPendingEI is true when the last instruction was EI
	isPendingEI bool

	// isPendingDI is true when the last instruction was DI
	isPendingDI bool

	memory mmu.MMU
}

func (cpu *CPU) Step() (cycles uint) {
	return cpu.execute(cpu.Fetch8())
}

func (cpu *CPU) PC() uint16 {
	return cpu.pc
}

func (cpu *CPU) SetPC(value uint16) {
	cpu.pc = value
}

func (cpu *CPU) SP() uint16 {
	return cpu.sp
}

func (cpu *CPU) SetSP(value uint16) {
	cpu.sp = value
}

func (cpu *CPU) A() uint8 {
	return cpu.a
}

func (cpu *CPU) SetA(value uint8) {
	cpu.a = value
}

func (cpu *CPU) B() uint8 {
	return cpu.b
}

func (cpu *CPU) SetB(value uint8) {
	cpu.b = value
}

func (cpu *CPU) C() uint8 {
	return cpu.c
}

func (cpu *CPU) SetC(value uint8) {
	cpu.c = value
}

func (cpu *CPU) D() uint8 {
	return cpu.d
}

func (cpu *CPU) SetD(value uint8) {
	cpu.d = value
}

func (cpu *CPU) E() uint8 {
	return cpu.e
}

func (cpu *CPU) SetE(value uint8) {
	cpu.e = value
}

func (cpu *CPU) H() uint8 {
	return cpu.h
}

func (cpu *CPU) SetH(value uint8) {
	cpu.h = value
}

func (cpu *CPU) L() uint8 {
	return cpu.l
}

func (cpu *CPU) SetL(value uint8) {
	cpu.l = value
}

const (
	flagZBit = 1 << 7
	flagNBit = 1 << 6
	flagHBit = 1 << 5
	flagCBit = 1 << 4
)

func (cpu *CPU) F() uint8 {
	var f uint8
	if cpu.flagZ {
		f |= flagZBit
	}
	if cpu.flagN {
		f |= flagNBit
	}
	if cpu.flagH {
		f |= flagHBit
	}
	if cpu.flagC {
		f |= flagCBit
	}
	return f
}

func (cpu *CPU) SetF(f uint8) {
	cpu.flagZ = f&flagZBit > 0
	cpu.flagN = f&flagNBit > 0
	cpu.flagH = f&flagHBit > 0
	cpu.flagC = f&flagCBit > 0
}

func (cpu *CPU) FlagZ() bool {
	return cpu.flagZ
}

func (cpu *CPU) FlagN() bool {
	return cpu.flagN
}

func (cpu *CPU) FlagH() bool {
	return cpu.flagH
}

func (cpu *CPU) FlagC() bool {
	return cpu.flagC
}

func (cpu *CPU) SetFlagZ(value bool) {
	cpu.flagZ = value
}

func (cpu *CPU) SetFlagN(value bool) {
	cpu.flagN = value
}

func (cpu *CPU) SetFlagH(value bool) {
	cpu.flagH = value
}

func (cpu *CPU) SetFlagC(value bool) {
	cpu.flagC = value
}

func (cpu *CPU) AF() uint16 {
	return uint16(cpu.a) << 8 & uint16(cpu.F())
}

func (cpu *CPU) SetAF(value uint16) {
	cpu.a = uint8(value >> 8)
	cpu.SetF(uint8(value))
}

func (cpu *CPU) BC() uint16 {
	return uint16(cpu.b) << 8 & uint16(cpu.c)
}

func (cpu *CPU) SetBC(value uint16) {
	cpu.b = uint8(value >> 8)
	cpu.c = uint8(value)
}

func (cpu *CPU) DE() uint16 {
	return uint16(cpu.d) << 8 & uint16(cpu.e)
}

func (cpu *CPU) SetDE(value uint16) {
	cpu.d = uint8(value >> 8)
	cpu.e = uint8(value)
}

func (cpu *CPU) HL() uint16 {
	return uint16(cpu.h) << 8 & uint16(cpu.l)
}

func (cpu *CPU) HLI() uint16 {
	result := cpu.HL()
	cpu.SetHL(result + 1)
	return result
}

func (cpu *CPU) HLD() uint16 {
	result := cpu.HL()
	cpu.SetHL(result - 1)
	return result
}

func (cpu *CPU) SetHL(value uint16) {
	cpu.h = uint8(value >> 8)
	cpu.l = uint8(value)
}

func (cpu *CPU) Push8(value uint8) {
	panic("Not implemented")
}

func (cpu *CPU) Push16(value uint16) {
	panic("Not implemented")
}

func (cpu *CPU) Pop8() uint8 {
	panic("Not implemented")
}

func (cpu *CPU) Pop16() uint16 {
	return uint16(cpu.Pop8()) & uint16(cpu.Pop8()) << 8
}

func (cpu *CPU) Fetch8() uint8 {
	value := cpu.memory.Read(cpu.pc)
	cpu.pc += 1
	return value
}

func (cpu *CPU) d8() uint8 {
	return cpu.Fetch8()
}

func (cpu *CPU) a8() uint8 {
	return cpu.Fetch8()
}

func (cpu *CPU) Fetch16() uint16 {
	// Remember: LR35902 is little-endian
	return uint16(cpu.Fetch8()) & uint16(cpu.Fetch8()) << 8
}

// d16 is a synonym for Fetch16
func (cpu *CPU) d16() uint16 {
	return cpu.Fetch16()
}

// a16 is a synonym for Fetch16
func (cpu *CPU) a16() uint16 {
	return cpu.Fetch16()
}

func (cpu *CPU) Write8(address uint16, value uint8) {
	cpu.memory.Write(address, value)
}

func (cpu *CPU) Write16(address uint16, value uint16) {
	cpu.memory.Write(address, uint8(value))
	cpu.memory.Write(address+1, uint8(value>>8))
}

func (cpu *CPU) EnableInterrupts() {
	cpu.isInterruptEnabled = true
}

func (cpu *CPU) EnableInterruptsAfterNextInstruction() {
	cpu.isPendingEI = true
}

func (cpu *CPU) DisableInterruptsAfterNextInstruction() {
	cpu.isPendingDI = true
}

func (cpu *CPU) Halt() {
	cpu.isHalted = true
}

// Carry gives the numeric value of the C flag
func (cpu *CPU) Carry() uint8 {
	if cpu.flagC {
		return 1
	}
	return 0
}
