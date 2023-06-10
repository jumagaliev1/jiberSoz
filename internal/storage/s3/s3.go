package s3

import (
	"bytes"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/spf13/viper"
)

type AmazonS3 struct {
	Service *s3.S3
	bucket  string
}

func NewAmazonS3() *AmazonS3 {
	region := viper.GetString("s3.region")
	creds := credentials.NewStaticCredentials(viper.GetString("aws_access_key_id"), viper.GetString("aws_secret_access_key"), "")

	sess := session.Must(session.NewSessionWithOptions(session.Options{
		Config:            *aws.NewConfig().WithRegion(region).WithCredentials(creds),
		SharedConfigState: session.SharedConfigEnable,
	}))

	return &AmazonS3{
		Service: s3.New(sess),
		bucket:  viper.GetString("s3.bucket"),
	}
}

func (s *AmazonS3) Upload(key string, body string) error {
	_, err := s.Service.PutObject(&s3.PutObjectInput{
		Bucket: aws.String(s.bucket),
		Key:    aws.String(key),
		Body:   bytes.NewReader([]byte(body)),
	})

	return err
}

func (s *AmazonS3) Download(key string) (string, error) {
	obj, err := s.Service.GetObject(&s3.GetObjectInput{
		Bucket: aws.String(s.bucket),
		Key:    aws.String(key),
	})
	if err != nil {
		return "", err
	}
	defer obj.Body.Close()

	buf := new(bytes.Buffer)
	_, err = buf.ReadFrom(obj.Body)
	if err != nil {
		return "", err
	}

	content := buf.Bytes()

	return string(content), nil
}
