package repository

import (
	"github.com/chirag3003/ecommerce-golang-api/config"
	"go.mongodb.org/mongo-driver/mongo"
)

type ImagesRepository interface {
}

func NewImagesRepo(db *mongo.Database) ImagesRepository {

	return &imageRepo{
		db.Collection(config.IMAGES_COLLECTION),
	}
}

type imageRepo struct {
	db *mongo.Collection
}
