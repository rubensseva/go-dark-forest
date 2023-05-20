package main

import (
	"fmt"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/rubensseva/go-dark-forest/darkforest"
)

var (
	WindowWidthAndHeight = 1000
)

func main() {
	systems := []*darkforest.System{}
	for i := 0; i < 200; i++ {
		s := darkforest.GenSystem(systems)
		systems = append(systems, &s)
	}
	for _, s := range systems {
		fmt.Println(s)
	}

	g := &darkforest.Game{
		Systems: systems,
	}

	darkforest.AddCiv(g, systems[0], "schmorp", color.RGBA{
		R: 255,
		A: 255,
	})
	darkforest.AddCiv(g, systems[1], "glorp", color.RGBA{
		G: 200,
		A: 255,
	})
	darkforest.AddCiv(g, systems[2], "larp", color.RGBA{
		B: 200,
		A: 255,
	})
	darkforest.AddCiv(g, systems[3], "larp", color.RGBA{
		R: 100,
		B: 100,
		A: 255,
	})
	darkforest.AddCiv(g, systems[4], "fjolp", color.RGBA{
		R: 100,
		G: 100,
		A: 255,
	})

	renderer := Renderer{
		game: g,
	}

	ebiten.SetWindowSize(WindowWidthAndHeight, WindowWidthAndHeight)
	ebiten.SetWindowTitle("Hello, World!")
	if err := ebiten.RunGame(&renderer); err != nil {
		log.Fatal(err)
	}

}
