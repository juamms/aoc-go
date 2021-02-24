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
	return 0
}
