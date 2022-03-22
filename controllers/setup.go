package controllers

import (
	"github.com/chirag3003/ecommerce-golang-api/db"
)

var conn db.Connection

type Controllers struct {
	Products   Products
	Categories Categories
}

func NewControllers(conf db.Connection) *Controllers {
	conn = conf
	controllers := &Controllers{
		Products:   ProductsControllers(),
		Categories: CategoryControllers(),
	}
	return controllers
}
