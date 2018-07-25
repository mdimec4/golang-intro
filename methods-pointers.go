package main

import "fmt"

type myFloat float64

func (f *myFloat) Scale(s float64) { //Methods with pointer receivers can modify the value
	*f = *f * myFloat(s)
}

func main() {
	v := myFloat(3.14159265)
	v.Scale(100.00)
	fmt.Println(v)
}
