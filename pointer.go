package main

import "fmt"

func main ()  {
	
	x,y := 10,10

	fmt.Println("Value of x before any function call is ",x)
	fmt.Println("Value of y before any function call is ",y)

	// Call by value example
	changeValueByValue(x)
	changeValueByValue(y)
	fmt.Println("Value of x after the call by value function call is ",x)
	fmt.Println("Value of y after the call by value function call is ",y)

	// Call by reference example
	changeValueByReference(&x)
	changeValueByReference(&y)
	fmt.Println("Value of x after the call by reference function call is ",x)
	fmt.Println("Value of y after the call by reference function call is ",y)
}

func changeValueByValue(a int)  {
	a = 5
}

func  changeValueByReference(b *int)  {
	*b = 5
}