package admindostup

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

var Meny PostMeny

type GetOrder []struct{
	ID int `json:"id"`
	Foods []string `json:"foods"`
	Price []int `json:"price"`
}


type PostMeny []struct{
	ID int `json:"id"`
	Name string `json:"name"`
	Price int `json:"price"`
}

func GetInfoOrder(){
	var orders GetOrder

	data,err:=ioutil.ReadFile("jsonData/orders.json")
	if err!=nil{
		panic(err)
	}

	err=json.Unmarshal(data,&orders)
	if err!=nil{
		panic(err)
	}

	for _,v:=range orders{
		fmt.Println("===================================")
		fmt.Printf("Buyurtma raqami: %v\n",v.ID)
		fmt.Printf("Buyurtmalar: %v\n",v.Foods)
		fmt.Printf("Buyurtma uchun to'lov: %v\n",v.Price)

	}
}

func PostNewFood(foodName string, price int){
	data,err:=ioutil.ReadFile("jsonData/meny.json")
	if err!=nil{
		panic(err)
	}
	err=json.Unmarshal(data,&Meny)
	if err!=nil{
		panic(err)
	}
	newFood := struct {
        ID int `json:"id"`
        Name string `json:"name"`
        Price int `json:"price"`
    }{
        ID: Meny[len(Meny)-1].ID + 1,
        Name: foodName,
        Price: price,
    }

	Meny = append(Meny, newFood)
	data,err=json.Marshal(&Meny)

	err = ioutil.WriteFile("jsonData/meny.json", data, 0644)
	if err != nil {
		fmt.Println("Error writing JSON to file:", err)
		return
	}


	fmt.Println("Taomnomaga yangi taom kiritildi ")
}