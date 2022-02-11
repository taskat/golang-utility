package array

import "github.com/taskat/golang-utility/intmath"

type UnaryPredicateString func(s string) bool

type UnaryModifierString func(s string) string

//Appends n to arr and returns it
func AppendString(arr []string, s string) []string {
	return append(arr, s)
}

//Return true if arr contains n
func ContainsString(arr []string, s string) bool {
	for _, value := range arr {
		if value == s {
			return true
		}
	}
	return false
}

//Returns how many elements satisfy the predicate
func CountString(arr []string, pred UnaryPredicateString) int {
	count := 0
	for _, value := range arr {
		if pred(value) {
			count++
		}
	}
	return count
}

//Returns the array of elements, that satisfies the predicate
func FilterString(arr []string, pred UnaryPredicateString) []string {
	result := make([]string, 0)
	for _, value := range arr {
		if pred(value) {
			result = append(result, value)
		}
	}
	return result
}

//Returns the length of arr
func LenString(arr []string) int {
	return len(arr)
}

//MapString calls f for every string in arr
func MapString(arr []string, f UnaryModifierString) []string {
	result := make([]string, len(arr))
	for i, value := range arr {
		result[i] = f(value)
	}
	return result
}

//Returns all the possible permutations of arr
func PermutateString(arr []string) [][]string {
	permutations := make([][]string, 0, intmath.Factorial(len(arr)))
	var helper func([]string, int)
	helper = func(arr []string, n int) {
		if n == 1 {
			tmp := make([]string, len(arr))
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

//Removes all instances of s from arr
func RemoveAllString(arr []string, s string) []string {
	newArr := make([]string, len(arr))
	copy(newArr, arr)
	arr = newArr
	for i := 0; i < len(arr); i++ {
		if arr[i] == s {
			arr = append(arr[:i], arr[i+1:]...)
		}
	}
	return arr
}

//Removes the first instance of s from arr
func RemoveFirstString(arr []string, s string) []string {
	newArr := make([]string, len(arr))
	copy(newArr, arr)
	arr = newArr
	for i := 0; i < len(arr); i++ {
		if arr[i] == s {
			return append(arr[:i], arr[i+1:]...)
		}
	}
	return arr
}
