package shapes

import "testing"

func TestArea(t *testing.T) {
	shapeTests := []struct {
		name    string
		shape   Shape
		hasArea float64
	}{
		{name: "Circle", shape: Circle{10.0}, hasArea: 314.1592653589793},
		{name: "Triangle", shape: Triangle{12, 6}, hasArea: 36.0},
		{name: "Rectangle", shape: Rectangle{12.0, 6.0}, hasArea: 72.0},
	}

	for _, test := range shapeTests {
		t.Run(test.name, func(t *testing.T) {
			if got := test.shape.Area(); got != test.hasArea {
				t.Errorf("%#v.Area(): got %.2f, want %.2f", test.shape, got, test.hasArea)
			}
		})
	}
}

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	got := rectangle.Perimeter()
	want := 40.0

	if got != want {
		t.Errorf("%#v.Perimeter(): got %.2f, want %.2f", rectangle, got, want)
	}
}
