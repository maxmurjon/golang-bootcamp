package adminvalidation

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Admin []struct{
	ID int `json:"id"`
	UserName string `json:"username"`
	Password string `json:"password"`
}

func CheckAdmin(username, password string) bool{
	defer func(){
		if r:=recover(); r!=nil{
			fmt.Println("Recover",r)
		}
	}()
	var adm Admin
	data, err := ioutil.ReadFile("jsonData/admins.json")
	if err != nil {
		panic(err)
	}
	err=json.Unmarshal(data,&adm)
	if err!=nil{
		panic(err)
	}
	for _,v:=range adm{
		if v.UserName==username&&v.Password==password{
			return true
		}else{
			return false
		}
	}
	return false

}
