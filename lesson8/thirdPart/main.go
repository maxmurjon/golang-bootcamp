package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Car struct{
	Name string `json:"name"`
	Model string `json:"model"`
	Year int `json:"year"`
}


func main(){
	var car []Car
	data ,err:=ioutil.ReadFile("cars.json")
	if err!=nil{
		fmt.Println(err)
		return
	}
	err=json.Unmarshal(data,&car)
	if err!=nil{
		fmt.Println(err)
		return
	}
	for _,v:=range car{
		fmt.Printf("Car name: %v\n",v.Name)
		fmt.Printf("Car model: %v\n",v.Model)
		fmt.Printf("Car year: %v\n",v.Year)
	}
}
