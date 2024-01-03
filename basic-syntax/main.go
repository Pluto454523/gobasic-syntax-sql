package main

import (
	"fmt"
	syntax "gobasic/basic-syntax/src"
	// syntax "gobasic/basic-syntax/src"
	// ? "syntax" is alias package and default is package name.
	// ? "gobasic" is name of go module.
	// ? "basic-syntax/src" is path folder in module.
)

var testGlobal string = "Golang basic Tutorial" // ! package variable ประกาศแบบ short declared ไม่ได้

func main() {

	fmt.Printf("========== Welcome to %v ==========\n", testGlobal) //**แสดง package variable

	// syntax.FormatPrint()
	syntax.Function()
	// syntax.ArrayLoopTutorial()
	// syntax.MapTutorial()
	// syntax.PointerTutorial()

	// ** Struct ธรรมดา
	// syntax.StructTutorial()
	// ** Struct แบบ OOP
	// constructorTutorial()

	// ** Interfact
	// interfaceStructTutorial()
	// syntax.InterfactCast()

}

func constructorTutorial() {
	x := syntax.NewPerson("PLUTO K", 21) // Constructor Function
	fmt.Printf("\n%#v\n", x)

	x.SetName("Nawin K") // TODO => set ผ่าน reciver method setName
	fmt.Printf("Name: %#v\n", x.GetName())
}

func interfaceStructTutorial() {
	info := &syntax.Account{Name: "PLUTO", Age: 21}
	x := syntax.NewProfile(info)
	pf := x.GetProfile()
	pf.Age = 50
	fmt.Printf("%#v", pf)
}
