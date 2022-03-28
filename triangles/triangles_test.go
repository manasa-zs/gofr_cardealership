package triangles

import (
	"reflect"
	"testing"
)

type Testcase []struct {
	des      string
	input    []int
	expected string
}

/* TestPrime is used to write the testing for Prime function*/
func TestTriangles(t *testing.T) {
	testcases := Testcase{
		{"passed case for ", []int{3, 4, 5}, "Scalene Triangle"},
		{des: "passed case 1", input: []int{-1, -1, -1}, expected: "Not a Triangle"},
		{des: "passed case 2", input: []int{6, 6, 6}, expected: "Equilateral Triangle"},
		{des: "passed case 3", input: []int{6, 5, 6}, expected: "Isosceles Triangle"},
	}

	for _, i := range testcases {
		output := Triangles(i.input)

		if !(reflect.DeepEqual(output, i.expected)) {

			t.Errorf("expected output %v not equal to output %v", i.expected, output)
		}
	}
}

func BenchmarkTriangles(b *testing.B) {
	testcases := Testcase{
		{des: "passed case 0", input: []int{1, 2, 3}, expected: "Scalene Triangle"},
		{des: "passed case 1", input: []int{1, 2, 2}, expected: "Isosceles Triangle"},
		{des: "passed case 2", input: []int{1, 1, 1}, expected: "Equilateral Triangle"},
		{des: "passed case 3", input: []int{1, 2, 1}, expected: "Isosceles Triangle"},
	}
	for i := 0; i < b.N; i++ {
		for _, v := range testcases {
			Triangles(v.input)
		}
	}
}
