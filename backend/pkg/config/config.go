package config

import (
	"EduConnect/pkg/logger"
	"EduConnect/pkg/mongodb"
	"EduConnect/pkg/redis"
	"fmt"
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	Http             Http               `mapstructure:"http" validate:"required"`
	Logger           *logger.Config     `mapstructure:"logger" validate:"required"`
	Redis            *redis.RedisConfig `mapstructure:"redis" validate:"required"`
	MongoCollections *MongoCollections  `mapstructure:"mongoCollections" validate:"required"`
	Mongo            *mongodb.Config    `mapstructure:"mongo" validate:"required"`
	AccessTokenExp   time.Duration      `mapstructure:"accessTokenExp" validate:"required"`
	RefreshTokenExp  time.Duration      `mapstructure:"refreshTokenExp" validate:"required"`
	PublicKeyPath    string             `mapstructure:"publicKeyPath" validate:"required"`
	PrivateKeyPath   string             `mapstructure:"privateKeyPath" validate:"required"`
}

type MongoCollections struct {
	Users         string `mapstructure:"users" validate:"required"`
	RefreshTokens string `mapstructure:"refreshTokens" validate:"required"`
	S3            string `mapstructure:"s3" validate:"required"`
	Courses       string `mapstructure:"courses" validate:"required"`
}

type Http struct {
	Port string `mapstructure:"port" validate:"required"`
}

func LoadConfig(pathToConfig string) (*Config, error) {
	if pathToConfig == "" {
		return nil, fmt.Errorf("path to cfg is empty")
	}

	cfg := &Config{}

	viper.SetConfigType("yml")
	viper.SetConfigFile(pathToConfig)

	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("(ReadInConfig) error: %v", err)
	}

	if err := viper.Unmarshal(cfg); err != nil {
		return nil, fmt.Errorf("failed to unmarshal cfg")
	}

	return cfg, nil
}
