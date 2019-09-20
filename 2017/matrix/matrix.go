package matrix

import (
	"fmt"
	"math"
)

// Matrix represents a Matrix
type Matrix struct {
	Size int
	area int
	data []int
}

// NewMatrix returns a Matrix with `size`x`size` area
func NewMatrix(size int) *Matrix {
	if size <= 0 {
		size = 1
	}

	area := size * size

	return &Matrix{
		Size: size,
		area: area,
		data: make([]int, area),
	}
}

func (matrix *Matrix) isPositionInRange(pos Position) bool {
	return pos.X >= 0 && pos.X < matrix.Size && pos.Y >= 0 && pos.Y < matrix.Size
}

// Get returns the value for the given position
func (matrix *Matrix) Get(pos Position) int {
	if !matrix.isPositionInRange(pos) {
		return 0
	}

	return matrix.data[(pos.X*matrix.Size)+pos.Y]
}

// Set sets the given `val` in the given position
func (matrix *Matrix) Set(pos Position, val int) {
	if matrix.isPositionInRange(pos) {
		matrix.data[(pos.X*matrix.Size)+pos.Y] = val
	}
}

// Center returns the Matrix's center position
func (matrix *Matrix) Center() Position {
	size := float64(matrix.Size)
	center := int(math.Floor(size / 2))

	return Position{center, center}
}

// Max returns the largest value in the Matrix
func (matrix *Matrix) Max() int {
	max := 0

	for _, v := range matrix.data {
		if v > max {
			max = v
		}
	}

	return max
}

// Fill fills the Matrix using a Filler object, with the value given by `valueForPosition` until `maxValue` or greater is reached
func (matrix *Matrix) Fill(filler Filler, maxValue int, valueForPosition func(Position, *Matrix) int) {
	pos := filler.Origin()

	for i := 0; i < matrix.area; i++ {
		val := valueForPosition(pos, matrix)
		matrix.Set(pos, val)

		if val >= maxValue {
			break
		}

		pos = filler.Next()
	}
}

// from http://cavaliercoder.com/blog/optimized-abs-for-int64-in-go.html
func abs(n int) int {
	y := n >> 63
	return (n ^ y) - y
}

// DistanceBetween returns the Manhattan Distance between `pos1` and `pos2`
func (matrix *Matrix) DistanceBetween(pos1, pos2 Position) int {
	if matrix.isPositionInRange(pos1) && matrix.isPositionInRange(pos2) {
		return abs(pos1.X-pos2.X) + abs(pos1.Y-pos2.Y)
	}

	return -1
}

func (matrix *Matrix) indexOf(n int) int {
	for i, v := range matrix.data {
		if v == n {
			return i
		}
	}

	return -1
}

// PositionOf returs the position of `n` in the Matrix
func (matrix *Matrix) PositionOf(n int) Position {
	i := matrix.indexOf(n)
	y := i % matrix.Size
	x := int(i / matrix.Size)

	return Position{x, y}
}

// AdjacentSum returns the um of all values adjacent to `pos`
func (matrix *Matrix) AdjacentSum(pos Position) int {
	sum := 0

	for _, p := range AdjacentPositions {
		adj := pos.Adding(p)
		sum += matrix.Get(adj)
	}

	return sum
}

func (matrix *Matrix) String() string {
	str := ""

	pos := Position{0, 0}

	for pos.X < matrix.Size {
		for pos.Y < matrix.Size {
			str += fmt.Sprintf("%3d ", matrix.Get(pos))

			pos.Y++
		}

		str += "\n"
		pos.X++
		pos.Y = 0
	}

	return str
}
