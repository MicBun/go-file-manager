package models

import (
	gorm "gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserInterface interface {
	ResetUserDatabase(db *gorm.DB) error
	Login(db *gorm.DB, user *User) error
	GetUsernameByID(db *gorm.DB, id uint) (string, error)
}

func NewUser() UserInterface {
	return &User{}
}

func (u *User) ResetUserDatabase(db *gorm.DB) error {
	db.Migrator().DropTable(&User{})
	db.Migrator().CreateTable(&User{})

	users := []User{
		{
			Username: "admin",
			Password: "admin",
		},
		{
			Username: "user1",
			Password: "password1",
		},
		{
			Username: "user2",
			Password: "password2",
		},
	}

	for _, user := range users {
		err := db.Create(&user).Error
		if err != nil {
			return err
		}
	}

	return nil
}

func (u *User) Login(db *gorm.DB, user *User) error {
	err := db.Where("username = ? AND password = ?", user.Username, user.Password).First(&user).Error
	if err != nil {
		return err
	}
	return nil
}

func (u *User) GetUsernameByID(db *gorm.DB, id uint) (string, error) {
	err := db.Where("id = ?", id).First(&u).Error
	if err != nil {
		return "", err
	}
	return u.Username, nil
}
