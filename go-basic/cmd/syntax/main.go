package main

import (
	"fmt"

	// ? "bs" is alias package and default is package name.
	// ? "format" is name of go module.
	// ? "github.com/pluto454523/basic-syntax/internal/syntax" is path folder in module.

	bs "github.com/pluto454523/basic-syntax/internal/syntax"
)

// ! package variable ประกาศแบบ short declared ไม่ได้
var message string = "The Fundamentals Golang"

func main() {

	fmt.Printf(">>-<< %v >>-<<\n", message)

	// bs.VariableTutorial()
	// bs.ControlStructureTutorial()
	// bs.FunctionTutorial()
	// bs.ArraySliceTutorial()
	// bs.MapTutorial()
	// bs.PointerTutorial()
	// bs.StructTutorial()
	// bs.MethodTutorial()
	// bs.InterfaceTutorial()
	bs.InterfaceTutorial2()

}
