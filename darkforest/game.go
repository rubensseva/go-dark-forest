package darkforest

import (
	"image/color"
)

type Game struct {
	Systems []*System
	Civs    []*Civ
}

func AddCiv(game *Game, firstSys *System, civName string, col color.Color) *Civ {
	firstSys.Population = 1
	firstCiv := Civ{
		Name:             civName,
		Color:            col,
		TechnologyLevel:  0,
		TechnologyGrowth: 0,
		Population:       1,
		Cohesion:         float64(randRange(100, 1000)),
		OwnedSystems:     []*System{firstSys},
	}
	firstSys.Civ = &firstCiv

	game.Civs = append(game.Civs, &firstCiv)

	return &firstCiv
}
