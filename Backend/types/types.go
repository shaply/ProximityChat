// This package contains universal types that are used across the application.

package types

import (
	"context"
	"time"
)

type RegisterUserPayload struct {
	FirstName string `bson:"firstName" json:"firstName"`
	LastName  string `bson:"lastName" json:"lastName"`
	Email     string `bson:"email" json:"email"`
	Password  string `bson:"password" json:"password"`
}

type User struct {
	ID        int       `bson:"id" json:"id"`
	FirstName string    `bson:"firstName" json:"firstName"`
	LastName  string    `bson:"lastName" json:"lastName"`
	Email     string    `bson:"email" json:"email"`
	Password  string    `bson:"password" json:"password"`
	CreatedAt time.Time `bson:"createdAt" json:"createdAt"`
}

type UserStore interface {
	GetUserByEmail(ctx context.Context, email string) (*User, error)
	GetUserByID(ctx context.Context, id int) (*User, error)
	CreateUser(ctx context.Context, user *User) error
}
