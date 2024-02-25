package syntax

import (
	"fmt"
	"reflect"
)

type shape interface {
	area() float64
}

// ** circle struct
type circle struct {
	radius float64
}

// ** circle's implementation of the shape interface
func (c circle) area() float64 {
	return 3.14 * c.radius * c.radius
}

// ** rectangle struct
type rectangle struct {
	width, height float64
}

// ** rectangle's implementation of the shape interface
func (r rectangle) area() float64 {
	return r.width * r.height
}

// ** function that accepts shape interface
func getArea(s shape) {

	//? ใช้ type assertion เพื่อแปลง shape เป็นชนิด rectangle หรือ circle
	//? -> shape, ok := s.(rectangle);
	//? Ref: https://go-tour-th.appspot.com/tour/methods/15

	//* ใช้ type assertion เพื่อ cast interface Shape
	//* เป็นรูปทรงต่างๆตามประเภทของ shape
	shapeName := reflect.TypeOf(s).Name()
	switch shape := s.(type) {
	case circle:
		fmt.Printf("%v(%v): %.2f\n", shapeName, shape.radius, s.area())
	case rectangle:
		fmt.Printf("%v(%v*%v): %.2f\n", shapeName, shape.height, shape.width, s.area())
	}
}

// ** function that accepts shape interface
func getArea2(s shape) {

	shapeName := reflect.TypeOf(s).Name()
	if shape, ok := s.(rectangle); ok {

		fmt.Printf("%v(%v*%v): %.2f\n", shapeName, shape.height, shape.width, shape.area())

	} else if shape, ok := s.(circle); ok {

		fmt.Printf("%v(%v): %.2f\n", shapeName, shape.radius, shape.area())

	}
}

func InterfaceTutorial() {

	//* สร้าง slice ของ interface Shape เพื่อเก็บวงกลมและสี่เหลี่ยมผืนผ้า
	shapes := []shape{
		circle{radius: 5},
		rectangle{width: 3, height: 4},
	}

	getArea(shapes[0])
	getArea2(shapes[1])

}
