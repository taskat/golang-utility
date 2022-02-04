package intmath

//Returns the sum of a and b
func Add(a, b int) int {
	return a + b
}

//Returns the fraction of a and b, rounded towards 0
func Divide(a, b int) int {
	return a / b
}

//Returns true if a equals b
func Equals(a, b int) bool {
	return a == b
}

//Returns true if a is greater than b
func Greater(a, b int) bool {
	return a > b
}

//Returns true if a is greater than, or equals b
func GreaterOrEquals(a, b int) bool {
	return a >= b
}

//Returns the a mod b 
func Modulo(a, b int) int {
	return a % b
}

//Returns the product of a and b
func Multiply(a, b int) int {
	return a * b
}

//Returns true if a is smaller than b
func Smaller(a, b int) bool {
	return a < b
}

//Returns true if a is smaller than, or equals b
func SmallerOrEquals(a, b int) bool {
	return a <= b
}

//Returns the difference of a and b
func Subtract(a, b int) int {
	return a - b
}

