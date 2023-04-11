package database

import (
	"context"
	"goapi/pkg/errors"
	"goapi/pkg/utils"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// MongodbConnection func for connection to mongodb database.
func MongodbConnection() (*mongo.Client, error) {
	url, err := utils.ConnectionURLBuilder("mongodb")

	if err != nil {
		return nil, errors.InternalServerError(errors.Message{
			"msg":     "Failed connect from provided database",
			"address": url,
		})
	}

	ctxTimeout, cancel := context.WithTimeout(context.Background(), time.Second*20)
	defer cancel()

	client, err := mongo.NewClient(options.Client().ApplyURI(url))

	if err = client.Connect(ctxTimeout); err != nil {
		return nil, errors.InternalServerError(errors.Message{
			"msg":     "timout error on connect with database",
			"address": url,
		})
	}

	// Try to ping database.
	if err := client.Ping(ctxTimeout, readpref.Primary()); err != nil {
		// close database connection
		defer client.Disconnect(ctxTimeout)
		return nil, errors.InternalServerError(errors.Message{
			"msg":   "Couldn't connect to database! Server timeout connection",
			"error": true,
		})
	}

	return client, nil
}
