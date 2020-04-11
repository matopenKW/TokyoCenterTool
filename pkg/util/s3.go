package util

import (
	"NameIdentification/pkg"
	"errors"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"gopkg.in/ini.v1"
)

func GetS3Session() (*session.Session, error) {
	c, _ := ini.Load(pkg.CONFIG_PATH)
	region := c.Section("s3").Key("region").String()
	accessKey := c.Section("s3").Key("accessKey").String()
	secretKey := c.Section("s3").Key("secretKey").String()

	creds := credentials.NewStaticCredentials(accessKey, secretKey, "")
	sess, err := session.NewSession(&aws.Config{
		Credentials: creds,
		Region:      aws.String(region)},
	)
	if err != nil {
		return nil, err
	}
	return sess, nil
}

func GetS3Client() (*s3.S3, error) {

	sess, err := GetS3Session()
	if err != nil {
		return nil, err
	}

	svc := s3.New(sess)
	if svc == nil {
		return nil, errors.New("Fail S3 Clieint")
	}
	return svc, nil
}

func GetObject(svc *s3.S3, bucktet, objectKey string) (*s3.GetObjectOutput, error) {
	if svc == nil {
		return nil, errors.New("Not S3 Clieint")
	}

	obj, err := svc.GetObject(&s3.GetObjectInput{
		Bucket: aws.String(bucktet),
		Key:    aws.String(objectKey),
	})
	if err != nil {
		return nil, err
	}

	return obj, nil
}

func Upload(sess *session.Session, bucket, objectKey string, file *os.File) error {
	uploader := s3manager.NewUploader(sess)
	_, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(objectKey),
		Body:   file,
	})

	if err != nil {
		return err
	}
	return nil
}
