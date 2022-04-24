package repository

import (
	"context"
	"github.com/chirag3003/ecommerce-golang-api/config"
	"github.com/chirag3003/ecommerce-golang-api/models"
	"go.mongodb.org/mongo-driver/mongo"
)

type ImagesRepository interface {
	NewImage(data models.Image) (*mongo.InsertOneResult, error)
}

func NewImagesRepo(db *mongo.Database) ImagesRepository {

	return &imageRepo{
		db.Collection(config.IMAGES_COLLECTION),
	}
}

type imageRepo struct {
	db *mongo.Collection
}

func (i *imageRepo) NewImage(data models.Image) (*mongo.InsertOneResult, error) {
	result, err := i.db.InsertOne(context.TODO(), data)
	if err != nil {
		return nil, err
	}
	return result, nil
}
