package repository

import (
	"context"
	"errors"

	"github.com/ebcardoso/api-clean-golang/domain/entities"
	"github.com/ebcardoso/api-clean-golang/domain/repository_interfaces"
	"github.com/ebcardoso/api-clean-golang/infrastructure/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type usersMongo struct {
	collection *mongo.Collection
	configs    *config.Config
}

func NewUsersMongoRepository(configs *config.Config) repository_interfaces.UsersRepository {
	return &usersMongo{
		collection: configs.Database.Collection("users"),
		configs:    configs,
	}
}

func (rep *usersMongo) ListUsers() ([]entities.User, error) {
	result, err := rep.collection.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, rep.configs.Exceptions.ErrUserList
	}

	defer result.Close(context.Background())

	items := []entities.User{}
	for result.Next(context.Background()) {
		var item entities.User
		if err := result.Decode(&item); err != nil {
			return nil, rep.configs.Exceptions.ErrUserFetch
		}
		items = append(items, item)
	}
	return items, nil
}

func (rep *usersMongo) CreateUser(user entities.User) (entities.User, error) {
	result, err := rep.collection.InsertOne(context.Background(), user)
	if err != nil {
		return entities.User{}, rep.configs.Exceptions.ErrUserCreate
	}
	user.ID = result.InsertedID.(primitive.ObjectID)
	return user, nil
}

func (rep *usersMongo) GetUserByID(id string) (entities.User, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return entities.User{}, rep.configs.Exceptions.ErrParseId
	}

	var result entities.User
	err = rep.collection.
		FindOne(context.Background(), bson.M{"_id": objID}).
		Decode(&result)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return entities.User{}, rep.configs.Exceptions.ErrUserNotFound
		}
		return entities.User{}, rep.configs.Exceptions.ErrUserGet
	}
	return result, nil
}

func (rep *usersMongo) GetUserByEmail(email string) (entities.User, error) {
	var result entities.User
	err := rep.collection.
		FindOne(context.Background(), bson.M{"email": email}).
		Decode(&result)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return entities.User{}, rep.configs.Exceptions.ErrUserNotFound
		}
		return entities.User{}, rep.configs.Exceptions.ErrUserGet
	}
	return result, nil
}

func (rep *usersMongo) UpdateUser(id string, user entities.User) error {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return rep.configs.Exceptions.ErrParseId
	}

	object := bson.M{}
	if user.Name != "" {
		object["name"] = user.Name
	}
	if user.TokenResetPassword != "" {
		object["tokenResetPassword"] = user.TokenResetPassword
	}
	if user.Password != "" {
		object["password"] = user.Password
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

func (rep *usersMongo) BlockUnblockUser(id string, isBlocked bool) error {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return rep.configs.Exceptions.ErrParseId
	}

	object := bson.M{}
	object["isBlocked"] = isBlocked

	result, err := rep.collection.
		UpdateOne(context.Background(), bson.M{"_id": objID}, bson.M{"$set": object})
	if err != nil {
		if isBlocked {
			return rep.configs.Exceptions.ErrUserBlock
		} else {
			return rep.configs.Exceptions.ErrUserUnblock
		}
	}
	if result.MatchedCount == 0 {
		return rep.configs.Exceptions.ErrUserNotFound
	}
	return nil
}
