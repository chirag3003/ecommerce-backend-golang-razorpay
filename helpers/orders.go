package helpers

import (
	"github.com/chirag3003/ecommerce-golang-api/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func ValidateNewOrderInput(input *models.NewOrderInput) bool {
	if input.Address == primitive.NilObjectID {
		return false
	}
	for _, prod := range input.Products {
		if prod.Quantity < 1 {
			return false
		}

	}
	return true
}
