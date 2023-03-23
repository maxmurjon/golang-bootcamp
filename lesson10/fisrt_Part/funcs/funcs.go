package funcs

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type JsonData []struct {
	ID       int    `json:"id"`
	Code     string `json:"Code"`
	Ccy      string `json:"Ccy"`
	CcyNmRU  string `json:"CcyNm_RU"`
	CcyNmUZ  string `json:"CcyNm_UZ"`
	CcyNmUZC string `json:"CcyNm_UZC"`
	CcyNmEN  string `json:"CcyNm_EN"`
	Nominal  string `json:"Nominal"`
	Rate     string `json:"Rate"`
	Diff     string `json:"Diff"`
	Date     string `json:"Date"`
}

func ParseJsontoStruct()JsonData{
	defer func (){
		if r:=recover();r!=nil{
			fmt.Println("Xato ",r)
		}
	}()


	var newData JsonData
	data,err:=ioutil.ReadFile("jsonData/data.json")
	if err!=nil{
		panic(err)
	}
	err=json.Unmarshal(data,&newData)
	if err!=nil{
		panic(err)
	}
	return newData
}

func FindData(money string)[]int{
	var list []int
	data:=ParseJsontoStruct()
	for i,v:=range data{
		if v.Ccy==money{
			list = append(list, i)
		}
	}
	return list
}