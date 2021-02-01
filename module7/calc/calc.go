package calc

import "fmt"

type calculator struct {
	a, b float64
	sign string
}

// Constructor create calculator
func Constructor(x, y float64, sig string) calculator {
	return calculator{
		a:    x,
		b:    y,
		sign: sig}
}

func (c *calculator) sum() float64 {
	return c.a + c.b
}

func (c *calculator) sub() float64 {
	return c.a - c.b
}

func (c *calculator) mult() float64 {
	return c.a * c.b
}

func (c *calculator) div() (float64, bool) {
	if c.b == 0 {
		return 0, true
	}
	return c.a / c.b, false
}

// Calculate calculieren
func (c *calculator) Calculate() (res float64, errFlag bool) {
	switch c.sign {
	case "+":
		res = c.sum()
	case "-":
		res = c.sub()
	case "*":
		res = c.mult()
	case "/":
		res, errFlag = c.div()
	default:
		errFlag = true
	}
	return
}

type somecalc interface{ Calculate() (float64, bool) }

//Someassert assert var to type
func Someassert(l somecalc) {
	fmt.Println(l)
	abc := l.(*calculator)
	fmt.Println(abc)
}
