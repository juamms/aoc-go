package days

import (
	"math"

	"github.com/juamms/aoc-go/2019/utils"
)

type Day1 struct{}

func fuelFor(mass int) int {
	fuel := int(math.Floor(float64(mass)/3) - 2)

	if fuel < 0 {
		return 0
	}

	return fuel
}

func (day Day1) Part1() interface{} {
	data, err := utils.GetInputIntSlice(1, "\n")
	utils.Handle(err)

	fuel := 0
	for _, mass := range data {
		fuel += fuelFor(mass)
	}

	return fuel
}

func (day Day1) Part2() interface{} {
	data, err := utils.GetInputIntSlice(1, "\n")
	utils.Handle(err)

	totalFuel := 0
	for _, mass := range data {
		tempFuel := fuelFor(mass)
		totalFuel += tempFuel

		for tempFuel > 0 {
			tempFuel = fuelFor(tempFuel)
			totalFuel += tempFuel
		}
	}

	return totalFuel
}
