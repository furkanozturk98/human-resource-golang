package models

import (
	"errors"
	"human-resources-backend/configs/messages"
	"time"

	"gorm.io/gorm"
)

type EmployeeModel struct {
	Database *gorm.DB
}

func NewEmployeeModel(database *gorm.DB) *EmployeeModel {
	return &EmployeeModel{
		Database: database,
	}
}

type Employee struct {
	Id           int    `gorm:"primaryKey"`
	FirstName    string `gorm:"not null"`
	LastName     string `gorm:"not null"`
	EmailAddress string `gorm:"not null,unique"`
	Phone        string `gorm:"not null"`
	CompanyId    int    `gorm:"not null"`
	Company      Company
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func (r *EmployeeModel) GetEmployeeList() ([]Employee, error) {

	Employees := []Employee{}

	response := r.Database.Table("employees").Find(&Employees)

	if response.Error != nil {
		if errors.Is(response.Error, gorm.ErrRecordNotFound) {
			return Employees, messages.NO_RECORD_FOUND
		}
		return nil, response.Error
	}

	return Employees, nil
}

func (r *EmployeeModel) GetEmployeeById(id int) (Employee, error) {

	Employee := Employee{}
	//fmt.Printf("%d", id)
	response := r.Database.Table("employees").Where("id = ?", id).First(&Employee)

	//fmt.Printf("%+v\n", response)

	if response.Error != nil {
		if errors.Is(response.Error, gorm.ErrRecordNotFound) {
			return Employee, messages.NO_RECORD_FOUND
		}
		return Employee, response.Error
	}

	return Employee, nil
}

func (r *EmployeeModel) CreateEmployee(employee *Employee) error {
	var EmployeeFind Employee
	response := r.Database.
		Where("email_address = ?", employee.EmailAddress).
		First(&EmployeeFind)

	if response.Error != nil && errors.Is(response.Error, gorm.ErrRecordNotFound) {
		response = r.Database.Create(&employee)

		if response.Error != nil {
			return response.Error
		}

		return nil
	} else {
		return errors.New("The email address has already been taken")
	}
}

func (r *EmployeeModel) UpdateEmployee(employee *Employee, body *Employee) error {
	var EmployeeFind Employee
	response := r.Database.
		Where("email_address = ?", body.EmailAddress).
		Where("id != ?", employee.Id).
		First(&EmployeeFind)

	if response.Error != nil && errors.Is(response.Error, gorm.ErrRecordNotFound) {
		response = r.Database.
			Model(&employee).
			Updates(&body)

		if response.Error != nil {
			return response.Error
		}

		return nil
	} else {
		return errors.New("The email address has already been taken")
	}
}

func (r *EmployeeModel) DeleteEmployee(id int) error {
	response := r.Database.
		Delete(&Employee{}, id)

	if response.Error != nil {
		return response.Error
	}

	return nil
}
