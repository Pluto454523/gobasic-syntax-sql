package basicSyntax

import "fmt"

type Person struct {
	name string
	age  int
}

func (p Person) GetName() string { //** ใช้หลักการ Reciver ในการทำ Method แบบ OOP
	return p.name
}

func (p *Person) SetName(name string) { //** ใช้หลักการ Reciver ในการทำ Method แบบ OOP
	p.name = name
}

func StructTutorial() {

	x := Person{
		name: "PLUTO",
		age:  20,
	}

	// y := x //** Pass by Vaule
	// println(&y)
	// println(&x)

	// z := &x //** Pass by Ref
	// println(z)
	// println(&x)

	fmt.Printf("%#v ", x.name)

}
