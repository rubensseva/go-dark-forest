package civ

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/rubensseva/go-dark-forest/point"
)

var (
	names1 = []string{"glorp", "schmorp", "floop", "gloop", "schmerp"}
	names2 = []string{"flatul", "narbgslag", "TZZKTZ", "uuuundulgltl"}

	MaxXAndY = int64(100)
	MinXAndY = int64(-100)
)

func randRange(min int64, max int64) int64 {
	return rand.Int63n(max-min) + min
}

func GenSystem(systems []*System) System {
	var pnt point.Point

	for {
		newP := point.Point{
			X: randRange(MinXAndY, MaxXAndY),
			Y: randRange(MinXAndY, MaxXAndY),
		}
		toclose := false
		for _, s := range systems {
			lenn := newP.Sub(s.Point).VecLen()
			if lenn < 10.0 {
				fmt.Printf("%v is too close to %v\n", newP, s.Point)
				toclose = true
				break
			}
		}

		if toclose {
			continue
		}

		pnt = newP
		break
	}

	name := fmt.Sprintf("%s-%s", names1[rand.Intn(len(names1))], names2[rand.Intn(len(names2))])

	return System{
		Name:            name,
		Resources:       rand.Intn(1000),
		Discoverability: rand.Intn(1000),
		Point:           pnt,

		LastUpdate: time.Now(),
	}
}
