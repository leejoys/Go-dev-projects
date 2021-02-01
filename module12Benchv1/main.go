package main

import (
	"math/rand"
	"sort"
)

func bubbleSort(ar []int) {
	arLen := len(ar)
	for i := 0; true; i++ {
		sorted := true
		for i := 1; i < arLen; i++ {
			if ar[i-1] > ar[i] {
				ar[i-1], ar[i] = ar[i], ar[i-1]
				sorted = false
			}
		}
		if sorted {
			return
		}
		arLen--
	}
}
func sortSort(ar []int) {
	sort.Ints(ar)
}

func generateSlice(max, size int) []int {
	ar := make([]int, size)
	for i := range ar {
		ar[i] = rand.Intn(max*2) - max
	}

	return ar
}
