package main

import "fmt"

type Persons struct {
	name string
}

func main() {
	var i interface{}
	fmt.Printf("%v %T", i, i)
	i = 10
	fmt.Printf("%v %T", i, i)
	i = "abcd"
	fmt.Printf("%v %T", i, i)
	var k interface{} = "Hello"
	fmt.Printf("%v %T", k, k)
	a := k.(string)
	fmt.Printf("%v %T", a, a)
	m, n := k.(int)
	fmt.Printf("%v %T", m, m)
	fmt.Println(m, n)

	var p Persons
	fmt.Printf("%v %T\n", p, p)
	var p1 *Persons
	fmt.Printf("%v %T", p1, p1)
	var a1 *int
	fmt.Printf("%v %T", a1, a1)
}
