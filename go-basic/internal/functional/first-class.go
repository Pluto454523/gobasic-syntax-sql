package functional

import "fmt"

// ** First-Class Function
// ** คือการที่ฟังก์ชันที่สร้างขึ้นนั้นไม่มีข้อจำกัดในการนำไปใช้งานส่วนต่างๆของโปรแกรม
// ** รวมไปถึงการนำไปกำหนดเป็นค่าให้กับตัวแปรก็ยังได้ ส่วน

func add(a, b int) int { // add is First-Class Function
	return a + b
}

func subtract(a, b int) int { // subtract is First-Class Function
	return a - b
}

func FirstClassFunction() {
	var operation func(int, int) int // operation is anonymous function
	operation = add
	result := operation(5, 3)
	fmt.Println("Result of addition:", result)

	operation = subtract
	result = operation(5, 3)
	fmt.Println("Result of subtraction:", result)
}
