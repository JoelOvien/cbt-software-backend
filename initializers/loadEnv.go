package initializers

import (
	"time"

	"github.com/spf13/viper"
)

// Config is a struct and uses the mapstructure tags to list the environment variables we want Viper to load.
type Config struct {
	DBHost         string `mapstructure:"POSTGRES_HOST"`
	DBUserName     string `mapstructure:"POSTGRES_USER"`
	DBUserPassword string `mapstructure:"POSTGRES_PASSWORD"`
	DBName         string `mapstructure:"POSTGRES_DB"`
	DBPort         string `mapstructure:"POSTGRES_PORT"`
	ServerPort     string `mapstructure:"PORT"`

	ClientOrigin string `mapstructure:"CLIENT_ORIGIN"`

	AccessTokenPrivateKey  string        `mapstructure:"ACCESS_TOKEN_PRIVATE_KEY"`
	AccessTokenPublicKey   string        `mapstructure:"ACCESS_TOKEN_PUBLIC_KEY"`
	RefreshTokenPrivateKey string        `mapstructure:"REFRESH_TOKEN_PRIVATE_KEY"`
	RefreshTokenPublicKey  string        `mapstructure:"REFRESH_TOKEN_PUBLIC_KEY"`
	AccessTokenExpiresIn   time.Duration `mapstructure:"ACCESS_TOKEN_EXPIRED_IN"`
	RefreshTokenExpiresIn  time.Duration `mapstructure:"REFRESH_TOKEN_EXPIRED_IN"`
	AccessTokenMaxAge      int           `mapstructure:"ACCESS_TOKEN_MAXAGE"`
	RefreshTokenMaxAge     int           `mapstructure:"REFRESH_TOKEN_MAXAGE"`
}

// LoadConfig is the function that loads tthe config from our app.env file
func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	// Tells viper the type of config we want to load
	viper.SetConfigType("env")
	// Tells viper the name of the config file
	viper.SetConfigName("app")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	// unmarshal the values into the Config struct
	err = viper.Unmarshal(&config)
	return
}
