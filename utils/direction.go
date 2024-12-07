package utils

import (
	"container/ring"
)

type Direction struct {
	Row, Col int
}

func (d *Direction) X() int {
	return d.Col
}

func (d *Direction) Y() int {
	return d.Row
}

var (
	East      = &Direction{Row: 0, Col: 1}
	North     = &Direction{Row: -1, Col: 0}
	NorthEast = &Direction{Row: -1, Col: 1}
	NorthWest = &Direction{Row: -1, Col: -1}
	South     = &Direction{Row: 1, Col: 0}
	SouthEast = &Direction{Row: 1, Col: 1}
	SouthWest = &Direction{Row: 1, Col: -1}
	West      = &Direction{Row: 0, Col: -1}

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
