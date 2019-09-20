package days

import (
	"github.com/juamms/aoc-go/2017/matrix"
	"github.com/juamms/aoc-go/2017/utils"

	u "github.com/juamms/go-utils"
)

type Day3 struct{}

var i = 1

func incrementalFiller(_ matrix.Position, _ *matrix.Matrix) int {
	val := i
	i++
	return val
}

func adjacentSumFiller(pos matrix.Position, matrix *matrix.Matrix) int {
	curr := matrix.Get(pos)
	if curr != 0 {
		return curr
	}

	return matrix.AdjacentSum(pos)
}

func (day Day3) Part1() interface{} {
	data, err := utils.GetInputString(3)
	utils.Handle(err)
	n := u.StringToInt(data, 0)

	// n/4 clips for smaller Matrices, but since n is a big number it works fine
	size := int(n / 4)
	m := matrix.NewMatrix(size)
	filler := matrix.NewSpiralFiller(m, false)
	m.Fill(filler, n, incrementalFiller)

	pos := m.PositionOf(n)
	d := m.DistanceBetween(m.Center(), pos)

	return d
}

func (day Day3) Part2() interface{} {
	data, err := utils.GetInputString(3)
	utils.Handle(err)
	n := u.StringToInt(data, 0)
	// n := 100

	size := int(n / 10)
	m := matrix.NewMatrix(size)
	m.Set(m.Center(), 1)
	filler := matrix.NewSpiralFiller(m, false)
	m.Fill(filler, n, adjacentSumFiller)

	return m.Max()
}
