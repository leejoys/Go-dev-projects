package main

import (
	"calculator-struct/calc"
	"fmt"
)

func main() {
	var number1, number2 float64
	var operator string
	var err error

	fmt.Println("Введите первое число:")
	_, err = fmt.Scanln(&number1)
	if err != nil {
		fmt.Printf("ошибка при сканировании первого числа: %v", err)
	}

	fmt.Println("Введите оператор:")
	_, err = fmt.Scanln(&operator)
	if err != nil {
		fmt.Printf("ошибка при сканировании оператора: %v", err)
	}

	fmt.Println("Введите второе число:")
	_, err = fmt.Scanln(&number2)
	if err != nil {
		fmt.Printf("ошибка при сканировании второго числа: %v", err)
	}

	calculator := calc.NewCalculator()
	result := calculator.Calculate(number1, number2, operator)

	fmt.Println(result)
}
