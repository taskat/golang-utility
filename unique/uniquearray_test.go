package unique

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testItem struct {
	value int
}

func (t testItem) Equals(other Item) bool {
	otherTestItem, ok := other.(testItem)
	return ok && t.value == otherTestItem.value
}

func TestCreate(t *testing.T) {
	testCases := []struct {
		name        string
		items       []Item
		expectedArr UniqueArray
	}{
		{
			name:        "Simple",
			items:       []Item{testItem{1}, testItem{2}, testItem{3}},
			expectedArr: UniqueArray{data: []Item{testItem{1}, testItem{2}, testItem{3}}},
		},
		{
			name:        "Redundant items",
			items:       []Item{testItem{1}, testItem{2}, testItem{3}, testItem{2}},
			expectedArr: UniqueArray{data: []Item{testItem{1}, testItem{2}, testItem{3}}},
		},
		{
			name:        "Nil items",
			items:       nil,
			expectedArr: UniqueArray{data: []Item{}},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.name, func(t *testing.T) {
			uniqueArr := Create(tC.items)
			assert.Equal(t, tC.expectedArr, uniqueArr)
		})
	}
}

func TestContains(t *testing.T) {
	testCases := []struct {
		name           string
		arr            UniqueArray
		item           Item
		expectedResult bool
	}{
		{
			name:           "Contains",
			arr:            Create([]Item{testItem{1}, testItem{2}, testItem{3}}),
			item:           testItem{1},
			expectedResult: true,
		},
		{
			name:           "Not contains",
			arr:            Create([]Item{testItem{1}, testItem{2}, testItem{3}}),
			item:           testItem{4},
			expectedResult: false,
		},
		{
			name:           "Nil data",
			arr:            Create(nil),
			item:           testItem{4},
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
		result Item
	}{
		{
			name:   "0",
			index:  0,
			result: testItem{1},
		},
		{
			name:   "1",
			index:  1,
			result: testItem{2},
		},
	}
	arr := Create([]Item{testItem{1}, testItem{2}, testItem{3}})
	for _, tC := range testCases {
		t.Run(tC.name, func(t *testing.T) {
			result := arr.Get(tC.index)
			assert.Equal(t, tC.result, result)
		})
	}
}

func TestLen(t *testing.T) {
	testCases := []struct {
		name        string
		arr         UniqueArray
		expectedLen int
	}{
		{
			name:        "Simple",
			arr:         Create([]Item{testItem{1}, testItem{2}, testItem{3}}),
			expectedLen: 3,
		},
		{
			name:        "Empty",
			arr:         Create([]Item{}),
			expectedLen: 0,
		},
		{
			name:        "Nil",
			arr:         Create(nil),
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

func TestPush(t *testing.T) {
	testCases := []struct {
		name           string
		arr            UniqueArray
		item           Item
		expectedArr    UniqueArray
		expectedResult bool
	}{
		{
			name:           "Simple",
			arr:            Create([]Item{testItem{1}, testItem{2}, testItem{3}}),
			item:           testItem{4},
			expectedArr:    Create([]Item{testItem{1}, testItem{2}, testItem{3}, testItem{4}}),
			expectedResult: true,
		},
		{
			name:           "AlreadyContains",
			arr:            Create([]Item{testItem{1}, testItem{2}, testItem{3}}),
			item:           testItem{1},
			expectedArr:    Create([]Item{testItem{1}, testItem{2}, testItem{3}}),
			expectedResult: false,
		},
		{
			name:           "Nil",
			arr:            Create(nil),
			item:           testItem{1},
			expectedArr:    Create([]Item{testItem{1}}),
			expectedResult: true,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.name, func(t *testing.T) {
			result := tC.arr.Push(tC.item)
			assert.Equal(t, tC.expectedArr, tC.arr)
			assert.Equal(t, tC.expectedResult, result)
		})
	}
}

func TestRemove(t *testing.T) {
	testCases := []struct {
		name           string
		arr            UniqueArray
		item           Item
		expectedArr    UniqueArray
		expectedResult bool
	}{
		{
			name:           "Simple",
			arr:            Create([]Item{testItem{1}, testItem{2}, testItem{3}}),
			item:           testItem{2},
			expectedArr:    Create([]Item{testItem{1}, testItem{3}}),
			expectedResult: true,
		},
		{
			name:           "AlreadyRemoved",
			arr:            Create([]Item{testItem{1}, testItem{2}, testItem{3}}),
			item:           testItem{4},
			expectedArr:    Create([]Item{testItem{1}, testItem{2}, testItem{3}}),
			expectedResult: false,
		},
		{
			name:           "Nil",
			arr:            Create(nil),
			item:           testItem{1},
			expectedArr:    Create(nil),
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
		arr            UniqueArray
		item           Item
		expectedArr    UniqueArray
		expectedResult bool
	}{
		{
			name:           "Simple",
			arr:            Create([]Item{testItem{1}, testItem{2}, testItem{3}}),
			item:           testItem{4},
			expectedArr:    Create([]Item{testItem{4}, testItem{2}, testItem{3}}),
			expectedResult: true,
		},
		{
			name:           "Fails",
			arr:            Create([]Item{testItem{1}, testItem{2}, testItem{3}}),
			item:           testItem{2},
			expectedArr:    Create([]Item{testItem{1}, testItem{2}, testItem{3}}),
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

