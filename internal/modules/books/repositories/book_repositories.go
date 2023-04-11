package repositories

import (
	"context"
	"fmt"
	"goapi/internal/modules/books/entities"
	"goapi/pkg/errors"
	"goapi/pkg/pagination"
	"time"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Index method for getting all books.
func (repo *BookRepository) Index(limit, page int64) ([]entities.Book, int64, error) {
	ctxTimeout, cancel := context.WithTimeout(context.Background(), time.Second*20)
	defer cancel()

	// Define books variable.
	books := []entities.Book{}

	// Length of books document.
	count, err := repo.BooksCollection.CountDocuments(ctxTimeout, bson.D{})

	// Error when cant count items into document
	if err != nil {
		return nil, 0, err
	}

	// Generate pagination to mongodb find method
	paginate := pagination.NewMongoPaginate(limit, page, count)

	// Send query to database.
	cur, err := repo.BooksCollection.Find(ctxTimeout, bson.D{}, paginate.Options())

	// Error when cant find any books into database
	if err != nil {
		// TODO tratar esse erro
		return nil, 0, err
	}

	// Error when marshal to books struct
	if err := cur.All(ctxTimeout, &books); err != nil {
		// TODO tratar esse erro
		return nil, 0, err
	}

	// Return query result.
	return books, count, nil
}

// Show method for getting one book by given ID.
func (repo *BookRepository) Show(id uuid.UUID) (*entities.Book, error) {
	ctxTimeout, cancel := context.WithTimeout(context.Background(), time.Second*20)
	defer cancel()

	// Define book variable.
	book := entities.Book{}

	err := repo.BooksCollection.FindOne(ctxTimeout, bson.D{{Key: "id", Value: id}}).Decode(&book)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			// This error means your query did not match any documents.
			return nil, errors.NotFound(errors.Message{
				"error": true,
				"msg":   fmt.Sprintf("book with the given ID: %s is not found", id),
			})

		}

		return nil, err
	}

	// Return query result.
	return &book, nil
}

func (repo *BookRepository) Store(b *entities.Book) (*entities.Book, error) {
	ctxTimeout, cancel := context.WithTimeout(context.Background(), time.Second*20)
	defer cancel()

	// Convert struct to bson structure.
	value := b.Value()

	// Insert book to database.
	_, err := repo.BooksCollection.InsertOne(ctxTimeout, value)

	// If dont create the book, return 400
	if err != nil {
		// Return empty object and error.
		return nil, errors.BadRequest(errors.Message{
			"error": true,
			"msg":   "Can't create this book!",
		})
	}

	// Return query result.
	return b, nil
}

// Update method for updating book by given Book object.
func (repo *BookRepository) Update(id *uuid.UUID, book *entities.Book) (*entities.Book, error) {
	ctxTimeout, cancel := context.WithTimeout(context.Background(), time.Second*20)
	defer cancel()

	newBook := entities.Book{}

	filter := bson.D{{Key: "id", Value: id}}
	update := bson.M{"$set": (*book).Value()}

	err := repo.BooksCollection.
		FindOneAndUpdate(ctxTimeout,
			filter,
			update,
			options.FindOneAndUpdate().SetReturnDocument(options.After), // options for find and decode after update book
		).Decode(&newBook)

	if err != nil {
		return nil, errors.NotFound(errors.Message{"msg": err.Error()})
	}

	return &newBook, nil
}

// Delete method for delete book by given ID.
func (repo *BookRepository) Delete(id uuid.UUID) error {
	ctxTimeout, cancel := context.WithTimeout(context.Background(), time.Second*20)
	defer cancel()

	filter := bson.D{{Key: "id", Value: id}}
	_, err := repo.BooksCollection.DeleteOne(ctxTimeout, filter)

	if err != nil {
		return errors.NotFound(errors.Message{"msg": err.Error()})
	}

	return nil
}
