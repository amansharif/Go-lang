package main

import "fmt"

func main ()  {

	x := 5
	fmt.Println("Factorial is ",factorial(x))
	fmt.Println()
	fmt.Println(addMeUp(10,20,30,40,50,60,70))

}

func factorial(num int) int{
	
	if num == 0 {
		return 1
	}

	return num * factorial(num - 1)
}

// Function taking multiple arguments
func addMeUp(args ...int) int {
	sum := 0
	for _,value := range args {
		sum += value
	}
	return sum
}