package redis

import (
	"fmt"
	"sync"

	"github.com/redis/go-redis/v9"
)

var (
	connectors = make(map[string]*RedisConnector)
	mu         sync.Mutex
)

type RedisConnector struct {
	client *redis.Client
}

func GetRedisConnector(settings map[string]interface{}) (*RedisConnector, error) {
	key := generateSettingsKey(settings)

	mu.Lock()
	defer mu.Unlock()

	if connector, exists := connectors[key]; exists {
		return connector, nil
	}

	client := redis.NewClient(&redis.Options{
		Addr:     settings["address"].(string),
		Password: settings["password"].(string),
		DB:       int(settings["db"].(int)),
	})

	// Optionally, test the connection
	// err := client.Ping(context.Background()).Err()
	// if err != nil {
	//     return nil, err
	// }

	connector := &RedisConnector{
		client: client,
	}
	connectors[key] = connector

	return connector, nil
}

func (r *RedisConnector) GetClient() *redis.Client {
	return r.client
}

func generateSettingsKey(settings map[string]interface{}) string {
	return fmt.Sprintf("%v_%v_%v", settings["address"], settings["password"], settings["db"])
}
