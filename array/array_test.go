package array

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAppend(t *testing.T) {
	testCases := []struct {
		name        string
		inputArr    []int
		inputN      int
		expectedArr []int
	}{
		{
			name:        "Simple",
			inputArr:    []int{1, 2, 3},
			inputN:      4,
			expectedArr: []int{1, 2, 3, 4},
		},
		{
			name:        "ArrIsNil",
			inputArr:    nil,
			inputN:      1,
			expectedArr: []int{1},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.name, func(t *testing.T) {
			arr := Append(tC.inputArr, tC.inputN)
			assert.Equal(t, tC.expectedArr, arr)
		})
	}
}

func TestContains(t *testing.T) {
	testCases := []struct {
		name     string
		inputArr []int
		inputN   int
		expected bool
	}{
		{
			name:     "Contains",
			inputArr: []int{1, 2, 3},
			inputN:   1,
			expected: true,
		},
		{
			name:     "NotContains",
			inputArr: []int{1, 2, 3},
			inputN:   4,
			expected: false,
		},
		{
			name:     "ArrIsNil",
			inputArr: nil,
			inputN:   1,
			expected: false,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.name, func(t *testing.T) {
			contains := Contains(tC.inputArr, tC.inputN)
			assert.Equal(t, tC.expected, contains)
		})
	}
}

func TestCount(t *testing.T) {
	testCases := []struct {
		name      string
		inputArr  []int
		inputPred UnaryPredicate[int]
		expected  int
	}{
		{
			name:      "Even",
			inputArr:  []int{1, 2, 3},
			inputPred: func(i int) bool { return i%2 == 0 },
			expected:  1,
		},
		{
			name:      "Odd",
			inputArr:  []int{1, 2, 3},
			inputPred: func(i int) bool { return i%2 == 1 },
			expected:  2,
		},
		{
			name:      "ArrIsNil",
			inputArr:  nil,
			inputPred: func(i int) bool { return i%2 == 0 },
			expected:  0,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.name, func(t *testing.T) {
			count := Count(tC.inputArr, tC.inputPred)
			assert.Equal(t, tC.expected, count)
		})
	}
}

func TestFilter(t *testing.T) {
	testCases := []struct {
		name      string
		inputArr  []int
		inputPred UnaryPredicate[int]
		expected  []int
	}{
		{
			name:      "Even",
			inputArr:  []int{1, 2, 3},
			inputPred: func(i int) bool { return i%2 == 0 },
			expected:  []int{2},
		},
		{
			name:      "Odd",
			inputArr:  []int{1, 2, 3},
			inputPred: func(i int) bool { return i%2 == 1 },
			expected:  []int{1, 3},
		},
		{
			name:      "ArrIsNil",
			inputArr:  nil,
			inputPred: func(i int) bool { return i%2 == 0 },
			expected:  []int{},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.name, func(t *testing.T) {
			result := Filter(tC.inputArr, tC.inputPred)
			assert.Equal(t, tC.expected, result)
		})
	}
}

func TestLen(t *testing.T) {
	testCases := []struct {
		name        string
		inputArr    []int
		expectedLen int
	}{
		{
			name:        "Simple",
			inputArr:    []int{1, 2, 3},
			expectedLen: 3,
		},
		{
			name:        "ArrIsNil",
			inputArr:    nil,
			expectedLen: 0,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.name, func(t *testing.T) {
			length := Len(tC.inputArr)
			assert.Equal(t, tC.expectedLen, length)
		})
	}
}

func TestMerge(t *testing.T) {
	testCases := []struct {
		name        string
		inputArr1   []int
		inputArr2   []int
		expectedArr []int
	}{
		{
			name:        "Simple",
			inputArr1:   []int{1, 2, 3},
			inputArr2:   []int{4, 5, 6},
			expectedArr: []int{1, 2, 3, 4, 5, 6},
		},
		{
			name:        "Arr1IsNil",
			inputArr1:   nil,
			inputArr2:   []int{4, 5, 6},
			expectedArr: []int{4, 5, 6},
		},
		{
			name:        "Arr2IsNil",
			inputArr1:   []int{1, 2, 3},
			inputArr2:   nil,
			expectedArr: []int{1, 2, 3},
		},
		{
			name:        "BothArrIsNil",
			inputArr1:   nil,
			inputArr2:   nil,
			expectedArr: []int{},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.name, func(t *testing.T) {
			result := Merge(tC.inputArr1, tC.inputArr2)
			assert.Equal(t, tC.expectedArr, result)
		})
	}
}

func TestMap(t *testing.T) {
	testCases := []struct {
		name     string
		inputArr []int
		inputF   UnaryModifier[int, int]
		expected []int
	}{
		{
			name:     "Square",
			inputArr: []int{1, 2, 3},
			inputF:   func(i int) int { return i * i },
			expected: []int{1, 4, 9},
		},
		{
			name:     "ArrIsNil",
			inputArr: nil,
			inputF:   func(i int) int { return i * i },
			expected: []int{},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.name, func(t *testing.T) {
			result := Map(tC.inputArr, tC.inputF)
			assert.Equal(t, tC.expected, result)
		})
	}
}

func TestPermutate(t *testing.T) {
	testCases := []struct {
		name     string
		inputArr []int
		expected [][]int
	}{
		{
			name:     "Simple",
			inputArr: []int{1, 2, 3},
			expected: [][]int{{1, 2, 3}, {2, 1, 3}, {3, 2, 1}, {2, 3, 1}, {3, 1, 2}, {1, 3, 2}},
		},
		{
			name:     "ArrIsNil",
			inputArr: nil,
			expected: [][]int{},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.name, func(t *testing.T) {
			result := Permutate(tC.inputArr)
			assert.Equal(t, tC.expected, result)
		})
	}
}

func TestRemoveAll(t *testing.T) {
	testCases := []struct {
		name        string
		inputArr    []int
		inputN      int
		expectedArr []int
	}{
		{
			name:        "Simple",
			inputArr:    []int{1, 2, 3, 2},
			inputN:      2,
			expectedArr: []int{1, 3},
		},
		{
			name:        "ArrIsNil",
			inputArr:    nil,
			inputN:      1,
			expectedArr: []int{},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.name, func(t *testing.T) {
			result := RemoveAll(tC.inputArr, tC.inputN)
			assert.Equal(t, tC.expectedArr, result)
		})
	}
}

func TestRemoveFirst(t *testing.T) {
	testCases := []struct {
		name        string
		inputArr    []int
		inputN      int
		expectedArr []int
	}{
		{
			name:        "Simple",
			inputArr:    []int{1, 2, 3, 2},
			inputN:      2,
			expectedArr: []int{1, 3, 2},
		},
		{
			name:        "ArrIsNil",
			inputArr:    nil,
			inputN:      1,
			expectedArr: []int{},
		},
		{
			name:        "NIsNotInArr",
			inputArr:    []int{1, 2, 3},
			inputN:      4,
			expectedArr: []int{1, 2, 3},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.name, func(t *testing.T) {
			result := RemoveFirst(tC.inputArr, tC.inputN)
			assert.Equal(t, tC.expectedArr, result)
		})
	}
}

func TestReverse(t *testing.T) {
	testCases := []struct {
		name        string
		inputArr    []int
		expectedArr []int
	}{
		{
			name:        "Odd",
			inputArr:    []int{1, 2, 3},
			expectedArr: []int{3, 2, 1},
		},
		{
			name:        "Even",
			inputArr:    []int{1, 2, 3, 4},
			expectedArr: []int{4, 3, 2, 1},
		},
		{
			name:        "ArrIsNil",
			inputArr:    nil,
			expectedArr: []int{},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.name, func(t *testing.T) {
			result := Reverse(tC.inputArr)
			assert.Equal(t, tC.expectedArr, result)
		})
	}
}
