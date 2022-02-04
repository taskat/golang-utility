package intmath

//Returns the maximal value of int64
func MaxInt() int {
	return int(^uint(0) >> 1)
}

//Returns the minimal value of int64
func MinInt() int {
	return -MaxInt() - 1
}

//Returns 1
func One() int {
	return 1
}

//Returns 0
func Zero() int {
	return 0
}