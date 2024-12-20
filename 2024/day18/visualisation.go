package main

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"

	"github.com/digitalcrab/adventofcode/utils"
)

const (
	tileWidth  = 32
	tileHeight = 32
)

var (
	emptyImage     = createTile(color.RGBA{R: 0xCC, G: 0xFF, B: 0xCC, A: 0xFF})
	corruptedImage = createTile(color.RGBA{R: 0xFF, G: 0x45, B: 0x45, A: 0xFF})
	pathImage      = createTile(color.RGBA{R: 0xCC, G: 0xCC, B: 0xCC, A: 0xFF})
)

func createTile(c color.Color) *ebiten.Image {
	img := ebiten.NewImage(tileWidth, tileHeight)
	img.Fill(c)
	return img
}

type Visualisation struct {
	tileStates     map[utils.Pos]int // current status of the tile
	fallingBytes   []utils.Pos       // list of all bytes that are about to fall
	fallingByteIdx int               // current falling byte
	done           bool
	lastSteps      int
}

func NewVisualisation(fallingBytes []utils.Pos) *Visualisation {
	tileStates := make(map[utils.Pos]int, mapHeight*mapWidth)
	for p := range utils.PositionsForHeightWidth(mapHeight, mapWidth) {
		tileStates[p] = StateEmpty
	}

	v := &Visualisation{
		tileStates:     tileStates,
		fallingBytes:   fallingBytes,
		fallingByteIdx: 1,
	}
	return v
}

func (v *Visualisation) Update() error {
	if v.done {
		return nil
	}

	if v.fallingByteIdx >= len(v.fallingBytes) {
		fmt.Printf("[Vis] After %d bytes fall, path len %d\n", v.fallingByteIdx, v.lastSteps)
		v.done = true
		return nil
	}

	// cleanup state
	for p := range utils.PositionsForHeightWidth(mapHeight, mapWidth) {
		v.tileStates[p] = StateEmpty
	}

	// mark corrupted tiles
	fallingBytes := v.fallingBytes[:v.fallingByteIdx]
	for _, bp := range fallingBytes {
		v.tileStates[bp] = StateCorrupted
	}

	path := FindShortestPath(v.tileStates)
	if len(path) > 0 {
		v.lastSteps = len(path) - 1
		for _, p := range path {
			v.tileStates[p] = StatePath
		}
	}

	v.fallingByteIdx++
	return nil
}

func (v *Visualisation) Draw(screen *ebiten.Image) {
	// map size
	mh := mapHeight * tileHeight
	mw := mapWidth * tileWidth

	screenSize := screen.Bounds().Size()

	scaleX := float64(screenSize.X) / float64(mw)
	scaleY := float64(screenSize.Y) / float64(mh)

	scale := scaleX
	if scaleY < scaleX {
		scale = scaleY
	}

	mapImage := ebiten.NewImage(mw, mh)

	for p := range utils.PositionsForHeightWidth(mapHeight, mapWidth) {
		var tileImg *ebiten.Image

		switch v.tileStates[p] {
		case StateCorrupted:
			tileImg = corruptedImage
		case StatePath:
			tileImg = pathImage
		default:
			tileImg = emptyImage
		}

		tileOp := &ebiten.DrawImageOptions{}
		tileOp.GeoM.Translate(float64(p.X()*tileWidth), float64(p.Y()*tileHeight))

		mapImage.DrawImage(tileImg, tileOp)
	}

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(scale, scale)
	screen.DrawImage(mapImage, op)
}

func (v *Visualisation) Layout(outsideWidth, outsideHeight int) (int, int) {
	return outsideWidth, outsideHeight
}
