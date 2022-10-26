package models

import (
	"errors"
	"fmt"
	"golang.org/x/crypto/bcrypt"
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
	Id        int    `gorm:"primaryKey"`
	FirstName string `gorm:"not null"`
	LastName  string `gorm:"not null"`
	Email     string `gorm:"not null,unique"`
	Password  string `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (r *UserModel) GetUserList() ([]User, error) {
	var users []User

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

func (r *UserModel) GetUserByEmail(email string) (User, error) {

	user := User{}

	response := r.Database.Table("users").Where("email = ?", email).First(&user)

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
	response := r.Database.Where("email = ?", user.Email).First(&userFind)

	if response.Error != nil && errors.Is(response.Error, gorm.ErrRecordNotFound) {
		hashedPassword, err := r.HashPassword(user.Password)

		if err != nil {
			return err
		}

		user.Password = hashedPassword

		response := r.Database.Create(&user)

		if response.Error != nil {
			return response.Error
		}

		fmt.Print(response)

		return nil
	} else {
		return errors.New("The email address has already been taken")
	}
}

func (r *UserModel) UpdateUser(user *User, body *User) error {
	var userFind User
	response := r.Database.
		Where("email = ?", body.Email).
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

func (r *UserModel) HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return "", err
	}

	//user.Password = string(bytes)
	return string(bytes), nil
}

func (r *UserModel) CheckPassword(userPassword string, providedPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(userPassword), []byte(providedPassword))
	if err != nil {
		return err
	}
	return nil
}
