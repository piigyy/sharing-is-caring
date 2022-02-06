package repository

import (
	"context"
	"errors"
	"log"

	"github.com/piigyy/sharing-is-caring/internal/auth/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type authRepository struct {
	collection *mongo.Collection
}

func NewUserMongoDB(collection *mongo.Collection) *authRepository {
	options := []mongo.IndexModel{
		{
			Keys:    bson.D{{Key: "email", Value: 1}},
			Options: options.Index().SetUnique(true),
		},
		{
			Keys:    bson.D{{Key: "phone", Value: 1}},
			Options: options.Index().SetUnique(true),
		},
	}

	collection.Indexes().CreateMany(context.Background(), options)
	return &authRepository{
		collection: collection,
	}
}

func (r *authRepository) GetUserByEmail(ctx context.Context, email string) (user model.User, err error) {
	log.Printf("authRepository.GetUserByEmail")
	err = r.collection.FindOne(ctx, bson.M{"email": email}).Decode(&user)
	if err != nil {
		err = errors.New("internal server error")
		log.Printf("error authRepository.GetUserByEmail: %v\n", err)
		if errors.Is(err, mongo.ErrNoDocuments) {
			err = model.ErrUserNotFound
			return
		}

		return
	}

	return
}

func (r *authRepository) CreateUser(ctx context.Context, entity model.User) (userID string, err error) {
	var (
		result   *mongo.InsertOneResult
		oid      primitive.ObjectID
		oidValid bool
	)

	log.Println("inserting to mongodb")
	result, err = r.collection.InsertOne(ctx, &entity)
	if err != nil {
		log.Printf("insert user failed: %v\n", err)
		return
	}

	log.Println("converting inserted user id")
	oid, oidValid = result.InsertedID.(primitive.ObjectID)
	if !oidValid {
		log.Printf("user object id invalid\n")
		err = errors.New("invalid object id")
		return
	}

	return oid.Hex(), nil
}

func (r *authRepository) DuplicateError(ctx context.Context, err error) bool {
	return mongo.IsDuplicateKeyError(err)
}

func (r *authRepository) UpdatePassword(ctx context.Context, email, password string) (err error) {
	var result *mongo.UpdateResult

	filter := bson.D{{"email", email}}
	update := bson.D{{"$set", bson.D{{"password", password}}}}

	result, err = r.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			err = model.ErrUnAuthorized
			return
		}
	}

	if result.MatchedCount < 1 {
		log.Printf("not documents updated")
		err = model.ErrUnAuthorized
		return
	}

	return nil
}
