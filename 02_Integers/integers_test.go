package integers

import (
	"fmt"
	"testing"
)

func TestInteger(t *testing.T) {
	assert := func(t *testing.T, received int, expected int) {
		t.Helper()
		if received != expected {
			t.Errorf("❌ received %d expected %d", received, expected)
		}
	}

	t.Run("should add two numbers and return the result", func(t *testing.T) {
		received := Add(2, 2)
		expected := 4

		assert(t, received, expected)
	})
}

func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// Output: 6
}
