package intmath

import "math"

// Returns the absolute value of n
func Abs(n int) int {
	if n < 0 {
		return -1 * n
	}
	return n
}

//Return the cube of n
func Cube(n int) int {
	return n * n * n
}

//Decreases the value of n by 1
func Decrease(n int) int {
	return n - 1
}

//Returns true if n is even
func Even(n int) bool {
	return n % 2 == 0
}

//Return the factorial of n
func Factorial(n int) int {
	product := 1
	for i := 2; i <= n; i++ {
		product *= i
	}
	return product
}

//Increases the value of n by 1
func Increase(n int) int {
	return n + 1
}

//Returns true if n < 0
func Negative(n int) bool {
	return n < 0
}

//Returns true if n is odd
func Odd(n int) bool {
	return n % 2 == 1
}

//Returns true if n > 0
func Positive(n int) bool {
	return n > 0
}

//Returns the square root of n, rounded towards 0
func Sqrt(n int) int {
	return int(math.Sqrt(float64(n)))
}

//Returns the square of n
func Square(n int) int {
	return n * n
}