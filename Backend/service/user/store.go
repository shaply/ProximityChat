package user

import (
	"context"

	"github.com/shaply/ProximityChat/Backend/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Store struct {
	db *mongo.Database
}

func NewStore(db *mongo.Database) *Store {
	return &Store{db: db}
}

func (s *Store) GetUserByEmail(ctx context.Context, email string) (*types.User, error) {
	var user types.User
	err := s.db.Collection("Users").FindOne(ctx, bson.M{"email": email}).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func (s *Store) GetUserByID(ctx context.Context, id int) (*types.User, error) {
	var user types.User
	err := s.db.Collection("Users").FindOne(ctx, bson.M{"id": id}).Decode(user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func (s *Store) CreateUser(ctx context.Context, user *types.User) error {
	_, err := s.db.Collection("Users").InsertOne(ctx, user)
	if err != nil {
		return err
	}
	return nil
}
