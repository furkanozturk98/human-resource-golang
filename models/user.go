package models

import (
	"errors"
	"fmt"
	"human-resources-backend/configs/messages"
	"time"

	"gorm.io/gorm"
)

type UserModel struct {
	Database *gorm.DB
}

func NewUserModel(database *gorm.DB) *UserModel {
	return &UserModel{
		Database: database,
	}
}

type User struct {
	Id           int    `gorm:"primaryKey"`
	FirstName    string `gorm:"not null"`
	LastName     string `gorm:"not null"`
	EmailAddress string `gorm:"not null,unique"`
	Password     string `gorm:"not null"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func (r *UserModel) GetUserList() ([]User, error) {

	users := []User{}

	response := r.Database.Table("users").Find(&users)

	if response.Error != nil {
		if errors.Is(response.Error, gorm.ErrRecordNotFound) {
			return users, messages.NO_RECORD_FOUND
		}
		return nil, response.Error
	}

	return users, nil
}

func (r *UserModel) GetUserById(id int) (User, error) {

	user := User{}
	//fmt.Printf("%d", id)
	response := r.Database.Table("users").Where("id = ?", id).First(&user)

	//fmt.Printf("%+v\n", response)

	if response.Error != nil {
		if errors.Is(response.Error, gorm.ErrRecordNotFound) {
			return user, messages.NO_RECORD_FOUND
		}
		return user, response.Error
	}

	return user, nil
}

func (r *UserModel) CreateUser(user *User) error {
	var userFind User
	response := r.Database.Where("email_address = ?", user.EmailAddress).First(&userFind)

	if response.Error != nil && errors.Is(response.Error, gorm.ErrRecordNotFound) {
		result := r.Database.Create(&user)

		fmt.Print(result)

		return nil
	} else {
		return errors.New("The email address has already been taken")
	}
}

func (r *UserModel) UpdateUser(user *User, body *User) error {
	var userFind User
	response := r.Database.
		Where("email_address = ?", body.EmailAddress).
		Where("id != ?", user.Id).
		First(&userFind)

	if response.Error != nil && errors.Is(response.Error, gorm.ErrRecordNotFound) {
		result := r.Database.
			Model(&user).
			Updates(&body)

		fmt.Print(result)

		return nil
	} else {
		return errors.New("The email address has already been taken")
	}
}

func (r *UserModel) DeleteUser(id int) {
	r.Database.
		Delete(&User{}, id)
}
