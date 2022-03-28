package project2

import (
	"fmt"
	"strings"
)

// Bob function takes the input
func Bob(input string) {
	//converting input string to uppercase and storing the value in s
	s := strings.ToUpper(input)
	//trimming the spaces in input and storing the value in k
	k := strings.TrimSpace(input)
	//checking whether the input is upppercase and is a question
	if s == input && input[len(input)-1] == '?' {
		fmt.Println("Calm down, I know what I'm doing!")
	} else /*checking whether the input is nothing*/ if k == "" {
		fmt.Println("Fine. Be that way!")
	} else /*checking whether the input is a question*/ if input[len(input)-1] == '?' {
		fmt.Println("Sure")
	} else /*checking whether input is only uppercase*/ if s == input {
		fmt.Println("Whoa, chill out!")
	} else /*input can be anything else*/ {
		fmt.Println("Whatever.")
	}

}
