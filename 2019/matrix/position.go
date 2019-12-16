package matrix

// Directions
var (
	N                 = Position{-1, 0}
	S                 = Position{1, 0}
	W                 = Position{0, -1}
	E                 = Position{0, 1}
	NW                = Position{-1, -1}
	NE                = Position{-1, 1}
	SW                = Position{1, -1}
	SE                = Position{1, 1}
	AdjacentPositions = []Position{N, S, W, E, NW, NE, SW, SE}
)

// Position represents an (X,Y) location in a Matrix
type Position struct {
	X, Y int
}

// Adding returns a new Position with the sum of self and other's X and Y coordinates
func (pos Position) Adding(other Position) Position {
	return Position{
		X: pos.X + other.X,
		Y: pos.Y + other.Y,
	}
}
