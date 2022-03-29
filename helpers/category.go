package helpers

import "github.com/chirag3003/ecommerce-golang-api/models"

func ValidateCategory(data *models.Category) Errors {
	err := Errors(map[string][]string{})
	err.CheckLen(data.Title, 5, "title")
	if err.IsValid() {
		return nil
	}
	return err
}

func ValidateSubCategory(data *models.Subcategory) Errors {
	err := Errors(map[string][]string{})
	err.CheckLen(data.Title, 5, "title")
	if err.IsValid() {
		return nil
	}
	return err
}
