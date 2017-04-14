package main

import "fmt"

func spilt(sum int) (x int, y int)  {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {
	fmt.Println(spilt(20))
}
