package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// ProductsModel model to store product data in the database
type ProductsModel struct {
	ID          primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"` //ID: The Doc id of mongodb
	Title       string             `json:"title"`                              //Title  of the Product
	Description string             `json:"description"`                        //Description The Product Description
	Details     string             `json:"details"`                            //Details about the product
	Images      []string           `json:"images"`                             //Images of the product
	Price       float64            `json:"price"`                              //Price of the Product
	Discount    float64            `json:"discount"`                           //Discount on the product in percent
	Highlights  []string           `json:"highlights"`                         // Highlights about the product
	Tags        []string           `json:"tags"`                               //Tags The tags associated with the products which helps in better search results
	Collection  []string           `json:"collection"`                         //Collection : The collections the product is part of (eg: New Collections, Trending)
	Sizes       []ProductSize      `json:"sizes"`                              // Sizes available for the product
	Category    string             `json:"category"`                           //Category The category of the product
	Subcategory string             `json:"subcategory"`                        //Subcategory The SubCategory of the products
	DesignID    string             `json:"designID" bson:"designID"`           //DesignID of the product
	Slug        string             `json:"slug"`                               //Slug of the product which shows on the url
	Public      bool               `json:"public"`                             //Public controls the visibility of the product

}
type ProductSize struct {
	Name    string `json:"name"`
	InStock bool   `json:"inStock" bson:"inStock"`
	Stock   int    `json:"stock"`
}

func (p *ProductsModel) SetDefaults() {
	if p.Tags == nil {
		p.Tags = []string{}
	}
	if p.Collection == nil {
		p.Collection = []string{}
	}
	if p.Images == nil {
		p.Images = []string{}
	}
	if p.Highlights == nil {
		p.Highlights = []string{}
	}
	if p.Sizes == nil {
		p.Sizes = []ProductSize{
			{Name: "XXS"},
			{Name: "XS"},
			{Name: "S"},
			{Name: "M"},
			{Name: "L"},
			{Name: "XL"},
			{Name: "2XL"},
			{Name: "3XL"},
		}
	}
	p.Public = false
}

type CartSearchResult struct {
	Product  *ProductsModel `json:"product"`
	Size     string         `json:"size"`
	Stock    int            `json:"stock"`
	Quantity int            `json:"quantity"`
}
