package helpers

import (
	"bytes"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/h2non/bimg"
	"io"
	"log"
	"mime/multipart"
	"os"
)

func OptimizeImage(file *multipart.File) (*bytes.Reader, error) {
	buffer, err := io.ReadAll(*file)
	if err != nil {
		panic(err)
	}

	newImage, err := bimg.NewImage(buffer).Convert(bimg.WEBP)
	return bytes.NewReader(newImage), err
}

func UploadFile(awsSession *session.Session, file *multipart.FileHeader, name string, bucketName ...string) (string, error) {

	uploader := s3manager.NewUploader(awsSession)
	open, err := file.Open()
	defer func(open multipart.File) {
		err := open.Close()
		if err != nil {
			log.Println("Error: ", err)
		}
	}(open)
	image, err := OptimizeImage(&open)
	if err != nil {
		return "", err
	}

	if err != nil {
		return "", err
	}
	bucket := os.Getenv("S3_BUCKET")
	if len(bucketName) > 0 {
		bucket = bucketName[0]
	}

	key := fmt.Sprintf("images/%s%s", name, ".webp")
	_, err = uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(bucket),
		ACL:    aws.String("public-read"),
		Key:    aws.String(key),
		Body:   image,
	})
	if err != nil {
		return "", err
	}
	imgUrl := fmt.Sprintf("%s/%s", os.Getenv("S3_OBJECT_URL"), key)
	return imgUrl, nil
}
