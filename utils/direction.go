package utils

// Direction represents the irreverence in Y and X
type Direction [2]int

func (d Direction) X() int {
	return d[1]
}

func (d Direction) Y() int {
	return d[0]
}

var (
	East      = Direction{0, 1}
	North     = Direction{-1, 0}
	NorthEast = Direction{-1, 1}
	NorthWest = Direction{-1, -1}
	South     = Direction{1, 0}
	SouthEast = Direction{1, 1}
	SouthWest = Direction{1, -1}
	West      = Direction{0, -1}

	AzimuthDirections = []Direction{North, East, South, West}
	AllDirections     = []Direction{North, NorthEast, East, SouthEast, South, SouthWest, West, NorthWest}

	SymbolDirectionIdx = map[byte]int{
		'^': 0, // North
		'>': 1, // East
		'v': 2, // South
		'<': 3, // West
	}
)

func RotateAzimuthRight(idx int) int {
	return (idx + 1) % len(AzimuthDirections)
}

func RotateAzimuthLeft(idx int) int {
	return (idx - 1 + len(AzimuthDirections)) % len(AzimuthDirections)
}

type Pos [2]int

func NewPos(y, x int) Pos {
	return [2]int{y, x}
}

func (p Pos) Eq(v Pos) bool {
	return p[1] == v[1] && p[0] == v[0]
}

func (p Pos) X() int {
	return p[1]
}

func (p Pos) Y() int {
	return p[0]
}

func (p Pos) Next(d Direction) Pos {
	return NewPos(p.Y()+d.Y(), p.X()+d.X())
}

func (p Pos) Values() (int, int) {
	return p.Y(), p.X()
}
