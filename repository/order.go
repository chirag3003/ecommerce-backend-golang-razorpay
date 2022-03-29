package repository

import (
	"context"
	"github.com/chirag3003/ecommerce-golang-api/config"
	"github.com/chirag3003/ecommerce-golang-api/models"
	"go.mongodb.org/mongo-driver/mongo"
)

type OrderRepository interface {
	Save(data *models.Order) error
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
