package database

import (
	"fmt"
	"human-resources-backend/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

/* var (
	host     = os.Getenv("POSTGRES_HOST")
	port     = os.Getenv("POSTGRES_PORT")
	user     = os.Getenv("POSTGRES_USER")
	dbname   = os.Getenv("POSTGRES_DB")
	password = os.Getenv("POSTGRES_PASS")
) */

var (
	host     = "localhost"
	port     = "5432"
	user     = "root"
	dbname   = "human_resource_golang"
	password = "root"
)

// ConnectDB is used to connect to database
func ConnectDB() *gorm.DB {

	connectionString := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		host, user, password, dbname, port,
	)

	DB, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	return DB
}

func Disconnect(DB *gorm.DB) {
	connection, err := DB.DB()
	if err != nil {
		panic(err)
	}
	connection.Close()
}

func Migration(DB *gorm.DB) {
	DB.AutoMigrate(&models.User{}, &models.Company{}, &models.Employee{})
}
