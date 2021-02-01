package main

import (
	"fmt"
)

var frsArLen, sndArLen, mapLen int

func main() {
	fmt.Println("Программа сравнения строковых массивов")
	fmt.Println("Задайте длину первого массива:")
	fmt.Scan(&frsArLen)
	fmt.Println("Задайте длину второго массива:")
	fmt.Scan(&sndArLen)

	fmt.Println("Введите элементы первого массива по одному:")
	frsAr := make([]string, frsArLen, frsArLen)
	for i := 0; i < frsArLen; i++ {
		var el string
		fmt.Scan(&el)
		frsAr[i] = el
	}

	fmt.Println("Введите элементы второго массива по одному:")
	sndAr := make([]string, sndArLen, sndArLen)
	for i := 0; i < sndArLen; i++ {
		var el string
		fmt.Scan(&el)
		sndAr[i] = el
	}
	if frsArLen < sndArLen {
		mapLen = sndArLen
	} else {
		mapLen = frsArLen
	}

	someMap := make(map[string]int, mapLen)
	for _, el := range frsAr {
		someMap[el]++
	}
	for _, el := range sndAr {
		someMap[el]++
	}

	resultAr := make([]string, 0)
	for key, val := range someMap {
		if val > 1 {
			resultAr = append(resultAr, key)
		}
	}
	fmt.Printf("Такие элементы присутствуют в обоих массивах: %q\n", resultAr)
}
