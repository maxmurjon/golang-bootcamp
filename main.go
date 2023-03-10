package main

import "fmt"

func main() {
	var boshjoy int = 50
	malumotlar :=[]map[string]interface{}{}
	firstnames:=make([]string,0)
	req:=0
	for true{
		fmt.Printf("Nechta bilet kerak ")
		fmt.Scanln(&req)
		if boshJoy(&boshjoy, req){
			firsName:=""
			lastName:=""
			email:=""
			tickedCount:=req
			fmt.Printf("Ismingizni kiriting ")
			fmt.Scanln(&firsName)
			fmt.Printf("Familiyangizni kiriting ")
			fmt.Scanln(&lastName)
			fmt.Printf("Emailingizni kiriting ")
			fmt.Scanln(&email)
			user1:=People{FirsName: firsName,LastName: lastName,Email: email,TickedCount: tickedCount}
			malumotlar=postPlace(malumotlar,user1)
			firstnames=append(firstnames,firsName)
		}else{
			fmt.Printf("JOylar soni %v \n",boshjoy)
			fmt.Printf("Iltimos %v kamroq bilet kiriting \n",req)
			fmt.Println("===============================")
	}
	fmt.Println(firstnames)
}
}

type People struct {
	FirsName    string
	LastName    string
	Email       string
	TickedCount int
}

func boshJoy(boshjoy *int, joy int) bool {
	if joy <= *boshjoy {
		*boshjoy -= joy
		return true
	}
	return false
}
func postPlace(arr []map[string]interface{}, people People) []map[string]interface{} {
	user := make(map[string]interface{})
	user["FirstName"] = people.FirsName
	user["LastName"] = people.LastName
	user["Email"] = people.Email
	user["TickedCount"] = people.TickedCount
	arr = append(arr, user)
	return arr
}
