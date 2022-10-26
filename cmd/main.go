package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"human-resources-backend/database"
	"human-resources-backend/routes"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("Application is starting")

	godotenv.Load()

	db := database.ConnectDB()
	database.Migration(db)

	awsSession := ConnectAWS()

	routes.RegisterRoutes(db, awsSession)
}

func ConnectAWS() *session.Session {
	AccessKeyID := os.Getenv("AWS_ACCESS_KEY_ID")
	SecretAccessKey := os.Getenv("AWS_SECRET_ACCESS_KEY")
	MyRegion := os.Getenv("AWS_REGION")

	sess, err := session.NewSession(
		&aws.Config{
			Region: aws.String(MyRegion),
			Credentials: credentials.NewStaticCredentials(
				AccessKeyID,
				SecretAccessKey,
				"",
			),
		})
	if err != nil {
		panic(err)
	}

	return sess
}
