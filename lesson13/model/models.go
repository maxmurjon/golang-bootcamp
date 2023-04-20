package model

type Methods interface{
	GetProduct()[]Products
	
}


type Products []struct{
	ID         string `json:"id"`
	Name       string `json:"name"`
	Price      int    `json:"price"`
	CategoryID string `json:"category_id"`
}

type Users []struct{
	ID string `json:"id"`
	Name       string `json:"name"`
	Surname      int    `json:"surname"`
	Balance string `json:"balance"`
}