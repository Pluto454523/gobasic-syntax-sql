package syntax

import (
	"fmt"
	"os"
	"strings"
	"unicode/utf8"
)

func VariableTutorial() {

	//** ใน go ไม่มี nil ของแต่ละ type โดยค่าเริ่มต้นของแต่ละ type มีแต่ zero value
	var numVar int // ? ไม่ได้กำหนดค่าให้, จะมีค่า zero value ของ int คือ 0
	numFloat := 15.589
	numVar = 10
	fmt.Printf("numFloat = %.0f\n", numFloat) //** formath percision
	fmt.Printf("numVar = %v\n", numVar)

	//** CONSTAIN : ค่าคงที่
	const NumberConst int = 50
	fmt.Printf("NumberConst = %v\n", NumberConst)

	//** การกำหนดค่า CONSTAIN สำหรับวันในสัปดาห์โดยใช้ iota
	const (
		Sunday    = iota // 0
		Monday           // 1 (iota ถูกใช้โดยอัตโนมัติ จะเพิ่มขึ้นทีละ 1)
		Tuesday          // 2
		Wednesday        // 3
		Thursday         // 4
		Friday           // 5
		Saturday         // 6
	)

	fmt.Println("Sunday :", Sunday)
	fmt.Println("Monday :", Monday)
	fmt.Println("Tuesday :", Tuesday)
	fmt.Println("Wednesday :", Wednesday)
	fmt.Println("Thursday :", Thursday)
	fmt.Println("Friday :", Friday)
	fmt.Println("Saturday :", Saturday)

	//** การนับ string ภาษาไทยและอื่นๆ

	//** นับได้ แต่ๆ ใช้กับภาษาไทยไม่เวิร์ค
	str := "สวัสดีวันจันทร์"
	fmt.Println(len(str))

	//** ใช้วิธีนี้ในการนับ String ภาษาไทย
	fmt.Printf(str+" Count is %v\n\n", utf8.RuneCountInString(str))

	//** การแยก String เป็น Array
	arrStr := strings.Split(str, "")
	fmt.Println(len(arrStr))

	// ** รูปแบบต่างๆของการ format
	p := [2]int{1, 2}
	fmt.Printf("struct1: %v\n", p)

	fmt.Printf("struct2: %+v\n", p)

	fmt.Printf("struct3: %#v\n", p)

	fmt.Printf("type: %T\n", p)

	fmt.Printf("bool: %t\n", true)

	fmt.Printf("int: %d\n", 123)

	fmt.Printf("bin: %b\n", 14)

	fmt.Printf("char: %c\n", 33)

	fmt.Printf("hex: %x\n", 456)

	fmt.Printf("float1: %f\n", 78.9)

	fmt.Printf("float2: %e\n", 123400000.0)
	fmt.Printf("float3: %E\n", 123400000.0)

	fmt.Printf("str1: %s\n", "\"string\"")

	fmt.Printf("str2: %q\n", "\"string\"")

	fmt.Printf("str3: %x\n", "hex this")

	fmt.Printf("pointer: %p\n", &p)

	fmt.Printf("width1: |%6d|%6d|\n", 12, 345)

	fmt.Printf("width2: |%6.2f|%6.2f|\n", 1.2, 3.45)

	fmt.Printf("width3: |%-6.2f|%-6.2f|\n", 1.2, 3.45)

	fmt.Printf("width4: |%6s|%6s|\n", "foo", "b")

	fmt.Printf("width5: |%-6s|%-6s|\n", "foo", "b")

	s := fmt.Sprintf("sprintf: a %s", "string")
	fmt.Println(s)
	fmt.Fprintf(os.Stderr, "io: an %s\n", "error")
}
