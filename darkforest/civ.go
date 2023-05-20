package darkforest

import (
	"image/color"
	"time"
)

type Civ struct {
	Name             string
	Color            color.Color
	TechnologyLevel  float64
	TechnologyGrowth float64
	Population       int
	Cohesion         float64
	OwnedSystems     []*System

	LastUpdate       time.Time
}

func (c *Civ) totalResources() int {
	sum := 0
	for _, s := range c.OwnedSystems {
		sum += s.Resources
	}
	return sum
}
