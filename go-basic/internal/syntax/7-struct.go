package syntax

import (
	"fmt"
)

type Student struct {
	firstame string
	lastname string
	course   Course
}

type Course struct {
	subject string
	score   int
}

func StructTutorial() {

	// ** Create an instance of the Student struct.
	var student Student
	student.firstame = "Firstfame"
	student.lastname = "Lastname"

	student = Student{
		firstame: "Nawin",
		lastname: "K",
		course: Course{
			subject: "Math",
			score:   80,
		},
	}

	fmt.Printf("student\n")
	fmt.Printf(" type  -> %T\n value -> %v\n\n", student, student)

	// ** Struct passed by value by default.
	// ? student2 is "Student Struct" is not "pointer struct".
	student2 := student
	student2.firstame = "Thanapat"
	fmt.Printf("student2\n")
	fmt.Printf(" type  -> %T\n value -> %v\n\n", student2, student2)

	// ** Struct passed by Reference.
	// ? student3 is pointer of student(Struct Student).
	student3 := &student
	student3.firstame = "Nawin"
	fmt.Printf("student2 after change student3\n")
	fmt.Printf(" type  -> %T\n value -> %v\n\n", student2, student2)

	structArrayTutorial()
	structMapTutorial()
}

func structArrayTutorial() {
	var students [3]Student

	students[0] = Student{firstame: "Nawin", lastname: "K"}
	students[1] = Student{firstame: "Thanapat", lastname: "R"}
	students[2] = Student{firstame: "Pannawat", lastname: "I"}

	fmt.Println("Students Array:", students)

}

func structMapTutorial() {

	var students = make(map[string]Student)

	students["st01"] = Student{firstame: "Nawin", lastname: "K"}
	students["st02"] = Student{firstame: "Thanapat", lastname: "R"}
	students["st03"] = Student{firstame: "Pannawat", lastname: "I"}

	fmt.Println("Students Map:", students)
}
