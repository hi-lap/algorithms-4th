package sorting

import (
	"fmt"
	"strconv"
	"strings"
)

type Ints []int

func (a Ints) Len() int {
	return len(a)
}

func (a Ints) Swap(i, j int) {
	fmt.Print("[")
	for idx, elm := range a {
		switch idx {
		case i, j:
			fmt.Printf("%4s", ">"+strconv.Itoa(elm))
		default:
			fmt.Printf("%4d", elm)
		}
	}
	fmt.Print("]\n")
	a[i], a[j] = a[j], a[i]

	fmt.Print("[")
	for idx, elm := range a {
		switch idx {
		case i, j:
			fmt.Printf("%4s", ">"+strconv.Itoa(elm))
		default:
			fmt.Printf("%4d", elm)
		}
	}
	fmt.Print("]")
	fmt.Println()
	fmt.Println()
}
func (a Ints) Less(i, j int) bool {
	return a[i] < a[j]
}

func (a Ints) IsSorted() bool {
	for i := 1; i < a.Len(); i++ {
		if a.Less(i, i-1) {
			return false
		}
	}
	return true
}

func (a Ints) String() string {
	texts := []string{}
	for _, item := range a {
		texts = append(texts, strconv.Itoa(item))
	}

	text := strings.Join(texts, " ")
	return "[" + text + "]"
}
