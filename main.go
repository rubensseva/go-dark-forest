package main

import (
	"fmt"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/rubensseva/go-dark-forest/civ"
)

var (
	WindowWidthAndHeight = 800
)

type Game struct {
	Systems []*civ.System
	Civs    []*civ.Civ
}

func addCiv(game *Game, firstSys *civ.System, civName string, col color.Color) {
	firstSys.Population = 1
	firstCiv := civ.Civ{
		Name:             civName,
		Color:            col,
		TechnologyLevel:  0,
		TechnologyGrowth: 0,
		Population:       1,
		OwnedSystems:     []*civ.System{firstSys},
	}
	firstSys.Civ = &firstCiv

	game.Civs = append(game.Civs, &firstCiv)
}

func main() {
	systems := []*civ.System{}
	for i := 0; i < 200; i++ {
		s := civ.GenSystem(systems)
		systems = append(systems, &s)
	}
	for _, s := range systems {
		fmt.Println(s)
	}

	game := &Game{
		Systems: systems,
	}

	addCiv(game, systems[0], "schmorp", color.RGBA{
		R: 255,
		A: 255,
	})
	addCiv(game, systems[1], "glorp", color.RGBA{
		G: 200,
		A: 255,
	})
	addCiv(game, systems[2], "larp", color.RGBA{
		B: 200,
		A: 255,
	})
	addCiv(game, systems[3], "larp", color.RGBA{
		R: 100,
		B: 100,
		A: 255,
	})
	addCiv(game, systems[4], "fjolp", color.RGBA{
		R: 100,
		G: 100,
		A: 255,
	})

	renderer := Renderer{
		game: game,
	}

	ebiten.SetWindowSize(WindowWidthAndHeight, WindowWidthAndHeight)
	ebiten.SetWindowTitle("Hello, World!")
	if err := ebiten.RunGame(&renderer); err != nil {
		log.Fatal(err)
	}

}
