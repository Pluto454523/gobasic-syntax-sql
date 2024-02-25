package syntax

import (
	"fmt"
)

type Person struct {
	firstname string
	lastname  string
}

// @value receiver perso
// @return string
func (p *Person) GetFullName() string {
	return p.firstname + " " + p.lastname
}

// @pointer receiver *Person
// @argument string
func (p *Person) changeName(x string) {
	p.firstname = x
}

func MethodTutorial() {
	person := Person{
		firstname: "Nawin",
		lastname:  "K",
	}
	fmt.Println("Before 	-> ", person.GetFullName())
	person.changeName("Thanapat")
	fmt.Println("After 	-> ", person.GetFullName())

	sliceMethodTutorial()
}

type intSlice []int

func (is intSlice) sum() int {

	a := 0
	for _, v := range is {
		a += v
	}
	return a
}

func sliceMethodTutorial() {

	numbers := intSlice{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println("sum 	-> ", numbers.sum())

}
