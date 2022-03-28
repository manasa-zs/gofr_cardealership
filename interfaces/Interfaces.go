package main

import "fmt"

type ReadWriter interface {
	Read() string
	Write(string)
}

type Person struct {
	Name string
}

func (p Person) Read() string {
	return p.Name
}
func (p Person) Write(input string) {
	p.Name = input
}

func main() {
	var r ReadWriter = &Person{"abc"}
	var rw ReadWriter = &Person{"bcd"}
	var k = Person{}
	fmt.Println(r.Read())
	rw.Write("Honey")
	fmt.Printf("%T %v", k, k)
}
