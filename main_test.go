package main

// Пишите тесты в этом файле
import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateRandomElements(t *testing.T) {
	tests := []struct {
		name   string
		size   int
		result int
	}{
		{name: "nil", size: 0, result: 0},
		{name: "one", size: 1, result: 1},
		{name: "negative", size: -15, result: 0},
		{name: "normal", size: 354, result: 354},
	}

	for _, test := range tests {
		testSlice := generateRandomElements(test.size)
		assert.Equal(t, test.result, len(testSlice))
	}
}

func TestMaximum(t *testing.T) {
	tests := []struct {
		data   []int
		result int
	}{
		{data: nil, result: 0},
		{data: []int{}, result: 0},
		{data: []int{49}, result: 49},
		{data: []int{33, 86}, result: 86},
		{data: []int{53, 423336, 12, 1, 9, 56, 17, 2}, result: 423336},
		{data: []int{10, 6, 41, 26, 7227, 876, 240, 2, 223, 4871, 6161, 8}, result: 7227},
	}

	for _, test := range tests {
		testMax := maximum(test.data)
		assert.Equal(t, test.result, testMax)
	}
}
