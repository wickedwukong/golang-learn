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

	return l * w
}

func calculateArea(circle *Circle) float64{
	return math.Pi * circle.r * circle.r
}

func main() {
	circle := Circle{x: 0, y: 0, r: 10}
	fmt.Println(calculateArea(&circle))

	rectangle := Rectangle{0, 0, 10, 10}
	fmt.Println(rectangle.area())
}