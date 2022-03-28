package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type UserAddress struct {
	ID           primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	UserID       primitive.ObjectID `json:"userID" bson:"userID"`
	Name         string             `json:"name" bson:"name,omitempty"`
	PhoneNo      int64              `json:"phoneNo" bson:"phoneNo,omitempty"`
	AddressLine1 string             `json:"addressLine1" bson:"addressLine1,omitempty"`
	AddressLine2 string             `json:"addressLine2" bson:"addressLine2,omitempty"`
	Country      string             `json:"country" bson:"country,omitempty"`
	City         string             `json:"city" bson:"city,omitempty"`
	Zipcode      string             `json:"zipcode" bson:"zipcode,omitempty"`
}

type UserAddressInput struct {
	Name         string `bson:"name,omitempty"`
	PhoneNo      int64  `bson:"phoneNo,omitempty"`
	AddressLine1 string `bson:"addressLine1,omitempty"`
	AddressLine2 string `bson:"addressLine2,omitempty"`
	Country      string `bson:"country,omitempty"`
	City         string `bson:"city,omitempty"`
	Zipcode      string `bson:"zipcode,omitempty"`
}
