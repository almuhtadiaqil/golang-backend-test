package config

import (
	"github.com/go-playground/validator/v10"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type appContext struct {
	db               *gorm.DB
	redisClient      *redis.Client
	cfg              *Configuration
	requestValidator *validator.Validate
}

var appCtx appContext

func Init() error {

	cfg, err := InitConfig()
	if err != nil {
		return err
	}

	db, err := ConnectDatabase(cfg.DBConfig)
	if err != nil {
		return err
	}

	redisClient := RedisConnect(cfg.RedisConfig)

	appCtx = appContext{
		db:          db,
		cfg:         cfg,
		redisClient: redisClient,
	}

	return nil
}

func RequestValidator() *validator.Validate {
	return appCtx.requestValidator
}

func DB() *gorm.DB {
	return appCtx.db
}

func REDIS() *redis.Client {
	return appCtx.redisClient
}

func Config() Configuration {
	return *appCtx.cfg
}
