package geometry

type Geometry interface {
	Perimeter() float64
	Area() float64
}

type Rectangle struct {
	length  int
	breadth int
}
type Square struct {
	side int
}
type Circle struct {
	radius int
}

// Perimeter is used to calculate perimeter of rectangle.
func (r Rectangle) Perimeter() float64 {
	var k float64
	if r.length >= 0 && r.breadth >= 0 {
		k = float64(2 * (r.length + r.breadth))
	}
	return k
}

// Area is used to calculate area of rectangle.
func (r Rectangle) Area() float64 {
	var k float64
	if r.length >= 0 && r.breadth >= 0 {
		k = float64(r.length * r.breadth)
	}
	return k
}

// Perimeter is used to calculate perimeter of square.
func (s Square) Perimeter() float64 {
	var k float64
	if s.side >= 0 {
		k = float64(4 * s.side)
	}
	return k
}

// Area is used to calculate area of square.
func (s Square) Area() float64 {
	var k float64
	if s.side >= 0 {
		k = float64(s.side * s.side)
	}
	return k
}

// Area is used to calculate area of the circle.
func (c Circle) Area() float64 {
	var k float64
	if c.radius >= 0 {
		k = 3.14 * float64(c.radius) * float64(c.radius)
	}
	return k
}

// Perimeter is used to calculate circumference of circle.
func (c Circle) Perimeter() float64 {
	var k float64
	if c.radius >= 0 {
		k = float64(2) * 3.14 * float64(c.radius)
	}
	return k
}
