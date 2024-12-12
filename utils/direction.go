package utils

import (
	"container/ring"
)

type Direction struct {
	Row, Col int
	name     string
}

func (d *Direction) X() int {
	return d.Col
}

func (d *Direction) Y() int {
	return d.Row
}

func (d *Direction) String() string {
	return d.name
}

var (
	East      = &Direction{Row: 0, Col: 1, name: "East"}
	North     = &Direction{Row: -1, Col: 0, name: "North"}
	NorthEast = &Direction{Row: -1, Col: 1, name: "NorthEast"}
	NorthWest = &Direction{Row: -1, Col: -1, name: "NorthWest"}
	South     = &Direction{Row: 1, Col: 0, name: "South"}
	SouthEast = &Direction{Row: 1, Col: 1, name: "SouthEast"}
	SouthWest = &Direction{Row: 1, Col: -1, name: "SouthWest"}
	West      = &Direction{Row: 0, Col: -1, name: "West"}

	AzimuthDirections = []*Direction{North, East, South, West}
	AllDirections     = []*Direction{North, NorthEast, East, SouthEast, South, SouthWest, West, NorthWest}
)

func NewAzimuthRing(pointingTo *Direction) *ring.Ring {
	r := ring.New(len(AzimuthDirections))

	for _, d := range AzimuthDirections {
		r.Value = d
		r = r.Next()
	}

	if pointingTo != nil {
		for {
			if r.Value.(*Direction) == pointingTo {
				break
			}
			r = r.Next()
		}
	}

	return r
}

type Position [2]int

func NewPosition(r, c int) Position {
	return [2]int{r, c}
}
