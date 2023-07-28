package lottery

import (
	"math/rand"
	"time"
)

type Prize struct {
	Name   string
	Weight float64
}

func Draw(prizes []Prize) string {
	
	rand.Seed(time.Now().UnixNano())
	totalWeight := 0.0
	for _, prize := range prizes {
		totalWeight += prize.Weight
	}

	r := rand.Float64() * totalWeight
	for _, prize := range prizes {
		r -= prize.Weight
		if r <= 0 {
			return prize.Name
		}
	}

	return ""
}
