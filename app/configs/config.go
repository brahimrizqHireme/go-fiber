package configs

import (
	"github.com/spf13/viper"
)

type Config struct {
	JWTSecretKey           string
	MongoDBURI             string
	MongodbDatabase        string
	SecretKeyExpiresMinute string
	ServerTimeOut          string
	AppVersion             string
	ServerUrl              string
}

var (
	AppConfig *Config
)

func LoadConfig() {
	viper.SetDefault("JWT_SECRET_KEY", "top-secret-key")
	viper.SetDefault("MONGODB_URI", "mongodb://mongodb:27017/")
	viper.SetDefault("MONGODB_DATABASE", "ci")
	viper.SetDefault("JWT_SECRET_KEY_EXPIRE_MINUTES_COUNT", 30)
	viper.SetDefault("SERVER_READ_TIMEOUT", 60)
	viper.SetDefault("APP_VERSION", "1.0.0")
	viper.SetDefault("SERVER_URL", ":3000")

	viper.AutomaticEnv()

	config := &Config{
		JWTSecretKey:           viper.GetString("JWT_SECRET_KEY"),
		MongoDBURI:             viper.GetString("MONGODB_URI"),
		MongodbDatabase:        viper.GetString("MONGODB_DATABASE"),
		ServerTimeOut:          viper.GetString("SERVER_READ_TIMEOUT"),
		AppVersion:             viper.GetString("APP_VERSION"),
		ServerUrl:              viper.GetString("SERVER_URL"),
		SecretKeyExpiresMinute: string(viper.GetString("JWT_SECRET_KEY_EXPIRE_MINUTES_COUNT")),
	}

	AppConfig = config
}
