package repositories

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/Systenix/fizzbuzz/internal/infrastructures/redis"
	"github.com/Systenix/fizzbuzz/internal/models"
)

type IStatisticsRepository interface {
	Increment(
		ctx context.Context,
		request *models.FizzBuzzRequest,
	) (
		err error,
	)
	GetMostFrequent(
		ctx context.Context,
	) (
		result *models.StatsResponse,
		err error,
	)
}

type StatisticsRepository struct {
	connector *redis.RedisConnector
}

func NewStatisticsRepository(settings map[string]interface{}) (*StatisticsRepository, error) {
	connector, err := redis.GetRedisConnector(settings)
	if err != nil {
		return nil, err
	}
	return &StatisticsRepository{
		connector: connector,
	}, nil
}

func (r *StatisticsRepository) Increment(
	ctx context.Context,
	request *models.FizzBuzzRequest,
) (
	err error,
) {
	key := "fizzbuzz:stats"
	member, err := json.Marshal(request)
	if err != nil {
		return err
	}

	err = r.connector.GetClient().ZIncrBy(ctx, key, 1, string(member)).Err()
	if err != nil {
		return fmt.Errorf("failed to increment count in Redis: %w", err)
	}

	return nil
}

func (r *StatisticsRepository) GetMostFrequent(
	ctx context.Context,
) (
	result *models.StatsResponse,
	err error,
) {
	key := "fizzbuzz:stats"
	members, err := r.connector.GetClient().ZRevRangeWithScores(ctx, key, 0, 0).Result()
	if err != nil {
		return nil, fmt.Errorf("failed to get statistics from Redis: %w", err)
	}

	if len(members) == 0 {
		return nil, fmt.Errorf("no statistics available")
	}

	var mostFrequentRequest models.FizzBuzzRequest
	err = json.Unmarshal([]byte(members[0].Member.(string)), &mostFrequentRequest)
	if err != nil {
		return nil, fmt.Errorf("failed to decode request data: %w", err)
	}

	result = &models.StatsResponse{
		MostFrequent: &mostFrequentRequest,
		Hits:         int(members[0].Score),
	}

	return result, nil
}
