package intcode

import "fmt"

// CPU represents an Intcode CPU
type CPU struct {
	memory []int
	ip     int
}

// NewCPU returns a new CPU with the given program loaded into memory
func NewCPU(program []int) *CPU {
	cpu := &CPU{
		memory: program,
		ip:     0,
	}

	return cpu
}

func (cpu *CPU) currentOpcode() *int {
	return &cpu.memory[cpu.ip]
}

// IsHalted checks if the CPU is in the halted state
func (cpu *CPU) IsHalted() bool {
	return *cpu.currentOpcode() == 99
}

// Step decodes and executes the current instruction and jumps the
// instruction pointer to the next instruction
func (cpu *CPU) Step() {
	op := cpu.currentOpcode()

	switch *op {
	case 1:
		cpu.add()
	case 2:
		cpu.mul()
	default:
		panic(fmt.Sprintf("Unsuported opcode: %d", *op))
	}

	cpu.ip++
}

// Run performs all instrunctions until the CPU is halted
func (cpu *CPU) Run() {
	for !cpu.IsHalted() {
		cpu.Step()
	}
}

// Dump prints a representation of the CPU's internal state
func (cpu *CPU) Dump() {
	fmt.Printf("%v\nOpcode [%d]: %d\n", cpu.memory, cpu.ip, *cpu.currentOpcode())
}

// Result returns the value at the first position in the program
func (cpu *CPU) Result() int {
	return cpu.memory[0]
}

func (cpu *CPU) getParameters(amount int) []int {
	ip := cpu.ip
	params := make([]int, 0)

	for idx := 1; idx <= amount; idx++ {
		params = append(params, cpu.memory[ip+idx])
		cpu.ip++
	}

	return params
}

func (cpu *CPU) add() {
	params := cpu.getParameters(3)
	lv, rv, out := params[0], params[1], params[2]

	cpu.memory[out] = cpu.memory[lv] + cpu.memory[rv]
}

func (cpu *CPU) mul() {
	params := cpu.getParameters(3)
	lv, rv, out := params[0], params[1], params[2]

	cpu.memory[out] = cpu.memory[lv] * cpu.memory[rv]
}
