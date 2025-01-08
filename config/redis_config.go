package config

import (
	"context"
	"log"
	"os"
	"strconv"

	"github.com/redis/go-redis/v9"
)

var RedisClient *redis.Client

// InitRedis initializes the Redis client with configurable options.
func InitRedis() {
	// Fetch connection details from environment variables
	redisAddr := os.Getenv("REDIS_ADDR")

	redisPassword := os.Getenv("REDIS_PASSWORD")
	redisDB := 0
	if dbEnv := os.Getenv("REDIS_DB"); dbEnv != "" {
		if db, err := strconv.Atoi(dbEnv); err == nil {
			redisDB = db
		} else {
			log.Printf("Invalid REDIS_DB value: %s, using default DB 0", dbEnv)
		}
	}

	// Create the Redis client
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     redisAddr,
		Password: redisPassword,
		DB:       redisDB,
	})

	// Test the Redis connection
	_, err := RedisClient.Ping(context.Background()).Result()
	if err != nil {
		log.Fatalf("Failed to connect to Redis: %v", err)
	}

	log.Println("Connected to Redis successfully")
}
