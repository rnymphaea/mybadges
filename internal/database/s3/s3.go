package s3

import (
	"io"

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

func (s *Storage) UploadFile(file io.Reader, key string) (string, error) {
	sess, err := session.NewSession(&aws.Config{
		Region:           aws.String(s.Region),
		Credentials:      credentials.NewStaticCredentials(s.AccessKey, s.SecretKey, ""),
		Endpoint:         aws.String(s.Endpoint),
		S3ForcePathStyle: aws.Bool(true)})
	if err != nil {
		return "", errors.ErrNoSession
	}

	uploader := s3manager.NewUploader(sess)

	result, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(s.Bucket),
		Key:    aws.String(key),
		Body:   file,
	})
	if err != nil {
		return "", err
	}

	return result.Location, nil
}
