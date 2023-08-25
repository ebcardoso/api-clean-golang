package config

import (
	"github.com/ebcardoso/api-clean-golang/infrastructure/database"
	"github.com/ebcardoso/api-clean-golang/infrastructure/translations"
	"go.mongodb.org/mongo-driver/mongo"
)

type Config struct {
	Env          *Env
	Database     *mongo.Database
	Translations *translations.Translations
}

func SetConfigs(file string) (*Config, error) {
	//Loading Env Vars
	env, err := LoadEnvs(file)
	if err != nil {
		return &Config{}, err
	}

	//Loading Database
	db, err := database.LoadMongoDB(env.MONGO_URI, env.MONGO_DATABASE)
	if err != nil {
		return &Config{}, err
	}

	//Loading Translations
	translations, err := translations.Load(env.DEFAULT_TRANSLATION)
	if err != nil {
		return &Config{}, err
	}

	c := &Config{
		Env:          env,
		Database:     db,
		Translations: translations,
	}
	return c, nil
}
