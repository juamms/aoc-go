package days

import (
	"github.com/juamms/aoc-go/2017/utils"
)

type Day5 struct{}

func (day Day5) Part1() interface{} {
	data, err := utils.GetInputIntSlice(5, "\n")
	utils.Handle(err)

	pointer := 0
	steps := 0

	for pointer >= 0 && pointer < len(data) {
		jump := data[pointer]
		data[pointer]++
		pointer += jump
		steps++
	}

	return steps
}

func (day Day5) Part2() interface{} {
	data, err := utils.GetInputIntSlice(5, "\n")
	utils.Handle(err)

	pointer := 0
	steps := 0

	for pointer >= 0 && pointer < len(data) {
		jump := data[pointer]

		if jump >= 3 {
			data[pointer]--
		} else {
			data[pointer]++
		}

		pointer += jump
		steps++
	}

	return steps
}
