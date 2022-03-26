package repository

import (
	"context"
	"github.com/chirag3003/ecommerce-golang-api/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository interface {
	Register(data *models.User) (*mongo.InsertOneResult, error)
	Login(email, password string) (bool, error)
	GetUser(email string) (*models.User, error)
	UpdateName(name string, id primitive.ObjectID) (*mongo.UpdateResult, error)
}

type userRepo struct {
	db *mongo.Collection
}

func (c userRepo) Register(data *models.User) (*mongo.InsertOneResult, error) {
	one, err := c.db.InsertOne(context.TODO(), data)
	if err != nil {
		return nil, err
	}
	return one, nil
}

func (c userRepo) Login(email, password string) (bool, error) {
	data := &models.User{}
	err := c.db.FindOne(context.TODO(), bson.M{"email": email}).Decode(data)
	if err != nil {
		return false, err
	}
	return data.CheckPass(password), nil
}

func (c userRepo) GetUser(email string) (*models.User, error) {
	data := &models.User{}
	err := c.db.FindOne(context.TODO(), bson.M{"email": email}).Decode(data)
	if err != nil {
		return nil, err
	}
	return data, nil

}

func (c userRepo) UpdateName(name string, id primitive.ObjectID) (*mongo.UpdateResult, error) {
	data, err := c.db.UpdateByID(context.TODO(), id, bson.M{"$set": bson.M{"name": name}})
	if err != nil {
		return nil, err
	}
	return data, nil
}

func NewUserRepository(col *mongo.Collection) UserRepository {
	return &userRepo{
		db: col,
	}
}
