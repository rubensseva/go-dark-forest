package point

import (
	"math"
)

type Point struct {
	X int64
	Y int64
}

func (p Point) Sub(pp Point) Point {
	dx := p.X - pp.X
	dy := p.Y - pp.Y
	return Point{
		X: dx,
		Y: dy,
	}
}

func (p Point) VecLen() float64 {
	return math.Sqrt(float64(p.X*p.X + p.Y*p.Y))
}
