package main

import (
	"bytes"

	"github.com/IgorRSGraziano/mysql-to-s3-backup/models"
	"github.com/IgorRSGraziano/mysql-to-s3-backup/services"
	"github.com/IgorRSGraziano/mysql-to-s3-backup/utils/compress"
	"github.com/IgorRSGraziano/mysql-to-s3-backup/utils/logger"
)

func main() {
	config := models.LoadConfig()

	dumpService := services.NewDump(config.Dump.Command)

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
	s3Service, err := services.NewS3Service(config.S3.Region, config.S3.Bucket, &config.S3.AccessKey, &config.S3.SecretKey)

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
