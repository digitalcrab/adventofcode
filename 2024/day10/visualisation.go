package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"

	"github.com/digitalcrab/adventofcode/utils"
)

const (
	tileWidth  = 32 // isometric, tileWidth is often double the tileHeight
	tileHeight = 32
)

var (
	greyPalette = []color.Color{
		color.RGBA{0xFF, 0xFF, 0xFF, 0xFF},
		color.RGBA{0xF2, 0xF2, 0xF2, 0xFF},
		color.RGBA{0xE6, 0xE6, 0xE6, 0xFF},
		color.RGBA{0xD9, 0xD9, 0xD9, 0xFF},
		color.RGBA{0xCC, 0xCC, 0xCC, 0xFF},
		color.RGBA{0xB3, 0xB3, 0xB3, 0xFF},
		color.RGBA{0x99, 0x99, 0x99, 0xFF},
		color.RGBA{0x80, 0x80, 0x80, 0xFF},
		color.RGBA{0x66, 0x66, 0x66, 0xFF},
		color.RGBA{0x4D, 0x4D, 0x4D, 0xFF},
	}
	greenPalette = []color.Color{
		color.RGBA{0xCC, 0xFF, 0xCC, 0xFF},
		color.RGBA{0xB3, 0xEE, 0xB3, 0xFF},
		color.RGBA{0x99, 0xDD, 0x99, 0xFF},
		color.RGBA{0x80, 0xCC, 0x80, 0xFF},
		color.RGBA{0x66, 0xBB, 0x66, 0xFF},
		color.RGBA{0x4D, 0xAA, 0x4D, 0xFF},
		color.RGBA{0x33, 0x99, 0x33, 0xFF},
		color.RGBA{0x1A, 0x88, 0x1A, 0xFF},
		color.RGBA{0x00, 0x77, 0x00, 0xFF},
		color.RGBA{0x00, 0x66, 0x00, 0xFF},
	}

	black = color.RGBA{0x00, 0x00, 0x00, 0xFF}

	greyTileImages  [10]*ebiten.Image
	greenTileImages [10]*ebiten.Image
)

func init() {
	for i := 0; i < 10; i++ {
		greyTileImages[i] = createTile(greyPalette[i])
		greenTileImages[i] = createTile(greenPalette[i])
	}
}

func createTile(c color.Color) *ebiten.Image {
	img := ebiten.NewImage(tileWidth, tileHeight)
	img.Fill(c)
	return img
}

const (
	StateNotVisited = iota
	StateVisiting
	StateVisited
)

type step struct {
	row, col      int
	nextDirection int
	seen          map[utils.Position]struct{}
}

type Visualisation struct {
	Landscape  [][]byte
	tileStates [][]int

	dfsStack []step
	done     bool
}

func NewVisualisation(landscape [][]byte) *Visualisation {
	// init tile states
	h := len(landscape)
	w := len(landscape[0])
	tileStates := make([][]int, h)
	for i := range tileStates {
		tileStates[i] = make([]int, w)
	}

	// create all starting points
	var stack []step
	for row := 0; row < h; row++ {
		for col := 0; col < w; col++ {
			if landscape[row][col] == '0' {
				stack = append(stack, step{
					row:  row,
					col:  col,
					seen: make(map[utils.Position]struct{}),
				})
			}
		}
	}

	v := &Visualisation{
		Landscape:  landscape,
		tileStates: tileStates,
		dfsStack:   stack,
	}
	return v
}

func (v *Visualisation) Update() error {
	if v.done {
		return nil
	}

	// all starting points are gone
	if len(v.dfsStack) == 0 {
		v.done = true
		return nil
	}

	stackIdx := len(v.dfsStack) - 1
	curStep := v.dfsStack[stackIdx]
	pos := utils.NewPosition(curStep.row, curStep.col)

	if _, found := curStep.seen[pos]; found { // going back?
		// seen this tile, pop and mark visited
		v.tileStates[curStep.row][curStep.col] = StateVisited
		v.dfsStack = v.dfsStack[:stackIdx]
		return nil
	} else {
		// mark as seen
		curStep.seen[pos] = struct{}{}
		v.tileStates[curStep.row][curStep.col] = StateVisiting
	}

	// check if this tile is a top point
	if v.Landscape[curStep.row][curStep.col] == '9' {
		v.tileStates[curStep.row][curStep.col] = StateVisited
		v.dfsStack = v.dfsStack[:stackIdx]
		return nil
	}

	// next directions
	for curStep.nextDirection < len(utils.AzimuthDirections) {
		dir := utils.AzimuthDirections[curStep.nextDirection]
		// move to the next direction
		curStep.nextDirection++

		newR, newC := curStep.row+dir.Row, curStep.col+dir.Col

		// check boundaries
		if newR < 0 || newR >= len(v.Landscape) || newC < 0 || newC >= len(v.Landscape[newR]) {
			continue
		}

		// we can move only if next step is 1 more then previous
		if v.Landscape[curStep.row][curStep.col]+1 != v.Landscape[newR][newC] {
			continue
		}

		// add next step onto stack
		newStep := step{
			row:  newR,
			col:  newC,
			seen: curStep.seen,
		}

		// put current step back to stack with updated next direction
		v.dfsStack[stackIdx] = curStep
		v.dfsStack = append(v.dfsStack, newStep)

		return nil
	}

	// basically visualisation only, if nothing found we go backwards
	v.tileStates[curStep.row][curStep.col] = StateVisited
	v.dfsStack = v.dfsStack[:stackIdx]

	return nil
}

func (v *Visualisation) Draw(screen *ebiten.Image) {
	// map size
	mapHeight := len(v.Landscape) * tileHeight
	mapWidth := len(v.Landscape[0]) * tileWidth

	screenSize := screen.Bounds().Size()

	scaleX := float64(screenSize.X) / float64(mapWidth)
	scaleY := float64(screenSize.Y) / float64(mapHeight)

	scale := scaleX
	if scaleY < scaleX {
		scale = scaleY
	}

	mapImage := ebiten.NewImage(mapWidth, mapHeight)

	for row := range v.Landscape {
		for col := range v.Landscape[row] {
			ch := v.Landscape[row][col] - '0'
			var tileImg *ebiten.Image

			switch v.tileStates[row][col] {
			case StateNotVisited:
				tileImg = greyTileImages[ch]
			case StateVisiting:
				tileImg = createTile(black)
			case StateVisited:
				tileImg = greenTileImages[ch]
			}

			tileOp := &ebiten.DrawImageOptions{}
			tileOp.GeoM.Translate(float64(col*tileWidth), float64(row*tileHeight))

			mapImage.DrawImage(tileImg, tileOp)
		}
	}

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(scale, scale)
	screen.DrawImage(mapImage, op)
}

func (v *Visualisation) Layout(outsideWidth, outsideHeight int) (int, int) {
	return outsideWidth, outsideHeight
}
