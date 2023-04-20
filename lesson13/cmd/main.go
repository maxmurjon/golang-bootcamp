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
	name := "((([{]})()]]}"

	// count:=0
	indexs:=make(map[int]bool)
	map1 := make(map[string]string)
	map1["("] = ")"
	map1["{"] = "}"
	map1["["] = "]"
	for i := 0; i < len(name)-1; i++ {
		// if map1[string(name[i])]==string(name[i+1]){
		// 	fmt.Println(true)
		// }

		if !indexs[i]{
			for j := i; j < len(name); j++ {
				if map1[string(name[i])] == string(name[j]) {	
					indexs[j]=true
					break
				}
			}
		}else{
			continue
		}
	}
	fmt.Println(len(indexs))
	fmt.Println(indexs)
}
