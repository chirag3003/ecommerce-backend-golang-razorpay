package controllers

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/chirag3003/ecommerce-golang-api/helpers"
	"github.com/chirag3003/ecommerce-golang-api/models"
	"github.com/chirag3003/ecommerce-golang-api/repository"
	"github.com/gofiber/fiber/v2"
	gonanoid "github.com/matoous/go-nanoid/v2"
	"log"
	"strings"
)

type Images interface {
	Upload(ctx *fiber.Ctx) error
	GalleryUpload(ctx *fiber.Ctx) error
	GetGallery(ctx *fiber.Ctx) error
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
		name, err := gonanoid.New(30)
		if err != nil {
			return ctx.SendStatus(fiber.StatusInternalServerError)
		}

		uploadURL, err := helpers.UploadFile(i.awsSession, file, name)
		if err != nil {
			log.Println(err)
			return ctx.SendStatus(fiber.StatusInternalServerError)
		}
		imageUrl = append(imageUrl, uploadURL)
		_, _ = i.Images.NewImage(models.Image{
			Src: uploadURL,
			Key: fmt.Sprintf("images/%s", name),
		})

	}

	return ctx.Status(fiber.StatusOK).JSON(imageUrl)
}

func (i *imagesRoutes) GalleryUpload(ctx *fiber.Ctx) error {
	file, err := ctx.FormFile("image")
	if err != nil || file == nil || strings.TrimSpace(ctx.FormValue("name")) == "" {
		return ctx.SendStatus(fiber.StatusBadRequest)
	}
	log.Println(file.Filename)

	uploadURL, err := helpers.UploadFile(i.awsSession, file, ctx.FormValue("name"))
	if err != nil {
		return err
	}
	i.Images.NewGalleryImage(models.GalleryImage{
		Src:  uploadURL,
		Key:  fmt.Sprintf("images/%s", ctx.FormValue("name")),
		Name: ctx.FormValue("name"),
	})
	log.Println(uploadURL)
	return ctx.SendString(uploadURL)
}

func (i *imagesRoutes) GetGallery(ctx *fiber.Ctx) error {
	images, err := i.Images.GetGalleryImages()
	if err != nil {
		log.Println(err)
		return ctx.SendStatus(fiber.StatusInternalServerError)
	}
	return ctx.JSON(images)
}
