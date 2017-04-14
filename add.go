package main

import (
	"fmt"
)

func add(x,y int) int  {
	return x+y
}

func main() {
	t := add(2,3)
	fmt.Println(t)
}
