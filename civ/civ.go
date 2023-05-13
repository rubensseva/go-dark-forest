package civ

import (
	"fmt"
	"image/color"
	"log"

	"github.com/rubensseva/go-dark-forest/point"
	"golang.org/x/exp/slices"
)

type Civ struct {
	Name             string
	Color            color.Color
	TechnologyLevel  int
	TechnologyGrowth int
	Population       int
	OwnedSystems     []*System
}

type System struct {
	Name            string
	Resources       int
	Discoverability int
	Point               point.Point

	// If nil, indicates unowned
	Civ *Civ
	Population int

	// Cached results for rendering, should never be
	// used in deciding anything
	Cached struct {
		BestSys *System
		BestSysScore float64
		NeedForExpansion float64
	}
}

func (c *Civ) totalResources() int {
	sum := 0
	for _, s := range c.OwnedSystems {
		sum += s.Resources
	}
	return sum
}

// CivTic represents a tic in time for a Civ.
// The main part is deciding what the Civ should do,
// which should go something like this:
//
// 1. The Civ is composed of many Systems, in some sense, each system
// decides what to do on its own, but keeping the entirety of the Civ in
// mind.
// 2. For each system that the Civ owns, it looks at it's own situation,
// the Civ's situation, and the neighboring planets. Then it decides what
// to do.
// 3. If the system has a population surplus, it will try to send population
// to neighboring Systems. It can also try to colonize a new System if any are
// available for colonization, but doing so is risky, which is the part where
// the situation of the Civ is taken into account.
// 4. The stats of the Civ will determine how willing it is to collinize a
// new system. The stats of the possibly colonized System is also taken into
// account.
//
// Thoughts:
// We could do this in a way that regards each system as its own AI. We could
// assign a state to each system, and then a set of behaviours that is triggered
// only from that state. This is complicated by the fact that the "state" encompasses
// all the neighboring systems, and the state of the entire Civ actually.
func (c *Civ) CivTic(allSystems []*System) {
	c.TechnologyLevel += c.TechnologyGrowth

	// Store a nice slice so we don't modify the slice while looping
	tmp := []*System{}
	for _, s := range c.OwnedSystems {
		tmp = append(tmp, s)
	}

	// Grow the population of the system
	for _, owned := range tmp {
		owned.OwnedSystemTic(allSystems)
	}
}

func (s *System) OwnedSystemTic(allSystems []*System) {
	if s.Population == 0 {
		panic(fmt.Sprintf("pop was zero for system %+v", s))
	}

	pop := s.Population
	resources := s.Resources

	// First, grow the population
	diff := resources - pop
	growth := diff / 2
	if growth <= 0 {
		growth = 1
	}
	s.Population += growth

	// Now we decide what to do
	// First: does this system even have a need for emigration?
	wantsToExpand := (resources - pop) > 0
	if !wantsToExpand {
		return
	}

	// Now we need sort all the non-owned systems based on systemscore
	nonOwnedSystems := []*System{}
	for _, ss := range allSystems {
		if ss.Civ != nil {
			continue
		}
		nonOwnedSystems = append(
			nonOwnedSystems,
			ss,
		)
	}
	SortSystems(*s, nonOwnedSystems)

	// Are there any systems available at all?
	// If this there isnt, it means the whole universe is
	// currently colonized
	if len(nonOwnedSystems) == 0 {
		log.Printf("the entire universe is colonized!")
		return
	}
	// Let's get the best candidate for emigration
	best := nonOwnedSystems[0]

	systemScore := SystemScore(*s, *best)
	expandThreshold := 1000.0
	popresfac := float64(pop) / float64(resources)
	needForExpansion := popresfac * expandThreshold

	if needForExpansion + systemScore >= expandThreshold {
		fmt.Printf("expanding civ %v... \n", s.Civ)
		best.Civ = s.Civ
		s.Civ.OwnedSystems = append(s.Civ.OwnedSystems, best)
		colonizingPop := s.Population / 2
		s.Population -= colonizingPop
		best.Population = colonizingPop
	}

	// Update cached data
	s.Cached.BestSys = best
	s.Cached.BestSysScore = systemScore
	s.Cached.NeedForExpansion = needForExpansion
}

// SystemScore calculates a value for a System.
func SystemScore(o System, s System) float64 {
	distance := o.Point.Sub(s.Point).VecLen()
	resources := float64(s.Resources)
	discoverability := float64(s.Discoverability)

	return resources - (distance * (distance / 4)) - discoverability
}

func SortSystems(o System, systems []*System) {
	sorty := func (s1, s2 *System) bool {
		score1 := SystemScore(o, *s1)
		score2 := SystemScore(o, *s2)
		return score1 > score2
	}
	slices.SortFunc(systems, sorty)
}
