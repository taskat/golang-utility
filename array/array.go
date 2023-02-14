package array

import "github.com/taskat/golang-utility/intmath"

type UnaryPredicate[T any] func(T) bool

type UnaryModifier[T, U any] func(T) U

// Appends elem to arr and returns it
// It is a wrapper around the append function
func Append[T any] (arr []T, elem T) []T {
	return append(arr, elem)
}

// Return true if arr contains elem
// A nil array does not contain any element, so it will return false
func Contains[T comparable] (arr []T, n T) bool {
	for _, value := range arr {
		if value == n {
			return true
		}
	}
	return false
}

// Returns how many elements satisfy the predicate
// A nil array does not contain any element, so it will return 0
func Count[T any] (arr []T, pred UnaryPredicate[T]) int {
	count := 0
	for _, value := range arr {
		if pred(value) {
			count++
		}
	}
	return count
}

// Returns the array of elements that satisfies the predicate
// A nil array does not contain any element, so it will return an empty array
func Filter[T any] (arr []T, pred UnaryPredicate[T]) []T {
	result := make([]T, 0)
	for _, value := range arr {
		if pred(value) {
			result = append(result, value)
		}
	}
	return result
}

// Returns the length of arr
// A nil array has length 0
func Len[T any] (arr []T) int {
	return len(arr)
}

// Merge merges two arrays
// It modifies the original array
// A nil array does not contain any element, so it will return the other array
// If both arrays are nil, it will return an empty array
func Merge[T any] (arr, other []T) []T {
	if arr == nil {
		if other == nil {
			return make([]T, 0)
		}
		return other
	}
	return append(arr, other...)
}

// Map calls f for every element in arr
// A nil array does not contain any element, so it will return an empty array of the new type
func Map[T, U any] (arr []T, f UnaryModifier[T, U]) []U {
	result := make([]U, len(arr))
	for i, value := range arr {
		result[i] = f(value)
	}
	return result
}

// Returns all the possible permutations of arr
// A nil array does not contain any element, so it will return an empty array
func Permutate[T any] (arr []T) [][]T {
	permutations := make([][]T, 0, intmath.Factorial(len(arr)))
	var helper func([]T, int)
	helper = func(arr []T, length int){
        if length == 1 {
            tmp := make([]T, len(arr))
            copy(tmp, arr)
            permutations = append(permutations, tmp)
        } else {
            for i := 0; i < length; i++{
                helper(arr, length - 1)
                if length % 2 == 1 {
					arr[i], arr[length - 1] = arr[length - 1], arr[i]
                } else {
					arr[0], arr[length - 1] = arr[length - 1], arr[0]
                }
            }
        }
    }
    helper(arr, len(arr))
    return permutations
}

// Removes all instances of elem from arr
// It modifies the original array
// A nil array does not contain any element, but it will return an empty array
func RemoveAll[T comparable] (arr []T, elem T) []T {
	if arr == nil {
		return make([]T, 0)
	}
	for i := 0; i < len(arr); i++ {
		if arr[i] == elem {
			arr = append(arr[:i], arr[i+1:]...)
		}
	}
	return arr
}

// Removes the first instance of elem from arr
// It modifies the original array
// A nil array does not contain any element, but it will return an empty array
func RemoveFirst[T comparable] (arr []T, n T) []T {
	if arr == nil {
		return make([]T, 0)
	}
	for i := 0; i < len(arr); i++ {
		if arr[i] == n {
			return append(arr[:i], arr[i+1:]...)
		}
	}
	return arr
}

// Reverse reverses the order of the elements in arr
// It modifies the original array
// A nil array does not contain any element, but it will return an empty array
func Reverse[T any] (arr []T) []T {
	if arr == nil {
		return make([]T, 0)
	}
	for i := 0; i < len(arr)/2; i++ {
		arr[i], arr[len(arr)-i-1] = arr[len(arr)-i-1], arr[i]
	}
	return arr
}
