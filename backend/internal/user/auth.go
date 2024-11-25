package user

import (
	"chess-backend/config"
	"chess-backend/models"
	"fmt"
)

type UserData struct {
	Name  string
	Email string
	ID    uint
}

func Login(email string, password string) (UserData, error) {
	var user models.User 
	var userData UserData

	res := config.DB.Where("email = ?", email).First(&user)

	if res.Error != nil {
		fmt.Printf("Error fetching user: %v\n", res.Error)
		return userData, res.Error
	}

	userData = UserData{
		Name:  user.Username,
		Email: user.Email,
		ID:    user.ID,
	}


	return userData, nil
}
