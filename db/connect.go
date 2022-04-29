package db

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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
	return c.session.Database(os.Getenv("MONGO_DB"))
}

func ConnectMongo() Connection {

	var c conn
	var err error

	uri := os.Getenv("MONGO_URI")
	c.session, err = mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}
	fmt.Println("Database connected")
	return &c
}
