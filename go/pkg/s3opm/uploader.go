package s3opm

import (
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

func S3UpLoader(s *session.Session, bucketName, objPath string, f *os.File) (*s3manager.UploadOutput, error) {
	u := s3manager.NewUploader(s)
	m, err := u.Upload(&s3manager.UploadInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(objPath),
		Body:   f,
	})
	if err != nil {
		return nil, err
	}
	return m, nil
}
