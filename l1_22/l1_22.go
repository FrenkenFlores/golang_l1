package l1_22

func Sum(a, b int64) int64 {
	return a + b
}

func Sub(a, b int64) int64 {
	return a - b
}

func Mult(a, b int64) int64 {
	return a * b
}

func Div(a, b int64) int64 {
	return a / b
}

func Operate(a, b int64, operator func(int64, int64) int64) int64 {
	return operator(a, b)
}
