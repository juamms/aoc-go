package matrix

// Filler is an interface representing a filling algorithm implementation
type Filler interface {
	Origin() Position
	Next() Position
}

// SpiralFiller fills the Matrix starting from the center and walking either
// clockwise or counter-clockwise
type SpiralFiller struct {
	clockwise         bool
	previousDirection Position
	current           Position
	matrix            *Matrix
}

// NewSpiralFiller returns a new SpiralFiller
func NewSpiralFiller(matrix *Matrix, clockwise bool) *SpiralFiller {
	filler := &SpiralFiller{
		clockwise: clockwise,
		matrix:    matrix,
	}
	filler.current = filler.Origin()

	if clockwise {
		filler.previousDirection = Up
	} else {
		filler.previousDirection = Down
	}

	return filler
}

// Origin returns this filler's starting position
func (filler SpiralFiller) Origin() Position {
	return filler.matrix.Center()
}

func (filler *SpiralFiller) nextDirection() Position {
	if filler.clockwise {
		switch filler.previousDirection {
		case Up:
			return Right
		case Right:
			return Down
		case Down:
			return Left
		default:
			return Up
		}
	} else {
		switch filler.previousDirection {
		case Up:
			return Left
		case Left:
			return Down
		case Down:
			return Right
		default:
			return Up
		}
	}
}

// Next returns this filler's next position
func (filler *SpiralFiller) Next() Position {
	nextDirection := filler.nextDirection()
	nextPos := filler.current.Adding(nextDirection)

	if filler.matrix.Get(nextPos) == 0 {
		filler.current = nextPos
		filler.previousDirection = nextDirection
	} else {
		nextPos = filler.current.Adding(filler.previousDirection)
		filler.current = nextPos
	}

	return nextPos
}
