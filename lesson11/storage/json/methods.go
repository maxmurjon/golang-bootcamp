package methods

import (
	""
	"encoding/json"
	"io/ioutil"
)

func (data repo.Kutubxona) GetBooks() []repo.Kitob {
	return data.Kitoblar
}

func (data repo.Kutubxona) GetBooksById(id int) repo.Kitob {
	var kitob1 repo.Kitob
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

func (data *repo.Kutubxona) AddBook(kitob repo.Kitob) []repo.Kitob {
	data.Kitoblar = append(data.Kitoblar, kitob)
	kitoblar := data.Kitoblar

	jsonData, _ := json.Marshal(kitoblar)

	ioutil.WriteFile("books.json", jsonData, 0644)
	return data.Kitoblar
}

func (data *repo.Kutubxona) RemoveBook(kitob repo.Kitob) []repo.Kitob {
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

func (data repo.Kutubxona) UpdateBook(kitob repo.Kitob) repo.Kitob {
	for i, _ := range data.Kitoblar {
		if data.Kitoblar[i].ID == kitob.ID {
			data.Kitoblar[i].ID = kitob.ID
			data.Kitoblar[i].Title = kitob.Title
			data.Kitoblar[i].Author = kitob.Author
			data.Kitoblar[i].Year = kitob.Year
			data.Kitoblar[i].Status = kitob.Status
			data.Kitoblar[i].Price = kitob.Price
			data.Kitoblar[i].Period = kitob.Period
			data.Kitoblar[i].Category = kitob.Category
			data.Kitoblar[i].Page = kitob.Page
			kitoblar := data.Kitoblar
			jsonData, _ := json.Marshal(kitoblar)
			ioutil.WriteFile("books.json", jsonData, 0644)
			return data.Kitoblar[i]
		}
	}
	return data.Kitoblar[kitob.ID]
}
