package project2

func Add(x int, y int) int {
	return x + y
}
func Sub(x int, y int) int {
	return x - y
}
func Mul(x int, y int) int {
	return x * y
}
func Div(x int, y int) int {
	return x / y
}
func Calculator(x int, y int, op func(x int, y int) int) int {
	return op(x, y)
}
