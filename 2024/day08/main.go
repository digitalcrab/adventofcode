package main

import (
	_ "embed"
	"fmt"
	"math"
	"strings"

	"github.com/digitalcrab/adventofcode/utils"
)

func CalculateUniqueAntiNodes(in [][]byte, distance float64, limit int) int {
	if limit < 0 {
		limit = math.MaxInt
	}

	// we collect all the antennas and their locations
	antennas := make(map[byte][]utils.Pos)

	for y, row := range in {
		for x := range row {
			ch := in[y][x]
			if ch == '.' {
				continue
			}
			antennas[ch] = append(antennas[ch], utils.NewPos(y, x))
		}
	}

	antiNodes := make(map[utils.Pos]struct{})

	// create all anti-nodes for the same type of antennas
	for _, positions := range antennas {
		CreateAntiNodes(positions, antiNodes, distance, limit, len(in)-1, len(in[0])-1)
	}

	return len(antiNodes)
}

func CreateAntiNodes(positions []utils.Pos, antiNodes map[utils.Pos]struct{}, t float64, limit int, maxY, maxX int) {
	// now we need to calculate the distances between antennas of the same group
	// and place the anti-node on the same line at the same distance

	for i := 0; i < len(positions); i++ {
		for j := i + 1; j < len(positions); j++ {
			antennaA := positions[i]
			antennaB := positions[j]

			// from A in direction of B
			for step := range limit {
				newPos := CreateNode(antennaA, antennaB, t*float64(step+1))

				if newPos.Y() < 0 || newPos.Y() > maxY || newPos.X() < 0 || newPos.X() > maxX {
					break
				}

				antiNodes[newPos] = struct{}{}
			}

			// from B in direction of A (use reverse unit vector)
			for step := range limit {
				newPos := CreateNode(antennaB, antennaA, t*float64(step+1))

				if newPos.Y() < 0 || newPos.Y() > maxY || newPos.X() < 0 || newPos.X() > maxX {
					break
				}

				antiNodes[newPos] = struct{}{}
			}
		}
	}
}

func CreateNode(a, b utils.Pos, t float64) utils.Pos {
	// direction vector (from A to B)
	dr := float64(b.Y() - a.Y())
	dc := float64(b.X() - a.X())

	// distance (sqrt(dx^2+dy^2))
	distance := math.Sqrt(dr*dr + dc*dc)

	// calculate the direction (unit vector)
	ur := dr / distance
	uc := dc / distance

	newY := int(math.Round(float64(a.Y()) + ur*t*distance))
	newX := int(math.Round(float64(a.X()) + uc*t*distance))

	return utils.NewPos(newY, newX)
}

//go:embed "example.txt"
var exampleInput string

//go:embed "input.txt"
var DayInput string

func main() {
	matrix, err := utils.ReadFileIntoBytesMatrix(strings.NewReader(exampleInput))
	if err != nil {
		panic(err)
	}

	// twice the distance
	uniqueLocations := CalculateUniqueAntiNodes(matrix, 2.0, 1)
	fmt.Printf("Unique locations of anti-nodes: %d\n", uniqueLocations)

	// same distance
	uniqueLocationsSameDistance := CalculateUniqueAntiNodes(matrix, 1.0, -1)
	fmt.Printf("Unique locations of anti-nodes (distance 1): %d\n", uniqueLocationsSameDistance)
}
