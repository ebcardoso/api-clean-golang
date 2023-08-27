package config

import (
	"github.com/ebcardoso/api-clean-golang/infrastructure/database"
	"github.com/ebcardoso/api-clean-golang/infrastructure/exceptions"
	"github.com/ebcardoso/api-clean-golang/infrastructure/translations"
	"go.mongodb.org/mongo-driver/mongo"
)

type Config struct {
	Env          *Env
	Database     *mongo.Database
	Translations *translations.Translations
	Exceptions   *exceptions.Exceptions
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

	//Loading Exceptions
	exceptions := exceptions.NewExceptions(translations)

	c := &Config{
		Env:          env,
		Database:     db,
		Translations: translations,
		Exceptions:   exceptions,
	}
	return c, nil
}
