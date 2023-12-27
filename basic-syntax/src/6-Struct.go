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

type course struct {
	name, instructor string
	price            int
}

// ** Reciver(c *course) (nil) (nil)
func (c *course) discount() int {
	c.price = c.price - 599
	fmt.Println(" discount:", c.price)
	return c.price
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

	// ********************
	// ** Pointer Struct **
	// ********************
	fmt.Println("\n[ Pointer Struct ]")
	d := &course{" Basic Go", "Nawin K", 9999}

	e := d.discount() // เข้าถึงค่าในตัวแปรเหมือนเดิมไม่ต้อง de-reference อีกที
	fmt.Println(" discount price:", e)
	fmt.Println(" price:", d.price)

}
