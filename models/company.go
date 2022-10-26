package models

import (
	"errors"
	"fmt"
	"human-resources-backend/configs/messages"
	"human-resources-backend/validators"
	"time"

	"gorm.io/gorm"
)

type Company struct {
	Id        int    `gorm:"primaryKey"`
	Name      string `gorm:"not null"`
	Address   string `gorm:"not null"`
	Phone     string `gorm:"not null"`
	Email     string `gorm:"not null,email"`
	Logo      string `gorm:"not null"`
	Website   string `gorm:"not null,url"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type CompanyModel struct {
	Database *gorm.DB
}

func NewCompanyModel(database *gorm.DB) *CompanyModel {
	return &CompanyModel{
		Database: database,
	}
}

func (r *CompanyModel) GetCompanyList() ([]Company, error) {

	companies := []Company{}

	response := r.Database.Table("companies").Find(&companies)

	if response.Error != nil {
		if errors.Is(response.Error, gorm.ErrRecordNotFound) {
			return companies, messages.NO_RECORD_FOUND
		}
		return nil, response.Error
	}

	return companies, nil
}

func (r *CompanyModel) GetCompanyById(id int) (Company, error) {

	Company := Company{}
	//fmt.Printf("%d", id)
	response := r.Database.Table("companies").Where("id = ?", id).First(&Company)

	//fmt.Printf("%+v\n", response)

	if response.Error != nil {
		if errors.Is(response.Error, gorm.ErrRecordNotFound) {
			return Company, messages.NO_RECORD_FOUND
		}
		return Company, response.Error
	}

	return Company, nil
}

func (r *CompanyModel) CreateCompany(data *validators.Company, company *Company, aswKey, imageURL string) error {
	response := r.Database.Create(&company)

	fmt.Print(response)

	if response.Error != nil {
		return response.Error
	}

	return nil
}

func (r *CompanyModel) UpdateCompany(Company *Company, body *Company) error {

	response := r.Database.
		Model(&Company).
		Updates(&body)

	fmt.Print(response)

	return nil
}

func (r *CompanyModel) DeleteCompany(id int) {
	r.Database.
		Delete(&Company{}, id)
}
