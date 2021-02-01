package main

import (
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano()) // необходимо для того, чтобы рандом был похож на рандомный
}

func main() {
	ar := make([]int, 50)
	for i := range ar {
		ar[i] = rand.Intn(200) - 100 // ограничиваем случайное значение от [-100;100]
	}

	fmt.Println("Вперед")
	bubbleSort(ar)
	fmt.Println(ar)

	fmt.Println("Назад")
	bubbleSortReversed(ar)
	fmt.Println(ar)

	fmt.Println("Рекурсия")
	bubbleSortRecursive(ar)
	fmt.Println(ar)
}

// ваш код здесь

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
