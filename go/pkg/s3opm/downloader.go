package s3opm

import (
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

func S3DownLoader(s *session.Session, f *os.File, bucketName, objPath string) (int64, error) {
	d := s3manager.NewDownloader(s)
	m, err := d.Download(f, &s3.GetObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(objPath),
	})
	if err != nil {
		return 0, err
	}
	return m, nil
}
