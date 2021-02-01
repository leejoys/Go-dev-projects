package main

import (
	"fmt"
	"math"
	"os"
)

func main() {

	fmt.Println("Программа расчета средней температуры по городу")
	fmt.Println("Для выхода введите значение температуры больше 99")

	var stationTemp, stationNum int
	arrTemp := [10]int{10, 10, 10, 10, 10, 10, 10, 10, 10, 10}

	for {
		fmt.Print("Введите значение температуры: ")
		fmt.Fscan(os.Stdin, &stationTemp)
		if stationTemp > 99 {
			break
		}

		fmt.Print("Введите номер станции: ")
		fmt.Fscan(os.Stdin, &stationNum)
		if stationNum < 1 || stationNum > 10 {
			fmt.Println("Неправильный номер станции")
			continue
		}

		arrTemp[stationNum-1] = stationTemp

		sum := 0

		for i := 0; i < len(arrTemp); i++ {
			sum = sum + arrTemp[i]
		}
		averageTemp := int(math.Round(float64(sum) / float64(len(arrTemp))))
		fmt.Printf("Средняя температура по городу: %d\n", averageTemp)
	}

}
