package days

import (
	"github.com/juamms/aoc-go/2019/utils"
)

type Day4 struct{}

func hasDoubles(pw int) bool {
	prev := -1

	for pw != 0 {
		n := pw % 10
		pw = pw / 10

		if prev == n {
			return true
		}

		prev = n
	}

	return false
}

func doesNotDecrease(pw int) bool {
	prev := 9

	for pw != 0 {
		n := pw % 10
		pw = pw / 10

		if n > prev {
			return false
		}

		prev = n
	}

	return true
}

func hasValidDoubles(pw int) bool {
	doubles := make(map[int]int)

	for pw != 0 {
		n := pw % 10
		pw = pw / 10

		if _, ok := doubles[n]; ok {
			doubles[n]++
		} else {
			doubles[n] = 1
		}
	}

	for _, v := range doubles {
		if v == 2 {
			return true
		}
	}

	return false
}

func (day Day4) Part1() interface{} {
	passwordRange, err := utils.GetInputIntSlice(4, "-")
	utils.Handle(err)

	validPasswords := 0
	for i := passwordRange[0]; i <= passwordRange[1]; i++ {
		if hasDoubles(i) && doesNotDecrease(i) {
			validPasswords++
		}

	}

	return validPasswords
}

func (day Day4) Part2() interface{} {
	passwordRange, err := utils.GetInputIntSlice(4, "-")
	utils.Handle(err)

	validPasswords := 0
	for i := passwordRange[0]; i <= passwordRange[1]; i++ {
		if doesNotDecrease(i) && hasValidDoubles(i) {
			validPasswords++
		}
	}

	return validPasswords
}
