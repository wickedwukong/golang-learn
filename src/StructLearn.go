package main

import (
	"fmt"
	"math"
)

type Rectangle struct {
  x1, y1, x2, y2 float64
} 

type Circle struct {
	x, y, r float64
}

func (c *Rectangle) area() float64 {
    l := c.y2 - c.y1
	w := c.x2 - c.x1
	*c = Rectangle{0, 0, 1, 99}

    fmt.Println("inside area()", c)
	return l * w
}

func calculateArea(circle *Circle) float64{
	return math.Pi * circle.r * circle.r
}

func main() {
	circle := Circle{x: 0, y: 0, r: 10}

    calculateArea(&circle)
	rectangle := Rectangle{0, 0, 10, 10}
	fmt.Println(rectangle.area())
}