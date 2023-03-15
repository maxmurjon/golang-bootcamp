package main

import (
	"github.com/Golang-bootcamp/lesson8/function"
	"github.com/Golang-bootcamp/lesson8/validation"
)

func main() {

	for *&function.AllTickets != 0 {
		//Get Users Data
		userInfo := function.GetUserData()

		//User information validation
		validName, validEmail, validPhoneNumber, validTicked := validation.InfoValidation(userInfo, function.AllTickets)

		//if validation is true
		if validName && validEmail && validPhoneNumber && validTicked {
			function.SendTicked(userInfo, &function.AllTickets, function.UsersData)
		} else { //if validation is false
			validation.ErrorSended(validName, validEmail, validPhoneNumber, validTicked)
		}
	}

}
