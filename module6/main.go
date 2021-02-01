package main

import (
	"fmt"
	"module6/maths"
)

func main() {
	var sign string
	var a, b, res float32
	var errFlag int

	fmt.Println("Калькулятор для простейших арифметических операций: + - * /")
	fmt.Println("Введите первое число, пробел, знак операции, пробел и второе число:")
	fmt.Scanf("%g %s %g", &a, &sign, &b)
	switch sign {
	case "+":
		res = maths.Sum(a, b)
	case "-":
		res = maths.Sub(a, b)
	case "*":
		res = maths.Mult(a, b)
	case "/":
		res, errFlag = maths.Div(a, b)
	default:
		errFlag = 1
	}

	if errFlag != 0 {
		fmt.Printf("Ошибка при вводе")
	} else {
		fmt.Printf("Результат равен %g", res)
	}
}
