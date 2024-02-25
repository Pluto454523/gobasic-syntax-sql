package syntax

import "fmt"

// ** signature หรือ function type
// ** Parameter(int, int) return(int, int, int)
func sum(a, b int) (int, int, int) {
	return a, b, a + b
}

// ** Parameter(int, int) return(int, int)
func swap(x, y int) (a int, b int) {
	a = y
	b = x
	return a, b
}

func FunctionTutorial() {

	// ** => Multiple return
	// ** การรับค่า return จาก function สามารถรับได้หลายค่า return

	// ? sum ที่ return มา 3 ค่า
	a, b, c := sum(10, 5)
	fmt.Printf("a = %#v | b = %#v | c = %#v\n", a, b, c) // ** a=10 b=5 c=15

	// ? swap ที่ return มา 2 ค่า
	a, b = swap(a, b)
	fmt.Printf("a = %#v | b = %#v | c = %#v\n", a, b, c) // ** a=5 b=10 c=15

	// ** => Anonymous Function
	// ** คือ ฟังก์ชันที่ไม่ต้องมีชื่อ
	// ** มักจะใช้ในกรณีที่ต้องการสร้างฟังก์ชันในขณะเดียวกับการใช้งาน
	// ** หรือส่งไปเป็นพารามิเตอร์ให้กับ function ต่างๆก็ได้

	// ** => First-Class Function
	// ** คือการที่ฟังก์ชันที่สร้างขึ้นนั้นไม่มีข้อจำกัดในการนำไปใช้งานส่วนต่างๆของโปรแกรม
	// ** รวมไปถึงการนำไปกำหนดเป็นค่าให้กับตัวแปรก็ยังได้

	// ? add เป็น First-Class Function
	// ? add ไม่ได้เก็บ address ของ anonymous function
	// ? add เก็บ anonymous function ไว้ในตัวแปรเลย
	// ? (int, int) int
	add := func(a, b int) int {
		return a + b
	}
	fmt.Printf("add(5, 5) = %#v\n", add(5, 5))

	// ** Callback function
	// ** การที่ First-Class Function ถูกส่งไปเป็นพารามิเตอร์ให้กับฟังก์ชันอื่น
	// ** เพื่อเรียกใช้ในชังก์ชั่นนั้นๆ ในขณะที่มีการเรียกทำงานเราเรียกว่า Callback function

	// ** Higher-Order Function
	// ** การที่ “ฟังก์ชัน” สามารถรับ “อีกฟังชัน” มาเป็น Parameter/Argument และนำไปจัดการได้
	// ** หรือ/และ การที่รีเทิร์นค่ากลับไปเป็นฟังก์ชันก็ได้เช่นกัน
	// ** ซึ่งจะทำให้ลด side effect ต่างๆที่จะเกิดขึ้นจากตัวแปรทั่วๆไปได้เยอะมากๆ

	// ? (func(int, int) int) int
	compute := func(fn func(int, int) int) int {

		// ? เมื่อส่ง function ใดๆที่ function type เหมือนกันกับ parameter fn
		// ? compute จะเรียก fn และส่งค่า 10,6 ไปให้ fn
		// ? เรียก fn ว่า callback function
		fn(10, 6)

		// ? return callback function ไปที่ตัวแปร compute
		return fn(10, 6)

		// ? หากไม่ return callback function
		// ? compute ก็ยังคงเป็น Higher-order function ตามนิยาม
	}

	// ? เรียกใช้งาน compute และส่ง add ที่เป็น first-class function
	// ? เพื่อใช้เป็น callback function ใน compute function
	x := compute(add)
	fmt.Printf("compute(add function): %#v\n", x)

	// ? เรียกใช้งาน compute และส่ง anonymous function ไม่เป็น first-class function
	// ? เพื่อใช้เป็น callback function ใน compute function
	x = compute(func(a, b int) int {
		return a * b
	})
	fmt.Printf("compute(anonymous function): %#v\n", x)

	arrayParameter1()
	arrayParameter2()
}

// ** Array parameter แบบที่ 1 กำหนดจำนวนตัวแปรไว้แล้ว.
// ** เมื่อเรียกใช้ฟังก์ชันที่รับพารามิเตอร์เป็น []int, จำเป็นต้องส่ง slice ของ int เท่านั้น.
func arrayParameter1() {
	// ? ([]int) int
	sumArray := func(a []int) int {
		x := 0
		for _, v := range a {
			x += v
		}
		return x
	}
	a := sumArray([]int{1, 2, 3, 4, 5, 6, 7, 8, 9})
	fmt.Printf("sumArray = %#v\n", a)
}

// ** Array parameter แบบที่ 2
// ** จำนวนตัวแปรที่ไม่แน่นอน (variadic parameter)
func arrayParameter2() {
	// ? (...int) int
	sumArray := func(a ...int) int {
		x := 0
		for _, v := range a {
			x += v
		}
		return x
	}
	a := sumArray(1, 2, 3, 4, 5, 6, 7, 8, 9)
	fmt.Printf("sumArray2 = %#v\n", a)
}
