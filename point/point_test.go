package point

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSub(t *testing.T) {
	p1 := Point{
		X: 4,
		Y: 3,
	}
	p2 := Point{
		X: 5,
		Y: 6,
	}

	expected := Point{
		X: -1,
		Y: -3,
	}

	res := p1.Sub(p2)
	require.Equal(t, expected, res)
}

func TestLen(t *testing.T) {
	p := Point{
		X: 12,
		Y: -16,
	}
	// sqrt(12 * 12 + -16 * -16)
	expected := 20.0
	require.Equal(t, expected, p.VecLen())
}
