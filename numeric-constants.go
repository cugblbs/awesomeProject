package main

import "fmt"

const (
	Big = 1 << 32
	Small = Big >> 32
)

func needInt(x int) int  {
	return x*10 + 1
}
func needFloat(x float64) float64  {
	return x*1.0
}

func main() {
	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
}
