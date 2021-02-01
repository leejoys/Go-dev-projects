package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano()) // необходимо для того, чтобы рандом был похож на рандомный
}

func makeAr() []int {
	ar := make([]int, 50)
	for i := range ar {
		ar[i] = rand.Intn(200) - 100 //добавить отрицательных
	}
	return ar
}

func findMaxNegative(array []int) (maxNeg int, err error) {
	found := false
	maxNeg = math.MinInt64
	for _, val := range array {
		if (val > maxNeg) && (val < 0) {
			maxNeg = val
			found = true
		}
	}
	if !found {
		return 0, fmt.Errorf("input slice doesn't contain negative numbers")
	}
	return maxNeg, nil
}

func trimNegative(ar []int) []int {
	resultAr := make([]int, 0)
	for i := range ar {
		if ar[i] > -1 {
			resultAr = append(resultAr, ar[i])
		}
	}
	return resultAr
}

func trimLessAverage(ar []int) []int {
	sumAr := 0
	for i := range ar {
		sumAr = sumAr + ar[i]
	}
	averageAr := sumAr / len(ar)
	resultAr := make([]int, 0)
	for i := range ar {
		if ar[i] >= averageAr {
			resultAr = append(resultAr, ar[i])
		}
	}
	return resultAr
}

func main() {
	ar := makeAr()
	fmt.Println(ar)
	fmt.Println(findMaxNegative(ar))
	fmt.Println(trimNegative(ar))
	fmt.Println(trimLessAverage(ar))
}
