package unique

// This type is contained by UniqueArray
type Item interface {
	Equals(other Item) bool
}

// Array type, it guarantees that the contained values are unique
type UniqueArray[T comparable] struct {
	data []T
}

// New a new UniiqueArray object from items
func New[T comparable] (items []T) UniqueArray[T] {
	u := UniqueArray[T]{data: make([]T, 0)}
	for _, item := range items {
		u.Insert(item)
	}
	return u
}

// Contains returns true if the array contains the element
func (u UniqueArray[T]) Contains(element T) bool {
	for _, value := range u.data {
		if value == element {
			return true
		}
	}
	return false
}

// Get returns the i. element of arr
func (u UniqueArray[T]) Get(index int) T {
	return u.data[index]
}

// GetData returns a copy of the contained array
func (u UniqueArray[T]) GetData() []T {
	dataCopy := make([]T, len(u.data))
	copy(dataCopy, u.data)
	return dataCopy
}

// GetIndex returns the index of item. If the item is not contained, it returns -1
func (u UniqueArray[T]) GetIndex(item T) int {
	for i, elem := range u.data {
		if elem == item {
			return i
		}
	}
	return -1
}

// Insert inserts the element into the UniqeArray if it is unique. It returns true if
// the element is inserted
func (u *UniqueArray[T]) Insert(element T) bool {
	for _, value := range u.data {
		if value == element {
			return false
		}
	}
	u.data = append(u.data, element)
	return true
}

// Len returns the length of the array
func (u UniqueArray[T]) Len() int {
	return len(u.data)
}

// Remove removes the element if it is contained. It returns true
// if the element is removed
func (u *UniqueArray[T]) Remove(element T) bool {
	for i, value := range u.data {
		if value == element {
			u.data = append(u.data[:i], u.data[i+1:]...)
			return true
		}
	}
	return false
}

// Set sets the i. element to element. It returns false if the set failed,
// because the new value would violate the uniqueness
func (u *UniqueArray[T]) Set(element T, index int) bool {
	for i, value := range u.data {
		if i == index {
			continue
		}
		if value == element {
			return false
		}
	}
	u.data[index] = element
	return true
}

