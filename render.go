package main

import (
	"hash/fnv"
	"image/color"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/rubensseva/go-dark-forest/civ"
)

type Renderer struct{
	game Game
}

func (g *Renderer) Update() error {
	for _, c := range g.game.Civs {
		c.CivTic(g.game.Systems)
	}
	time.Sleep(500 * time.Millisecond)
	return nil
}

func (g *Renderer) Draw(screen *ebiten.Image) {
	for _, sys := range g.game.Systems {
		renderSystem(screen, *sys)
	}
}

func (g *Renderer) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}


func hash(s string) uint32 {
        h := fnv.New32a()
        h.Write([]byte(s))
        return h.Sum32()
}


func convertPoints(x, y int64) (float64, float64) {
	// xdiff := civ.MinX * (-1)
	// ydiff := civ.MinY * (-1)
	return float64(x + (civ.MinX * (-1))), float64(y + (civ.MinY * (-1)))
}

func renderSystem(screen *ebiten.Image, sys civ.System) {
	// Assuming minx and miny is negative
	newX, newY := convertPoints(sys.Point.X, sys.Point.Y)

	var col color.Color
	if sys.Civ != nil {
		col = sys.Civ.Color
		// col = color.RGBA64{
		// 	R: uint16(hash(sys.Civ.Name)),
		// 	G: uint16(hash(sys.Civ.Name)) % 32,
		// 	B: uint16(hash(sys.Civ.Name)) % 10000,
		// 	A: 255,
		// }
	} else {
		col = color.Gray{
			Y: 255,
		}
	}

	ebitenutil.DrawRect(screen, float64(newX), float64(newY), 5.0, 5.0, col)

	if sys.Cached.BestSys != nil {
		xx, yy := convertPoints(sys.Cached.BestSys.Point.X, sys.Cached.BestSys.Point.Y)
		ebitenutil.DrawLine(
			screen,
			newX,
			newY,
			xx,
			yy,
			col,
		)
	}
}
