package calc

import "log"

const (
	operatorAddition       = "+"
	operatorSubtraction    = "-"
	operatorMultiplication = "*"
	operatorDivision       = "/"
)

type calculator struct{}

func NewCalculator() calculator {
	return calculator{}
}

func (c *calculator) Calculate(number1, number2 float64, operator string) float64 {
	switch operator {
	case operatorAddition:
		return c.add(number1, number2)
	case operatorSubtraction:
		return c.subtract(number1, number2)
	case operatorMultiplication:
		return c.multiply(number1, number2)
	case operatorDivision:
		return c.divide(number1, number2)
	default:
		log.Printf("несуществующий оператор: %s\n", operator)
		return 0
	}
}

func (c *calculator) add(number1, number2 float64) float64 {
	return number1 + number2
}
func (c *calculator) subtract(number1, number2 float64) float64 {
	return number1 - number2
}
func (c *calculator) multiply(number1, number2 float64) float64 {
	return number1 * number2
}
func (c *calculator) divide(number1, number2 float64) float64 {
	if number2 == 0 {
		log.Println("ошибка, деление на 0")
		return 0
	}

	return number1 / number2
}
