package array

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAppend(t *testing.T) {
	testCases := []struct {
		name        string
		inputArr    []interface{}
		inputN      interface{}
		expectedArr []interface{}
	}{
		{
			name:        "Simple",
			inputArr:    []interface{}{1, 2, 3},
			inputN:      4,
			expectedArr: []interface{}{1, 2, 3, 4},
		},
		{
			name:        "ArrIsNil",
			inputArr:    nil,
			inputN:      1,
			expectedArr: []interface{}{1},
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
		inputArr []interface{}
		inputN   interface{}
		expected bool
	}{
		{
			name:     "Contains",
			inputArr: []interface{}{1, 2, 3},
			inputN:   1,
			expected: true,
		},
		{
			name:     "NotContains",
			inputArr: []interface{}{1, 2, 3},
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
		inputArr  []interface{}
		inputPred UnaryPredicate
		expected  interface{}
	}{
		{
			name:      "Even",
			inputArr:  []interface{}{1, 2, 3},
			inputPred: func(i interface{}) bool { return i.(int)%2 == 0 },
			expected:  1,
		},
		{
			name:      "Odd",
			inputArr:  []interface{}{1, 2, 3},
			inputPred: func(i interface{}) bool { return i.(int)%2 == 1 },
			expected:  2,
		},
		{
			name:      "ArrIsNil",
			inputArr:  nil,
			inputPred: func(i interface{}) bool { return i.(int)%2 == 0 },
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
		inputArr  []interface{}
		inputPred UnaryPredicate
		expected  []interface{}
	}{
		{
			name:      "Even",
			inputArr:  []interface{}{1, 2, 3},
			inputPred: func(i interface{}) bool { return i.(int)%2 == 0 },
			expected:  []interface{}{2},
		},
		{
			name:      "Odd",
			inputArr:  []interface{}{1, 2, 3},
			inputPred: func(i interface{}) bool { return i.(int)%2 == 1 },
			expected:  []interface{}{1, 3},
		},
		{
			name:      "ArrIsNil",
			inputArr:  nil,
			inputPred: func(i interface{}) bool { return i.(int)%2 == 0 },
			expected:  []interface{}{},
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
		inputArr    []interface{}
		expectedLen int
	}{
		{
			name:        "Simple",
			inputArr:    []interface{}{1, 2, 3},
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
		inputArr []interface{}
		inputF   UnaryModifier
		expected []interface{}
	}{
		{
			name:     "Square",
			inputArr: []interface{}{1, 2, 3},
			inputF:   func(i interface{}) interface{} { return i.(int) * i.(int) },
			expected: []interface{}{1, 4, 9},
		},
		{
			name:     "ArrIsNil",
			inputArr: nil,
			inputF:   func(i interface{}) interface{} { return i.(int) * i.(int) },
			expected: []interface{}{},
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
		inputArr []interface{}
		expected [][]interface{}
	}{
		{
			name:     "Simple",
			inputArr: []interface{}{1, 2, 3},
			expected: [][]interface{}{{1, 2, 3}, {2, 1, 3}, {3, 2, 1}, {2, 3, 1}, {3, 1, 2}, {1, 3, 2}},
		},
		{
			name:     "ArrIsNil",
			inputArr: nil,
			expected: [][]interface{}{},
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
		inputArr    []interface{}
		originalArr []interface{}
		inputN      interface{}
		expectedArr []interface{}
	}{
		{
			name:        "Simple",
			inputArr:    []interface{}{1, 2, 3, 2},
			originalArr: []interface{}{1, 2, 3, 2},
			inputN:      2,
			expectedArr: []interface{}{1, 3},
		},
		{
			name:        "ArrIsNil",
			inputArr:    nil,
			originalArr: nil,
			inputN:      1,
			expectedArr: []interface{}{},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.name, func(t *testing.T) {
			result := RemoveAll(tC.inputArr, tC.inputN)
			assert.Equal(t, tC.expectedArr, result)
			assert.Equal(t, tC.originalArr, tC.inputArr)
		})
	}
}

func TestRemoveFirst(t *testing.T) {
	testCases := []struct {
		name        string
		inputArr    []interface{}
		originalarr []interface{}
		inputN      interface{}
		expectedArr []interface{}
	}{
		{
			name:        "Simple",
			inputArr:    []interface{}{1, 2, 3, 2},
			originalarr: []interface{}{1, 2, 3, 2},
			inputN:      2,
			expectedArr: []interface{}{1, 3, 2},
		},
		{
			name:        "ArrIsNil",
			inputArr:    nil,
			originalarr: nil,
			inputN:      1,
			expectedArr: []interface{}{},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.name, func(t *testing.T) {
			result := RemoveFirst(tC.inputArr, tC.inputN)
			assert.Equal(t, tC.expectedArr, result)
			assert.Equal(t, tC.originalarr, tC.inputArr)
		})
	}
}
