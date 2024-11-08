package database

import (
	"context" // Import context for Redis operations
	"fmt"
	"log"
	"os"

	"github.com/go-redis/redis/v8" // Import Redis package
)

var (
	rdb *redis.Client          // Redis client
	ctx = context.Background() // Context for Redis operations
)

// InitializeRedis initializes the Redis client
func InitializeRedis() {
	redisHost := os.Getenv("REDIS_HOST")
	redisPort := os.Getenv("REDIS_PORT")
	rdb = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", redisHost, redisPort), // Redis server address
		Password: "",                                         // No password set
		DB:       0,                                          // Use default DB
	})

	// Test the connection
	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("Could not connect to Redis: %v", err) // Log fatal error if connection fails
	}

	log.Println("Connected to Redis successfully") // Log success message
}

// GetRedisClient returns the Redis client
func GetRedisClient() *redis.Client {
	return rdb
}

// CloseRedis closes the Redis client connection
func CloseRedis() {
	if rdb != nil {
		err := rdb.Close()
		if err != nil {
			log.Printf("Error closing Redis connection: %v", err) // Log error if closing fails
		} else {
			log.Println("Redis connection closed successfully") // Log success message
		}
	}
}
