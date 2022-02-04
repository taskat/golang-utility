package unique

//This type is contained by UniqueArray
type Item interface {
	Equals(other Item) bool
}

//Array type, it guarantees that the contained values are unique
type UniqueArray struct {
	data []Item
}

//Creates a new UniiqueArray object from items
func Create(items []Item) UniqueArray {
	u := UniqueArray{data: make([]Item, 0)}
	for _, item := range items {
		u.Push(item)
	}
	return u
}

//Returns true if the array contains the element
func (u UniqueArray) Contains(element Item) bool {
	for _, value := range u.data {
		if value.Equals(element) {
			return true
		}
	}
	return false
}

//Returns the i. element of arr
func (u UniqueArray) Get(index int) Item {
	return u.data[index]
}

//Returns the length of the array
func (u UniqueArray) Len() int {
	return len(u.data)
}

//Inserts the element into the UniqeArray if it is unique. It returns true if
//the element is inserted
func (u *UniqueArray) Push(element Item) bool {
	for _, value := range u.data {
		if value.Equals(element) {
			return false
		}
	}
	u.data = append(u.data, element)
	return true
}

//Removes the element if it is contained. It returns true
//if the element is removed
func (u *UniqueArray) Remove(element Item) bool {
	for i, value := range u.data {
		if value.Equals(element) {
			u.data = append(u.data[:i], u.data[i+1:]...)
			return true
		}
	}
	return false
}

//Sets the i. element to element. It returns false if the set failed,
//because the new value would violate the uniqueness
func (u *UniqueArray) Set(element Item, index int) bool {
	for i, value := range u.data {
		if i == index {
			continue
		}
		if value.Equals(element) {
			return false
		}
	}
	u.data[index] = element
	return true
}

