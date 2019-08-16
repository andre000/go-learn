package iteration

import "testing"

func TestRepeat(t *testing.T) {
	assert := func(t *testing.T, received string, expected string) {
		t.Helper()
		if received != expected {
			t.Errorf("❌ received %q expected %q", received, expected)
		}
	}

	t.Run("should repeat the character 'a' fice times", func(t *testing.T) {
		received := Repeat("a")
		expected := "aaaaa"

		assert(t, received, expected)
	})
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a")
	}
}
