package main

import (
	"fmt"
	"github.com/lesson5/functions"
)

func main() {
	
	for functions.Boshjoy > 0 {
		fmt.Printf("Nechta bilet kerak ")
		fmt.Scanln(&functions.Ticket)
		if functions.BoshJoy(&functions.Boshjoy, functions.Ticket) {
			functions.StructPeople()
		} else {
			fmt.Printf("JOylar soni %v \n", functions.Boshjoy)
			fmt.Printf("Iltimos %v kamroq bilet kiriting \n", functions.Ticket)
			fmt.Println("===============================")
		}
	}
	functions.GetAllUsers(functions.Malumotlar)
}