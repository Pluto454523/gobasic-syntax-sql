package basicSyntax

import (
	"fmt"
)

type person struct {
	name string
	age  int
}

func NewPerson(name string, age int) person {
	return person{name, age}
}

// ** Receiver method has argument type Person
// ** Argument(c *Person) parameter(string) return(nil)
func (p person) GetName() string {
	return p.name
}

// ** Receiver method has argument type Person
// ** Argument(c *Person) parameter(string) return(nil)
func (p *person) SetName(name string) {
	p.name = name
}

func StructTutorial() {

	fmt.Println("\n[ Struct V.1 ]")
	x := person{
		name: "xPLUTo K",
		age:  20,
	}
	fmt.Printf("x = %T\n", x)
	fmt.Printf(" x.name = %v\n", x.name)
	x.SetName("xNawin K")
	fmt.Printf(" x.name = %v\n\n", x.GetName())

	// ** Pass by Value
	// y := x // ** y is Struct Person is not pointer struct
	// fmt.Printf("y = %T\n", y)
	// y.SetName("yPLUTO K")
	// fmt.Printf(" y.name = %v\n", y.name)
	// fmt.Printf(" x.name = %v\n", x.GetName())
	// fmt.Printf(" y.name = %v\n\n", y.GetName())

	// ** Pass by Ref
	// z := &y // ** Z is pointer of y(Struct Person).
	// fmt.Printf("z = %T(%p)\n", z, z)
	// z.name = "zNawin K"

	// println(x.name)
	// println(y.name)
	// println(z.name)

	fmt.Printf(" address x = %p\n", &x)

	//** Pass by Vaule
	a := x // TODO => a store value form x
	fmt.Printf(" address a(value) = %p\n", &a)

	// ** Pass by Ref
	b := &x // TODO => b is pointer store address x
	fmt.Printf(" address b(ref) = %p\n", b)
	fmt.Println("")

}
