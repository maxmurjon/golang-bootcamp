package requests

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	admindostup "github.com/Golang-bootcamp/lesson9/admins/adminDostup"
	adminvalidation "github.com/Golang-bootcamp/lesson9/admins/adminValidation"
	"github.com/Golang-bootcamp/lesson9/client"
)

type NewOrder struct {
	ID    int      `json:"id"`
	Foods []string `json:"foods"`
	Price []int    `json:"price"`
}

func Requests() {
	var orni string
	fmt.Printf("Assalomu alaykum \nO'rningizni tanlang (client,admin ) ")
	fmt.Scanln(&orni)
	if orni == "admin" {
		username := ""
		password := ""
		fmt.Println("Shaxsingizni tasdiqlang")
		fmt.Printf("Username ingizni kiriting ")
		fmt.Scanln(&username)
		fmt.Printf("Password ingizni kiriting ")
		fmt.Scanln(&password)
		if adminvalidation.CheckAdmin(username, password) {
			var request string
			fmt.Printf("getInfoOrder/postNewFood ")
			fmt.Scanln(&request)
			if request == "getInfoOrder" {
				admindostup.GetInfoOrder()
			} else if request == "postNewFood" {
				var fooName string
				var foodPrice int
				fmt.Printf("Taomning nomini kiriting ")
				fmt.Scanln(&fooName)
				fmt.Printf("Taomning narxini kiriting ")
				fmt.Scanln(&foodPrice)

				admindostup.PostNewFood(fooName, foodPrice)
			}
		}

	} else {
		var orderNum int = 0
		fmt.Println("Buyurtmalarni qabul qilishga tayyormiz ")
		fmt.Println("Bizning taomnomamiz ")
		meny := client.GetMeny()
		for _, v := range meny {
			fmt.Println("==================")
			fmt.Printf("Name: %v \n", v.Name)
			fmt.Printf("Price: %v \n", v.Price)
		}
		order := true
		var newOrder NewOrder
		for order {
			foods := ""
			var price int

			fmt.Printf("Buyurtmani kiriting ")
			fmt.Scanln(&foods)
			for _, v := range meny {
				if v.Name == foods {
					price = v.Price
				}
				orderNum=v.ID
			}
			newOrder.ID = orderNum+1
			newOrder.Foods = append(newOrder.Foods, foods)
			newOrder.Price = append(newOrder.Price, price)
			fmt.Println("Yana taom buyurtma qilasizmi? (true/false)")
			fmt.Scanln(&order)

		}
		var orders admindostup.GetOrder
		data, err := ioutil.ReadFile("jsonData/orders.json")
		if err != nil {
			panic(err)
		}
		err = json.Unmarshal(data, &orders)
		if err != nil {
			panic(err)
		}
		orders = append(orders, newOrder)
		jsonData, err := json.Marshal(orders)

		err = ioutil.WriteFile("jsonData/orders.json", jsonData, 0644)
		if err != nil {
			panic(err)
		}

	}
}
