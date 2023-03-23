package main

import (
	"fmt"

	"github.com/Golang-bootcamp/lesson10/funcs"
)

func main() {
	data:=funcs.ParseJsontoStruct()
	kupyura:=""
	fmt.Printf("Kupyurani kiriting ")
	fmt.Scanln(&kupyura)
	var list []int
	list=funcs.FindData(kupyura)
	for _,v:=range list{
		fmt.Println(data[v])
	}

	
}
