package unique

//This is a wrapper around ints, to satisfy the Item interface
type myInt int

//Return true if other is an instance of myInt, and equals with m
func (m myInt) Equals(other Item) bool {
	otherMyInt, ok := other.(myInt)
	return ok && m == otherMyInt
}

//Specialization of UniqueArray for ints
type UniqueIntArray struct {
	UniqueArray
}

//Creates a new UniqueIntArray object
func CreateInts(ints []int) UniqueIntArray {
	u := UniqueIntArray{UniqueArray: Create(nil)}
	for _, num := range ints {
		u.UniqueArray.Push(myInt(num))
	}
	return u
}

//Returns true if u contains i
func (u UniqueIntArray) Contains(i int) bool {
	return u.UniqueArray.Contains(myInt(i))
}

//Returns the i-th element of u
func (u UniqueIntArray) Get(index int) int {
	return int(u.UniqueArray.Get(index).(myInt))
}

//Returns a copy of the contained array
func (u UniqueIntArray) GetData() []int {
	arr := u.UniqueArray.GetData()
	ints := make([]int, len(arr))
	for i, value := range arr {
		ints[i] = int(value.(myInt))
	}
	return ints
}

//Returns the index of num. If num is not contained it return -1
func (u UniqueIntArray) GetIndex(num int) int {
	return u.UniqueArray.GetIndex(myInt(num))
}

//Returns true if number is inserted into array, false if it already contained
func (u *UniqueIntArray) Push(number int) bool {
	return u.UniqueArray.Push(myInt(number))
}

//Returns true if number is removed, false if it is not contained
func (u *UniqueIntArray) Remove(number int) bool {
	return u.UniqueArray.Remove(myInt(number))
}

//Returns true if the i-th element of array is set succesfully to num, 
//false if it cannot be set
func (u *UniqueIntArray) Set(number, index int) bool {
	return u.UniqueArray.Set(myInt(number), index)
}