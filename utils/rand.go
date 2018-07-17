package utils

import (
	"math"
	"math/rand"
	"time"
)

var r = rand.New(rand.NewSource(time.Now().UTC().UnixNano()))

func IFeelLucky(oneof int, botsAmount int) bool {
	return rand.Float64() < (float64(oneof) / float64(botsAmount))
}

func RandomId() int64 {
	return r.Int63n(math.MaxInt64)
}

func RandomInt(max int) int {
	if max == 0 {
		return 0
	} else {
		return r.Intn(max)
	}
}

func Random(min, max int64) int64 {
	if min == max {
		return min
	} else {
		return r.Int63n(max-min) + min
	}
}
