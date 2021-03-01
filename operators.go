package main

import "fmt"

func main ()  {
	
	x,y := 5,6

	fmt.Println("x + y = ",x+y)
	fmt.Println("x - y = ",x-y)
	fmt.Println("x * y = ",x*y)
	fmt.Println("x % y = ",x%y)
	fmt.Println("x / y = ",x/y)

	love, hate := false,true

	fmt.Println("love && hate = ", love && hate)
	fmt.Println("love || hate = ", love || hate)
	fmt.Println("! love = ", !love)
}