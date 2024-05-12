package main

import (
	"bytes"
	"log"
	"mysql-gdrive-backup/services"
	"mysql-gdrive-backup/utils/compress"
	"os"

	"github.com/joho/godotenv"
)

func Setup() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
		panic(err)
	}
}

func main() {
	Setup()

	dumpCommand := os.Getenv("DUMP_COMMAND")
	dumpPath := os.Getenv("DUMP_PATH")
	s3Region := os.Getenv("S3_REGION")
	s3Bucket := os.Getenv("S3_BUCKET")
	s3AccessKey := os.Getenv("S3_ACCESS_KEY")
	s3SecretKey := os.Getenv("S3_SECRET_KEY")

	dumpService := services.NewDump(dumpCommand, dumpPath)

	err := dumpService.GenerateDumpFile()

	if err != nil {
		//TODO: Add email handler
		log.Fatal("Error generating dump:", err)
	}
	defer dumpService.DeleteDumpFile()

	if err != nil {
		log.Fatal("Error sending email:", err)
	}

	buf := new(bytes.Buffer)

	err = compress.CompressFile(dumpService.GetFullFilePath(), buf)

	if err != nil {
		log.Fatal("Error compressing dump:", err)
	}

	s3Service, err := services.NewS3Service(s3Region, s3Bucket, &s3AccessKey, &s3SecretKey)

	if err != nil {
		log.Fatal("Error creating S3 service:", err)
	}

	bufBytes := buf.Bytes()

	err = s3Service.Upload(*dumpService.FileName+"tar.gz", &bufBytes)

	if err != nil {
		log.Fatal("Error uploading dump to S3:", err)
	}

}
