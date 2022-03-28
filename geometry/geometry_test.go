package geometry

import (
	"reflect"
	"testing"
)

func TestPerimeterGeometry(t *testing.T) {
	Testcases := []struct {
		des      string
		input    Geometry
		expected float64
	}{
		{"test case for perimeter of rectangle", Rectangle{3, 4}, 14},
		{"test case for perimeter of rectangle", Rectangle{10, 20}, 60},
		{"test case for perimeter of circle", Circle{4}, 25.12},
		{"test case for perimeter of circle", Circle{8}, 50.24},
		{"test case for perimeter of square", Square{3}, 12},
		{"test case for perimeter of square", Square{10}, 40},
	}
	for _, v := range Testcases {
		output := (v.input).Perimeter()
		if !(reflect.DeepEqual(output, v.expected)) {
			t.Errorf("%v have \n %v want", output, v.expected)
		}
	}
}

func TestAreaGeometry(t *testing.T) {
	Testcases := []struct {
		des      string
		input    Geometry
		expected float64
	}{
		{"test case for area of rectangle", Rectangle{3, 4}, 12},
		{"test case for area of rectangle", Rectangle{10, 20}, 200},
		{"test case for area of circle", Circle{4}, 50.24},
		{"test case for area of circle", Circle{8}, 200.96},
		{"test case for area of square", Square{3}, 9},
		{"test case for area of square", Square{10}, 100},
	}
	for _, v := range Testcases {
		output := (v.input).Area()
		if !(reflect.DeepEqual(output, v.expected)) {
			t.Errorf("%v have \n %v want", output, v.expected)
		}
	}
}

func BenchmarkAreaGeometry(b *testing.B) {
	for i := 0; i < b.N; i++ {
		(Circle{4}).Area()
		(Rectangle{4, 3}).Area()
		(Square{4}).Area()
	}
}

func BenchmarkCircle_Perimeter(b *testing.B) {
	for i := 0; i < b.N; i++ {
		(Circle{4}).Perimeter()
		(Rectangle{4, 3}).Perimeter()
		(Square{4}).Perimeter()
	}
}
