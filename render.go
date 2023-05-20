package main

import (
	"fmt"
	"hash/fnv"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"

	"github.com/rubensseva/go-dark-forest/civ"
)

var (
	ScreenWidth = 1600
	ScreenHeight = 1600

	mplusNormalFont font.Face
	mplusBigFont    font.Face
)

func init() {
	tt, err := opentype.Parse(fonts.MPlus1pRegular_ttf)
	if err != nil {
		log.Fatal(err)
	}

	const dpi = 72
	mplusNormalFont, err = opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    24,
		DPI:     dpi,
		Hinting: font.HintingVertical,
	})
	if err != nil {
		log.Fatal(err)
	}
	mplusBigFont, err = opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    32,
		DPI:     dpi,
		Hinting: font.HintingVertical,
	})
	if err != nil {
		log.Fatal(err)
	}
}

type Renderer struct{
	game *Game
}

func (g *Renderer) Update() error {
	for _, c := range g.game.Civs {
		c.CivTic(g.game.Systems)
	}
	// time.Sleep(100 * time.Millisecond)
	return nil
}

func (g *Renderer) Draw(screen *ebiten.Image) {
	for _, sys := range g.game.Systems {
		renderSystem(screen, *sys)
	}
}

func (g *Renderer) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return ScreenWidth, ScreenHeight
}


func hash(s string) uint32 {
        h := fnv.New32a()
        h.Write([]byte(s))
        return h.Sum32()
}


func convertPoints(x, y int64) (float64, float64) {
	// Assuming minx and miny is negative
	newX := float64(x + (civ.MinX * (-1)))
	newY := float64(y + (civ.MinY * (-1)))

	facX := newX / (float64(civ.MaxX) + (float64(civ.MinX) * (-1)))
	facY := newY / (float64(civ.MaxY) + (float64(civ.MinY) * (-1)))

	resX, resY := float64(ScreenWidth) * facX, float64(ScreenHeight) * facY
	return resX, resY
}

func renderSystem(screen *ebiten.Image, sys civ.System) {
	newX, newY := convertPoints(sys.Point.X, sys.Point.Y)

	var col color.Color
	var lowA color.Color
	if sys.Civ != nil {
		col = sys.Civ.Color
		r, g, b, _ := sys.Civ.Color.RGBA()
		lowA = color.RGBA{
			R: uint8(r),
			G: uint8(g),
			B: uint8(b),
			A: 1,
		}
	} else {
		col = color.Gray{
			Y: 255,
		}
		lowA = color.Gray{
			Y: 100,
		}
	}


	ebitenutil.DrawRect(screen, float64(newX), float64(newY), 5.0, 5.0, col)

	sr := sys.ScanRange()
	ebitenutil.DrawCircle(
		screen,
		newX,
		newY,
		sr,
		lowA,
	)


	if sys.Cached.BestSys != nil {
		text.Draw(
			screen,
			fmt.Sprintf("%v", int(sys.Cached.BestSysScore)),
			mplusNormalFont,
			int(newX),
			int(newY),
			color.White,
		)
		text.Draw(
			screen,
			fmt.Sprintf("%v", int(sys.Population)),
			mplusNormalFont,
			int(newX),
			int(newY) + 30,
			color.White,
		)

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
