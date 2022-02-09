package unique

//This is a wrapper around strings, to satisfy the Item interface
type myString string

//Return true if other is an instance of myInt, and equals with m
func (m myString) Equals(other Item) bool {
	otherMyString, ok := other.(myString)
	return ok && m == otherMyString
}

//Specialization of UniqueArray for strings
type UniqueStringArray struct {
	UniqueArray
}

//Creates a new UniqueStringArray object
func CreatStrings(strings []string) UniqueStringArray {
	u := UniqueStringArray{UniqueArray: Create(nil)}
	for _, s := range strings {
		u.UniqueArray.Push(myString(s))
	}
	return u
}

//Returns true if u contains s
func (u UniqueStringArray) Contains(s string) bool {
	return u.UniqueArray.Contains(myString(s))
}

//Returns the i-th element of u
func (u UniqueStringArray) Get(index int) string {
	return string(u.UniqueArray.Get(index).(myString))
}

//Returns a copy of the contained array
func (u UniqueStringArray) GetData() []string {
	arr := u.UniqueArray.GetData()
	strings := make([]string, len(arr))
	for i, value := range arr {
		strings[i] = string(value.(myString))
	}
	return strings
}

//Returns the index of s. If s is not contained it return -1
func (u UniqueStringArray) GetIndex(s string) int {
	return u.UniqueArray.GetIndex(myString(s))
}

//Returns true if s is inserted into array, false if it already contained
func (u *UniqueStringArray) Push(s string) bool {
	return u.UniqueArray.Push(myString(s))
}

//Returns true if s is removed, false if it is not contained
func (u *UniqueStringArray) Remove(s string) bool {
	return u.UniqueArray.Remove(myString(s))
}

//Returns true if the i-th element of array is set succesfully to s, 
//false if it cannot be set
func (u *UniqueStringArray) Set(s string, index int) bool {
	return u.UniqueArray.Set(myString(s), index)
}