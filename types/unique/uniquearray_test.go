package unique

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreate(t *testing.T) {
	testCases := []struct {
		name        string
		items       []int
		expectedArr UniqueArray[int]
	}{
		{
			name:        "Simple",
			items:       []int{1, 2, 3},
			expectedArr: UniqueArray[int]{data: []int{1, 2, 3}},
		},
		{
			name:        "Redundant items",
			items:       []int{1, 2, 3, 1, 2, 3},
			expectedArr: UniqueArray[int]{data: []int{1, 2, 3}},
		},
		{
			name:        "Nil items",
			items:       nil,
			expectedArr: UniqueArray[int]{data: []int{}},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.name, func(t *testing.T) {
			uniqueArr := New(tC.items)
			assert.Equal(t, tC.expectedArr, uniqueArr)
		})
	}
}

func TestContains(t *testing.T) {
	testCases := []struct {
		name           string
		arr            UniqueArray[int]
		item           int
		expectedResult bool
	}{
		{
			name:           "Contains",
			arr:            New([]int{1, 2, 3}),
			item:           2,
			expectedResult: true,
		},
		{
			name:           "Not contains",
			arr:            New([]int{1, 2, 3}),
			item:           4,
			expectedResult: false,
		},
		{
			name:           "Nil data",
			arr:            UniqueArray[int]{data: nil},
			item:           4,
			expectedResult: false,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.name, func(t *testing.T) {
			result := tC.arr.Contains(tC.item)
			assert.Equal(t, tC.expectedResult, result)
		})
	}
}

func TestGet(t *testing.T) {
	testCases := []struct {
		name   string
		index  int
		result int
	}{
		{
			name:   "0",
			index:  0,
			result: 1,
		},
		{
			name:   "1",
			index:  1,
			result: 2,
		},
	}
	arr := New([]int{1, 2, 3})
	for _, tC := range testCases {
		t.Run(tC.name, func(t *testing.T) {
			result := arr.Get(tC.index)
			assert.Equal(t, tC.result, result)
		})
	}
}

func TestGetData(t *testing.T) {
	testCases := []struct {
		name   string
		arr    []int
		result []int
	}{
		{
			name:   "Simple",
			arr:    []int{1, 2, 3},
			result: []int{1, 2, 3},
		},
		{
			name:   "Nil",
			arr:    nil,
			result: []int{},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.name, func(t *testing.T) {
			arr := New(tC.arr)
			result := arr.GetData()
			assert.Equal(t, tC.result, result)
		})
	}
}

func TestGetIndex(t *testing.T) {
	testCases := []struct {
		name          string
		item          int
		expectedIndex int
	}{
		{
			name:          "0",
			item:          1,
			expectedIndex: 0,
		},
		{
			name:          "1",
			item:          2,
			expectedIndex: 1,
		},
		{
			name:          "-1",
			item:          4,
			expectedIndex: -1,
		},
	}
	arr := New([]int{1, 2, 3})
	for _, tC := range testCases {
		t.Run(tC.name, func(t *testing.T) {
			result := arr.GetIndex(tC.item)
			assert.Equal(t, tC.expectedIndex, result)
		})
	}
}

func TestInsert(t *testing.T) {
	testCases := []struct {
		name           string
		arr            UniqueArray[int]
		item           int
		expectedArr    UniqueArray[int]
		expectedResult bool
	}{
		{
			name:           "Simple",
			arr:            New([]int{1, 2, 3}),
			item:           4,
			expectedArr:    New([]int{1, 2, 3, 4}),
			expectedResult: true,
		},
		{
			name:           "AlreadyContains",
			arr:            New([]int{1, 2, 3}),
			item:           2,
			expectedArr:    New([]int{1, 2, 3}),
			expectedResult: false,
		},
		{
			name:           "Nil",
			arr:            New[int](nil),
			item:           2,
			expectedArr:    New([]int{2}),
			expectedResult: true,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.name, func(t *testing.T) {
			result := tC.arr.Insert(tC.item)
			assert.Equal(t, tC.expectedArr, tC.arr)
			assert.Equal(t, tC.expectedResult, result)
		})
	}
}

func TestLen(t *testing.T) {
	testCases := []struct {
		name        string
		arr         UniqueArray[int]
		expectedLen int
	}{
		{
			name:        "Simple",
			arr:         New([]int{1, 2, 3}),
			expectedLen: 3,
		},
		{
			name:        "Empty",
			arr:         New([]int{}),
			expectedLen: 0,
		},
		{
			name:        "Nil",
			arr:         New[int](nil),
			expectedLen: 0,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.name, func(t *testing.T) {
			length := tC.arr.Len()
			assert.Equal(t, tC.expectedLen, length)
		})
	}
}

func TestRemove(t *testing.T) {
	testCases := []struct {
		name           string
		arr            UniqueArray[int]
		item           int
		expectedArr    UniqueArray[int]
		expectedResult bool
	}{
		{
			name:           "Simple",
			arr:            New([]int{1, 2, 3}),
			item:           2,
			expectedArr:    New([]int{1, 3}),
			expectedResult: true,
		},
		{
			name:           "AlreadyRemoved",
			arr:            New([]int{1, 2, 3}),
			item:           4,
			expectedArr:    New([]int{1, 2, 3}),
			expectedResult: false,
		},
		{
			name:           "Nil",
			arr:            New[int](nil),
			item:           2,
			expectedArr:    New[int](nil),
			expectedResult: false,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.name, func(t *testing.T) {
			result := tC.arr.Remove(tC.item)
			assert.Equal(t, tC.expectedArr, tC.arr)
			assert.Equal(t, tC.expectedResult, result)
		})
	}
}

func TestSet(t *testing.T) {
	testCases := []struct {
		name           string
		arr            UniqueArray[int]
		item           int
		expectedArr    UniqueArray[int]
		expectedResult bool
	}{
		{
			name:           "Simple",
			arr:            New([]int{1, 2, 3}),
			item:           4,
			expectedArr:    New([]int{4, 2, 3}),
			expectedResult: true,
		},
		{
			name:           "Fails",
			arr:            New([]int{1, 2, 3}),
			item:           2,
			expectedArr:    New([]int{1, 2, 3}),
			expectedResult: false,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.name, func(t *testing.T) {
			result := tC.arr.Set(tC.item, 0)
			assert.Equal(t, tC.expectedArr, tC.arr)
			assert.Equal(t, tC.expectedResult, result)
		})
	}
}
