package repository

import (
	"context"
	"errors"
	"log"

	"github.com/piigyy/sharing-is-caring/internal/auth/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type authRepository struct {
	collection *mongo.Collection
}

func NewUserMongoDB(collection *mongo.Collection) *authRepository {
	return &authRepository{
		collection: collection,
	}
}

func (r *authRepository) GetUserByEmail(ctx context.Context, email string) (user model.User, err error) {
	log.Printf("authRepository.GetUserByEmail")
	err = r.collection.FindOne(ctx, bson.M{"email": email}).Decode(&user)
	if err != nil {
		log.Printf("error authRepository.GetUserByEmail: %v\n", err)
		if errors.Is(err, mongo.ErrNoDocuments) {
			err = model.ErrUserNotFound
			return
		}

		return
	}

	return
}
