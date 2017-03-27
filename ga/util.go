package ga

import (
	"math"
	"math/rand"
	"time"
)

func RandomFloat() float64 {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	return r.Float64()
}

func RoundFloat(f float64) int {
	whole, frac := math.Modf(f)
	if frac >= 0.5 {
		whole++
	}

	return int(whole)
}
