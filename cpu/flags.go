package cpu

const (
	ZeroFlagBit      = 7
	SubtractFlagBit  = 6
	HalfCarryFlagBit = 5
	CarryFlagBit     = 4
)

type Flags struct {
	// Zero is set to 1 iff the last instruction resulted in 0
	Zero bool

	// Subtract is set to 1 when doing subtraction
	Subtract bool

	// HalfCarry is set to 1 when the 3rd bit overflows
	HalfCarry bool

	// Carry is set to 1 when the 7th bit overflows
	Carry bool
}

func (f Flags) Read() byte {
	var result byte
	if f.Zero {
		result &= (1 << ZeroFlagBit)
	}
	if f.Subtract {
		result &= (1 << SubtractFlagBit)
	}
	if f.HalfCarry {
		result &= (1 << HalfCarryFlagBit)
	}
	if f.Carry {
		result &= (1 << CarryFlagBit)
	}
	return result
}
