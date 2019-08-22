package arrayslice

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	assert := func(t *testing.T, received int, expected int, numbers []int) {
		t.Helper()
		if received != expected {
			t.Errorf("❌ received %d expected %d with %v", received, expected, numbers)
		}
	}

	t.Run("should sum all numbers of the given array", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		received := Sum(numbers)
		expected := 15

		assert(t, received, expected, numbers)
	})
}

func TestSumAllTails(t *testing.T) {
	assert := func(t *testing.T, received []int, expected []int, arr1 []int, arr2 []int) {
		t.Helper()
		if !reflect.DeepEqual(received, expected) {
			t.Errorf("❌ received %d expected %v with %v and %v", received, expected, arr1, arr2)
		}
	}

	t.Run("should sum the total of the two given arrays", func(t *testing.T) {
		received := SumAllTails([]int{1, 2}, []int{0, 9})
		expected := []int{2, 9}

		assert(t, received, expected, []int{1, 2}, []int{0, 9})
	})

	t.Run("should safely sum empty slices", func(t *testing.T) {
		received := SumAllTails([]int{}, []int{3, 4, 5})
		expected := []int{0, 9}

		assert(t, received, expected, []int{1, 2}, []int{0, 9})
	})
}
