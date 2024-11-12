package services

import (
	"context"
	"fmt"

	"github.com/Systenix/fizzbuzz/internal/infrastructures/repositories"
	"github.com/Systenix/fizzbuzz/internal/models"
)

type IFizzBuzzService interface {
	FizzBuzz(
		ctx context.Context,
		request *models.FizzBuzzRequest,
	) (
		result []string,
		err error,
	)
	GetStatistics(
		ctx context.Context,
	) (
		result *models.StatsResponse,
		err error,
	)
}

type FizzBuzzService struct {
	statisticsRepository *repositories.StatisticsRepository
}

func NewFizzBuzzService(
	statisticsRepository *repositories.StatisticsRepository,
) *FizzBuzzService {
	return &FizzBuzzService{
		statisticsRepository: statisticsRepository,
	}
}

func (s *FizzBuzzService) FizzBuzz(
	ctx context.Context,
	req *models.FizzBuzzRequest,
) (
	result []string,
	err error,
) {
	if req.Int1 <= 0 || req.Int2 <= 0 || req.Limit <= 0 {
		return nil, fmt.Errorf("int1, int2, and limit must be positive integers")
	}
	if req.Str1 == "" || req.Str2 == "" {
		return nil, fmt.Errorf("str1 and str2 cannot be empty")
	}

	result = make([]string, 0, req.Limit)
	for i := 1; i <= req.Limit; i++ {
		var entry string

		switch {
		case i%req.Int1 == 0 && i%req.Int2 == 0:
			entry = req.Str1 + req.Str2
		case i%req.Int1 == 0:
			entry = req.Str1
		case i%req.Int2 == 0:
			entry = req.Str2
		default:
			entry = fmt.Sprintf("%d", i)
		}

		result = append(result, entry)
	}

	// Record the request in statistics
	err = s.statisticsRepository.Increment(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("error recording statistics: %v", err)
	}
	return
}

func (s *FizzBuzzService) GetStatistics(
	ctx context.Context,
) (
	result *models.StatsResponse,
	err error,
) {
	return s.statisticsRepository.GetMostFrequent(ctx)
}
