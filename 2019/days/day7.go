package days

import (
	"strings"

	"github.com/juamms/aoc-go/2019/intcode"
	"github.com/juamms/aoc-go/2019/utils"
)

type Day7 struct{}

func (day Day7) Part1() interface{} {
	source, err := utils.GetInputIntSlice(7, ",")
	utils.Handle(err)

	program := make([]int, len(source))
	prevOutput := 0
	maxOutput := 0

	utils.ScanInputFile(71, func(raw string) {
		sequence := utils.ToIntSlice(strings.Split(raw, ","))

		for _, v := range sequence {
			copy(program, source)

			cpu := intcode.NewCPU(program)
			cpu.Input = &[]int{v, prevOutput}
			cpu.Output = &prevOutput
			cpu.Run()
		}

		if prevOutput > maxOutput {
			maxOutput = prevOutput
		}

		prevOutput = 0
	})

	return maxOutput
}

func (day Day7) Part2() interface{} {
	source, err := utils.GetInputIntSlice(7, ",")
	utils.Handle(err)

	maxOutput := 0

	utils.ScanInputFile(72, func(raw string) {
		sequence := utils.ToIntSlice(strings.Split(raw, ","))

		output := 0
		cpus := make([]*intcode.CPU, 5)

		for i := range cpus {
			program := make([]int, len(source))
			copy(program, source)
			cpus[i] = intcode.NewCPU(program)

			input := []int{sequence[i]}
			cpus[i].Input = &input
			cpus[i].Output = &output
		}

		curCPU := 0

		for {
			cpu := cpus[curCPU]
			input := append(*cpu.Input, output)
			cpu.Input = &input

			cpu.RunUntil(4) // output
			curCPU++

			if curCPU == 5 {
				curCPU = 0
			}

			if cpu.IsHalted() {
				break
			}
		}

		if output > maxOutput {
			maxOutput = output
		}
	})

	return maxOutput
}
