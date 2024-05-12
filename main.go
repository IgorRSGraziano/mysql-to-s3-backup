package main

import (
	"bytes"
	"mysql-gdrive-backup/services"
	"mysql-gdrive-backup/utils/compress"
	"mysql-gdrive-backup/utils/logger"
	"os"

	"github.com/joho/godotenv"
)

func Setup() {
	err := godotenv.Load()
	if err != nil {
		logger.Fatal("Error loading .env file")
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

	logger.Info("Generating dump file")
	err := dumpService.GenerateDumpFile()

	if err != nil {
		logger.Fatal("Error generating dump file:" + err.Error())
	}

	defer dumpService.DeleteDumpFile()

	if err != nil {
		logger.Fatal("Error deleting dump file:" + err.Error())
	}

	buf := new(bytes.Buffer)

	logger.Info("Compressing dump file")
	err = compress.CompressFile(dumpService.GetFullFilePath(), buf)
	logger.Info("Compressed dump file")

	if err != nil {
		logger.Fatal("Error compressing dump file:" + err.Error())
	}

	logger.Info("Creating S3 service")
	s3Service, err := services.NewS3Service(s3Region, s3Bucket, &s3AccessKey, &s3SecretKey)

	if err != nil {
		logger.Fatal("Error creating S3 service:" + err.Error())
	}

	bufBytes := buf.Bytes()

	logger.Info("Uploading file to S3")
	err = s3Service.Upload(*dumpService.FileName+"tar.gz", &bufBytes)

	if err != nil {
		logger.Fatal("Error uploading file to S3:" + err.Error())
	}

	logger.Success("Backup Success")
	services.SendWarnEmail("Backup Success", "Backup success")

}
