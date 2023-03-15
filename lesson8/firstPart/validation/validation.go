package validation

import (
	"fmt"
	"github.com/Golang-bootcamp/lesson8/function"
	"strings"
)

func InfoValidation(user function.Users, ticket int) (bool, bool, bool, bool) {
	isValidName := len(user.UserName) > 3 && len(user.UserLastName) > 3
	isValidEmail := strings.Contains(user.UserEmail, "@")
	isValidPhoneNumber := len(user.PhoneNumber) == 13
	if isValidPhoneNumber {
		if user.PhoneNumber[:4] == "+998" {
			isValidPhoneNumber = true
		} else {
			isValidPhoneNumber = false
		}
	}
	isTicked := user.TicketNumber <= uint(ticket)
	return isValidName, isValidEmail, isValidPhoneNumber, isTicked
}
func ErrorSended(validName, validEmail, validPhoneNumber, validTicked bool) {
	if !validName {
		fmt.Println("Ism familiyangizni to'g'ri kiriting ")
	}
	if !validEmail {
		fmt.Println("Emailingizni to'g'ri kiriting")
	}
	if !validPhoneNumber {
		fmt.Println("Telefon raqamingizni to'g'ri kiriting ")
	}
	if !validTicked {
		fmt.Printf("Chiptalarimiz %v ta iltimos shunga ko'ra kiriting\n", function.AllTickets)
	}
}
