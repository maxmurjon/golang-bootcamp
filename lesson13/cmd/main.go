package main

import (
	"fmt"
)

// type Products model.Products
func main() {

	// user := ""
	// fmt.Printf("Ismingizni kiriting ")
	// fmt.Scanln(&user)
	// if validation.ValidUser(user) {
	// 	fmt.Println("Bizning Tovarlarimiz")
	// 	data := storage.ReadFile("../datas/product.json")

	// 	var product Products
	// 	json.Unmarshal(data,&product)
	// 	product.GetProducts()

	// } else {
	// 	fmt.Println("Siz sahifamizda login qilmagansiz\nXizmatlardan foydalanish uchun login qiling ")
	// }

	name := "ai1e1f1h1i1oehaife"
	bir := 0
	for i := range name {
		if name[i] == '1' {
			bir += 1
		}
	}
	fmt.Println(bir)
}
