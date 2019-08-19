package shapes

import "testing"

func TestPerimeter(t *testing.T) {
	assert := func(t *testing.T, received float64, expected float64) {
		t.Helper()
		if received != expected {
			t.Errorf("❌ received %.2f expected %.2f", received, expected)
		}
	}

	t.Run("should be able to correctly calculate the perimeter of rectangles", func(t *testing.T) {
		rectangle := Rectangle{10.0, 10.0}

		received := Perimeter(rectangle)
		expected := 40.0

		assert(t, received, expected)
	})

}

func TestArea(t *testing.T) {
	assert := func(t *testing.T, shape Shape, received float64, expected float64) {
		t.Helper()
		if received != expected {
			t.Errorf("❌ %#v: received %.2f expected %.2f", shape, received, expected)
		}
	}

	checkArea := func(t *testing.T, shape Shape, expected float64) {
		t.Helper()
		received := shape.Area()
		assert(t, shape, received, expected)
	}

	t.Run("should be able to calculate the area of any shape", func(t *testing.T) {
		areaTests := []struct {
			shape    Shape
			expected float64
		}{
			{shape: Circle{Radius: 10}, expected: 314.1592653589793},
			{shape: Triangle{Base: 12, Height: 6}, expected: 36.0},
			{shape: Rectangle{Width: 12.0, Height: 6.0}, expected: 72.0},
		}

		for _, tt := range areaTests {
			checkArea(t, tt.shape, tt.expected)
		}
	})

}
