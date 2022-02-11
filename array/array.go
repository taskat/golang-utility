package array

import "github.com/taskat/golang-utility/intmath"

type UnaryPredicate func(i interface{}) bool

type UnaryModifier func(i interface{}) interface{}

//Appends n to arr and returns it
func Append(arr []interface{}, n interface{}) []interface{} {
	return append(arr, n)
}

//Return true if arr contains n
func Contains(arr []interface{}, n interface{}) bool {
	for _, value := range arr {
		if value == n {
			return true
		}
	}
	return false
}

//Returns how many elements satisfy the predicate
func Count(arr []interface{}, pred UnaryPredicate) int {
	count := 0
	for _, value := range arr {
		if pred(value) {
			count++
		}
	}
	return count
}

//Returns the array of elements, that satisfies the predicate
func Filter(arr []interface{}, pred UnaryPredicate) []interface{} {
	result := make([]interface{}, 0)
	for _, value := range arr {
		if pred(value) {
			result = append(result, value)
		}
	}
	return result
}

//Returns the length of arr
func Len(arr []interface{}) int {
	return len(arr)
}

//Map calls f for every element in arr
func Map(arr []interface{}, f UnaryModifier) []interface{} {
	result := make([]interface{}, len(arr))
	for i, value := range arr {
		result[i] = f(value)
	}
	return result
}

//Returns all the possible permutations of arr
func Permutate(arr []interface{}) [][]interface{} {
	permutations := make([][]interface{}, 0, intmath.Factorial(len(arr)))
	var helper func([]interface{}, int)
	helper = func(arr []interface{}, n int){
        if n == 1 {
            tmp := make([]interface{}, len(arr))
            copy(tmp, arr)
            permutations = append(permutations, tmp)
        } else {
            for i := 0; i < n; i++{
                helper(arr, n - 1)
                if n % 2 == 1 {
					arr[i], arr[n - 1] = arr[n - 1], arr[i]
                } else {
					arr[0], arr[n - 1] = arr[n - 1], arr[0]
                }
            }
        }
    }
    helper(arr, len(arr))
    return permutations
}

//Removes all instances of n from arr
func RemoveAll(arr []interface{}, n interface{}) []interface{} {
	newArr := make([]interface{}, len(arr))
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
func RemoveFirst(arr []interface{}, n interface{}) []interface{} {
	newArr := make([]interface{}, len(arr))
	copy(newArr, arr)
	arr = newArr
	for i := 0; i < len(arr); i++ {
		if arr[i] == n {
			return append(arr[:i], arr[i+1:]...)
		}
	}
	return arr
}
