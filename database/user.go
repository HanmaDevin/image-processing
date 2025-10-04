package database

import (
	"errors"

	"github.com/HanmaDevin/image-processing/types"
	"golang.org/x/crypto/bcrypt"
)

func RegisterUser(user *types.User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)
	return DB.Create(user).Error
}

func LoginUser(username, password string) (types.User, error) {
	var user types.User
	if err := DB.Where("username = ?", username).First(&user).Error; err != nil {
		return user, errors.New("user not found")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return types.User{}, errors.New("invalid password")
	}

	return user, nil
}
