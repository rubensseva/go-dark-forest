package main

import (
	"hash/fnv"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"

	"github.com/rubensseva/go-dark-forest/darkforest"
)

var (
	ScreenWidthAndHeight = 1600

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

type Renderer struct {
	game *darkforest.Game
}

func (g *Renderer) Update() error {
	newCivs := make([]*darkforest.Civ, 0, len(g.game.Civs))
	for _, c := range g.game.Civs {
		if len(c.OwnedSystems) == 0 {
			// fmt.Printf("civ %v was exterminated\n", c.Name)
			continue
		}
		newCivs = append(newCivs, c)
	}
	g.game.Civs = newCivs

	for _, c := range g.game.Civs {
		c.CivTic(g.game)
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
	return ScreenWidthAndHeight, ScreenWidthAndHeight
}

func hash(s string) uint32 {
	h := fnv.New32a()
	h.Write([]byte(s))
	return h.Sum32()
}

func scalePoint(p float64) float64 {
	facP := p / (float64(darkforest.MaxXAndY) + (float64(darkforest.MinXAndY) * (-1)))
	return float64(ScreenWidthAndHeight) * facP
}

func convertPoints(x, y int64) (float64, float64) {
	// Assuming minx and miny is negative
	newX := float64(x + (darkforest.MinXAndY * (-1)))
	newY := float64(y + (darkforest.MinXAndY * (-1)))

	return scalePoint(newX), scalePoint(newY)
}

func renderSystem(screen *ebiten.Image, sys darkforest.System) {
	newX, newY := convertPoints(sys.Point.X, sys.Point.Y)

	var col color.Color
	var lowA color.Color
	if sys.Civ != nil {
		col = sys.Civ.Color
		r, g, b, _ := sys.Civ.Color.RGBA()
		lowA = color.RGBA{
			R: uint8(r) / 5,
			G: uint8(g) / 5,
			B: uint8(b) / 5,
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

	// vector.DrawFilledRect(screen, float32(newX), float32(newY), 1.0, 1.0, col, false)

	if sys.Civ != nil {

		sr := scalePoint(sys.ScanRange())
		vector.StrokeCircle(
			screen,
			float32(newX),
			float32(newY),
			float32(sr),
			10,
			lowA,
			false,
		)

		sp := scalePoint(sys.Power()) * 0.001
		vector.DrawFilledCircle(
			screen,
			float32(newX),
			float32(newY),
			float32(sp),
			col,
			false,
		)

		// text.Draw(
		// 	screen,
		// 	fmt.Sprintf("%v", sys.Population),
		// 	mplusNormalFont,
		// 	int(newX),
		// 	int(newY)+30,
		// 	color.White,
		// )
	}
	if sys.Cached.BestSys != nil {
		// text.Draw(
		// 	screen,
		// 	fmt.Sprintf("%v", sys.Cached.BestSysScore),
		// 	mplusNormalFont,
		// 	int(newX),
		// 	int(newY),
		// 	color.White,
		// )

		xx, yy := convertPoints(sys.Cached.BestSys.Point.X, sys.Cached.BestSys.Point.Y)

		vector.StrokeLine(
			screen,
			float32(newX),
			float32(newY),
			float32(xx),
			float32(yy),
			1,
			lowA,
			false,
		)
	}
}
