package functional

import "fmt"

func factorial(n int) int {
	if n <= 1 {
		return 1
	}
	return n * factorial(n-1)
}

func RecursionFunction() {
	fmt.Println("Factorial of 5 is:", factorial(5))
}
