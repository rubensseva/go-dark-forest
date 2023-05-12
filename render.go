package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/rubensseva/go-dark-forest/civ"
)

type Renderer struct{
	game Game
}

func (g *Renderer) Update() error {
	return nil
}

func (g *Renderer) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, "Hello, World!")
	for _, sys := range g.game.Systems {
		renderSystem(screen, sys)
	}
}

func (g *Renderer) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}

func renderSystem(screen *ebiten.Image, sys civ.System) {
	// Assuming minx and miny is negative
	xdiff := civ.MinX * (-1)
	ydiff := civ.MinY * (-1)

	newX := sys.Point.X + xdiff
	newY := sys.Point.Y + ydiff

	ebitenutil.DrawRect(screen, float64(newX), float64(newY), 5.0, 5.0, color.Gray{Y: 255})
}
