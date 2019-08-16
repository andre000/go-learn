package iteration

import (
	"fmt"
	"strings"
	"testing"
)

func TestRepeat(t *testing.T) {
	assert := func(t *testing.T, received string, expected string) {
		t.Helper()
		if received != expected {
			t.Errorf("❌ received %q expected %q", received, expected)
		}
	}

	t.Run("should repeat the character 'a' five times", func(t *testing.T) {
		received := Repeat("a", 20)
		expected := strings.Repeat("a", 20)

		assert(t, received, expected)
	})
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}

func ExampleRepeat() {
	str := Repeat("Z", 3)
	fmt.Printf(str)
	// Output: ZZZ
}
