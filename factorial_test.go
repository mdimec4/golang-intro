package factorial

import "testing"

func TestFactorial(t *testing.T) {
	check := func(in uint, expected uint) {
		if out := Factorial(in); out != expected {
			t.Errorf("factorial of %d is %d and not %d", in, expected, out)
		}
	}
	check(0, 1)
	check(1, 1)
	check(2, 2)
	check(3, 6)
	check(4, 24)
}
