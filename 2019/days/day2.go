package days

import (
	"github.com/juamms/aoc-go/2019/intcode"
	"github.com/juamms/aoc-go/2019/utils"
)

type Day2 struct{}

func (day Day2) Part1() interface{} {
	program, err := utils.GetInputIntSlice(2, ",")
	utils.Handle(err)

	program[1] = 12
	program[2] = 2

	cpu := intcode.NewCPU(program)
	cpu.Run()

	return cpu.Result()
}

func (day Day2) Part2() interface{} {
	data, err := utils.GetInputIntSlice(2, ",")
	utils.Handle(err)

	program := make([]int, len(data))
	for noun := 0; noun <= 99; noun++ {
		for verb := 0; verb <= 99; verb++ {
			copy(program, data)

			program[1] = noun
			program[2] = verb

			cpu := intcode.NewCPU(program)
			cpu.Run()

			if cpu.Result() == 19690720 {
				return 100*noun + verb
			}
		}
	}

	return -1
}
