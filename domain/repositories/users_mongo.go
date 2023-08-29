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

func (rep *usersMongo) ListUsers() ([]entities.User, error) {
	items := []entities.User{}

	result, err := rep.collection.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, rep.configs.Exceptions.ErrUserList
	}

	defer result.Close(context.Background())

	for result.Next(context.Background()) {
		var item entities.UserDB
		if err := result.Decode(&item); err != nil {
			return nil, rep.configs.Exceptions.ErrUserFetch
		}
		items = append(items, item.MapUserDB())
	}
	return items, nil
}

func (rep *usersMongo) CreateUser(input entities.UserDB) (entities.User, error) {
	result, err := rep.collection.InsertOne(context.Background(), input)
	if err != nil {
		return entities.User{}, rep.configs.Exceptions.ErrUserCreate
	}
	input.ID = result.InsertedID.(primitive.ObjectID)
	return input.MapUserDB(), nil
}

func (rep *usersMongo) GetUserByID(id string) (entities.UserDB, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return entities.UserDB{}, rep.configs.Exceptions.ErrParseId
	}

	var result entities.UserDB
	err = rep.collection.
		FindOne(context.Background(), bson.M{"_id": objID}).
		Decode(&result)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return entities.UserDB{}, rep.configs.Exceptions.ErrUserNotFound
		}
		return entities.UserDB{}, rep.configs.Exceptions.ErrUserGet
	}
	return result, nil
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

func (rep *usersMongo) UpdateUser(id string, input entities.UserDB) error {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return rep.configs.Exceptions.ErrParseId
	}

	object := bson.M{}
	if input.Name != "" {
		object["name"] = input.Name
	}
	if input.TokenResetPassword != "" {
		object["tokenResetPassword"] = input.TokenResetPassword
	}
	if input.Password != "" {
		object["password"] = input.Password
	}
	result, err := rep.collection.
		UpdateOne(context.Background(), bson.M{"_id": objID}, bson.M{"$set": object})
	if err != nil {
		return rep.configs.Exceptions.ErrUserUpdate
	}
	if result.MatchedCount == 0 {
		return rep.configs.Exceptions.ErrUserNotFound
	}
	return nil
}

func (rep *usersMongo) DestroyUser(id string) error {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return rep.configs.Exceptions.ErrParseId
	}

	result, err := rep.collection.DeleteOne(context.Background(), bson.M{"_id": objID})
	if err != nil {
		return rep.configs.Exceptions.ErrUserDestroy
	}
	if result.DeletedCount == 0 {
		return rep.configs.Exceptions.ErrUserNotFound
	}
	return nil
}
