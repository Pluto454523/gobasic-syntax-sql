package basicSyntax

import "fmt"

func ArrayLoopTutorial() {

	// x := [3]int{1, 2, 3} 	//** Array 1D
	// x := [...]int{4, 5, 6} 	//** Array 1D
	// x := []int{1, 2, 3} 		//** Array 1D แบบไม่มีขนาด
	// x := [][]int{ 			//** Array 2D แบบไม่มีขนาด
	// 	{1, 2, 3},
	// 	{4, 5, 6},
	// 	{7, 8},
	// }

	arrayX := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	arrayX = append(arrayX, 0) //** เพิ่มตำแหน่ง arrayX+1 และ assign ค่า 0 ไปที่ตำแหน่ง arrayX+1
	arrayX[8] = 9              //** Assign 9 ไปที่ arrayX ตำแหน่งที่ 8
	//arrayX[10] = 4 // ! ทำได้ แต่จะมีสามารถ assign ได้เนื่องจาก ยังไม่มี array ตำแหน่งที่ 10
	//len(arrayX) // ! ไม่ใช้นับจำนวน String เพราะภาษาอื่นมี byte ไม่เท่ากัน
	//tf8.RuneCountInString("ฟหกฟหก") // TODO แนะนำให้ใช้ utf8.RuneCountInString ในการนับ String
	fmt.Printf("arrayX = %#v \nlen arrayX = %#v \n", arrayX, len(arrayX))

	arrayY := arrayX[1:5] // ** Slices เริ่มตั่งแต่ 1 ถึง 4
	fmt.Print("arrayYํ = ")

	//** Like For loop
	for i := 0; i < len(arrayY); i++ {
		fmt.Printf("%v ", arrayY[i])
	}

	//** like while loop
	i := 0
	for i < len(arrayY) {
		fmt.Printf("%v ", arrayY[i])
		i++
	}

	for _, v := range arrayY { // ** For range
		fmt.Printf("%v ", v)
	}

	fmt.Println()

	var arrayZ = append(arrayX, []int{9, 10, 11}...)
	arrayZ = append(arrayX, 12, 13, 14, 15)
	fmt.Printf("arrayZ = %v\n", arrayZ)
}
