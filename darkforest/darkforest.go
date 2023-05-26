package darkforest

import (
	"fmt"
	"time"
)

var (
	MaxXAndY = int64(100)
	MinXAndY = int64(-100)
)

// CivTic represents a tic in time for a Civ.
func (c *Civ) CivTic(game *Game) {
	if c.LastUpdate.IsZero() {
		c.LastUpdate = time.Now()
	}
	dt := time.Since(c.LastUpdate)
	c.LastUpdate = time.Now()
	ds := dt.Seconds()

	// Grow technology
	c.TechnologyLevel += c.TechnologyGrowth * ds

	// Decrease cohesion
	c.Cohesion -= float64(len(c.OwnedSystems)) * ds

	// Store a nice slice so we don't modify the slice while looping
	tmp := []*System{}
	for _, s := range c.OwnedSystems {
		tmp = append(tmp, s)
	}

	for _, owned := range tmp {
		owned.OwnedSystemTic(game)
	}
}

func (s *System) OwnedSystemTic(game *Game) {
	if s.Population == 0 {
		panic(fmt.Sprintf("pop was zero for system %+v", s))
	}

	if s.LastUpdate.IsZero() {
		s.LastUpdate = time.Now()
	}
	dt := time.Since(s.LastUpdate)
	s.LastUpdate = time.Now()
	ds := dt.Seconds()

	pop := s.Population
	resources := s.Resources

	// First, grow the population
	diff := resources - pop
	growth := (float64(diff) / 4.0) * ds
	if growth <= 1.0 {
		growth = 1.0
	}
	s.Population += int(growth)

	if pop > ((resources * 1) + int(s.Civ.Cohesion)) {
		collapse(s, game)
		return
	}

	// Decrease discoverability
	s.Discoverability -= 100.0 * ds

	// Now we need sort all the non-owned systems based on systemscore
	nonOwnedSystems := []*System{}
	for _, ss := range game.Systems {
		if ss.Civ == s.Civ {
			continue
		}

		dis := s.Point.Sub(ss.Point).VecLen()
		if dis > s.ScanRange() {
			continue
		}

		nonOwnedSystems = append(
			nonOwnedSystems,
			ss,
		)
	}
	sortSystems(*s, nonOwnedSystems)

	// Are there any systems available at all?
	// If there is not, it means the whole universe is
	// currently colonized
	if len(nonOwnedSystems) == 0 {
		s.Cached = CachedSysVals{}
		return
	}
	// Let's get the best candidate for emigration
	best := nonOwnedSystems[0]

	// We found a civ! exterminate the system
	if best.Civ != nil {
		best.Civ.OwnedSystems = remove(best, best.Civ.OwnedSystems)

		best.Civ = nil
		best.Population = 1

		best.Discoverability += (best.Discoverability / 2)
		best.Cached = CachedSysVals{}
	}

	systemScore := systemScore(*s, *best)
	expandThreshold := 1000.0
	popresfac := float64(pop) / float64(resources)
	needForExpansion := popresfac * expandThreshold

	if needForExpansion+systemScore >= expandThreshold {
		expand(s, best)
	}

	// Update cached data
	s.Cached.BestSys = best
	s.Cached.BestSysScore = systemScore
	s.Cached.NeedForExpansion = needForExpansion
}
