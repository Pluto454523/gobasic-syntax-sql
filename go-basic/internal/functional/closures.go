package functional

import "fmt"

func outer() func(string) {
	message := "Hello, "
	inner := func(name string) {
		fmt.Println(message + name)
	}
	return inner
}

func Closures() {
	greeting := outer()
	greeting("Alice")
	greeting("Bob")
}
