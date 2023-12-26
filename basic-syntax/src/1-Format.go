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

	//** CONSTAIN : ค่าคงที่
	const NumberConst int = 50
	fmt.Printf("NumberConst = %v\n", NumberConst)
	const (
		Sunday = iota
		Monday
		Tuesday
		Wednesday
		Thursday
		Friday
		Saturday
	)

	fmt.Println("Sunday :", Sunday)
	fmt.Println("Monday :", Monday)
	fmt.Println("Tuesday :", Tuesday)
	fmt.Println("Wednesday :", Wednesday)
	fmt.Println("Thursday :", Thursday)
	fmt.Println("Friday :", Friday)
	fmt.Println("Saturday :", Saturday)

	// TODO: How to count lenght "UTF-8" string
	testString := "ฟหก"
	// **: นับได้ แต่ๆ ใช้กับภาษาไทยไม่เวิร์ค
	fmt.Println(len(testString))
	// **: ใช้วิธีนี้ในการนับ String ภาษาไทย
	fmt.Printf(testString+" Count is %v\n\n", utf8.RuneCountInString(testString))

	// ** รูปแบบต่างๆของการ format
	// p := [2]int{1, 2}
	// fmt.Printf("struct1: %v\n", p)

	// fmt.Printf("struct2: %+v\n", p)

	// fmt.Printf("struct3: %#v\n", p)

	// fmt.Printf("type: %T\n", p)

	// fmt.Printf("bool: %t\n", true)

	// fmt.Printf("int: %d\n", 123)

	// fmt.Printf("bin: %b\n", 14)

	// fmt.Printf("char: %c\n", 33)

	// fmt.Printf("hex: %x\n", 456)

	// fmt.Printf("float1: %f\n", 78.9)

	// fmt.Printf("float2: %e\n", 123400000.0)
	// fmt.Printf("float3: %E\n", 123400000.0)

	// fmt.Printf("str1: %s\n", "\"string\"")

	// fmt.Printf("str2: %q\n", "\"string\"")

	// fmt.Printf("str3: %x\n", "hex this")

	// fmt.Printf("pointer: %p\n", &p)

	// fmt.Printf("width1: |%6d|%6d|\n", 12, 345)

	// fmt.Printf("width2: |%6.2f|%6.2f|\n", 1.2, 3.45)

	// fmt.Printf("width3: |%-6.2f|%-6.2f|\n", 1.2, 3.45)

	// fmt.Printf("width4: |%6s|%6s|\n", "foo", "b")

	// fmt.Printf("width5: |%-6s|%-6s|\n", "foo", "b")

	// s := fmt.Sprintf("sprintf: a %s", "string")
	// fmt.Println(s)
	// fmt.Fprintf(os.Stderr, "io: an %s\n", "error")
}
