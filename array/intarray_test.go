package array

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/taskat/golang-utility/intmath"
)

func TestAppendInt(t *testing.T) {
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
			arr := AppendInt(tC.inputArr, tC.inputN)
			assert.Equal(t, tC.expectedArr, arr)
		})
	}
}

func TestContainsInt(t *testing.T) {
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
			contains := ContainsInt(tC.inputArr, tC.inputN)
			assert.Equal(t, tC.expected, contains)
		})
	}
}

func TestCountInt(t *testing.T) {
	testCases := []struct {
		name      string
		inputArr  []int
		inputPred UnaryPredicateInt
		expected  int
	}{
		{
			name:      "Even",
			inputArr:  []int{1, 2, 3},
			inputPred: intmath.Even,
			expected:  1,
		},
		{
			name:      "Odd",
			inputArr:  []int{1, 2, 3},
			inputPred: intmath.Odd,
			expected:  2,
		},
		{
			name:      "ArrIsNil",
			inputArr:  nil,
			inputPred: intmath.Even,
			expected:  0,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.name, func(t *testing.T) {
			count := CountInt(tC.inputArr, tC.inputPred)
			assert.Equal(t, tC.expected, count)
		})
	}
}

func TestFilterInt(t *testing.T) {
	testCases := []struct {
		name      string
		inputArr  []int
		inputPred UnaryPredicateInt
		expected  []int
	}{
		{
			name:      "Even",
			inputArr:  []int{1, 2, 3},
			inputPred: intmath.Even,
			expected:  []int{2},
		},
		{
			name:      "Odd",
			inputArr:  []int{1, 2, 3},
			inputPred: intmath.Odd,
			expected:  []int{1, 3},
		},
		{
			name:      "ArrIsNil",
			inputArr:  nil,
			inputPred: intmath.Even,
			expected:  []int{},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.name, func(t *testing.T) {
			result := FilterInt(tC.inputArr, tC.inputPred)
			assert.Equal(t, tC.expected, result)
		})
	}
}

func TestLenInt(t *testing.T) {
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
			length := LenInt(tC.inputArr)
			assert.Equal(t, tC.expectedLen, length)
		})
	}
}

func TestMapInt(t *testing.T) {
	testCases := []struct {
		name     string
		inputArr []int
		inputF   UnaryModifierInt
		expected []int
	}{
		{
			name:     "Square",
			inputArr: []int{1, 2, 3},
			inputF:   intmath.Square,
			expected: []int{1, 4, 9},
		},
		{
			name:     "ArrIsNil",
			inputArr: nil,
			inputF:   intmath.Square,
			expected: []int{},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.name, func(t *testing.T) {
			result := MapInt(tC.inputArr, tC.inputF)
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
			inputArr:    []int{2, 1, 3},
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

func TestPermutateInt(t *testing.T) {
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
			result := PermutateInt(tC.inputArr)
			assert.Equal(t, tC.expected, result)
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

func TestRemoveAllInt(t *testing.T) {
	testCases := []struct {
		name        string
		inputArr    []int
		originalArr []int
		inputN      int
		expectedArr []int
	}{
		{
			name:        "Simple",
			inputArr:    []int{1, 2, 3, 2},
			originalArr: []int{1, 2, 3, 2},
			inputN:      2,
			expectedArr: []int{1, 3},
		},
		{
			name:        "ArrIsNil",
			inputArr:    nil,
			originalArr: nil,
			inputN:      1,
			expectedArr: []int{},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.name, func(t *testing.T) {
			result := RemoveAllInt(tC.inputArr, tC.inputN)
			assert.Equal(t, tC.expectedArr, result)
			assert.Equal(t, tC.originalArr, tC.inputArr)
		})
	}
}

func TestRemoveFirstInt(t *testing.T) {
	testCases := []struct {
		name        string
		inputArr    []int
		originalarr []int
		inputN      int
		expectedArr []int
	}{
		{
			name:        "Simple",
			inputArr:    []int{1, 2, 3, 2},
			originalarr: []int{1, 2, 3, 2},
			inputN:      2,
			expectedArr: []int{1, 3, 2},
		},
		{
			name:        "ArrIsNil",
			inputArr:    nil,
			originalarr: nil,
			inputN:      1,
			expectedArr: []int{},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.name, func(t *testing.T) {
			result := RemoveFirstInt(tC.inputArr, tC.inputN)
			assert.Equal(t, tC.expectedArr, result)
			assert.Equal(t, tC.originalarr, tC.inputArr)
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
