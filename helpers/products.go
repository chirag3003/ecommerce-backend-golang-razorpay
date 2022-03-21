package helpers

import (
	"github.com/chirag3003/ecommerce-golang-api/models"
)

func ValidateProductData(body *models.ProductsModel) Errors {
	var err = Errors(map[string][]string{})
	if !err.CheckLen(body.Title, 3) {
		err.Add("title", "min Length required is 3")
	}
	if !err.CheckLen(body.Slug, 3) {
		err.Add("slug", "min Length required is 3")
	}
	if !err.CheckMinValue(float64(body.Stock), 0) {
		err.Add("stock", "min stock required is 0")
	}
	if err.IsValid() {
		return nil
	}
	return err
}
