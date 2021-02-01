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
		ar[i] = rand.Intn(200) //- 100 //добавить отрицательных
	}
	return ar
}

func findMostOftenRepeated(array []int) (mostOften int, err error) {
	if len(array) == 0 {
		return 0, fmt.Errorf("could not found repeated numbers in empty slice")
	}
	var maxIndex, maxCount = 0, 0
	for i, number := range array {
		currentCount := 0
		for _, numberToCompare := range array[i+1:] {
			if number == numberToCompare {
				currentCount++
			}
		}
		if currentCount > maxCount {
			maxIndex = i
			maxCount = currentCount
		}
	}
	return array[maxIndex], nil
}

func findMostOftenRepeatedWithMap(array []int) (mostOften int, err error) {
	if len(array) == 0 {
		return 0, fmt.Errorf("could not found repeated numbers in empty slice")
	}

	oftenRepeated := make(map[int]int)

	for _, el := range array {
		oftenRepeated[el] = oftenRepeated[el] + 1
	}

	countMap := 0
	ok := false
	for el, count := range oftenRepeated {
		if count > 1 {
			ok = true
			if count > countMap {
				mostOften = el
				countMap = count
			}
		}
	}
	if !ok {
		return 0, fmt.Errorf("could not found repeated numbers in empty slice")
	}
	return mostOften, nil
}

func main() {
	ar := makeAr()
	//ar := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 1, 2, 3, 4, 5, 6, 7, 1, 2, 3, 4, 5, 1, 2, 3, 1}
	//ar := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	fmt.Println(findMostOftenRepeated(ar))
	fmt.Println(findMostOftenRepeatedWithMap(ar))
}
