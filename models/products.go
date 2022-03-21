package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// ProductsModel model to store product data in the database
type ProductsModel struct {
	ID          primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"` //ID: The Doc id of mongodb
	Title       string             `json:"title"`                              //Title  of the Product
	Description string             `json:"description"`                        //Description The Product Description
	Price       float64            `json:"price"`                              //Price of the Product
	Discount    float64            `json:"discount"`                           //Discount on the product in percent
	Tags        []string           `json:"tags"`                               //Tags The tags associated with the products which helps in better search results
	Category    string             `json:"category"`                           //Category The category of the product
	Subcategory string             `json:"subcategory"`                        //Subcategory The SubCategory of the products
	Slug        string             `json:"slug"`                               //Slug of the product which shows on the url
	Stock       int                `json:"stock"`                              //Stock of the product
	Public      bool               `json:"public"`                             //Public controls the visibility of the product
}

func (p *ProductsModel) SetDefaults() {
	if p.Tags == nil {
		p.Tags = []string{}
	}
	p.Public = false
}
