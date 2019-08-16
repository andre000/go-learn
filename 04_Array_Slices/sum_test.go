package array_slice

import "testing"

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
