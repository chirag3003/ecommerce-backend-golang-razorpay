package repository

import (
	"context"
	"github.com/chirag3003/ecommerce-golang-api/config"
	"github.com/chirag3003/ecommerce-golang-api/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type ImagesRepository interface {
	NewImage(data models.Image) (*mongo.InsertOneResult, error)
	NewGalleryImage(data models.GalleryImage) (*mongo.InsertOneResult, error)
	GetGalleryImages() (*[]*models.GalleryImage, error)
	GetGalleryImage(name string) (*models.GalleryImage, error)
}

func NewImagesRepo(db *mongo.Database) ImagesRepository {

	return &imageRepo{
		db.Collection(config.IMAGES_COLLECTION),
		db.Collection(config.GALLERY_COLLECTION),
	}
}

type imageRepo struct {
	db      *mongo.Collection
	gallery *mongo.Collection
}

func (i *imageRepo) NewImage(data models.Image) (*mongo.InsertOneResult, error) {
	result, err := i.db.InsertOne(context.TODO(), data)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (i *imageRepo) NewGalleryImage(data models.GalleryImage) (*mongo.InsertOneResult, error) {
	result, err := i.gallery.InsertOne(context.TODO(), data)
	if err != nil {
		return nil, err
	}
	return result, err
}

func (i *imageRepo) GetGalleryImages() (*[]*models.GalleryImage, error) {
	find, err := i.gallery.Find(context.TODO(), bson.D{})
	if err != nil {
		return nil, err
	}
	var data = &[]*models.GalleryImage{}
	err = find.All(context.TODO(), data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (i *imageRepo) GetGalleryImage(name string) (*models.GalleryImage, error) {
	data := &models.GalleryImage{}
	err := i.gallery.FindOne(context.TODO(), bson.D{{"name", name}}).Decode(data)
	if err != nil {
		return nil, err
	}
	if err != nil {
		return nil, err
	}
	return data, nil
}
