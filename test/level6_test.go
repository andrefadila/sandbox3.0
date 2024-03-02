package main

import (
	"reflect"
	"testing"

	"sandbox3.0/task"
)

func TestGeneratePrimes(t *testing.T) {
	var empty []int

	tests := []struct {
		limit    int
		expected []int
	}{
		{limit: 10, expected: []int{2, 3, 5, 7}},
		{limit: 20, expected: []int{2, 3, 5, 7, 11, 13, 17, 19}},
		{limit: 30, expected: []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29}},
		{limit: 1, expected: empty},
		{limit: 0, expected: empty},
	}

	for _, test := range tests {
		result := task.GeneratePrimes(test.limit)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("For limit %d, expected %v, but got %v", test.limit, test.expected, result)
		}
	}
}

func TestGeneratePrimesSieve(t *testing.T) {
	tests := []struct {
		limit    int
		expected []int
	}{
		{limit: 10, expected: []int{2, 3, 5, 7}},
		{limit: 20, expected: []int{2, 3, 5, 7, 11, 13, 17, 19}},
		{limit: 30, expected: []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29}},
		{limit: 1, expected: []int{}},
		{limit: 0, expected: []int{}},
	}

	for _, test := range tests {
		result := task.GeneratePrimesSieve(test.limit)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("For limit %d, expected %v, but got %v", test.limit, test.expected, result)
		}
	}
}

func BenchmarkGeneratePrimes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		task.GeneratePrimes(1000)
	}
}

func BenchmarkGeneratePrimesSieve(b *testing.B) {
	for i := 0; i < b.N; i++ {
		task.GeneratePrimesSieve(10000)
	}
}
