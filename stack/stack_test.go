package stack

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreate(t *testing.T) {
	testCases := []struct {
		name          string
		expectedStack Stack
	}{
		{
			name:          "Simple",
			expectedStack: Stack{data: []interface{}{}},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.name, func(t *testing.T) {
			s := Create()
			assert.Equal(t, tC.expectedStack, s)
		})
	}
}

func TestLen(t *testing.T) {
	testCases := []struct {
		name           string
		stack          Stack
		expectedLength int
	}{
		{
			name:          "Simple",
			stack: Stack{data: []interface{}{1, 2, 3}},
			expectedLength: 3,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.name, func(t *testing.T) {
			length := tC.stack.Len()
			assert.Equal(t, tC.expectedLength, length)
		})
	}
}

func TestPop(t *testing.T) {
	testCases := []struct {
		name           string
		stack          Stack
		expectedResult interface{}
		expectedStack  Stack
	}{
		{
			name:           "Simple",
			stack:          Stack{data: []interface{}{1, 2, 3}},
			expectedResult: 3,
			expectedStack:  Stack{data: []interface{}{1, 2}},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.name, func(t *testing.T) {
			result := tC.stack.Pop()
			assert.Equal(t, tC.expectedResult, result)
			assert.Equal(t, tC.expectedStack, tC.stack)
		})
	}
}

func TestPush(t *testing.T) {
	testCases := []struct {
		name          string
		stack         Stack
		item          interface{}
		expectedStack Stack
	}{
		{
			name:          "Simple",
			stack:         Stack{data: []interface{}{1, 2, 3}},
			item:          4,
			expectedStack: Stack{data: []interface{}{1, 2, 3, 4}},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.name, func(t *testing.T) {
			tC.stack.Push(tC.item)
			assert.Equal(t, tC.expectedStack, tC.stack)
		})
	}
}

func TestTop(t *testing.T) {
	testCases := []struct {
		name           string
		stack          Stack
		expectedResult interface{}
		expectedStack  Stack
	}{
		{
			name:           "Simple",
			stack:          Stack{data: []interface{}{1, 2, 3}},
			expectedResult: 3,
			expectedStack:  Stack{data: []interface{}{1, 2, 3}},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.name, func(t *testing.T) {
			result := tC.stack.Top()
			assert.Equal(t, tC.expectedResult, result)
			assert.Equal(t, tC.expectedStack, tC.stack)
		})
	}
}
