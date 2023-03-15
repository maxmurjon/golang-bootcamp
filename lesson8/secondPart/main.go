package main

import "fmt"


type Rect struct{
	width ,height int
}

func (r *Rect) apply()(int){
	return r.width*r.height
}
func (r Rect) perimetr()(int){
	return 2*(r.width+r.height)
}

func main(){
	kvadrat:=Rect{width: 2,height: 2}
	fmt.Println(kvadrat.apply())
	fmt.Println(kvadrat.perimetr())
}