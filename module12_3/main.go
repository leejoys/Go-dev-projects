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
	ar := make([]int, 50)
	for i := range ar {
		ar[i] = rand.Intn(200) - 100 // ограничиваем случайное значение от [-100;100]
	}
	return ar
}

func main() {
	ar := makeAr()
	insertionSort(ar)
	fmt.Println(ar)
}

func insertionSort(ar []int) {
	// ваш код здесь
	var tempCell int
	arLen := len(ar)

	for i := 1; i < arLen; i++ {
		tempCell = ar[i]
		iSort := i - 1
		for iSort >= -1 {
			if iSort == -1 {
				ar[iSort+1] = tempCell
				break
			} else if ar[iSort] > tempCell {
				ar[iSort+1] = ar[iSort]
			} else {
				ar[iSort+1] = tempCell
				break
			}
			iSort--
		}
	}
	return
}
