package entities

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID        string `json:"id,omitempty"`
	Name      string `json:"name,omitempty"`
	Email     string `json:"email,omitempty"`
	IsBlocked bool   `json:"isBlocked,omitempty"`
	Password  string `json:"password,omitempty"`
}

type UserDB struct {
	ID                 primitive.ObjectID `bson:"_id,omitempty"`
	Name               string             `bson:"name,omitempty"`
	Email              string             `bson:"email,omitempty"`
	IsBlocked          bool               `bson:"isBlocked,omitempty"`
	TokenResetPassword string             `bson:"tokenResetPassword,omitempty"`
	Password           string             `bson:"password,omitempty"`
}

func (u UserDB) MapUserDB() User {
	return User{
		ID:        u.ID.Hex(),
		Name:      u.Name,
		Email:     u.Email,
		IsBlocked: u.IsBlocked,
	}
}

func CheckPassword(user UserDB, password string) bool {
	currentPassword := []byte(user.Password)
	candidate := []byte(password)

	err := bcrypt.CompareHashAndPassword(currentPassword, candidate)
	if err != nil {
		return false
	} else {
		return true
	}
}
