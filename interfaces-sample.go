package main

import "fmt"
import "math"

func main() {
	rect := Rectangle{50,60}
	circle := Circle{7}
	
	fmt.Println("Area of rectangle: ", getArea(rect))
	fmt.Println("Area of circle: ", getArea(circle))

}

type Rectangle struct {
	height float64
	width float64
}

func (r Rectangle) area() float64 {
	return r.height * r.width
}
type Circle struct {
	raidus float64
}

func (c Circle) area() float64 {
	return math.Pi * math.Pow(c.raidus,2)
}

type Shape interface {
	area() float64
}

func getArea(shape Shape) float64 {
	return shape.area()
}
