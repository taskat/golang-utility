package intmath

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
		inputPred UnaryPredicate
		expected  int
	}{
		{
			name:      "Even",
			inputArr:  []int{1, 2, 3},
			inputPred: Even,
			expected:  1,
		},
		{
			name:      "Odd",
			inputArr:  []int{1, 2, 3},
			inputPred: Odd,
			expected:  2,
		},
		{
			name:      "ArrIsNil",
			inputArr:  nil,
			inputPred: Even,
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
		inputPred UnaryPredicate
		expected  []int
	}{
		{
			name:      "Even",
			inputArr:  []int{1, 2, 3},
			inputPred: Even,
			expected:  []int{2},
		},
		{
			name:      "Odd",
			inputArr:  []int{1, 2, 3},
			inputPred: Odd,
			expected:  []int{1, 3},
		},
		{
			name:      "ArrIsNil",
			inputArr:  nil,
			inputPred: Even,
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

func TestMap(t *testing.T) {
	testCases := []struct {
		name     string
		inputArr []int
		inputF   UnaryModifier
		expected []int
	}{
		{
			name:      "Square",
			inputArr:  []int{1, 2, 3},
			inputF: Square,
			expected:  []int{1, 4, 9},
		},
		{
			name:      "ArrIsNil",
			inputArr:  nil,
			inputF: Square,
			expected:  []int{},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.name, func(t *testing.T) {
			result := Map(tC.inputArr, tC.inputF)
			assert.Equal(t, tC.expected, result)
		})
	}
}

func TestMax(t *testing.T) {
	testCases := []struct {
		name        string
		inputArr    []int
		expectedMax int
	}{
		{
			name:        "Simple",
			inputArr:    []int{1, 2, 3},
			expectedMax: 3,
		},
		{
			name:        "ArrIsNil",
			inputArr:    nil,
			expectedMax: 0,
		},
		{
			name:        "ArrIsEmpty",
			inputArr:    []int{},
			expectedMax: 0,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.name, func(t *testing.T) {
			length := Max(tC.inputArr)
			assert.Equal(t, tC.expectedMax, length)
		})
	}
}

func TestMin(t *testing.T) {
	testCases := []struct {
		name        string
		inputArr    []int
		expectedMin int
	}{
		{
			name:        "Simple",
			inputArr:    []int{1, 2, 3},
			expectedMin: 1,
		},
		{
			name:        "ArrIsNil",
			inputArr:    nil,
			expectedMin: 0,
		},
		{
			name:        "ArrIsEmpty",
			inputArr:    []int{},
			expectedMin: 0,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.name, func(t *testing.T) {
			length := Min(tC.inputArr)
			assert.Equal(t, tC.expectedMin, length)
		})
	}
}

func TestProduct(t *testing.T) {
	testCases := []struct {
		name     string
		inputArr []int
		expected int
	}{
		{
			name:     "Simple",
			inputArr: []int{1, 2, 3},
			expected: 6,
		},
		{
			name:     "ArrIsNil",
			inputArr: nil,
			expected: 0,
		},
		{
			name:     "ArrIsEmpty",
			inputArr: []int{},
			expected: 0,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.name, func(t *testing.T) {
			product := Product(tC.inputArr)
			assert.Equal(t, tC.expected, product)
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
			expectedArr: nil,
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
			expectedArr: nil,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.name, func(t *testing.T) {
			result := RemoveFirst(tC.inputArr, tC.inputN)
			assert.Equal(t, tC.expectedArr, result)
		})
	}
}

func TestSum(t *testing.T) {
	testCases := []struct {
		name     string
		inputArr []int
		expected int
	}{
		{
			name:     "Simple",
			inputArr: []int{1, 2, 3},
			expected: 6,
		},
		{
			name:     "ArrIsNil",
			inputArr: nil,
			expected: 0,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.name, func(t *testing.T) {
			sum := Sum(tC.inputArr)
			assert.Equal(t, tC.expected, sum)
		})
	}
}
