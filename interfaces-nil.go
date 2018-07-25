package main

import "fmt"

// START OMIT
func main() {
	var i interface{}
	if i == nil {
		fmt.Printf("1 i == nil (%v, %T)\n", i, i)
	}

	var number int = 42
	i = number
	if i != nil {
		fmt.Printf("2 i != nil (%v, %T)\n", i, i)
	}

	var numberPtr *int = nil
	i = numberPtr
	if i != nil {
		fmt.Printf("3 WTF !!! i != nil (%v, %T)\n", i, i)
	}
	numberPtr = &number
	i = numberPtr
	if i != nil {
		fmt.Printf("4 i != nil (%v, %T)\n", i, i)
		fmt.Println("4 pointer points to value: ", *i.(*int))
	}
}

// END OMIT
