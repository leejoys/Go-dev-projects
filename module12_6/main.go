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
	quickSort(ar)
	fmt.Println(ar)
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
	return
}
