package array

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAppendString(t *testing.T) {
	testCases := []struct {
		name        string
		inputArr    []string
		inputS      string
		expectedArr []string
	}{
		{
			name:        "Simple",
			inputArr:    []string{"a", "b", "c"},
			inputS:      "d",
			expectedArr: []string{"a", "b", "c", "d"},
		},
		{
			name:        "ArrIsNil",
			inputArr:    nil,
			inputS:      "a",
			expectedArr: []string{"a"},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.name, func(t *testing.T) {
			arr := AppendString(tC.inputArr, tC.inputS)
			assert.Equal(t, tC.expectedArr, arr)
		})
	}
}

func TestContainsString(t *testing.T) {
	testCases := []struct {
		name     string
		inputArr []string
		inputS   string
		expected bool
	}{
		{
			name:     "Contains",
			inputArr: []string{"a", "b", "c"},
			inputS:   "a",
			expected: true,
		},
		{
			name:     "NotContains",
			inputArr: []string{"a", "b", "c"},
			inputS:   "d",
			expected: false,
		},
		{
			name:     "ArrIsNil",
			inputArr: nil,
			inputS:   "a",
			expected: false,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.name, func(t *testing.T) {
			contains := ContainsString(tC.inputArr, tC.inputS)
			assert.Equal(t, tC.expected, contains)
		})
	}
}

func TestCountString(t *testing.T) {
	testCases := []struct {
		name      string
		inputArr  []string
		inputPred UnaryPredicateString
		expected  int
	}{
		{
			name:      "Short",
			inputArr:  []string{"a", "bb", "cc"},
			inputPred: func(s string) bool {return len(s) < 2},
			expected:  1,
		},
		{
			name:      "Long",
			inputArr:  []string{"a", "bb", "cc"},
			inputPred: func(s string) bool {return len(s) >= 2},
			expected:  2,
		},
		{
			name:      "ArrIsNil",
			inputArr:  nil,
			inputPred: func(s string) bool {return len(s) <= 2},
			expected:  0,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.name, func(t *testing.T) {
			count := CountString(tC.inputArr, tC.inputPred)
			assert.Equal(t, tC.expected, count)
		})
	}
}

func TestFilterString(t *testing.T) {
	testCases := []struct {
		name      string
		inputArr  []string
		inputPred UnaryPredicateString
		expected  []string
	}{
		{
			name:      "Short",
			inputArr:  []string{"a", "bb", "cc"},
			inputPred: func(s string) bool {return len(s) < 2},
			expected:  []string{"a"},
		},
		{
			name:      "Long",
			inputArr:  []string{"a", "bb", "cc"},
			inputPred: func(s string) bool {return len(s) >= 2},
			expected:  []string{"bb", "cc"},
		},
		{
			name:      "ArrIsNil",
			inputArr:  nil,
			inputPred: func(s string) bool {return len(s) <= 2},
			expected:  []string{},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.name, func(t *testing.T) {
			result := FilterString(tC.inputArr, tC.inputPred)
			assert.Equal(t, tC.expected, result)
		})
	}
}

func TestLenString(t *testing.T) {
	testCases := []struct {
		name        string
		inputArr    []string
		expectedLen int
	}{
		{
			name:        "Simple",
			inputArr:    []string{"a", "b", "c"},
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
			length := LenString(tC.inputArr)
			assert.Equal(t, tC.expectedLen, length)
		})
	}
}

func TestMapString(t *testing.T) {
	testCases := []struct {
		name     string
		inputArr []string
		inputF   UnaryModifierString
		expected []string
	}{
		{
			name:     "Double",
			inputArr: []string{"a", "b", "c"},
			inputF:   func(s string) string {return s + s},
			expected: []string{"aa", "bb", "cc"},
		},
		{
			name:     "ArrIsNil",
			inputArr: nil,
			inputF:   func(s string) string {return s + s},
			expected: []string{},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.name, func(t *testing.T) {
			result := MapString(tC.inputArr, tC.inputF)
			assert.Equal(t, tC.expected, result)
		})
	}
}

func TestPermutateString(t *testing.T) {
	testCases := []struct {
		name     string
		inputArr []string
		expected [][]string
	}{
		{
			name:     "Simple",
			inputArr: []string{"a", "b", "c"},
			expected: [][]string{{"a", "b", "c"}, {"b", "a", "c"},
				{"c", "b", "a"}, {"b", "c", "a"}, {"c", "a", "b"},
				{"a", "c", "b"}},
		},
		{
			name:     "ArrIsNil",
			inputArr: nil,
			expected: [][]string{},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.name, func(t *testing.T) {
			result := PermutateString(tC.inputArr)
			assert.Equal(t, tC.expected, result)
		})
	}
}

func TestRemoveAllString(t *testing.T) {
	testCases := []struct {
		name        string
		inputArr    []string
		originalArr []string
		inputS      string
		expectedArr []string
	}{
		{
			name:        "Simple",
			inputArr:    []string{"a", "b", "c", "b"},
			originalArr: []string{"a", "b", "c", "b"},
			inputS:      "b",
			expectedArr: []string{"a", "c"},
		},
		{
			name:        "ArrIsNil",
			inputArr:    nil,
			originalArr: nil,
			inputS:      "a",
			expectedArr: []string{},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.name, func(t *testing.T) {
			result := RemoveAllString(tC.inputArr, tC.inputS)
			assert.Equal(t, tC.expectedArr, result)
			assert.Equal(t, tC.originalArr, tC.inputArr)
		})
	}
}

func TestRemoveFirstString(t *testing.T) {
	testCases := []struct {
		name        string
		inputArr    []string
		originalarr []string
		inputS      string
		expectedArr []string
	}{
		{
			name:        "Simple",
			inputArr:    []string{"a", "b", "c", "b"},
			originalarr: []string{"a", "b", "c", "b"},
			inputS:      "b",
			expectedArr: []string{"a", "c", "b"},
		},
		{
			name:        "ArrIsNil",
			inputArr:    nil,
			originalarr: nil,
			inputS:      "a",
			expectedArr: []string{},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.name, func(t *testing.T) {
			result := RemoveFirstString(tC.inputArr, tC.inputS)
			assert.Equal(t, tC.expectedArr, result)
			assert.Equal(t, tC.originalarr, tC.inputArr)
		})
	}
}
