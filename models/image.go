package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Image struct {
	ID  primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Src string             `json:"src"`
	Key string             `json:"key"`
}
