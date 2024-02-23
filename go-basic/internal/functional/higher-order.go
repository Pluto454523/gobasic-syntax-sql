package functional

import "fmt"

// ** Higher-Order Function
// ** จะเป็นการที่ “ฟังก์ชัน” สามารถรับ “อีกฟังชัน” มาเป็น parameter/Argument และนำไปจัดการได้
// ** รวมถึงการที่รีเทิร์นค่ากลับไปเป็นฟังก์ชันก็ได้เช่นกัน
// ** ซึ่งจะทำให้ลด side effect ต่างๆที่จะเกิดขึ้นจากตัวแปรทั่วๆไปได้เยอะมากๆ

func HigherOrderFunctions() {
	result := applyOperation(5, 3, sub)
	fmt.Println("Result of sub:", result)

	result = applyOperation(5, 3, multiply)
	fmt.Println("Result of multiplication:", result)
}


func applyOperation(a, b int, operation func(int, int) int) int {
	// operation 
	// operation is callback function 
	return operation(a, b)
}

func sub(a, b int) int {
	return a - b
}

func multiply(a, b int) int {
	return a * b
}
