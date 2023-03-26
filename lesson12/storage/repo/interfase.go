package repo

import (
	"github.com/Golang-bootcamp/lesson12/json/methods"
)


type Methodlari interface {
	GetBooks() []methods.Kitob
	AddBook(methods.itob) []methods.Kitob
	RemoveBook(methods.Kitob) []methods.Kitob
	UpdateBook(methods.Kitob) methods.Kitob
	GetBookById(methods.Kitob) methods.Kitob
}