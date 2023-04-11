package entities

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Book struct to describe book object.
type Book struct {
	ID         *uuid.UUID `bson:"id,omitempty" json:"id" validate:"required,uuid"`
	CreatedAt  *time.Time `bson:"created_at,omitempty" json:"created_at"`
	UpdatedAt  *time.Time `bson:"updated_at,omitempty" json:"updated_at"`
	UserID     *uuid.UUID `bson:"user_id,omitempty" json:"user_id" validate:"required,uuid"`
	Title      string     `bson:"title,omitempty" json:"title" validate:"required,lte=255"`
	Author     string     `bson:"author,omitempty" json:"author" validate:"required,lte=255"`
	BookStatus *int       `bson:"book_status,omitempty" json:"book_status" validate:"required,len=1"`
	BookAttrs  *BookAttrs `bson:"book_attrs,omitempty" json:"book_attrs" validate:"required,dive"`
}

// BookAttrs struct to describe book attributes.
type BookAttrs struct {
	Picture     string `bson:"picture,omitempty" json:"picture"`
	Description string `bson:"description,omitempty" json:"description"`
	Rating      *int   `bson:"rating,omitempty" json:"rating" validate:"min=1,max=10"`
}

// Value make the BookAttrs struct implement the driver.Valuer interface.
// This method simply returns the JSON-encoded representation of the struct.
func (b Book) Value() primitive.M {
	byte, _ := bson.Marshal(b)

	var updated bson.M
	bson.Unmarshal(byte, &updated)

	return updated
}

// Scan make the BookAttrs struct implement the sql.Scanner interface.
// This method simply decodes a JSON-encoded value into the struct fields.
func (b Book) Scan(value interface{}) error {
	j, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	return bson.Unmarshal(j, &b)
}
