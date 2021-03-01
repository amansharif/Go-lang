package main

import (
	"fmt"
	"math"
)

func main ()  {
	rect := rectangle{50,60}
	circ := circle{7}
	fmt.Println("The area of the rectangle is",getArea(rect))
	fmt.Println("The area of the circle is",getArea(circ))
}

// interface declaration
type Shape interface {
	area() float64
}

type rectangle struct {
	height float64
	width float64
}

type circle struct {
	radius float64
}

func (r rectangle) area() float64 {
	return r.height * r.width
}
func (c circle) area() float64 {
	return math.Pi * math.Pow(c.radius,2)
}
func getArea(shape Shape) float64 {
	return shape.area()
}