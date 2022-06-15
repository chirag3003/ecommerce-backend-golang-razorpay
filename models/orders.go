package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Order struct {
	ID            primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	OrderID       string             `json:"orderID" bson:"orderID"`
	UserID        primitive.ObjectID `json:"userID" bson:"userID"`
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
	Product  ProductsModel `json:"productID" bson:"productID"`
	Size     string        `json:"size" bson:"size"`
	Quantity int           `json:"quantity" bson:"quantity"`
}

type NewOrderInput struct {
	Products []NewOrderProductInput `json:"products"`
	Address  primitive.ObjectID
}

type NewOrderProductInput struct {
	Product  primitive.ObjectID `json:"product"`
	Size     string             `json:"size"`
	Quantity int                `json:"quantity"`
}

type NewOrderResponse struct {
	OrderID       string `json:"orderID" bson:"orderID"`
	TransactionID string `json:"transactionID" bson:"transactionID"`
}

func (o *Order) SetCreatedAt() {
	o.CreatedAt = time.Now().Unix()
	o.UpdatedAt = o.CreatedAt
}

func (o *Order) SetUpdatedAt() {
	o.UpdatedAt = time.Now().Unix()
}
