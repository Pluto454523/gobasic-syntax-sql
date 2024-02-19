package functional

import "fmt"

func add(a, b int) int {
	return a + b
}

func subtract(a, b int) int {
	return a - b
}

func FirstClassFunction() {
	var operation func(int, int) int
	operation = add
	result := operation(5, 3)
	fmt.Println("Result of addition:", result)

	operation = subtract
	result = operation(5, 3)
	fmt.Println("Result of subtraction:", result)
}
