package main

import "fmt"

func main ()  {
	
	rect := rectangle{10,5} // Structure value assignment
	fmt.Println("Height of the Rectangle is",rect.height)
	fmt.Println("Width of the Rectangle is",rect.width)
	fmt.Println("The area of the rectangle is",rect.area())
}

// Structure declaration
type rectangle struct {
	height float64
	width float64
}

// Custom function for the structure
func (x *rectangle) area() float64 {
	return x.height * x.width
	
}