package main

import (
	"github.com/hi-lap/algorithms-4th/sorting"
	"github.com/hi-lap/algorithms-4th/sorting/selection"
	"math/rand"
	"time"
)

func main() {
	ints := sorting.Ints{}

	rand.Seed(time.Now().UnixNano())
	for i := 0; i< 10; i++ {
		ints = append(ints, rand.Intn(100))
	}
	selection.Sort(ints)
}
