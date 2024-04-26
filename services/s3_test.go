package services

import (
	"log"
	"os"
	"testing"
	"time"

	"github.com/joho/godotenv"
)

func loadEnv() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
		panic(err)
	}
}

func createS3Service() (*S3Service, error) {
	loadEnv()
	region := os.Getenv("S3_REGION")
	bucket := os.Getenv("S3_BUCKET")
	accessKey := os.Getenv("S3_ACCESS_KEY")
	secretKey := os.Getenv("S3_SECRET_KEY")

	return NewS3Service(region, bucket, &accessKey, &secretKey)
}

func Test_NewS3Service(t *testing.T) {
	_, err := createS3Service()
	if err != nil {
		t.Error("Error creating session:", err)
	}
}

func Test_S3Service_Upload(t *testing.T) {
	service, err := createS3Service()
	if err != nil {
		t.Error("Error creating session:", err)
	}

	//file with current timestamp
	fileName := "test.txt"
	file := []byte(time.Now().String())

	err = service.Upload(fileName, &file)
	if err != nil {
		t.Error("Error uploading file:", err)
	}
}
