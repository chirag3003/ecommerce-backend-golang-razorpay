package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Order struct {
	ID            primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	OrderID       string             `json:"orderID" bson:"orderID"`
	Address       UserAddress        `json:"address" bson:"address"`
	Products      []OrderProduct     `json:"products" bson:"products"`
	OrderStatus   string             `json:"orderStatus" bson:"orderStatus"`
	PaymentMethod string             `json:"paymentMethod" bson:"paymentMethod"`
	PaymentStatus string             `json:"paymentStatus" bson:"paymentStatus"`
	TransactionID string             `json:"transactionID" bson:"transactionID"`
	CreatedAt     int64              `json:"createdAt" bson:"createdAt"`
	UpdatedAt     int64              `json:"updatedAt" bson:"updatedAt"`
}

type OrderProduct struct {
	ProductID primitive.ObjectID `json:"productID" bson:"productID"`
	Quantity  primitive.ObjectID `json:"quantity" bson:"quantity"`
}

func (o *Order) SetCreatedAt() {
	o.CreatedAt = time.Now().Unix()
}

func (o *Order) SetUpdatedAt() {
	o.UpdatedAt = time.Now().Unix()
}
