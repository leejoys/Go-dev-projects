package main

import (
	"fmt"
	"module7/calc"
)

func main() {
	var sign string
	var a, b float64

	fmt.Println("Калькулятор для простейших арифметических операций: + - * /")
	fmt.Println("Введите первое число, пробел, знак операции, пробел и второе число:")
	fmt.Scanf("%g %s %g", &a, &sign, &b)

	c := calc.Constructor(a, b, sign)

	res, errFlag := c.Calculate()

	if errFlag {
		fmt.Printf("Ошибка при вводе\n")
	} else {
		fmt.Printf("Результат равен %g\n", res)
	}

	fmt.Println(c)
	calc.Someassert(&c)

}
