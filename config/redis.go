package config

import (
	"backend-test/src/entities"
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

func RedisConnect(cfg RedisConfig) *redis.Client {
	ctx := context.Background()

	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", cfg.Host, cfg.Port),
		Username: cfg.Username,
		Password: cfg.Password,
		DB:       0,
	})

	defer client.Close()

	status, err := client.Ping(ctx).Result()
	if err != nil {
		fmt.Println("Redis connection was refused")
	}
	fmt.Println(status)

	// SaveDataFromRedisToDB(ctx)
	return client
}

func SaveDataFromRedisToDB(ctx context.Context) {
	for {
		select {
		case <-time.After(2 * time.Minute):
			// Get data from Redis
			keys := REDIS().Keys(ctx, "product_*").Val()
			for _, key := range keys {
				val, err := REDIS().Get(ctx, key).Result()
				if err != nil {
					fmt.Println("Error getting data from Redis:", err)
					continue
				}

				// Unmarshal JSON data
				var product entities.Product
				if err := json.Unmarshal([]byte(val), &product); err != nil {
					fmt.Println("Error unmarshalling JSON:", err)
					continue
				}

				// Save data to DB
				if err := DB().Create(&product).Error; err != nil {
					fmt.Println("Error saving data to DB:", err)
					continue
				}

				// Delete data from Redis
				if err := REDIS().Del(ctx, key).Err(); err != nil {
					fmt.Println("Error deleting data from Redis:", err)
					continue
				}

				fmt.Println("Data saved from Redis to DB:", product)
			}
		}
	}
}
