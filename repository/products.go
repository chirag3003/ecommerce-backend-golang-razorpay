package repository

import (
	"context"
	"github.com/chirag3003/ecommerce-golang-api/models"
	"go.mongodb.org/mongo-driver/mongo"
)

type ProductsRepository interface {
	Save(model *models.ProductsModel) (*mongo.InsertOneResult, error)
	Find() error
	FindAll() error
	Delete() error
}

type productRepo struct {
	db *mongo.Collection
}

func NewProductsRepository(Products *mongo.Collection) ProductsRepository {
	return &productRepo{
		db: Products,
	}
}

func (c *productRepo) Save(data *models.ProductsModel) (*mongo.InsertOneResult, error) {
	one, err := c.db.InsertOne(context.TODO(), data)
	if err != nil {
		return nil, err
	}

	return one, nil
}

func (c *productRepo) Find() error    { return nil }
func (c *productRepo) FindAll() error { return nil }
func (c *productRepo) Delete() error  { return nil }
