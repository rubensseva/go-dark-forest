package system

import (
	"fmt"
	"math/rand"

	"github.com/rubensseva/go-dark-forest/point"
)

type System struct {
	Name            string
	Resources       int
	Discoverability int
	Point               point.Point
}

func randRange(min int64, max int64) int64 {
	return rand.Int63n(max-min) + min
}

func GenSystem(systems []System) System {
    var pnt point.Point

	for {
		newP := point.Point{
			X: randRange(-100, 100),
			Y: randRange(-100, 100),
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

	return System{
		Name:            "Test",
		Resources:       1000,
		Discoverability: 1000,
		Point: pnt,
	}
}
