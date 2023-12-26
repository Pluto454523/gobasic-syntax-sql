package main

import (
	"fmt"
	syntax "gobasic/basic-syntax/src"
)

var testGlobal string = "Golang basic Tutorial" // ! package variable ประกาศแบบ short declared ไม่ได้

func main() {

	fmt.Printf("========== Welcome to %v ==========\n", testGlobal) //**แสดง package variable

	// syntax.FormatPrint()
	// syntax.Function()
	// syntax.ArrayLoopTutorial()
	syntax.MapTutorial()
	// syntax.PointerTutorial()

	// ** Struct ธรรมดา
	// syntax.StructTutorial()

	// ** Struct แบบ OOP
	// x := syntax.Person{}
	// x.SetName("PLUTo")
	// fmt.Printf("%#v", x.GetName())

}
