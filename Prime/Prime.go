package Prime

/*Prime is used to find the array of prime numbers within the specified limit*/
func Prime(n int) []int {
	var a []int
	for i := 2; i < n; i++ {
		if (i%2 != 0 && i%3 != 0 && i%5 != 0 && i%7 != 0) || (i == 2 || i == 3 || i == 5 || i == 7) {
			a = append(a, i)
		}
	}
	return a
}
