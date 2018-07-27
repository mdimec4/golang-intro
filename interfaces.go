package main

import (
	"fmt"
	"math"
)

// START ABSER OMIT
type Abser interface { // Interface definition
	Abs() float64 // requires Abs() method with this signature
}

//END ABSER OMIT

// START VERTEX OMIT

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 { // Vertex implemnts Abser
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// END VERTEX OMIT

// START MYFLOAT OMIT
type MyFloat float64

func (f MyFloat) Abs() float64 { // myFloat also implements Abser
	return math.Abs(float64(f))
}

// END MYFLOAT OMIT

// START MAIN OMIT
// PrintAbs accepts anything that implements Abser as a parameter
func PrintAbs(v Abser) {
	a := v.Abs()
	fmt.Println(a)
}

func main() {
	v := []Abser{Vertex{3, 4}, MyFloat(-3.14)}
	for _, a := range v {
		PrintAbs(a)
	}
}

// END MAIN OMIT
