package models

import (
)

type StatsResponse struct {
    MostFrequent *FizzBuzzRequest `json:"most_frequent"`
    Hits int `json:"hits"`
}
