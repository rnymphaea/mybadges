package s3

import (
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"

	"mybadges/internal/utils/errors"
)

type Storage struct {
	AccessKey string
	SecretKey string
	Endpoint  string
	Bucket    string
	Region    string
}

func New(akey, skey, ep, b, r string) *Storage {
	return &Storage{
		AccessKey: akey,
		SecretKey: skey,
		Endpoint:  ep,
		Bucket:    b,
		Region:    r,
	}
}

func (s *Storage) UploadFile(filename string) (string, error) {
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(s.Region),
		Credentials: credentials.NewStaticCredentials(s.AccessKey, s.SecretKey, ""),
		Endpoint:    aws.String(s.Endpoint)})
	if err != nil {
		return "", errors.ErrNoSession
	}

	uploader := s3manager.NewUploader(sess)

	file, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer file.Close()

	result, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(s.Bucket),
		Key:    aws.String(filename),
		Body:   file,
	})
	if err != nil {
		return "", err
	}

	return result.Location, nil
}
