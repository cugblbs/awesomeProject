package main

import "fmt"

type Person struct {
	Name string
	Age int
}

func (p Person) String() string  {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

func main() {
	a := Person{"zhudong", 24}
	z := Person{"wl", 22}
	fmt.Println(a, z)
}
