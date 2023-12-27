package basicSyntax

import (
	"fmt"
)

func PointerTutorial() {

	// *******************
	// ** Pass by vaule **
	// *******************
	a := 3
	b := a
	fmt.Println("[ Pass by vaule ]")
	fmt.Println(" a =", a)
	fmt.Println(" b =", b)

	// ****************************
	// ** Pass by vaule Function **
	// ****************************
	c := 0
	sumByValue(c)                               // ? => ถ้า pass by reference ตัวแปร c จะเท่ากับ 100
	fmt.Println(" Pass by vaule Function =", c) // * c = 0

	// ***********************
	// ** Pass by reference **
	// ***********************
	x := 0
	y := &x // ? => y จะถูกกำหนดค่าเป็น address ของ x และ type y = *int(pointer of int)
	x = 50
	*y = 20 // ? => assign ค่าของ address ที่ y เก็บไว้เป็น 20
	*y = *y - 10
	// y = 10 // ! ทำไม่ได้ เนื่องจาก type ของ y = *int(pointer ของ int) แต่ 10 เป็น int ธรรมดา
	fmt.Println("\n[ Pass by reference ]")
	fmt.Printf(" %#v | x = %#v\n", &x, x) // TODO => &x = address of x and x = 10
	fmt.Printf(" %#v | y = %#v\n", y, *y) // TODO => y = address of x and *y = 10
	// fmt.Println("", &x, "| x =", x)
	// fmt.Println("", y, "| y =", *y)

	// ? => (type of test2 = ***int) pointer ของ pointer ของ int.
	// ? (***int) pointer ของ pointer ของ int.
	var test **int = &y // TODO => การประกาศตัวแบบเต็ม
	fmt.Printf(" test  = %#v\n", test)

	// ? => (type of test2 = ***int) pointer ของ pointer ของ pointer ของ int.
	// ? => เป็น pointer ที่ชี้ไปยัง pointer ที่ชี้ไปยัง pointer ที่ชี้ไปยัง int
	var test2 ***int = &test // TODO => การประกาศตัวแบบเต็ม
	fmt.Printf(" test2 = %#v\n\n", test2)

	// ********************************
	// ** Pass by reference Function **
	// ********************************
	z := 0
	sumByRef(&z)
	fmt.Println(" Pass by reference Function =", z) // * z = 10
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
