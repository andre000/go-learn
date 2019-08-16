package shapes

import "testing"

func TestPerimeter(t *testing.T) {
	assert := func(t *testing.T, received float64, expected float64) {
		t.Helper()
		if received != expected {
			t.Errorf("❌ received %.2f expected %.2f", received, expected)
		}
	}

	t.Run("should be able to correctly calculate the perimeter", func(t *testing.T) {
		received := Perimeter(10.0, 10.0)
		expected := 40.0

		assert(t, received, expected)
	})
}
