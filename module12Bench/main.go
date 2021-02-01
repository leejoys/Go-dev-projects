package main

import (
	"fmt"
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

func bubbleSortReversed(ar []int) {
	arLen := len(ar)
	for i := 0; true; i++ {
		sorted := true
		for i := 1; i < arLen; i++ {
			if ar[i-1] < ar[i] {
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

func bubbleSortRecursive(ar []int) {
	arLen := len(ar)
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
	bubbleSortRecursive(ar[:len(ar)-1])
}

func selectionSort(ar []int) {
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

func insertionSort(ar []int) {
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

		if firstArray[0] <= secondArray[0] {
			if len(firstArray) <= 1 {
				el = firstArray[0]
				resultAr = append(resultAr, el)
				resultAr = append(resultAr, secondArray...)
				return resultAr
			}
			el, firstArray = firstArray[0], firstArray[1:]
		} else if firstArray[0] > secondArray[0] {
			if len(secondArray) <= 1 {
				el = secondArray[0]
				resultAr = append(resultAr, el)
				resultAr = append(resultAr, firstArray...)
				return resultAr
			}
			el, secondArray = secondArray[0], secondArray[1:]
		}

		resultAr = append(resultAr, el)
	}

	return resultAr
}

func quickSort(ar []int) {
	// ваш код здесь
	left := 0

	arLen := len(ar)
	if arLen <= 1 {
		return
	}

	//первый шаг, указанный в подсказке, необязателен
	//если опорный элемент последний
	//ar[arLen-1], ar[left] = ar[left], ar[arLen-1] // так опорным будет первый
	for i := 0; i < arLen-1; i++ {
		if ar[i] < ar[arLen-1] {
			ar[left], ar[i] = ar[i], ar[left]
			left++
		}
	}
	ar[left], ar[arLen-1] = ar[arLen-1], ar[left]

	quickSort(ar[:left])
	quickSort(ar[left+1:])
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

func generateUpSlice(size int) []int {
	ar := make([]int, size)
	for i := range ar {
		ar[i] = i
	}

	return ar
}

func generateDownSlice(size int) []int {
	ar := make([]int, size)
	for i := 0; i < size; i++ {
		ar[i] = size - i
	}

	return ar
}

func main() {
	fmt.Println(generateDownSlice(100))
	fmt.Println(generateUpSlice(100))
}
