package repo

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

type Methodlari interface {
	GetBooks() []Kitob
	AddBook(Kitob) []Kitob
	RemoveBook(Kitob) []Kitob
	UpdateBook(Kitob) Kitob
	GetBookById(Kitob) Kitob
	// GetBookByCategory(Book)
}