package syntax

import (
	"fmt"
)

func ArraySliceTutorial() {

	//** Static  Array 1D calling array
	arrayA := [3]int{1, 2, 3}
	_ = arrayA

	//** Dynamic Array 1D calling slice
	arrayB := [...]int{4, 5, 6}
	_ = arrayB

	//** Dynamic Array 1D calling slice
	arrayC := []int{1, 2, 3}
	_ = arrayC

	//** Dynamic Array 2D calling slice
	arrayD := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8},
	}
	_ = arrayD

	fmt.Printf("arrayA = %#v \n", arrayA)

	// ** Slice

	// ? => Convert array to slice
	sliceA := arrayA[:]

	// ? => เป็นการบวก 4,5,6,7,8,9,10 ต่อจาก sliceA แล้วเอาไปแทนที่ sliceA
	sliceA = append(sliceA, 4, 5, 6, 7, 8, 9, 10)

	// ? => Assign 9 ไปที่ sliceA ตำแหน่งที่ 8
	sliceA[8] = 9

	// ! => ทำไม่ได้เนื่องจาก ยังไม่มี array ตำแหน่งที่ 10
	// sliceA[10] = 4

	fmt.Printf("sliceA = %#v \n", sliceA)
	fmt.Printf("length sliceA = %#v \n", len(sliceA))
	fmt.Printf("capacity sliceA = %#v \n", cap(sliceA))

	// ** Append slice to slice
	var sliceB = append(sliceA, []int{11, 12, 13}...)
	sliceB = append(sliceA, 14, 15)
	fmt.Printf("sliceB = %#v \n", sliceB)
	fmt.Printf("length sliceB = %#v \n", len(sliceB))
	fmt.Printf("capacity sliceB = %#v \n", cap(sliceB))

	// ** Slice เริ่มตั่งแต่ 1 ถึง 4
	sliceC := sliceA[1:5]

	//** Like for loop
	fmt.Print("\nsliceC for loop = ")
	for i := 0; i < len(sliceC); i++ {
		fmt.Printf("%v ", sliceC[i])
	}

	//** like while loop
	fmt.Print("\nsliceC while loop = ")
	i := 0
	for i < len(sliceC) {
		fmt.Printf("%v ", sliceC[i])
		i++
	}

	// ** for range
	fmt.Print("\nsliceC for range = ")
	for _, v := range sliceC {
		fmt.Printf("%v ", v)
	}

	fmt.Println()
}
