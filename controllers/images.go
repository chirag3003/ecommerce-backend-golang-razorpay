package controllers

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/chirag3003/ecommerce-golang-api/helpers"
	"github.com/chirag3003/ecommerce-golang-api/models"
	"github.com/chirag3003/ecommerce-golang-api/repository"
	"github.com/gofiber/fiber/v2"
	gonanoid "github.com/matoous/go-nanoid/v2"
	"log"
	"os"
	"path/filepath"
	"strings"
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

func (i *imagesRoutes) Upload(ctx *fiber.Ctx) error {

	uploader := s3manager.NewUploader(i.awsSession)

	form, err := ctx.MultipartForm()
	if err != nil {

		return err
	}
	files := form.File["images"]
	var imageUrl []string
	for _, file := range files {
		if !strings.HasPrefix(file.Header["Content-Type"][0], "image/") {
			return ctx.Status(fiber.StatusBadRequest).JSON("file type not supported")
		}
		id, err := gonanoid.New(30)
		if err != nil {
			return ctx.SendStatus(fiber.StatusInternalServerError)
		}
		name := fmt.Sprintf("%s%s", id, filepath.Ext(file.Filename))
		open, err := file.Open()
		image, err := helpers.OptimizeImage(&open)
		if err != nil {
			return ctx.SendStatus(fiber.StatusInternalServerError)
		}

		if err != nil {
			return err
		}
		res, err := uploader.Upload(&s3manager.UploadInput{
			Bucket: aws.String(os.Getenv("S3_BUCKET")),
			ACL:    aws.String("public-read"),
			Key:    aws.String(fmt.Sprintf("images/%s", name)),
			Body:   image,
		})
		imageUrl = append(imageUrl, res.Location)
		_, _ = i.Images.NewImage(models.Image{
			Src: res.Location,
			Key: fmt.Sprintf("images/%s", name),
		})

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
