package repository

import (
	"context"
	"github.com/chirag3003/ecommerce-golang-api/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ProductsRepository interface {
	Save(model *models.ProductsModel) (*mongo.InsertOneResult, error)
	FindByID(ID primitive.ObjectID) (*models.ProductsModel, error)
	Find(slug string) (*models.ProductsModel, error)
	FindAll() ([]models.ProductsModel, error)
	Delete(string) (*mongo.DeleteResult, error)
	ChangeVisibility(id string, p bool) (*mongo.UpdateResult, error)
	Update(id string, data *models.ProductsModel) (*mongo.UpdateResult, error)
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
func (c *productRepo) FindByID(ID primitive.ObjectID) (*models.ProductsModel, error) {
	find := c.db.FindOne(context.TODO(), bson.M{"_id": ID})
	data := &models.ProductsModel{}
	err := find.Decode(data)
	if err != nil {
		return nil, err
	}
	return data, nil
}
func (c *productRepo) Find(slug string) (*models.ProductsModel, error) {
	find := c.db.FindOne(context.TODO(), bson.M{"slug": slug})
	data := &models.ProductsModel{}
	err := find.Decode(data)
	if err != nil {
		return nil, err
	}
	return data, nil
}
func (c *productRepo) FindAll() ([]models.ProductsModel, error) {
	find, err := c.db.Find(context.TODO(), bson.D{{}})
	if err != nil {
		return nil, err
	}
	var data []models.ProductsModel
	err = find.All(context.TODO(), &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}
func (c *productRepo) Delete(id string) (*mongo.DeleteResult, error) {
	ID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	one, err := c.db.DeleteOne(context.TODO(), bson.M{"_id": ID})
	if err != nil {
		return nil, err
	}
	return one, nil
}
func (c *productRepo) ChangeVisibility(id string, p bool) (*mongo.UpdateResult, error) {
	ID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	byID, err := c.db.UpdateByID(context.TODO(), ID,
		bson.D{{"$set", bson.D{{"public", p}}}})
	if err != nil {
		return nil, err
	}

	return byID, nil
}
func (c *productRepo) Update(id string, data *models.ProductsModel) (*mongo.UpdateResult, error) {
	ID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	byID, err := c.db.UpdateByID(context.TODO(), ID,
		bson.D{{"$set", data}})
	if err != nil {
		return nil, err
	}

	return byID, nil
}
