package main

import "fmt"
var r Rectangle
func main() {
	rect1 := Rectangle{width:5, height:10}
	rect2 := Rectangle{5,20}

	fmt.Println(rect1.height)
	fmt.Println(rect2.height)

	fmt.Println("Area of rect2", rect2.area())
	fmt.Println("Area of rect1", rect1.area())
}

type Rectangle struct {
	height float64
	width float64
}

func (rect *Rectangle) area() float64 {
	return rect.height * rect.width
}
