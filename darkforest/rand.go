package darkforest

import "math/rand"

func randRange(min int64, max int64) int64 {
	return rand.Int63n(max-min) + min
}
