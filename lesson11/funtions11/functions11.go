package funtions11

import (
	"fmt"
	"math"
)

type Shape interface {
	area()
	perimeter()
}

type Circle struct {
	radius float64
}

type Rectangle struct {
	width  float64
	height float64
}

type Uchburchak struct{
	height float64
	tomoni1 float64
	tomoni2 float64
}


func (c Circle) area() float64 {
	return math.Pi*math.Pow(c.radius,2)
}
func (c Circle) perimeter() float64{
	return 2*math.Pi*c.radius
}

func (r Rectangle) area() float64{
	return r.width*r.height
}
func (r Rectangle) perimeter() float64{
	return 2*(r.width+r.height)
}

func (u Uchburchak) area() float64{
	return u.height*u.tomoni1/2
}
func (u Uchburchak) perimeter() float64{
	return u.height+u.tomoni1+u.tomoni2
}


func Return1(){
	u:=Uchburchak{height:2.0,tomoni1:2,tomoni2:3}
	c := Circle{radius: 2.5}
	r := Rectangle{width: 3.0, height: 4.0}
	
	fmt.Println(c.area())
	fmt.Println(r.area())
	fmt.Println(u.area())

	fmt.Println(c.perimeter())
	fmt.Println(r.perimeter())
	fmt.Println(u.perimeter())
}
