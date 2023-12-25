package basicSyntax

import (
	"fmt"
	"unicode/utf8"
)

func FormatPrint() {

	var numVar int = 10
	numFloat := 15.589
	fmt.Printf("numFloat = %.0f\n", numFloat) //** formath percision
	fmt.Printf("numVar = %v\n", numVar)

	const NumberConst int = 50 //**ค่าคงที่
	fmt.Printf("NumberConst = %v\n", NumberConst)

	testString := "ฟหก"          // **: ใช้วิธีนี้ในการนับ String ภาษาไทย
	fmt.Println(len(testString)) // **: นับได้ แต่ๆ ใช้กับภาษาไทยไม่เวิร์ค
	fmt.Printf(testString+" Count is %v\n\n", utf8.RuneCountInString(testString))

}
