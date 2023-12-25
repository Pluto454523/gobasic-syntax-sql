package basicSyntax

func Function() {

	// a, _, b := sum(10, 5) //** การรับค่า return จาก function สามารถรับได้หลาย return
	// fmt.Printf("a = %#v \nb = %#v", a, b)

	// x := func(a, b int) int { //** Anonymous function
	// 	return a + b
	// }
	// sum := x(8, 6)
	// println(sum)

	// cal(func(a, b int) int { //** Lamda function
	// 	return a + b
	// })

	//** Array parameter แบบที่ 1
	ArrayP1 := sumArray([]int{1, 2, 3, 4, 5, 6, 7, 8, 9})
	println(ArrayP1)

	//** Array parameter แบบที่ 2
	ArrayP2 := sumArray2(1, 2, 3, 4, 5, 6, 7, 8, 9)
	println(ArrayP2)
}

// ** เมื่อส่ง function ใดๆที่ signature เหมือนกันกับ parameter
// ** cal จะเรียก function นั้นแล้ว ส่งค่า 10,6 ไป
// ** และ return ค่าไปที่ตัวแปร sum
func cal(f func(int, int) int) {
	sum := f(10, 6)
	println(sum)
}

// ** signature หรือ function type = ** (int,int) int
func add(a, b int) int {
	return a + b
}

// ** (int,int) int
func sub(a, b int) int {
	return a - b
}

// ** ([]int) int
func sumArray(a []int) int { //** Paramter Array แบบที่ 1

	x := 0
	for _, v := range a {
		x += v
	}
	return x
}

// ** ([]int) int
func sumArray2(a ...int) int { //** Paramter Array แบบที่ 2

	x := 0
	for _, v := range a {
		x += v
	}
	return x
}
