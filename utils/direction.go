package utils

import (
	"container/ring"
)

type Direction struct {
	Y, X int
	name string
}

func (d *Direction) String() string {
	return d.name
}

var (
	East      = &Direction{Y: 0, X: 1, name: "East"}
	North     = &Direction{Y: -1, X: 0, name: "North"}
	NorthEast = &Direction{Y: -1, X: 1, name: "NorthEast"}
	NorthWest = &Direction{Y: -1, X: -1, name: "NorthWest"}
	South     = &Direction{Y: 1, X: 0, name: "South"}
	SouthEast = &Direction{Y: 1, X: 1, name: "SouthEast"}
	SouthWest = &Direction{Y: 1, X: -1, name: "SouthWest"}
	West      = &Direction{Y: 0, X: -1, name: "West"}

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

type Pos [2]int

func NewPos(y, x int) Pos {
	return [2]int{y, x}
}

func (p Pos) X() int {
	return p[1]
}

func (p Pos) Y() int {
	return p[0]
}
