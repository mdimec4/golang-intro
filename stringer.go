package main

import "fmt"

// START OMIT
type Person struct {
	Name string
	Age  uint
}

func (p Person) String() string {
	return fmt.Sprintf("%s (%d years)", p.Name, p.Age)
}

func main() {
	a := Person{"Arthur Dent", 42}
	fmt.Printf("with fmt.Printf(): %s\n", a)
	fmt.Println("with fmt.Println(): ", a)
}

// END OMIT
