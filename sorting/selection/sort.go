package selection

import (
	"sort"
)

func Sort(s sort.Interface) {
	for i := 0; i < s.Len(); i++ {
		min := i
		for j := i + 1; j < s.Len(); j++ {
			if s.Less(j, min) {
				min = j
			}
		}
		s.Swap(i, min)
	}
}