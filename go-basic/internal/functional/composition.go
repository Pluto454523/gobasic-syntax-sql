package functional

import "fmt"

func addOne(x int) int {
	return x + 1
}

func multiplyByTwo(x int) int {
	return x * 2
}

func CompositionFunction() {
	compose := func(f, g func(int) int) func(int) int {
		return func(x int) int {
			return f(g(x))
		}
	}

	addThenMultiply := compose(multiplyByTwo, addOne)
	result := addThenMultiply(5)
	fmt.Println("Result:", result) // Output: 12 (5 + 1) * 2
}
