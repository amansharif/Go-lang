package main

import "fmt"

func main ()  {
	
	var Num[5] int // Declaration of an array

	// Value assign or Initialization or Definition whatever you like to call it.
	Num[0] = 11
	Num[1] = 22
	Num[2] = 33
	Num[3] = 44
	Num[4] = 55

	fmt.Println(Num[3])
	
	fmt.Println()

	EvenNum := [5]int{12,14,18,20,24} //Declaration plus definition. Second way.

	fmt.Println(EvenNum[3])
	
	fmt.Println()
	fmt.Println()
	
	// Array traversal
	for _, val := range Num {
		fmt.Println(val)
	}

	fmt.Println()
	fmt.Println()
	
	// Array traversal with the index
	for i, value := range EvenNum {
		fmt.Println(value,i)
	}

	fmt.Println()
	fmt.Println()

	// Array slicing
	numSlice := []int{1,2,3,4,5,6,7}

	sliced := numSlice[3:5]
	fmt.Println(sliced)
}