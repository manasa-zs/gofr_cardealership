package triangles

// Triangles used to find whether a triangle is equilateral or isosceles or scalene
func Triangles(a []int) string {
	k := "Not a Triangle"
	if a[0]+a[1] > a[2] && a[1]+a[2] > a[0] && a[0]+a[2] > a[1] && a[0] > 0 && a[1] > 0 && a[2] > 0 {
		if a[0] == a[1] && a[0] == a[2] && a[1] == a[2] {
			k = "Equilateral Triangle"
		} else if a[0] == a[1] || a[1] == a[2] || a[0] == a[2] {
			k = "Isosceles Triangle"
		} else {
			k = "Scalene Triangle"
		}
	}
	return k
}
