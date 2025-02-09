package connection

import (
	"io"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

const (
	HTTP_NO_FILE = "http: no such file"
)

type Aws struct {
	session       *session.Session
	BucketName    string
	CloudFrontURL string
}

func NewAws(region, accessKeyId, secretKey, bucketName, cloudFrontURL string) (*Aws, error) {
	newSession, err := session.NewSession(
		&aws.Config{
			Region: aws.String(region),
			Credentials: credentials.NewStaticCredentials(
				accessKeyId,
				secretKey,
				"",
			),
		},
	)

	if err != nil {
		return nil, err
	}
	return &Aws{
		session:       newSession,
		BucketName:    bucketName,
		CloudFrontURL: cloudFrontURL,
	}, nil
}

func (awsInstance *Aws) UploadFile(reader io.Reader, fileName, bucketName, key string) (string, error) {
	uploader := s3manager.NewUploader(awsInstance.session, func(u *s3manager.Uploader) {
		u.PartSize = 5 * 1024 * 1024 // The minimum/default allowed part size is 5MB
		u.Concurrency = 2
	})

	up, err := uploader.Upload(&s3manager.UploadInput{
		Bucket:          aws.String(bucketName),
		ACL:             aws.String("public-read"),
		Key:             aws.String(key + "/" + fileName),
		ContentType:     aws.String("image/png"),
		Body:            reader,
		ContentEncoding: aws.String("base64"),
	})

	if err != nil {
		return "", err
	}

	end := strings.Split(up.Location, ".com/")

	return end[1], nil
}
