package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	methods "github.com/Golang-bootcamp/lesson12/storage/json"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Sprintln(r)
		}
	}()
	var kitoblar []methods.Kitob
	data, err := ioutil.ReadFile("../storage/json/books.json")
	if err != nil {
		panic("Json datani o'qishda hato")
	}
	err = json.Unmarshal(data, &kitoblar)
	if err != nil {
		panic("Unmarshal qilishda muammo")
	}

	kutubxona := methods.Kutubxona{kitoblar}

	// newBook := methods.Kitob{
	// 	ID:     1,
	// 	Title:  "Jaloladdin",
	// 	Author: "hadfioh",
	// 	Year:   2004,
	// 	Status: "given",
	// 	Price:  890,
	// }

	// fmt.Println(kutubxona.GetBooks())
	// fmt.Println(kutubxona.RemoveBook(1))
	fmt.Println(kutubxona.GetBooksById(2))
	// fmt.Println(kutubxona.AddBook(newBook))
	// fmt.Println(kutubxona.UpdateBook(newBook))
}
