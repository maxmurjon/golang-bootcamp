package main

import (
	"github.com/Golang-bootcamp/lesson11/storage/repo"
	"encoding/json"
	"fmt"
	"io/ioutil"
)




func main() {
	var kitoblar []repo.Kitob
	data, err := ioutil.ReadFile("books.json")
	if err != nil {
		panic("Json datani o'qishda hato")
	}
	err = json.Unmarshal(data, &kitoblar)
	if err != nil {
		panic("Unmarshal qilishda muammo")
	}

	kutubxona := repo.Kutubxona{kitoblar}

	// newBook := repo.Kitob{
	// 	ID:     1,
	// 	Title:  "Jaloladdin",
	// 	Author: "hadfioh",
	// 	Year:   2004,
	// 	Status: "given",
	// 	Price:  890,
	// }

	// fmt.Println(kutubxona.GetBooks())
	// fmt.Println(kutubxona.RemoveBook(1))
	// fmt.Println(kutubxona.GetBooksById(2))
	fmt.Println(kutubxona.GetBooks(newBook))

}
// ok := Kitob{
// 	ID:     1,
// 	Title:  "Jaloladdin",
// 	Author: "hadfioh",
// 	Year:   2004,
// 	Status: "given",
// 	Price:  890,
// }