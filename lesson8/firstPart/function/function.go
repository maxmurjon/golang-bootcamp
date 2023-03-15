package function

import "fmt"

var AllTickets =50
var UsersData=[]map[string]interface{}{}

type Users struct{
	UserName string
	UserLastName string
	UserEmail string
	TicketNumber uint
	PhoneNumber string
}

func GetUserData()(Users){
	var userName string
	var userLastName string
	var userEmail string
	var ticketNumber uint
	var phoneNumber string

	fmt.Printf("Ismingizni kiriting ")
	fmt.Scanln(&userName)

	fmt.Printf("Familiyangizni kiriting ")
	fmt.Scanln(&userLastName)

	fmt.Printf("Emailingizni kiriting ")
	fmt.Scanln(&userEmail)
	
	fmt.Printf("Telefon raqamingizni kiriting ")
	fmt.Scanln(&phoneNumber)
	
	fmt.Printf("Nechta chipta olmoqchisiz ")
	fmt.Scanln(&ticketNumber)

	user:=Users{UserName: userName,UserLastName: userLastName,UserEmail: userEmail,TicketNumber: ticketNumber,PhoneNumber: phoneNumber}
	return user 
}	

func SendTicked(user Users,allticket *int,arr []map[string]interface{})(string){
	*allticket-=int(user.TicketNumber)
	user1:= make(map[string]interface{})
	user1["First Name"]=user.UserName
	user1["Last Name"]=user.UserLastName
	user1["Email"]=user.UserEmail
	user1["Phone number"]=user.PhoneNumber
	user1["Ticket number"]=user.TicketNumber
	arr = append(arr, user1)
	return "Succsessfully"
}