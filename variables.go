package main

import "fmt"

func main ()  {
	
	var a int = 5
	var b float32 = 4.32
	const pi float64 = 3.1415139475

	var (
		c = 5
		d = "I am a String literal."
	)

	x,y := 12,13

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(pi)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(x)
	fmt.Println(y)
}