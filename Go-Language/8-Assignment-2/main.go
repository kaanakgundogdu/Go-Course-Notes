package main

import (
	"fmt"
)

type triangle struct {
	base   float64
	height float64
}

type square struct {
	sideLength float64
}

type shape interface{
	getArea()float64
}

func main() {
	triangle1:=triangle{1.5,3.5}
	square1:=square{10}

	printArea(triangle1)
	printArea(square1)
}

func (t triangle) getArea() float64 {
	var val = float64(0.5)
	return (val * t.base * t.height)
}

func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}

func printArea(s shape) {
	fmt.Println("Area is:",s.getArea())
}
