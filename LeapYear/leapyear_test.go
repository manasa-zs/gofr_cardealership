package LeapYear

import (
	"fmt"
	"reflect"
	"testing"
)

type Testcase []struct {
	des      string
	input    int
	expected string
}

/* TestPrime is used to write the testing for Prime function*/
func TestLeapYear(t *testing.T) {
	Testcases := Testcase{
		{des: "passed case 0", input: 2000, expected: "Leap year"},
		{des: "passed case 1", input: 2017, expected: "Non Leap Year"},
		{des: "passed case 2", input: 2004, expected: "Leap year"},
		{des: "passed case 3", input: 76, expected: "Leap year"},
	}

	for _, i := range Testcases {
		output := LeapYear(i.input)

		if !(reflect.DeepEqual(output, i.expected)) {

			t.Errorf("expected output %v not equal to output %v", i.expected, output)
		}
	}
}

/* BenchmarkPrime is used to find the performance of the Prime function*/
func BenchmarkLeapYear(b *testing.B) {
	Testcases1 := Testcase{
		{des: "passed case 0", input: 2000, expected: "Leap year"},
		{des: "passed case 1", input: 2017, expected: "Non Leap year"},
		{des: "passed case 2", input: 2004, expected: "Leap year"},
		{des: "passed case 3", input: 2016, expected: "Leap year"},
	}

	for _, i := range Testcases1 {
		b.Run(fmt.Sprintf("input size_%d", i.input), func(b *testing.B) {
			for j := 0; j < b.N; j++ {
				LeapYear(i.input)
			}
		})

	}
}
