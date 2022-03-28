package project2

import (
	"fmt"
	"strconv"
	"unicode"
)

func Map() {
	a := "jm12jkashx93"
	sum := 0
	s := ""
	for i := 0; i < len(a); i++ {
		if i == len(a)-1 {
			if unicode.IsNumber(rune(a[i])) {
				s += string(a[i])
				k, _ := strconv.Atoi(string(s))
				sum += k
			}
		} else if unicode.IsNumber(rune(a[i])) {
			s += string(a[i])
		} else {
			k, _ := strconv.Atoi(string(s))
			sum += k
			s = ""
		}

	}
	fmt.Println(sum, s)
}


