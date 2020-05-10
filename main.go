package main

import (
	"io"
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

func main() {
	creds := credentials.NewStaticCredentials("AKIAQLVPVCZSTVH3EV4X", "cBEqkX6hmcAs9UOzDSIXp7o6N/W4cakdXbbU/NlH", "")
	sess, err := session.NewSession(&aws.Config{
		Credentials: creds,
		Region:      aws.String("ap-northeast-1")},
	)
	if err != nil {
		log.Println(err)
	}
	svc := s3.New(sess)

	bucket := "tokyocentertool"
	objectKey := "memberslist/test-list.csv"

	obj, err := svc.GetObject(&s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(objectKey),
	})
	if err != nil {
		log.Println(err)
	}

	filename := "filename.list"
	file, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	_, err = io.Copy(file, obj.Body)
	if err != nil {
		log.Fatal(err)
	}

	file, err = os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	uploader := s3manager.NewUploader(sess)
	_, err = uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(bucket),
		Key:    aws.String("test/test-list.csv"),
		Body:   file,
	})

	// file, err := os.Create("test.csv")
	// if err != nil {
	// 	log.Println(err)
	// }

	// defer file.Close()

}
