package maths

// Div C = A / B
func Div(a, b float32) (float32, int) {
	if b == 0 {
		return 0, 1
	}
	return a / b, 0
}
