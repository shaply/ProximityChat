// This package contains universal types that are used across the application.

package types

import (
	"context"
	"time"

	"github.com/gorilla/websocket"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type RegisterUserPayload struct {
	FirstName string `bson:"firstName" json:"firstName" validate:"required"`
	LastName  string `bson:"lastName" json:"lastName" validate:"required"`
	Email     string `bson:"email" json:"email" validate:"required,email"`
	Password  string `bson:"password" json:"password" validate:"required,min=3"`
}

type LoginUserPayload struct {
	Email    string `bson:"email" json:"email" validate:"required,email"`
	Password string `bson:"password" json:"password" validate:"required"`
}

type User struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	FirstName string             `bson:"firstName" json:"firstName"`
	LastName  string             `bson:"lastName" json:"lastName"`
	Email     string             `bson:"email" json:"email"`
	Password  string             `bson:"password" json:"password"`
	CreatedAt time.Time          `bson:"createdAt" json:"createdAt"`
}

type Client struct {
	Email string
	Conn  *websocket.Conn
}

type Message struct {
	Email   string `json:"email"`
	Message string `json:"message"`
}

type UserStore interface {
	GetUserByEmail(ctx context.Context, email string) (*User, error)
	GetUserByID(ctx context.Context, id primitive.ObjectID) (*User, error)
	CreateUser(ctx context.Context, user *User) error
}
