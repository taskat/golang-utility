package array

import "github.com/taskat/golang-utility/intmath"

type UnaryPredicateInt func(n int) bool

type UnaryModifierInt func(n int) int

//Appends n to arr and returns it
func AppendInt(arr []int, n int) []int {
	return append(arr, n)
}

//Return true if arr contains n
func ContainsInt(arr []int, n int) bool {
	for _, value := range arr {
		if value == n {
			return true
		}
	}
	return false
}

//Returns how many elements satisfy the predicate
func CountInt(arr []int, pred UnaryPredicateInt) int {
	count := 0
	for _, value := range arr {
		if pred(value) {
			count++
		}
	}
	return count
}

//Returns the array of elements, that satisfies the predicate
func FilterInt(arr []int, pred UnaryPredicateInt) []int {
	result := make([]int, 0)
	for _, value := range arr {
		if pred(value) {
			result = append(result, value)
		}
	}
	return result
}

//Returns the length of arr
func LenInt(arr []int) int {
	return len(arr)
}

//MapInt calls f for every number in arr
func MapInt(arr []int, f UnaryModifierInt) []int {
	result := make([]int, len(arr))
	for i, value := range arr {
		result[i] = f(value)
	}
	return result
}

//Returns the maximal value of arr
func Max(arr []int) int {
	if len(arr) == 0 {
		return 0
	}
	max := arr[0]
	for _, value := range arr {
		if value > max {
			max = value
		}
	}
	return max
}

//Returns the minimal value of arr
func Min(arr []int) int {
	if len(arr) == 0 {
		return 0
	}
	min := arr[0]
	for _, value := range arr {
		if value < min {
			min = value
		}
	}
	return min
}

//Returns all the possible permutations of arr
func PermutateInt(arr []int) [][]int {
	permutations := make([][]int, 0, intmath.Factorial(len(arr)))
	var helper func([]int, int)
	helper = func(arr []int, n int) {
		if n == 1 {
			tmp := make([]int, len(arr))
			copy(tmp, arr)
			permutations = append(permutations, tmp)
		} else {
			for i := 0; i < n; i++ {
				helper(arr, n-1)
				if n%2 == 1 {
					arr[i], arr[n-1] = arr[n-1], arr[i]
				} else {
					arr[0], arr[n-1] = arr[n-1], arr[0]
				}
			}
		}
	}
	helper(arr, len(arr))
	return permutations
}

//Returns the product of all values in arr
func Product(arr []int) int {
	if len(arr) == 0 {
		return 0
	}
	product := 1
	for _, value := range arr {
		product *= value
	}
	return product
}

//Removes all instances of n from arr
func RemoveAllInt(arr []int, n int) []int {
	newArr := make([]int, len(arr))
	copy(newArr, arr)
	arr = newArr
	for i := 0; i < len(arr); i++ {
		if arr[i] == n {
			arr = append(arr[:i], arr[i+1:]...)
		}
	}
	return arr
}

//Removes the first instance of n from arr
func RemoveFirstInt(arr []int, n int) []int {
	newArr := make([]int, len(arr))
	copy(newArr, arr)
	arr = newArr
	for i := 0; i < len(arr); i++ {
		if arr[i] == n {
			return append(arr[:i], arr[i+1:]...)
		}
	}
	return arr
}

//Returns the sum of values in arr
func Sum(arr []int) int {
	sum := 0
	for _, value := range arr {
		sum += value
	}
	return sum
}
