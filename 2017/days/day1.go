package days

import "github.com/juamms/aoc-go/2017/utils"

type Day1 struct{}

// Faster solution
func (day Day1) solveDirect(captcha []int) int {
	captcha = append(captcha, captcha[0])

	sum := 0
	i := 0

	for i < len(captcha)-1 {
		n := captcha[i]
		m := captcha[i+1]

		if n == m {
			sum += n
		}

		i++
	}

	return sum
}

// Slower, but uses the circular get method
func (day Day1) solveCircular(captcha []int) int {
	sum := 0
	i := 0

	for i < len(captcha) {
		n := utils.CircularGetInt(captcha, i)
		m := utils.CircularGetInt(captcha, i+1)

		if n == m {
			sum += n
		}

		i++
	}

	return sum
}

func (day Day1) Part1() interface{} {
	captcha, err := utils.GetInputIntSlice(1, "")
	utils.Handle(err)

	return day.solveDirect(captcha)
}

func (day Day1) Part2() interface{} {
	captcha, err := utils.GetInputIntSlice(1, "")
	utils.Handle(err)

	anchor := len(captcha) / 2
	sum := 0
	i := 0

	for i < len(captcha) {
		n := captcha[i]
		m := utils.CircularGetInt(captcha, i+anchor)

		if n == m {
			sum += n
		}

		i++
	}

	return sum
}
