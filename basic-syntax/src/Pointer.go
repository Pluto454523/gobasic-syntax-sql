package basicSyntax

import "fmt"

func PointerTutorial() {

	// ** Pass by vaule

	a := 3
	b := a
	fmt.Println("[ Pass by vaule ]")
	fmt.Println(" a =", a)
	fmt.Println(" b =", b)

	// ** Pass by vaule Function
	c := 0
	sumByValue(c)
	fmt.Println(" Pass by vaule Function =", c)

	// ** Pass by reference
	x := 0
	y := &x
	x = 50
	*y = 10
	fmt.Println("[ Pass by reference ]")
	fmt.Println("", &x, "| x =", x)
	fmt.Println("", y, "| y =", *y)

	// ** Pass by reference Function
	z := 0
	sumByRef(&z)
	fmt.Println(" Pass by reference Function =", z)

}

func sumByRef(result *int) {
	a := 80
	b := 20
	*result = a + b
}

func sumByValue(result int) {
	a := 80
	b := 20
	result = a + b
}
