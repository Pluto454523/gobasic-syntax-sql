package main

import (
	"fmt"

	// ? "basicSyntax" is alias package and default is package name.
	// ? "format" is name of go module.
	// ? "github.com/pluto454523/basic-syntax/internal" is path folder in module.

	basicSyntax "github.com/pluto454523/basic-syntax/internal/syntax"
)

// ! package variable ประกาศแบบ short declared ไม่ได้
var testGlobal string = "Golang basic Tutorial"

func main() {

	fmt.Printf("========== Welcome to %v ==========\n", testGlobal) //**แสดง package variable

	// basicSyntax.FormatTutorial()
	// basicSyntax.FunctionTutorial()
	// basicSyntax.ArrayLoopTutorial()
	// basicSyntax.MapTutorial()
	// basicSyntax.PointerTutorial()

	// ** Struct ธรรมดา
	// basicSyntax.StructTutorial()
	// ** Struct แบบ OOP
	// constructorTutorial()

	// ** Interfact
	// interfaceStructTutorial()
	// basicSyntax.InterfactCast()

}

func constructorTutorial() {
	x := basicSyntax.NewPerson("PLUTO K", 21) // Constructor Function
	fmt.Printf("\n%#v\n", x)

	x.SetName("Nawin K") // TODO => set ผ่าน reciver method setName
	fmt.Printf("Name: %#v\n", x.GetName())
}

func interfaceStructTutorial() {
	info := &basicSyntax.Account{Name: "PLUTO", Age: 21}
	x := basicSyntax.NewProfile(info)
	pf := x.GetProfile()
	pf.Age = 50
	fmt.Printf("%#v", pf)
}
