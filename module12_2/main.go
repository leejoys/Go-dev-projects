package main

import (
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano()) // необходимо для того, чтобы рандом был похож на рандомный
}

func makeAr() []int {
	ar := make([]int, 51)
	for i := range ar {
		ar[i] = rand.Intn(200) - 100 // ограничиваем случайное значение от [-100;100]
	}
	return ar
}

func main() {
	fmt.Println("По меньшему")
	ar := makeAr()
	selectionSort(ar)
	fmt.Println(ar)

	fmt.Println("По большему")
	ar = makeAr()
	selectionSortByMax(ar)
	fmt.Println(ar)

	fmt.Println("В обе стороны")
	ar = makeAr()
	selectionSortMinMax(ar)
	fmt.Println(ar)
}

func selectionSort(ar []int) {
	// ваш код здесь
	minIndex := 0
	arLen := len(ar)
	sorted := true
	for i := 0; i < arLen; i++ {
		if ar[i] < ar[minIndex] {
			minIndex = i
		}
		sorted = false
	}
	if sorted {
		return
	}
	ar[0], ar[minIndex] = ar[minIndex], ar[0]
	selectionSort(ar[1:])
}

func selectionSortByMax(ar []int) {

	arLen := len(ar)
	maxIndex := arLen - 1
	sorted := true
	for i := 0; i < arLen; i++ {
		if ar[i] > ar[maxIndex] {
			maxIndex = i
		}
		sorted = false
	}
	if sorted {
		return
	}
	ar[arLen-1], ar[maxIndex] = ar[maxIndex], ar[arLen-1]
	selectionSort(ar[:arLen-1])
}

func selectionSortMinMax(ar []int) {

	minIndex := 0
	arLen := len(ar)
	maxIndex := arLen - 1
	sorted := true
	for i := 0; i < arLen; i++ {
		if ar[i] > ar[maxIndex] {
			maxIndex = i
		}
		if ar[i] < ar[minIndex] {
			minIndex = i
		}
		sorted = false
	}

	if sorted {
		return
	}
	ar[0], ar[arLen-1], ar[minIndex], ar[maxIndex] = ar[minIndex], ar[maxIndex], ar[0], ar[arLen-1]
	selectionSort(ar[1 : arLen-1])
}
