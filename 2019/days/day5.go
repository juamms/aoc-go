package days

import (
	"github.com/juamms/aoc-go/2019/intcode"
	"github.com/juamms/aoc-go/2019/utils"
)

type Day5 struct{}

var output int

func (day Day5) Part1() interface{} {
	program, err := utils.GetInputIntSlice(5, ",")
	utils.Handle(err)

	cpu := intcode.NewCPU(program)
	cpu.Input = &[]int{1}
	cpu.Output = &output
	cpu.Run()

	return output
}

func (day Day5) Part2() interface{} {
	program, err := utils.GetInputIntSlice(5, ",")
	utils.Handle(err)

	cpu := intcode.NewCPU(program)
	cpu.Input = &[]int{5}
	cpu.Output = &output
	cpu.Run()

	return output
}
