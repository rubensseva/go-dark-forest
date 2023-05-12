package main

import (
	"fmt"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/rubensseva/go-dark-forest/civ"
)

type Game struct {
	Systems []civ.System
}

func main() {
    systems := []civ.System{}
    for i := 0; i < 3; i++ {
        systems = append(systems, civ.GenSystem(systems))
    }
    for _, s := range systems {
        fmt.Println(s)
    }

	game := Game{
		Systems: systems,
	}

	firstSys := systems[0]

	// TODO: Extract to function
	firstSys.Population = 1
	firstCiv := civ.Civ{
		Name:             "",
		TechnologyLevel:  0,
		TechnologyGrowth: 0,
		Population:       1,
		OwnedSystems:     []*civ.System{&firstSys},
	}
	firstSys.Civ = &firstCiv

	for i := 0; i < 10; i++ {
		fmt.Println("printinng...")
		for _, s := range firstCiv.OwnedSystems {
			fmt.Printf("sys: %+v\n", *s)
		}
		fmt.Println("done")
		firstCiv.CivTic(systems)
	}

	fmt.Printf("owned: %+v\n", firstCiv.OwnedSystems)

	renderer := Renderer{
		game: game,
	}

	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Hello, World!")
	if err := ebiten.RunGame(&renderer); err != nil {
		log.Fatal(err)
	}

}
