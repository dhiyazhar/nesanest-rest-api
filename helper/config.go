package helper

import (
	"os"

	_ "github.com/joho/godotenv/autoload"
)

var AppConfig *Config

type Config struct {
	DB         PostgresConfig
	AppBaseUrl string
}

type UploadBaseDirConfig struct {
	UploadBasePath string
}

type PostgresConfig struct {
	PostgresHost     string
	PostgresPort     string
	PostgresUsername string
	PostgresPassword string
	PostgresDBName   string
}

func (u UploadBaseDirConfig) GetUploadBasePath() string {
	return u.UploadBasePath
}

func (p PostgresConfig) GetHost() string {
	return p.PostgresHost
}

func (p PostgresConfig) GetPort() string {

	return p.PostgresPort
}

func (p PostgresConfig) GetUsername() string {
	return p.PostgresUsername
}

func (p PostgresConfig) GetPassword() string {
	return p.PostgresPassword
}

func (p PostgresConfig) GetDBName() string {
	return p.PostgresDBName
}

func init() {
	cfg := &Config{
		DB: PostgresConfig{
			PostgresHost:     os.Getenv("DBHOST"),
			PostgresPort:     os.Getenv("DBPORT"),
			PostgresUsername: os.Getenv("DBUSERNAME"),
			PostgresPassword: os.Getenv("DBPASSWORD"),
			PostgresDBName:   os.Getenv("DBNAME"),
		},
		AppBaseUrl: os.Getenv("APP_BASE_URL"),
	}

	AppConfig = cfg
}
