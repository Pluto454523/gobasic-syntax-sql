package syntax

import "fmt"

func ControlStructureTutorial() {

	var score int = 62
	var grade string

	// ** switch
	switch {
	case score >= 80:
		grade = "A"
	case score >= 70:
		grade = "B"
	case score >= 60:
		grade = "C"
	default:
		grade = "F"
	}
	fmt.Printf("SWITCH: Your grade is %s\n", grade)

	// ** if-else
	if score >= 80 {
		grade = "A"
	} else if score >= 70 {
		grade = "B"
	} else if score >= 60 {
		grade = "C"
	} else {
		grade = "F"
	}
	fmt.Printf("IF-ELSE: Your grade is %s\n", grade)

	// ** Pre process if else
	num1, num2 := 10, 20
	if sumNum := num1 + num2; sumNum >= 30 {
		fmt.Printf("Pre process: sumNum: %#v\n", sumNum)
	}

	// ** For Loop
	for a := 1; a < 10; a++ {
		fmt.Printf("For Loop: %d\n", a)
	}

	// ** Do While Loop
	b := 1
	for {
		fmt.Printf("Do While Loop: %d\n", b)
		b++
		if b >= 10 {
			break
		}
	}

	// ** While Loop
	c := 1
	for c < 10 {
		fmt.Printf("While Loop: %d\n", c)
		c++
	}
}
