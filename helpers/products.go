package helpers

import (
	"github.com/chirag3003/ecommerce-golang-api/models"
)

func ValidateProductData(body *models.ProductsModel) Errors {
	var err = Errors(map[string][]string{})
	err.CheckLen(body.Title, 3, "title")
	err.CheckLen(body.Slug, 3, "slug")
	if err.IsValid() {
		return nil
	}
	return err
}
