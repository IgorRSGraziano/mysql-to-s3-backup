package services

import (
	"bytes"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

type S3Service struct {
	Region string
	Bucket string
	client *s3.S3
}

func NewS3Service(region, bucket string, accessKey *string, secretKey *string) (*S3Service, error) {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(region),
		Credentials: credentials.NewStaticCredentials(
			*accessKey,
			*secretKey,
			"",
		),
	})

	if err != nil {
		fmt.Println("Error creating session:", err)
		return nil, err
	}

	svc := s3.New(sess)

	return &S3Service{
		Region: region,
		Bucket: bucket,
		client: svc,
	}, nil
}

func (s *S3Service) Upload(fileName string, file *[]byte) error {
	_, err := s.client.PutObject(&s3.PutObjectInput{
		Bucket: aws.String(s.Bucket),
		Key:    aws.String(fileName),
		Body:   bytes.NewReader(*file),
	})
	return err
}
