package repositories

import (
	"goapi/infrastructure/database"
	"goapi/internal/modules/books/entities"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/mongo"
)

type BookRepositoryImpl interface {
	Index(limit, page int64) ([]entities.Book, int64, error)
	Show(id uuid.UUID) (*entities.Book, error)
	Store(b *entities.Book) (*entities.Book, error)
	Update(id *uuid.UUID, book *entities.Book) (*entities.Book, error)
	Delete(id uuid.UUID) error
}

// BookRepository struct for queries from Book model.
type BookRepository struct {
	BooksCollection *mongo.Collection
}

func NewBookRepository() (BookRepositoryImpl, error) {
	db, err := database.OpenDBConnection("mongodb")
	if err != nil {
		// Return status 500 and database connection error.
		return nil, err
	}

	database := db.Mongo.Database("app")
	booksCollection := database.Collection("books")

	return &BookRepository{booksCollection}, nil
}
