package basicSyntax

func Function() {

	// ? การรับค่า return จาก function สามารถรับได้หลาย return
	// a, b, c := sum(10, 5)
	// fmt.Printf("a = %#v \nb = %#v", a, b)

	// ? Anonymous function
	// x := func(a, b int) int {
	// 	return a + b
	// }
	// println(x(4, 6))
	// sum := x(8, 6)
	// println(sum)

	// ? Lamda function Ex.1
	// cal(func(a, b int) int {
	// 	return a + b
	// })

	// ? Lamda function Ex.2
	// cal(add)

	// ? Lamda function EX.3
	// cal(sub)

	// ? Array parameter แบบที่ 1
	ArrayP1 := sumArray([]int{1, 2, 3, 4, 5, 6, 7, 8, 9})
	println(ArrayP1)

	// ? Array parameter แบบที่ 2
	ArrayP2 := sumArray2(1, 2, 3, 4, 5, 6, 7, 8, 9)
	println(ArrayP2)

	swap("a", "b")
}

// ** signature หรือ function type
// ** Parameter(string, string) return(string, string)
func swap(x, y string) (a string, b string) {
	a = y
	b = x
	return a, b
}

// ** (float64, float64) float64
func compute(fn func(float64, float64) float64) float64 {
	// ** Function Callback ส่ง Function กลับคืนไป
	return fn(3, 4)
}

// ** (func(int, int) int) (ไม่มี reuturn)
func cal(f func(int, int) int) {
	// ** เมื่อส่ง function ใดๆที่ signature เหมือนกันกับ parameter
	// ** cal จะเรียก function นั้นแล้ว ส่งค่า 10,6 ไป และ return ค่าไปที่ตัวแปร sum
	sum := f(10, 6)
	println(sum)
}

// ** (int,int) int
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
