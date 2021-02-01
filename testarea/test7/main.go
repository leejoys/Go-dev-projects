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

	fmt.Println(mergeSort(ar))
}

func mergeSort(ar []int) []int {
	// ваш код здесь
	var el int
	resultAr := make([]int, 0)
	arLen := len(ar)
	if arLen <= 1 {
		return ar
	}

	firstArray, secondArray := mergeSort(ar[:(arLen/2)]), mergeSort(ar[(arLen/2):])

	for i := 0; i < arLen; i++ {
		if (len(firstArray) == 0) || (len(secondArray) == 0) {
			if len(secondArray) == 0 {
				if len(firstArray) > 1 {
					el, firstArray = firstArray[0], firstArray[1:]
				} else {
					el = firstArray[0]
					firstArray = make([]int, 0)
				}
			} else if len(firstArray) == 0 {
				if len(secondArray) > 1 {
					el, secondArray = secondArray[0], secondArray[1:]
				} else {
					el = secondArray[0]
					secondArray = make([]int, 0)
				}
			}
		} else {
			if firstArray[0] <= secondArray[0] {
				if len(firstArray) > 1 {
					el, firstArray = firstArray[0], firstArray[1:]
				} else {
					el = firstArray[0]
					firstArray = make([]int, 0)
				}
			} else if firstArray[0] > secondArray[0] {
				if len(secondArray) > 1 {
					el, secondArray = secondArray[0], secondArray[1:]
				} else {
					el = secondArray[0]
					secondArray = make([]int, 0)
				}
			}
		}

		resultAr = append(resultAr, el)
	}

	return resultAr
}
