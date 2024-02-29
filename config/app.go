package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	//MYSQL
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	//Redis
	RedisHost string
	RedisPort string
	RedisPass string
	RedisDB   int
	//PORT
	PORT string
}

func LoadConfig() (*Config, error) {
	viper.SetConfigFile(".env")
	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	config := &Config{
		//MYSQL
		DBHost:     viper.GetString("DB_HOST"),
		DBPort:     viper.GetString("DB_PORT"),
		DBUser:     viper.GetString("DB_USER"),
		DBPassword: viper.GetString("DB_PASSWORD"),
		DBName:     viper.GetString("DB_NAME"),
		//Redis
		RedisHost: viper.GetString("REDIS_HOST"),
		RedisPort: viper.GetString("REDIS_PORT"),
		RedisPass: viper.GetString("REDIS_PASSWORD"),
		RedisDB:   viper.GetInt("REDIS_DB"),
		//端口
		PORT: viper.GetString("PORT"),
	}

	return config, nil
}
