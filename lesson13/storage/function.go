package storage

import (
	"io/ioutil"

	"github.com/Golang-bootcamp/lesson13/model"
)

type Products model.Products

func ReadFile(path string) []byte {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		panic("Read function error")
	}
	return data
}

func (p Products) GetProducts() []struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	Price      int    `json:"price"`
	CategoryID string `json:"category_id"`
} {
	return p
}
