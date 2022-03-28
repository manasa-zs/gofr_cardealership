package Strings

func Count(input string) map[string]int {
	m := make(map[string]int)
	for _, i := range input {
		m[string(i)] += 1
	}
	return m
}
