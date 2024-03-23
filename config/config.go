package config

import (
	"os"

	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
)

type (
	DBConfig struct {
		Username string `mapstructure:"DB_USERNAME" validate:"required"`
		Password string `mapstructure:"DB_PASSWORD" validate:"required"`
		Host     string `mapstructure:"DB_HOST" validate:"required"`
		Port     string `mapstructure:"DB_PORT" validate:"required"`
		Database string `mapstructure:"DB_DATABASE" validate:"required"`
	}

	RedisConfig struct {
		Username string `mapstructure:"REDIS_USERNAME"`
		Password string `mapstructure:"REDIS_PASSWORD"`
		Host     string `mapstructure:"REDIS_HOST"`
		Port     string `mapstructure:"REDIS_PORT"`
		Database string `mapstructure:"REDIS_DATABASE"`
	}

	Configuration struct {
		Environment string      `mapstructure:"ENV" validate:"required,oneof=development staging production"`
		BindAddress int         `mapstructure:"BIND_ADDRESS" validate:"required"`
		ServiceName string      `mapstructure:"SERVICE_NAME"`
		DBConfig    DBConfig    `mapstructure:",squash"`
		RedisConfig RedisConfig `mapstructure:",squash"`
	}
)

func InitConfig() (*Configuration, error) {
	var cfg Configuration

	viper.SetConfigType("env")
	envFile := os.Getenv("ENV_FILE")
	if envFile == "" {
		envFile = ".env"
	}

	_, err := os.Stat(envFile)
	if !os.IsNotExist(err) {
		viper.SetConfigFile(envFile)

		if err := viper.ReadInConfig(); err != nil {
			return nil, err
		}
	}

	viper.AutomaticEnv()

	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, err
	}

	validate := validator.New()
	if err := validate.Struct(cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
