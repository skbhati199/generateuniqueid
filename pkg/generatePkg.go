package generatePkg

import (
	"fmt"
	"github.com/go-redis/redis"
	"log"
	"os"
	"strconv"
	"time"
)

const (
	envRedisUrl      = "REDIS_URL"
	envRedisPassword = "REDIS_PASSWORD"
)

func NewGeneratePkg() (string, error) {
	// Create a Redis client
	redisUrl, ok := os.LookupEnv(envRedisUrl)
	if !ok {
		fmt.Printf("Failed to get redis url from .env file %v", ok)
		return "", fmt.Errorf("failed to get redis password from .env file %v", ok)
	}
	redisPassword, ok := os.LookupEnv(envRedisPassword)
	if !ok {
		fmt.Printf("Failed to get redis password from .env file %v", ok)
		return "", fmt.Errorf("failed to get redis password from .env file %v", ok)
	}
	client := redis.NewClient(&redis.Options{
		Addr:         redisUrl,
		Password:     redisPassword,
		DB:           0, // use default DB
		ReadTimeout:  time.Second * 30,
		WriteTimeout: time.Second * 30,
	})

	// Ping the Redis server to check if it's available
	pong, err := client.Ping().Result()
	if err != nil {
		log.Fatalf("Failed to ping Redis: %v", err)
		return "", err
	}
	log.Printf("Redis ping response: %s", pong)

	// Define a key prefix for the transaction IDs

	// Define the initial value for the transaction ID sequence
	initialValue := uint64(1000)

	// Get the current value of the transaction ID sequence from Redis
	currValueStr, err := client.Get("transactionIds:curr").Result()
	if err == redis.Nil {
		// The key doesn't exist, set it to the initial value
		err := client.Set("transactionIds:curr", strconv.FormatUint(initialValue, 10), 0).Err()
		if err != nil {
			log.Fatalf("Failed to set initial value for transaction ID sequence: %v", err)
			return "", err
		}
		currValueStr = strconv.FormatUint(initialValue, 10)
	} else if err != nil {
		log.Fatalf("Failed to get current value for transaction ID sequence: %v", err)
		return "", err
	}

	// Parse the current value as a uint64
	currValue, err := strconv.ParseUint(currValueStr, 10, 64)
	if err != nil {
		log.Fatalf("Failed to parse current value for transaction ID sequence: %v", err)
		return "", err
	}

	// Define a function to generate a new transaction ID
	generateTxID := func() string {
		// Increment the current value of the transaction ID sequence
		newValue := currValue + 1
		err := client.Set("transactionIds:curr", strconv.FormatUint(newValue, 10), 0).Err()
		if err != nil {
			log.Fatalf("Failed to update current value for transaction ID sequence: %v", err)
		}
		currValue = newValue

		// Generate the transaction ID
		return strconv.FormatUint(currValue, 10)
	}

	// Use the generateTxID function to generate a new transaction ID
	txID := generateTxID()
	return txID, nil
}
