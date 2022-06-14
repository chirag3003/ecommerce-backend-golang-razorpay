package repository

import (
	"context"
	"github.com/chirag3003/ecommerce-golang-api/config"
	"github.com/chirag3003/ecommerce-golang-api/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

type TransactionsRepository interface {
	NewTransaction(data *models.Transaction) (*mongo.InsertOneResult, error)
	AddEvent(transactionID string, event interface{}) *mongo.UpdateResult
	SetStatus(transactionID, status string) (*mongo.UpdateResult, error)
}

type transactionRepo struct {
	db *mongo.Collection
}

func NewTransactionRepo(col *mongo.Database) TransactionsRepository {
	return &transactionRepo{
		db: col.Collection(config.TRANSACTION_COLLECTION),
	}
}

func (t *transactionRepo) NewTransaction(data *models.Transaction) (*mongo.InsertOneResult, error) {
	res, err := t.db.InsertOne(context.TODO(), data)
	if err != nil {
		return nil, err
	}
	return res, nil

}

func (t *transactionRepo) AddEvent(transactionID string, event interface{}) *mongo.UpdateResult {
	result, err := t.db.UpdateOne(context.TODO(), bson.D{{"razorpay.razorpayID", transactionID}}, bson.D{{"$push", bson.D{{"razorpay.events", event}}}})
	if err != nil {
		log.Println(err)
		return nil
	}
	return result
}

func (t *transactionRepo) GetEvents(transactionID string) ([]interface{}, error) {
	trans := &models.Transaction{}
	err := t.db.FindOne(context.TODO(), bson.D{{"razorpay.razorpayID", transactionID}}).Decode(trans)
	if err != nil {
		return nil, err
	}
	return trans.Razorpay.Events, nil
}

func (t *transactionRepo) SetStatus(transactionID, status string) (*mongo.UpdateResult, error) {
	one, err := t.db.UpdateOne(context.TODO(), bson.D{{"razorpay.razorpayID", transactionID}}, bson.D{{"$set", bson.D{{"paymentStatus", status}}}})
	if err != nil {
		return nil, err
	}
	return one, nil
}
