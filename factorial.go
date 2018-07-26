package factorial

import "fmt"

func Factorial(n uint) uint {
	if n == 1 || n == 0 {
		return 1
	}
	return n * Factorial(n-1)
}
