package functional

import "fmt"

// ***************************************************
// ** Higher-order function เป็นฟังก์ชันที่ parameter เป็นฟังก์ชัน
// ** และ/หรือ return ฟังก์ชันออกไป
// ***************************************************

func HigherOrderFunctions() {
	result := applyOperation(5, 3, sub)
	fmt.Println("Result of sub:", result)

	result = applyOperation(5, 3, multiply)
	fmt.Println("Result of multiplication:", result)
}

func applyOperation(a, b int, operation func(int, int) int) int {
	return operation(a, b)
}

func sub(a, b int) int {
	return a - b
}

func multiply(a, b int) int {
	return a * b
}
