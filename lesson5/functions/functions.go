package functions

import "fmt"

var Boshjoy int = 50
var Malumotlar = []map[string]interface{}{}
var Ticket = 0

type People struct {
	FirsName    string
	LastName    string
	Email       string
	TickedCount int
}

func BoshJoy(boshjoy *int, joy int) bool {
	if joy <= *boshjoy {
		*boshjoy -= joy
		return true
	}
	return false
}

func StructPeople() {
	firsName := ""
	lastName := ""
	email := ""
	tickedCount := Ticket
	fmt.Printf("Ismingizni kiriting ")
	fmt.Scanln(&firsName)
	fmt.Printf("Familiyangizni kiriting ")
	fmt.Scanln(&lastName)
	fmt.Printf("Emailingizni kiriting ")
	fmt.Scanln(&email)
	user1 := People{FirsName: firsName, LastName: lastName, Email: email, TickedCount: tickedCount}
	Malumotlar = PostPlace(Malumotlar, user1)
}


func PostPlace(arr []map[string]interface{}, people People) []map[string]interface{} {
	user := make(map[string]interface{})
	user["FirstName"] = people.FirsName
	user["LastName"] = people.LastName
	user["Email"] = people.Email
	user["TickedCount"] = people.TickedCount
	arr = append(arr, user)
	return arr
}

func GetAllUsers(data []map[string]interface{}) {
	for _, v := range data {
		fmt.Println("=============================")
		fmt.Printf("First name %v\n", v["FirstName"])
		fmt.Printf("Last name %v\n", v["LastName"])
		fmt.Printf("Email name %v\n", v["Email"])
		fmt.Printf("Last name %v\n", v["TickedCount"])
		fmt.Println("=============================")
	}
}
