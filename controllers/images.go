package controllers

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/chirag3003/ecommerce-golang-api/helpers"
	"github.com/chirag3003/ecommerce-golang-api/repository"
	"github.com/gofiber/fiber/v2"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

type Images interface {
	Upload(ctx *fiber.Ctx) error
}

func ImagesControllers() Images {
	return &imagesRoutes{
		repository.NewImagesRepo(conn.DB()),
		helpers.ConnectS3(),
	}
}

type imagesRoutes struct {
	Images     repository.ImagesRepository
	awsSession *session.Session
}

func (i imagesRoutes) Upload(ctx *fiber.Ctx) error {

	uploader := s3manager.NewUploader(i.awsSession)

	form, err := ctx.MultipartForm()
	if err != nil {

		return err
	}
	files := form.File["images"]
	var imageUrl []string
	for i, file := range files {
		if !strings.HasPrefix(file.Header["Content-Type"][0], "image/") {
			return ctx.Status(fiber.StatusBadRequest).JSON("file type not supported")
		}
		rand.Seed(time.Now().Unix())
		name := fmt.Sprintf("%d%s", rand.Int(), file.Filename)
		open, err := file.Open()

		if err != nil {
			return err
		}
		res, err := uploader.Upload(&s3manager.UploadInput{
			Bucket: aws.String(os.Getenv("S3_BUCKET")),
			ACL:    aws.String("public-read"),
			Key:    aws.String(fmt.Sprintf("images/%d%s", i, name)),
			Body:   open,
		})
		imageUrl = append(imageUrl, res.Location)
		if err != nil {
			log.Println(err)
			return err
		}
		err = open.Close()
		if err != nil {
			return ctx.SendStatus(fiber.StatusInternalServerError)
		}

	}
	return ctx.Status(fiber.StatusOK).JSON(imageUrl)
}
