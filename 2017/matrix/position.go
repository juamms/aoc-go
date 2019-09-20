package matrix

// Directions
var (
	Up    = Position{-1, 0}
	Down  = Position{1, 0}
	Left  = Position{0, -1}
	Right = Position{0, 1}
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
