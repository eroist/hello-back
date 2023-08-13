package config

import (
	"fmt"
	"os"
	"reflect"
	"time"

	"github.com/joho/godotenv"
)

type EnvVar struct {
	// App env
	AppPort string
	AppEnv  string
	// token env
	TokenSymmetricKey    string
	AccessTokenDuration  time.Duration
	RefreshTokenDuration time.Duration
	DBHost               string
	DBName               string
	DBUser               string
	DBPassword           string
	DBPort               string
	// Filebase credential
	FilebaseBucket     string
	FilebaseAccessKey  string
	FilebaseSecretKey  string
	FilebasePinningKey string
	// Github OAuth credential
	GithubClientID     string
	GithubClientSecret string
}

var env EnvVar

func LoadEnv() (err error) {
	// skip load env when docker
	if os.Getenv("APP_PORT") == "" {
		err = godotenv.Load(".env")
		if err != nil {
			return err
		}
	}

	atd, err := time.ParseDuration(os.Getenv("ACCESS_TOKEN_DURATION"))
	if err != nil {
		return err
	}

	rtd, err := time.ParseDuration(os.Getenv("REFRESH_TOKEN_DURATION"))
	if err != nil {
		return err
	}

	env = EnvVar{
		// App env
		AppPort: os.Getenv("APP_PORT"),
		AppEnv:  os.Getenv("APP_ENV"),
		// token env
		TokenSymmetricKey:    os.Getenv("TOKEN_SYMMETRIC_KEY"),
		AccessTokenDuration:  atd,
		RefreshTokenDuration: rtd,
		// Postgres
		DBHost:     os.Getenv("POSTGRES_HOST"),
		DBName:     os.Getenv("POSTGRES_DB"),
		DBUser:     os.Getenv("POSTGRES_USER"),
		DBPassword: os.Getenv("POSTGRES_PASSWORD"),
		DBPort:     os.Getenv("POSTGRES_PORT"),
		// Filebase credential
		FilebaseBucket:     os.Getenv("FILEBASE_BUCKET"),
		FilebaseAccessKey:  os.Getenv("FILEBASE_ACCESS_KEY"),
		FilebaseSecretKey:  os.Getenv("FILEBASE_SECRET_KEY"),
		FilebasePinningKey: os.Getenv("FILEBASE_PINNING_KEY"),
		// Github OAuth credentail
		GithubClientID:     os.Getenv("GITHUB_CLIENT_ID"),
		GithubClientSecret: os.Getenv("GITHUB_CLIENT_SECRET"),
	}

	values := reflect.ValueOf(env)
	types := values.Type()
	for i := 0; i < values.NumField(); i++ {
		if values.Field(i).String() == "" {
			return fmt.Errorf("config: %s is missing", types.Field(i).Name)
		}
	}

	if err != nil {
		return
	}

	return
}

func Env() EnvVar {
	return env
}
