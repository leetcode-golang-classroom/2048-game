package internal

import (
	"math/rand"
	"time"
)

var defaultRandomPositioner randomPositoner = func(totalSize int) int {
	if totalSize == 1 {
		return 0
	}
	rand.New(rand.NewSource(time.Now().UnixNano()))
	return rand.Intn(totalSize)
}

var defaultRandomFunc randomGenerator = func() float64 {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	return rand.Float64()
}
