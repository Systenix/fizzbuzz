package services

import (
	"context"
	"testing"

	"github.com/Systenix/fizzbuzz/internal/models"
	"github.com/stretchr/testify/assert"
)

func TestFizzBuzzService_FizzBuzz(t *testing.T) {
	service := NewFizzBuzzService(nil) // Pass a mock repository if needed

	req := &models.FizzBuzzRequest{
		Int1:  3,
		Int2:  5,
		Limit: 15,
		Str1:  "Fizz",
		Str2:  "Buzz",
	}

	result, err := service.FizzBuzz(context.Background(), req)
	assert.NoError(t, err)
	expected := []string{"1", "2", "Fizz", "4", "Buzz", "Fizz", "7", "8", "Fizz", "Buzz", "11", "Fizz", "13", "14", "FizzBuzz"}
	assert.Equal(t, expected, result)
}
