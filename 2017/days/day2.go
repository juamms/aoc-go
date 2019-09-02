package days

import (
	"math"
	"regexp"
	"sort"

	"github.com/juamms/aoc-go/2017/utils"
)

type Day2 struct{}

var regex = regexp.MustCompile(`\s`)

func diff(data []int) int {
	max := 0
	min := math.MaxInt64

	for _, n := range data {
		if n > max {
			max = n
		}

		if n < min {
			min = n
		}
	}

	return max - min
}

func div(data []int) int {
	sort.Sort(sort.Reverse(sort.IntSlice(data)))
	len := len(data)

	for i, n := range data {
		for j := i + 1; j < len; j++ {
			if n%data[j] == 0 {
				return n / data[j]
			}
		}
	}

	return 0
}

func (day Day2) Part1() interface{} {
	sum := 0

	utils.ScanInputFile(2, func(line string) {
		data := regex.Split(line, -1)
		arr := utils.ToIntSlice(data)
		sum += diff(arr)
	})

	return sum
}

func (day Day2) Part2() interface{} {
	sum := 0

	utils.ScanInputFile(2, func(line string) {
		data := regex.Split(line, -1)
		arr := utils.ToIntSlice(data)
		sum += div(arr)
	})

	return sum
}
