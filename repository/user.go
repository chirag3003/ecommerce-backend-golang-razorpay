package repository

import (
	"context"
	"github.com/chirag3003/ecommerce-golang-api/config"
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
	AddAddress(ID primitive.ObjectID, data *models.UserAddress) (*mongo.InsertOneResult, error)
	GetAddresses(id primitive.ObjectID) (*[]models.UserAddress, error)
}

func NewUserRepository(col *mongo.Database) UserRepository {
	return &userRepo{
		User:    col.Collection(config.USER_COLLECTION),
		Address: col.Collection(config.ADDRESS_COLLECTION),
	}
}

type userRepo struct {
	User    *mongo.Collection
	Address *mongo.Collection
}

func (c userRepo) Register(data *models.User) (*mongo.InsertOneResult, error) {
	one, err := c.User.InsertOne(context.TODO(), data)
	if err != nil {
		return nil, err
	}
	return one, nil
}

func (c userRepo) Login(email, password string) (bool, error) {
	data := &models.User{}
	err := c.User.FindOne(context.TODO(), bson.M{"email": email}).Decode(data)
	if err != nil {
		return false, err
	}
	return data.CheckPass(password), nil
}

func (c userRepo) GetUser(email string) (*models.User, error) {
	data := &models.User{}
	err := c.User.FindOne(context.TODO(), bson.M{"email": email}).Decode(data)
	if err != nil {
		return nil, err
	}
	return data, nil

}

func (c userRepo) UpdateName(name string, id primitive.ObjectID) (*mongo.UpdateResult, error) {
	data, err := c.User.UpdateByID(context.TODO(), id, bson.M{"$set": bson.M{"name": name}})
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (c userRepo) AddAddress(ID primitive.ObjectID, data *models.UserAddress) (*mongo.InsertOneResult, error) {
	data.UserID = ID
	result, err := c.Address.InsertOne(context.TODO(), data)
	if err != nil {
		return nil, err
	}

	objectID := result.InsertedID.(primitive.ObjectID)
	_, err = c.User.UpdateByID(context.TODO(), ID, bson.M{"$push": bson.M{"address": objectID}})
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (c userRepo) GetAddresses(id primitive.ObjectID) (*[]models.UserAddress, error) {
	find, err := c.Address.Find(context.TODO(), bson.M{"userid": id})
	if err != nil {
		return nil, err
	}
	data := &[]models.UserAddress{}
	err = find.All(context.TODO(), data)
	if err != nil {
		return nil, err
	}
	return data, nil
}
