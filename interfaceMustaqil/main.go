package main

import (
    "fmt"
)

type Shape interface {
    area() float64
}

type Rectangle struct {
    width  float64
    height float64
}

type Circle struct {
    radius float64
}

func (r Rectangle) area() float64 {
    return r.width * r.height
}

func (c Circle) area() float64 {
    return 3.14 * c.radius * c.radius
}

func getArea(shape Shape) float64 {
    return shape.area()
}

func main() {
    var s Shape
    s = Rectangle{width: 5, height: 10}
    fmt.Println("Area of rectangle: ", getArea(s))
    s = Circle{radius: 7}
    fmt.Println("Area of circle: ", getArea(s))
}
