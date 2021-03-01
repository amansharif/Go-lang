package main

import "fmt"

func main ()  {
	
	// for style loop
	for i:=1; i<10; i++ {
		fmt.Println("Iteration No : ",i)
	}

	fmt.Println()

	// while style loop
	x := 1
	for x<10 {
		fmt.Println("Iteration No : ",x)
		x++
	}

	// Nested loop

    for i := 1; i < 10; i++ {
	    for j := 1; j < i; j++ {
		    fmt.Printf("*")
	    }
		fmt.Println()
    }
}