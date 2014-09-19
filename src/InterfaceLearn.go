package main

import (
	"fmt"
	"math"
)

type Circle struct {
	x, y, r float64
}

func (circle Circle) area() float64{
	return math.Pi * circle.r * circle.r
}

type Rectangle struct {
	x1, y1, x2, y2 float64
}

func (c Rectangle) area() float64 {
    l := c.y2 - c.y1
	w := c.x2 - c.x1

	return l * w
}

type Shape interface {
    area() float64
}

type MultiShape struct {
	shapes []Shape
}

func (multiShape MultiShape) area() float64 {
	var area float64

	for _, multiShape := range multiShape.shapes {
		area += multiShape.area()
	}

	return area
}

func totalArea(shapes ...Shape) float64 {
    var area float64
    for _, s := range shapes {
        area += s.area()
    }
    return area
}

func main() {
	multiShape := MultiShape{shapes: []Shape{Circle{x: 0, y: 0, r: 10}}}

	rectangle := Rectangle{x1: 0, y1: 0, x2: 10, y2: 10}
	circle := Circle{x: 0, y: 0, r: 10}

	fmt.Println("The area of a multi shape is", totalArea(multiShape))
	fmt.Println("The area of all shapes is", totalArea(rectangle, circle, multiShape))

}