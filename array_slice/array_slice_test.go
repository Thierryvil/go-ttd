package array_slice

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {

	t.Run("Should be return 15 with 5 numbers", func(t *testing.T) {
		nums := []int{1, 2, 3, 4, 5}
		result := Sum(nums)
		expected := 15

		if result != expected {
			t.Errorf("result: '%d', expected: '%d'", result, expected)
		}
	})

	t.Run("Should be return 6 with 3 numbers", func(t *testing.T) {
		nums := []int{1, 2, 3}
		result := Sum(nums)
		expected := 6

		if result != expected {
			t.Errorf("result: '%d', expected: '%d'", result, expected)
		}
	})
}

func TestSumAll(t *testing.T) {
	result := SumAll([]int{1, 1, 1}, []int{2, 9})
	expected := []int{3, 11}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("result: '%d', expected: '%d'", result, expected)
	}
}

func TestSumAllRest(t *testing.T) {
	checkSum := func(t *testing.T, result, expected []int) {
		t.Helper()
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("result: '%d', expected: '%d'", result, expected)
		}
	}
	t.Run("Should be return a slice []int{2, 9}", func(t *testing.T) {
		result := SumAllRest([]int{1, 2}, []int{0, 9})
		expected := []int{2, 9}
		checkSum(t, result, expected)

	})

	t.Run("Should be return a slice []int{0, 9}", func(t *testing.T) {
		result := SumAllRest([]int{}, []int{3, 4, 5})
		expected := []int{0, 9}

		checkSum(t, result, expected)

	})
}
