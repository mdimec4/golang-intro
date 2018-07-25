package main

import (
	"fmt"
	"math"
)

// START OMIT
type Abser interface { // Interface definition
	Abs() float64 // requires Abs() method with this signature
}

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 { // Vertex implemnts Abser
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

type MyFloat float64

func (f MyFloat) Abs() float64 { // myFloat also implements Abser
	return math.Abs(float64(f))
}

func main() {
	v := []Abser{Vertex{3, 4}, MyFloat(-3.14)}
	for _, a := range v {
		fmt.Println(a.Abs())
	}
}

//END OMIT
