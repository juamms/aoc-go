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

// Increment increments the value in the given position by `amount`
func (matrix *Matrix) Increment(pos Position, amount int) {
	matrix.Set(pos, matrix.Get(pos)+amount)
}

// IncrementLine icrements the value on an virtual line in the given direction by `amount`
// and returns the final position
func (matrix *Matrix) IncrementLine(start, direction Position, length, amount int) Position {
	pos := start

	matrix.WalkLine(pos, direction, length, func(p Position) bool {
		matrix.Increment(p, amount)
		pos = p

		return false
	})

	return pos
}

// WalkLine follows the line starting at `start`, going in `direction` for the given `length`
// and then executes the function at that position. Return true to stop the wal early.
func (matrix *Matrix) WalkLine(start, direction Position, length int, then func(Position) bool) {
	pos := start

	for ; length > 0; length-- {
		pos = pos.Adding(direction)

		if then(pos) {
			break
		}
	}
}

// Center returns the Matrix's center position
func (matrix *Matrix) Center() Position {
	size := float64(matrix.Size)
	center := int(math.Floor(size / 2))

	return Position{center, center}
}

// MaxPositions returns the positions with the largest value in the Matrix
func (matrix *Matrix) MaxPositions() []Position {
	max := 0

	for _, v := range matrix.data {
		if v > max {
			max = v
		}
	}

	positions := make([]Position, 0)

	for _, idx := range matrix.indicesOf(max) {
		positions = append(positions, matrix.positionForIndex(idx))
	}

	return positions
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

func (matrix *Matrix) indicesOf(n int) []int {
	indices := make([]int, 0)

	for i, v := range matrix.data {
		if v == n {
			indices = append(indices, i)
		}
	}

	return indices
}

func (matrix *Matrix) indexOf(n int) int {
	for i, v := range matrix.data {
		if v == n {
			return i
		}
	}

	return -1
}

// PositionOfIndex returns the position of `n` in the Matrix
func (matrix *Matrix) positionForIndex(idx int) Position {
	y := idx % matrix.Size
	x := int(idx / matrix.Size)

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
