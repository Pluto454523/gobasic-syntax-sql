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

	// ***************************************************
	// ** Multiple return
	// ** การรับค่า return จาก function สามารถรับได้หลายค่า return
	// ***************************************************

	// TODO: Example : การรับค่าจาก sum ที่ return มา 3 ค่า
	a, b, c := sum(10, 5)
	fmt.Printf("a = %#v | b = %#v | c = %#v\n", a, b, c) // ** a=10 b=5 c=15

	// TODO: Example : swap ที่ return มา 2 ค่า
	a, b = swap(a, b)
	fmt.Printf("a = %#v | b = %#v | c = %#v\n", a, b, c) // ** a=5 b=10 c=15

	// ***************************************************
	// ** Anonymous function คือ ฟังก์ชันที่ไม่ต้องมีชื่อ
	// ** มักจะใช้ในกรณีที่ต้องการสร้างฟังก์ชันในขณะเดียวกับการใช้งาน
	// ** หรือส่งไปเป็นพารามิเตอร์ให้กับ function ต่างๆก็ได้
	// ***************************************************

	// TODO: Example : add เก็บ anonymous function ไว้ในตัวแปรไม่ได้เก็บ address
	// ** (int, int) int
	add := func(a, b int) int {
		return a + b
	}
	fmt.Printf("add(5, 5) = %#v\n", add(5, 5))

	// TODO: Example : compute เป็น Higher-order function
	// ** (func(int, int) int) int
	compute := func(fn func(int, int) int) int {
		// ? เมื่อส่ง function ใดๆที่ signature เหมือนกันกับ parameter
		// ? cal จะเรียก parameter fn และส่งค่า 10,6 ไปให้ fn
		// ? และ return ค่าไปที่ตัวแปร sum
		sum := fn(10, 6) // TODO: เรียกใช้ fn function ที่รับค่า 10 และ 6 เป็นพารามิเตอร์
		fmt.Printf("Result sum: %#v\n", sum)

		// ? return ค่าไปที่ฟังก์ชัน
		// ! หากไม่ return fn(10, 6) ยังคงเป็น Higher-order function ตามนิยาม
		return fn(10, 6)
	}

	// ***************************************************
	// ** Callback function เป็นฟังก์ชันที่ถูกส่งไปเป็นพารามิเตอร์ให้กับฟังก์ชันอื่น
	// ** เพื่อให้ฟังก์ชันนั้นสามารถเรียกใช้ callback function นั้นในขณะที่ทำงาน
	// ***************************************************

	// ? เรียกใช้งาน compute และส่ง add เป็น callback function
	compute(add)

	// ? เรียกใช้งาน compute และส่ง anonymous callback function
	compute(func(a, b int) int {
		return a * b
	})

	// ***************************************************
	// ** Array parameter แบบที่ 1 กำหนดจำนวนตัวแปรไว้แล้ว.
	// ** เมื่อเรียกใช้ฟังก์ชันที่รับพารามิเตอร์เป็น []int, จำเป็นต้องส่ง slice ของ int เท่านั้น.
	// ***************************************************
	// ** ([]int) int
	sumArray := func(a []int) int {
		x := 0
		for _, v := range a {
			x += v
		}
		return x
	}
	arrayP1 := sumArray([]int{1, 2, 3, 4, 5, 6, 7, 8, 9})
	// ? 	([]int) int
	fmt.Printf("sumArray = %#v\n", arrayP1)

	// ***************************************************
	// ** Array parameter แบบที่ 2
	// ** จำนวนตัวแปรที่ไม่แน่นอน (variadic parameter)
	// ***************************************************
	// ** (...int) int
	sumArray2 := func(a ...int) int {
		x := 0
		for _, v := range a {
			x += v
		}
		return x
	}
	arrayP2 := sumArray2(1, 2, 3, 4, 5, 6, 7, 8, 9)
	fmt.Printf("sumArray2 = %#v\n", arrayP2)
}
