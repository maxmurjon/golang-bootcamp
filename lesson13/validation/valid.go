package validation

import (
	"encoding/json"

	"github.com/Golang-bootcamp/lesson13/model"
	"github.com/Golang-bootcamp/lesson13/storage"
)

func ValidUser(userName string)bool {
	var users model.Users
	data := storage.ReadFile("../datas/users.json")
	json.Unmarshal(data,&users)
	for _, v := range users {
		if v.Name == userName {
			return true
		}
	}
	return false
}
