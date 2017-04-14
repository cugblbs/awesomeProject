package main

import "fmt"

func main() {

	p := []string{"1","2", "3", "4", "5", "6"}

	fmt.Println("p ==", p[1:])

	for i := 0 ; i<len(p) ; i++ {
		fmt.Printf("p[%d] is %s\n", i, p[i])
	}
}
