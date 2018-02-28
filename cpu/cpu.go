package cpu

type CPU struct {
	// PC is the Program Counter
	PC uint16

	// SP is the Stack Pointer
	SP uint16

	// Registers are the registers
	Registers Registers
}
