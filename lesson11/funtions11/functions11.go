package funtions11

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Methods interface {
	GetBooks()
	RemoveBook()
	UpdateBook()
	AddBook()
	GetBooksById()
}

type Book struct{
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


type Books struct {
	Books []Book
}


func (data Books) GetBooks()Books{
	return data
}
// func (data *Books) AddBook(title ,author string ,year int,status string,price, period int, category string,page int){
// 	var newBook map[string]interface{}
// 	*data=append(*data,newBook)
// }

func ReadData() {
	books:=Books{Books:[]Book{}}
	data, err := ioutil.ReadFile("books.json")
	if err != nil {
		panic(err)
	}
	err=json.Unmarshal(data, &books)
	fmt.Println(books)

}
