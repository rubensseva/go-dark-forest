package darkforest

import (
	"fmt"
	"image/color"
	"math/rand"
	"time"

	"github.com/rubensseva/go-dark-forest/point"
	"golang.org/x/exp/slices"
)

var (
	names1 = []string{"glorp", "schmorp", "floop", "gloop", "schmerp"}
	names2 = []string{"flatul", "narbgslag", "TZZKTZ", "uuuundulgltl"}
)

type CachedSysVals struct {
	BestSys          *System
	BestSysScore     float64
	NeedForExpansion float64
}

type System struct {
	Name            string
	Resources       int
	Discoverability float64
	Point           point.Point

	LastUpdate time.Time

	// These fields are only relevant if the system is owned
	Civ        *Civ // If nil, indicates unowned
	Population int

	// Cached results for rendering, should never be
	// used in deciding anything
	Cached CachedSysVals
}

func (s *System) Power() float64 {
	if s.Civ == nil {
		return 0
	}
	return s.Civ.TechnologyLevel + float64(s.Population) + float64(s.Resources)
}

func (s *System) ScanRange() float64 {
	if s.Civ == nil {
		return 0
	}
	scanFactor := 0.02
	return s.Power() * scanFactor
}

func collapse(s *System, g *Game) {
	civ := s.Civ
	civ.OwnedSystems = remove(s, civ.OwnedSystems)

	s.Population = 1
	// s.Discoverability *= 2
	// s.Resources *= 4
	s.Civ = nil
	s.Cached = CachedSysVals{}

	newCiv := AddCiv(g, s, fmt.Sprintf("rand-civ-%d", randRange(1, 10000)), color.RGBA{
		R: uint8(randRange(10, 100)),
		G: uint8(randRange(10, 100)),
		B: uint8(randRange(10, 100)),
		A: 255,
	})
	newCiv.Cohesion *= 3
}

func expand(expanding *System, target *System) {
	target.Civ = expanding.Civ
	expanding.Civ.OwnedSystems = append(expanding.Civ.OwnedSystems, target)
	colonizingPop := expanding.Population / 2
	expanding.Population -= colonizingPop
	if colonizingPop < 1 {
		colonizingPop = 1
	}
	target.Population = colonizingPop

	target.LastUpdate = time.Time{}
}

// systemScore calculates a value for a System.
func systemScore(o System, s System) float64 {
	_distance := o.Point.Sub(s.Point).VecLen()
	resources := float64(s.Resources)
	discoverability := float64(s.Discoverability)

	distance := _distance * (_distance / 4)

	// Do some checks to avoid overflows
	if distance > resources {
		return 0
	}
	if discoverability > resources {
		return 0
	}
	if (distance + discoverability) > resources {
		return 0
	}

	res := resources - distance - discoverability
	if res < 0 {
		res = 0
	}

	return res
}

func sortSystems(o System, systems []*System) {
	sorty := func(s1, s2 *System) bool {
		score1 := systemScore(o, *s1)
		score2 := systemScore(o, *s2)
		return score1 > score2
	}
	slices.SortFunc(systems, sorty)
}

func GenSystem(systems []*System) System {
	var pnt point.Point

	for {
		newP := point.Point{
			X: randRange(MinXAndY, MaxXAndY),
			Y: randRange(MinXAndY, MaxXAndY),
		}


		// // Enforce some distance between systems
		// toclose := false
		// for _, s := range systems {
		// 	lenn := newP.Sub(s.Point).VecLen()
		// 	if lenn < 2.0 {
		// 		fmt.Printf("%v is too close to %v\n", newP, s.Point)
		// 		toclose = true
		// 		break
		// 	}
		// }

		// if toclose {
		// 	continue
		// }

		pnt = newP
		break
	}

	name := fmt.Sprintf("%s-%s", names1[rand.Intn(len(names1))], names2[rand.Intn(len(names2))])

	return System{
		Name:            name,
		Resources:       rand.Intn(1000),
		Discoverability: float64(rand.Intn(1000)),
		Point:           pnt,

		LastUpdate: time.Now(),
	}
}
