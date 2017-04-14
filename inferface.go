package main

import (
	"math"
	"fmt"
)

type MYFLOAT float64

type Abser interface {
	Abs() float64
}

func (f MYFLOAT) Abs() float64  {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Abs() float64  {
	return math.Sqrt(v.X*v.X + v.Y * v.Y)
}


func main() {
	var a Abser
	f := MYFLOAT(3)
	v := Vertex{3, 4}
	a = f
	a = &v
	fmt.Println(a.Abs())
}
