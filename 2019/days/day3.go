package days

import (
	"math"
	"strings"

	"github.com/juamms/aoc-go/2019/matrix"
	"github.com/juamms/aoc-go/2019/utils"
	u "github.com/juamms/go-utils"
)

type Day3 struct{}

func loadData() (*matrix.Matrix, []string, []string) {
	m := matrix.NewMatrix(999)
	m.Set(m.Center(), -1)

	wires := make([][]string, 0)

	utils.ScanInputFile(3, func(data string) {
		wires = append(wires, strings.Split(data, ","))
	})

	drawWire(wires[0], m, 1)
	drawWire(wires[1], m, 10)

	return m, wires[0], wires[1]
}

func drawWire(wire []string, m *matrix.Matrix, amount int) {
	pos := m.Center()
	for _, ins := range wire {
		var dir matrix.Position
		switch ins[0] {
		case 'R':
			dir = matrix.E
		case 'D':
			dir = matrix.S
		case 'L':
			dir = matrix.W
		case 'U':
			dir = matrix.N
		}

		length := u.StringToInt(ins[1:], 0)
		pos = m.IncrementLine(pos, dir, length, amount)
	}
}

func (day Day3) Part1() interface{} {
	matrix, _, _ := loadData()

	maxPositions := matrix.MaxPositions()
	min := math.MaxInt64

	for _, pos := range maxPositions {
		distance := matrix.DistanceBetween(matrix.Center(), pos)

		if distance < min {
			min = distance
		}
	}

	return min
}

func stepsToPosition(wire []string, m *matrix.Matrix, target matrix.Position) int {
	pos := m.Center()
	steps := 0
	found := false

	for _, ins := range wire {
		var dir matrix.Position
		switch ins[0] {
		case 'R':
			dir = matrix.E
		case 'D':
			dir = matrix.S
		case 'L':
			dir = matrix.W
		case 'U':
			dir = matrix.N
		}

		length := u.StringToInt(ins[1:], 0)

		m.WalkLine(pos, dir, length, func(p matrix.Position) bool {
			pos = p
			steps++

			if p == target {
				found = true
				return true
			}

			return false
		})

		if found {
			break
		}
	}

	return steps
}

func (day Day3) Part2() interface{} {
	mx, wire1, wire2 := loadData()

	steps := make([]int, 0)

	positions := mx.MaxPositions()

	for _, pos := range positions {
		stepsWire1 := stepsToPosition(wire1, mx, pos)
		stepsWire2 := stepsToPosition(wire2, mx, pos)

		steps = append(steps, stepsWire1+stepsWire2)
	}

	totalSteps := math.MaxInt64
	for _, steps := range steps {
		if steps < totalSteps {
			totalSteps = steps
		}
	}

	return totalSteps
}
