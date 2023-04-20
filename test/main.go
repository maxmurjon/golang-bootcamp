package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	// "sort"
)

type ShopCards struct {
	shopCard []ShopCard
}

type ShopCard struct {
	ID        string `json:"id"`
	ProductID string `json:"productId"`
	UserID    string `json:"userID"`
	Count     int    `json:"count"`
	Status    bool   `json:"status"`
	Date      string `json:"date"`
}

type Methods interface {
	SortDataWithDate()
}

func main() {

	data, err := ioutil.ReadFile("shopCard.json")
	if err != nil {
		fmt.Println("Failed to open file ", err)
		return
	}

	var shopcard []ShopCard

	err = json.Unmarshal(data, &shopcard)
	if err != nil {
		fmt.Println(err)
	}

	shopcards := ShopCards{shopcard}

	// fmt.Println(shopcards)

	fmt.Println(shopcards.SortDataWithDate("2022-02-13","2022-12-19"))

}

func (data ShopCards) SortDataWithDate(fromDate, toDate string) []ShopCard{
	// fmt.Println(sortedData)
	newDate := []ShopCard{}
	count := 0
		for i, v := range data.shopCard {
			if v.Date >=fromDate && v.Date<=toDate  {
				newDate = append(newDate, data.shopCard[i])
			}
		count += 1
	}
	return newDate
}
