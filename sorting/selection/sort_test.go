package selection

import (
	"github.com/hi-lap/algorithms-4th/sorting"
	"math/rand"
	"testing"
	"time"
)

func TestSort(t *testing.T) {
	ints := sorting.Ints{}

	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 10; i++ {
		ints = append(ints, rand.Intn(100))
	}
	Sort(ints)
	if !ints.IsSorted() {
		t.Error("Result is: ", ints.String())
	}
}
