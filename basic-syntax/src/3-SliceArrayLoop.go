package basicSyntax

import "fmt"

func showAll(ns []int) {
	fmt.Printf("show all : %#v\n", ns)
}

func ArrayLoopTutorial() {

	// x := [3]int{1, 2, 3} 	//** Array 1D
	// x := [...]int{4, 5, 6} 	//** Array 1D
	// x := []int{1, 2, 3} 		//** Array 1D แบบไม่มีขนาด
	// x := [][]int{ 			//** Array 2D แบบไม่มีขนาด
	// 	{1, 2, 3},
	// 	{4, 5, 6},
	// 	{7, 8},
	// }

	// ** Slice
	arrayA := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	showAll(arrayA)
	//** เพิ่มตำแหน่ง arrayA+1 และ assign ค่า 0 ไปที่ตำแหน่ง arrayA+1
	arrayA = append(arrayA, 0)
	//** Assign 9 ไปที่ arrayA ตำแหน่งที่ 8
	arrayA[8] = 9

	// ! ทำได้ แต่จะมีสามารถ assign ได้เนื่องจาก ยังไม่มี array ตำแหน่งที่ 10
	//arrayA[10] = 4
	// ! ไม่ใช้นับจำนวน String เพราะภาษาอื่นมี byte ไม่เท่ากัน
	//len(arrayA)
	// TODO แนะนำให้ใช้ utf8.RuneCountInString ในการนับ String
	//tf8.RuneCountInString("ฟหกฟหก")
	fmt.Printf("arrayA = %#v \nlen arrayA = %#v \n", arrayA, len(arrayA))

	// ** Append Slices
	var arrayB = append(arrayA, []int{9, 10, 11}...)
	arrayB = append(arrayA, 12, 13, 14, 15)
	fmt.Printf("arrayB = %v\n", arrayB)

	// ** Slices เริ่มตั่งแต่ 1 ถึง 4
	arrayC := arrayA[1:5]
	fmt.Print("arrayCํ = ")

	//** Like For loop
	for i := 0; i < len(arrayC); i++ {
		fmt.Printf("%v ", arrayC[i])
	}

	//** like while loop
	i := 0
	for i < len(arrayC) {
		fmt.Printf("%v ", arrayC[i])
		i++
	}

	// ** For range
	for _, v := range arrayC {
		fmt.Printf("%v ", v)
	}

	fmt.Println()
}
