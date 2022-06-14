package repository

import (
	"context"
	"github.com/chirag3003/ecommerce-golang-api/config"
	"github.com/chirag3003/ecommerce-golang-api/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type OrderRepository interface {
	Save(data *models.Order) error
	SetPaid(orderID string) (*mongo.UpdateResult, error)
	SetStatus(orderID, status string) (*mongo.UpdateResult, error)
}

func NewOrderRepository(col *mongo.Database) OrderRepository {
	return &orderRepo{
		Order: col.Collection(config.ORDER_COLLECTION),
	}
}

type orderRepo struct {
	Order *mongo.Collection
}

func (o *orderRepo) Save(data *models.Order) error {
	_, err := o.Order.InsertOne(context.TODO(), data)
	if err != nil {
		return err
	}
	return nil
}

func (o *orderRepo) SetPaid(orderID string) (*mongo.UpdateResult, error) {
	one, err := o.Order.UpdateOne(context.TODO(), bson.D{{"orderID", orderID}}, bson.D{{"$set", bson.D{{"paymentStatus", "paid"}}}})
	if err != nil {
		return nil, err
	}
	return one, nil
}

func (o *orderRepo) SetStatus(orderID, status string) (*mongo.UpdateResult, error) {
	one, err := o.Order.UpdateOne(context.TODO(), bson.D{{"orderID", orderID}}, bson.D{{"$set", bson.D{{"orderStatus", status}}}})
	if err != nil {
		return nil, err
	}
	return one, nil
}
