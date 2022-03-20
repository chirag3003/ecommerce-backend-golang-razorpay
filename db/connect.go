package db

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
)

type Connection interface {
	Close()
	DB() *mongo.Database
}

type conn struct {
	session *mongo.Client
}

func (c *conn) Close() {
	if err := c.session.Disconnect(context.TODO()); err != nil {
		panic(err)
	}
}

func (c *conn) DB() *mongo.Database {
	return c.session.Database("Ecommerce")
}

func ConnectMongo() Connection {

	var c conn
	var err error

	uri := os.Getenv("MONGO_URI")
	if uri == "" {
		log.Fatal("You must set your 'MONGODB_URI' environmental variable. See\n\t https://docs.mongodb.com/drivers/go/current/usage-examples/#environment-variable")
	}
	c.session, err = mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}
	return &c
}
