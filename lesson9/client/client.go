package client

import (
	"encoding/json"
	"io/ioutil"
)

type Meny []struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
}

func GetMeny() Meny {
	var meny Meny
	data, err := ioutil.ReadFile("jsonData/meny.json")
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(data, &meny)
	return meny
}
