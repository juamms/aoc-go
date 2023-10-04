package intcode

import "fmt"

// CPU represents an Intcode CPU
type CPU struct {
	memory []int
	Input  *[]int
	Output *int
	ins    *Instruction
	ip     int
}

// NewCPU returns a new CPU with the given program loaded into memory
func NewCPU(program []int) *CPU {
	ins := &Instruction{
		Modes: make([]Mode, 3),
	}
	cpu := &CPU{
		memory: program,
		Input:  nil,
		Output: nil,
		ins:    ins,
		ip:     0,
	}

	return cpu
}

func (cpu *CPU) rawInstruction() *int {
	return &cpu.memory[cpu.ip]
}

// IsHalted checks if the CPU is in the halted state
func (cpu *CPU) IsHalted() bool {
	return *cpu.rawInstruction() == 99
}

// Step decodes and executes the current instruction and jumps the
// instruction pointer to the next instruction
func (cpu *CPU) Step() {
	raw := cpu.rawInstruction()
	decode(*raw, cpu.ins)

	switch cpu.ins.Opcode {
	case 1:
		cpu.add()
	case 2:
		cpu.mul()
	case 3:
		cpu.in()
	case 4:
		cpu.out()
	case 5:
		cpu.jt()
	case 6:
		cpu.jf()
	case 7:
		cpu.lt()
	case 8:
		cpu.eq()
	default:
		panic(fmt.Sprintf("Unsuported opcode: %d", cpu.ins.Opcode))
	}

	cpu.ip++
}

// Run performs all instrunctions until the CPU is halted
func (cpu *CPU) Run() {
	for !cpu.IsHalted() {
		cpu.Step()
	}
}

// RunUntil performs all instructions until `opcode` is reached or the CPU is halted
func (cpu *CPU) RunUntil(opcode int) {
	for !cpu.IsHalted() {

		ins := cpu.ins
		cpu.Step()

		if ins.Opcode == opcode {
			break
		}
	}
}

// Dump prints a representation of the CPU's internal state
func (cpu *CPU) Dump() {
	fmt.Printf("%v\nInstruction [%d]: %d\n", cpu.memory, cpu.ip, cpu.ins)
}

// Result returns the value at the first position in the program
func (cpu *CPU) Result() int {
	return cpu.memory[0]
}

func (cpu *CPU) getParam() int {
	cpu.ip++
	return cpu.memory[cpu.ip]
}

func (cpu *CPU) getParams(amount int) []int {
	ip := cpu.ip
	params := make([]int, 0)

	for idx := 1; idx <= amount; idx++ {
		params = append(params, cpu.memory[ip+idx])
		cpu.ip++
	}

	return params
}

func (cpu *CPU) getValue(param int) int {
	if cpu.ins.Modes[0] == ModePosition {
		return cpu.memory[param]
	}

	return param
}

func (cpu *CPU) getValues(parameters []int) []int {
	values := make([]int, len(parameters))

	for idx, p := range parameters {
		val := p

		if cpu.ins.Modes[idx] == ModePosition {
			val = cpu.memory[val]
		}

		values[idx] = val
	}

	return values
}

func (cpu *CPU) add() {
	params := cpu.getParams(3)
	values := cpu.getValues(params)
	lv, rv := values[0], values[1]
	out := params[2]

	cpu.memory[out] = lv + rv
}

func (cpu *CPU) mul() {
	params := cpu.getParams(3)
	values := cpu.getValues(params)
	lv, rv := values[0], values[1]
	out := params[2]

	cpu.memory[out] = lv * rv
}

func (cpu *CPU) in() {
	param := cpu.getParam()
	var value int

	if cpu.Input != nil {
		value = (*cpu.Input)[0]
		*cpu.Input = (*cpu.Input)[1:]
	} else {
		_, err := fmt.Scanf("%d", &value)

		if err != nil {
			panic(err)
		}

	}

	cpu.memory[param] = value
}

func (cpu *CPU) out() {
	param := cpu.getParam()
	output := cpu.memory[param]

	if cpu.Output != nil {
		*cpu.Output = output
	} else {
		fmt.Println(cpu.memory[param])
	}
}

func (cpu *CPU) jt() {
	params := cpu.getParams(2)
	values := cpu.getValues(params)

	if values[0] != 0 {
		// We subtract 1 here cause the Step method automatically increments the ip by 1
		cpu.ip = values[1] - 1
	}
}

func (cpu *CPU) jf() {
	params := cpu.getParams(2)
	values := cpu.getValues(params)

	if values[0] == 0 {
		// We subtract 1 here cause the Step method automatically increments the ip by 1
		cpu.ip = values[1] - 1
	}
}

func (cpu *CPU) lt() {
	params := cpu.getParams(3)
	values := cpu.getValues(params)
	out := params[2]

	if values[0] < values[1] {
		cpu.memory[out] = 1
	} else {
		cpu.memory[out] = 0
	}
}

func (cpu *CPU) eq() {
	params := cpu.getParams(3)
	values := cpu.getValues(params)
	out := params[2]

	if values[0] == values[1] {
		cpu.memory[out] = 1
	} else {
		cpu.memory[out] = 0
	}
}
