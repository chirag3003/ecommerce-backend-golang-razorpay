package middlewares

import (
	"github.com/chirag3003/ecommerce-golang-api/db"
)

var conn db.Connection

func SetupMiddlewares(conf db.Connection) {
	conn = conf
}
