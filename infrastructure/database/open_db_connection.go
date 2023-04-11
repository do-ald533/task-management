package database

import (
	"github.com/jmoiron/sqlx"
	"go.mongodb.org/mongo-driver/mongo"
)

type Databases struct {
	Sql   *sqlx.DB
	Mongo *mongo.Client
}

// OpenDBConnection func for opening database connection.
// dbType must be:
// "mongodb" | "memory_mongodb" | "mysql" | "postgres"
func OpenDBConnection(dbType string) (*Databases, error) {
	var (
		dbSql   *sqlx.DB
		dbMongo *mongo.Client
		err     error
	)

	// Define a new Database connection with right DB type.
	switch dbType {
	case "postgres":
		dbSql, err = PostgreSQLConnection()
	case "mysql":
		dbSql, err = MysqlConnection()
	case "mongodb":
		dbMongo, err = MongodbConnection()
	case "memory_mongodb":
		dbMongo, err = MemoryMongodbConnection()
	}

	if err != nil {
		return nil, err
	}

	return &Databases{
		Sql:   dbSql,
		Mongo: dbMongo,
	}, nil
}
