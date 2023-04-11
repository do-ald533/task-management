package database

import (
	"context"
	"goapi/pkg/errors"
	"time"

	"github.com/tryvium-travels/memongo"
	"github.com/tryvium-travels/memongo/memongolog"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MongodbConnection func for connection to mongodb database.
func MemoryMongodbConnection() (*mongo.Client, error) {
	server, err := memongo.StartWithOptions(&memongo.Options{
		MongoVersion: "5.0.0",
		LogLevel:     memongolog.LogLevelSilent,
	})

	if err != nil {
		return nil, errors.InternalServerError(errors.Message{
			"msg":   "Failed to start memory mongodb server",
			"error": err.Error(),
		})
	}

	ctxTimeout, cancel := context.WithTimeout(context.Background(), time.Second*20)
	defer cancel()

	client, err := mongo.NewClient(options.Client().ApplyURI(server.URI()))
	if err != nil {
		return nil, errors.InternalServerError(errors.Message{
			"msg":   "Error when try connect to mongodb",
			"error": true,
		})

	}

	if err = client.Connect(ctxTimeout); err != nil {
		return nil, errors.InternalServerError(errors.Message{
			"msg":   "timout when connect to client",
			"error": true,
		})

	}

	// Try to ping database.
	if err := client.Ping(ctxTimeout, nil); err != nil {
		// close database connection
		return nil, errors.InternalServerError(errors.Message{
			"msg":   "Couldn't connect to client! Server timeout connection",
			"error": true,
		})
	}

	return client, nil
}
