package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Methodlari interface {
	GetBooks() []Kitob
	AddBook(Kitob) []Kitob
	RemoveBook(Kitob) []Kitob
	UpdateBook(Kitob) Kitob
	GetBookById(Kitob) Kitob
	// GetBookByCategory(Book)
}

type Kutubxona struct {
	Kitoblar []Kitob
}
type Kitob struct {
	ID       int    `json:"id"`
	Title    string `json:"title"`
	Author   string `json:"author"`
	Year     int    `json:"year"`
	Status   string `json:"status"`
	Price    int    `json:"price"`
	Period   int    `json:"period"`
	Category string `json:"category"`
	Page     int    `json:"page"`
}

func (data Kutubxona) GetBooks() []Kitob {
	return data.Kitoblar
}

func (data Kutubxona) GetBooksById(id int) Kitob {
	var kitob1 Kitob
	for _, v := range data.Kitoblar {
		if v.ID == id {
			kitob1.ID = v.ID
			kitob1.Title = v.Title
			kitob1.Author = v.Author
			kitob1.Year = v.Year
			kitob1.Status = v.Status
			kitob1.Price = v.Price
			kitob1.Period = v.Period
			kitob1.Category = v.Category
			kitob1.Page = v.Page
			return kitob1
		}
	}
	return kitob1
}

func (data *Kutubxona) AddBook(kitob Kitob) []Kitob {
	data.Kitoblar = append(data.Kitoblar, kitob)
	kitoblar := data.Kitoblar

	jsonData, _ := json.Marshal(kitoblar)

	ioutil.WriteFile("books.json", jsonData, 0644)
	return data.Kitoblar
}

func (data *Kutubxona) RemoveBook(kitob Kitob) []Kitob{
	for i, v := range data.Kitoblar {
		if v.ID == kitob.ID {
			data.Kitoblar = append(data.Kitoblar[:i], data.Kitoblar[i+1:]...)
			break
		}
	}

	kitoblar := data.Kitoblar
	jsonData, _ := json.Marshal(kitoblar)
	ioutil.WriteFile("books.json", jsonData, 0644)

	return data.Kitoblar
}

func (data Kutubxona) UpdateBook(kitob Kitob) Kitob {
	for i, _ := range data.Kitoblar {
		if data.Kitoblar[i].ID == kitob.ID {
			data.Kitoblar[i].ID=kitob.ID 
			data.Kitoblar[i].Title=kitob.Title 
			data.Kitoblar[i].Author=kitob.Author
			data.Kitoblar[i].Year=kitob.Year 
			data.Kitoblar[i].Status=kitob.Status  
			data.Kitoblar[i].Price=kitob.Price
			data.Kitoblar[i].Period=kitob.Period 
			data.Kitoblar[i].Category=kitob.Category 
			data.Kitoblar[i].Page=kitob.Page 
			kitoblar := data.Kitoblar
			jsonData, _ := json.Marshal(kitoblar)
			ioutil.WriteFile("books.json", jsonData, 0644)
			return data.Kitoblar[i]
		}
	}
	return data.Kitoblar[kitob.ID]
}


func main() {
	var kitoblar []Kitob
	data, err := ioutil.ReadFile("books.json")
	if err != nil {
		panic("Json datani o'qishda hato")
	}
	err = json.Unmarshal(data, &kitoblar)
	if err != nil {
		panic("Unmarshal qilishda muammo")
	}

	kutubxona := Kutubxona{kitoblar}

	newBook := Kitob{
		ID:     1,
		Title:  "Jaloladdin",
		Author: "hadfioh",
		Year:   2004,
		Status: "given",
		Price:  890,
	}

	// fmt.Println(kutubxona.GetBooks())
	// fmt.Println(kutubxona.RemoveBook(newBook))
	// fmt.Println(kutubxona.GetBooksById(2))
	fmt.Println(kutubxona.UpdateBook(newBook))

}
ok := Kitob{
	ID:     1,
	Title:  "Jaloladdin",
	Author: "hadfioh",
	Year:   2004,
	Status: "given",
	Price:  890,
}