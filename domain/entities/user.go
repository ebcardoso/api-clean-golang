package entities

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID                 primitive.ObjectID `bson:"_id,omitempty"`
	Name               string             `bson:"name,omitempty"`
	Email              string             `bson:"email,omitempty"`
	IsBlocked          bool               `bson:"isBlocked,omitempty"`
	TokenResetPassword string             `bson:"tokenResetPassword,omitempty"`
	Password           string             `bson:"password,omitempty"`
}

func (user User) CheckPassword(password string) bool {
	currentPassword := []byte(user.Password)
	candidate := []byte(password)

	err := bcrypt.CompareHashAndPassword(currentPassword, candidate)
	if err != nil {
		return false
	} else {
		return true
	}
}
