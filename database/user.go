package database

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `json:"username" gorm:"unique"`
	Password string `json:"password"`
}

type Image struct {
}

func RegisterUser(user *User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)
	return DB.Create(user).Error
}

func LoginUser(username, password string) (User, error) {
	var user User
	if err := DB.Where("username = ?", username).First(&user).Error; err != nil {
		return user, errors.New("user not found")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return User{}, errors.New("invalid password")
	}

	return user, nil
}
