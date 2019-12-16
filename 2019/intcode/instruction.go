package intcode

// Mode represents a paramater mode for the CPU
type Mode int

// The available parameter modes
const (
	ModePosition  = Mode(0)
	ModeImmediate = Mode(1)
)

// Instruction represents a Intcode CPU instruction
type Instruction struct {
	Opcode int
	Modes  []Mode
}

func decode(raw int, into *Instruction) {
	into.Opcode = raw % 100
	raw = raw / 100

	for i := 0; i < 3; i++ {
		into.Modes[i] = Mode(raw % 10)
		raw = raw / 10
	}
}
