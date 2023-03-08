package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Category model to store product data in the database
type Category struct {
	ID            primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"` //ID: The Doc id of mongodb
	Title         string             `json:"title"`                              //Title  of the category
	Description   string             `json:"description"`                        //Description The category
	Tags          []string           `json:"tags"`                               //Tags The tags associated with the category which helps in better search results
	Subcategories []Subcategory      `json:"subcategories"`                      //Subcategories in the category
	Public        bool               `json:"public"`                             //Public controls the visibility of the category
}

func (d *Category) SetDefaults() {
	if d.Tags == nil {
		d.Tags = []string{}
	}
}

type Subcategory struct {
	ID          primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"` //ID: The Doc id of mongodb
	Title       string             `json:"title"`                              //Title  of the subcategory
	Description string             `json:"description"`                        //Description The subcategory
	Tags        []string           `json:"tags"`                               //Tags The tags associated with the subcategory which helps in better search results
}

func (d *Subcategory) SetDefaults() {
	if d.Tags == nil {
		d.Tags = []string{}
	}
}

type CategoryUpdateInput struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	Title       string             `bson:"title,omitempty"`       //Title  of the category
	Description string             `bson:"description,omitempty"` //Description The category
	Tags        []string           `bson:"tags,omitempty"`
}
