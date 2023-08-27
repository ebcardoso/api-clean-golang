package repositories

import (
	"context"
	"errors"

	"github.com/ebcardoso/api-clean-golang/domain/entities"
	"github.com/ebcardoso/api-clean-golang/domain/interfaces"
	"github.com/ebcardoso/api-clean-golang/infrastructure/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type usersMongo struct {
	collection *mongo.Collection
	configs    *config.Config
}

func NewUsersMongoRepository(configs *config.Config) interfaces.UsersRepository {
	return &usersMongo{
		collection: configs.Database.Collection("users"),
		configs:    configs,
	}
}

func (rep *usersMongo) CreateUser(input entities.UserDB) (entities.User, error) {
	result, err := rep.collection.InsertOne(context.Background(), input)
	if err != nil {
		return entities.User{}, rep.configs.Exceptions.ErrUserCreate
	}
	input.ID = result.InsertedID.(primitive.ObjectID)
	return entities.MapUserDB(input), nil
}

func (rep *usersMongo) GetUserByEmail(email string) (entities.UserDB, error) {
	var result entities.UserDB
	err := rep.collection.
		FindOne(context.Background(), bson.M{"email": email}).
		Decode(&result)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return entities.UserDB{}, rep.configs.Exceptions.ErrUserNotFound
		}
		return entities.UserDB{}, rep.configs.Exceptions.ErrUserGet
	}
	return result, nil
}
