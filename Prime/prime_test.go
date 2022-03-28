package Prime

import (
	"fmt"
	"assignmentsgithub.com/stretchr/testify/assert"
	"testing"
)

type Testcase []struct {
	des      string
	input    int
	expected []int
}

/* TestPrime is used to write the testing for Prime function*/
func TestPrime(t *testing.T) {
	Testcases := Testcase{
		{des: "passed case 0", input: -1, expected: []int(nil)},
		{des: "passed case 1", input: 5, expected: []int{2, 3}},
		{des: "passed case 2", input: 6, expected: []int{2, 3, 5}},
		{des: "passed case 3", input: 10, expected: []int{2, 3, 5, 7}},
	}

	for _, i := range Testcases {
		output := Prime(i.input)

		if !(assert.Equal(t, output, i.expected)) {

			t.Errorf("expected output %v not equal to output %v", i.expected, output)
		}
	}
}

/* BenchmarkPrime is used to find the performance of the Prime function*/
func BenchmarkPrime(b *testing.B) {
	Testcases1 := Testcase{
		{des: "passed case 1", input: 5, expected: []int{2, 3}},
		{des: "passed case 2", input: 6, expected: []int{2, 3, 5}},
		{des: "passed case 3", input: 10, expected: []int{2, 3, 5, 7}},
	}

	for _, i := range Testcases1 {
		b.Run(fmt.Sprintf("input size_%d", i.input), func(b *testing.B) {
			for j := 0; j < b.N; j++ {
				Prime(i.input)
			}
		})

	}
}
