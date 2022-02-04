package intmath

type UnaryPredicate func(n int) bool

type UnaryModifier func(n int) int

//Appends n to arr and returns it
func Append(arr []int, n int) []int {
	return append(arr, n)
}

//Return true if arr contains n
func Contains(arr []int, n int) bool {
	for _, value := range arr {
		if value == n {
			return true
		}
	}
	return false
}

//Returns how many elemnts satisfy the predicate
func Count(arr []int, pred UnaryPredicate) int {
	count := 0
	for _, value := range arr {
		if pred(value) {
			count++
		}
	}
	return count
}

//Returns the array of elements, that satisfies the predicate
func Filter(arr []int, pred UnaryPredicate) []int {
	result := make([]int, 0)
	for _, value := range arr {
		if pred(value) {
			result = append(result, value)
		}
	}
	return result
}

//Returns the length of arr
func Len(arr []int) int {
	return len(arr)
}

//Map calls f for every number int arr
func Map(arr []int, f UnaryModifier) []int {
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
func RemoveAll(arr []int, n int) []int {
	for i := 0; i < len(arr); i++ {
		if arr[i] == n {
			arr = append(arr[:i], arr[i + 1:]...)
		}
	}
	return arr
}

//Removes the first instance of n from arr
func RemoveFirst(arr []int, n int) []int {
	for i := 0; i < len(arr); i++ {
		if arr[i] == n {
			return append(arr[:i], arr[i + 1:]...)
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
