package civ

import (
	"github.com/rubensseva/go-dark-forest/system"
	"golang.org/x/exp/slices"
)

type Civ struct {
	Name             string
	TechnologyLevel  int
	TechnologyGrowth int
	Population       int
	OwnedSystems     []OwnedSystem
}

type OwnedSystem struct {
	Civ *Civ
	System *system.System
	Population int
}

func (c *Civ) totalResources() int {
	sum := 0
	for _, s := range c.OwnedSystems {
		sum += s.System.Resources
	}
	return sum
}

func (c *Civ) ownsSystem(system *system.System) bool {
	for _, owned := range c.OwnedSystems {
		if system == owned.System {
			return true
		}
	}
	return false
}

func (c *Civ) CivTic() {
	c.TechnologyLevel += c.TechnologyGrowth

	for _, owned := range c.OwnedSystems {
		pop := owned.Population
		resources := owned.System.Resources

		diff := resources - pop
		growth := diff / 10
		owned.Population += growth
	}



}

func SystemScore(o OwnedSystem, s system.System) float64 {
	distance := o.System.Point.Sub(s.Point).VecLen()
	resources := float64(s.Resources)
	discoverability := float64(s.Discoverability)

	return resources - distance - discoverability
}

func SortSystems(o OwnedSystem, systems []system.System) []system.System {
	sorty := func (s1, s2 system.System) bool {
		score1 := SystemScore(o, s1)
		score2 := SystemScore(o, s2)
		return score1 < score2
	}
	slices.SortFunc(systems, sorty)
}
