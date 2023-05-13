package main

import (
	"fmt"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/rubensseva/go-dark-forest/civ"
)

type Game struct {
	Systems []*civ.System
	Civs    []*civ.Civ
}

func main() {
	systems := []*civ.System{}
	for i := 0; i < 100; i++ {
		s := civ.GenSystem(systems)
		systems = append(systems, &s)
	}
	for _, s := range systems {
		fmt.Println(s)
	}

	game := Game{
		Systems: systems,
	}

	{
		c := color.RGBA{
			R: 0,
			G: 0,
			B: 255,
			A: 255,
		}

		firstSys := systems[0]
		// TODO: Extract to function
		firstSys.Population = 1
		firstCiv := civ.Civ{
			Name:             "schmorp",
			Color:            c,
			TechnologyLevel:  0,
			TechnologyGrowth: 0,
			Population:       1,
			OwnedSystems:     []*civ.System{firstSys},
		}
		firstSys.Civ = &firstCiv

		game.Civs = append(game.Civs, &firstCiv)
	}

	{
		c := color.RGBA{
			R: 255,
			G: 0,
			B: 0,
			A: 255,
		}

		secondSys := systems[1]
		// TODO: Extract to function
		secondSys.Population = 1
		secondCiv := civ.Civ{
			Name:             "larppppp",
			Color:            c,
			TechnologyLevel:  0,
			TechnologyGrowth: 0,
			Population:       1,
			OwnedSystems:     []*civ.System{secondSys},
		}
		secondSys.Civ = &secondCiv

		game.Civs = append(game.Civs, &secondCiv)
	}

	// for i := 0; i < 100; i++ {
	// 	fmt.Println("printinng...")
	// 	for _, s := range firstCiv.OwnedSystems {
	// 		fmt.Printf("sys: %+v\n", *s)
	// 	}
	// 	fmt.Println("done")
	// 	firstCiv.CivTic(systems)
	// }

	renderer := Renderer{
		game: game,
	}

	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Hello, World!")
	if err := ebiten.RunGame(&renderer); err != nil {
		log.Fatal(err)
	}

}
