package main

import (
	"log"

	"github.com/Systenix/fizzbuzz/internal/infrastructures/repositories"
	"github.com/Systenix/fizzbuzz/internal/interfaces/handlers"
	"github.com/Systenix/fizzbuzz/internal/interfaces/middleware"
	"github.com/Systenix/fizzbuzz/internal/services"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	// Create a router without any middleware by default
	router := gin.Default()

	// Initialize repositories
	statisticsRepository, err := repositories.NewStatisticsRepository(map[string]interface{}{
		"address":  "redis:6379",
		"db":       0,
		"password": "",
	})
	if err != nil {
		log.Fatal(err)
	}

	// Initialize services
	fizzBuzzService := services.NewFizzBuzzService(
		statisticsRepository,
	)

	// Initialize handlers
	fizzBuzzHandler := handlers.NewFizzBuzzHandler(fizzBuzzService)

	// Register routes
	// Route: GET /fizzbuzz
	router.GET("/fizzbuzz",
		middleware.MetricsMiddleware(),
		fizzBuzzHandler.FizzBuzz,
	)
	// Route: GET /statistics
	router.GET("/statistics",
		fizzBuzzHandler.GetStatistics,
	)

	// New Group for Prometheus metrics
	metricsGroup := router.Group("/metrics")
	metricsGroup.GET("", gin.WrapH(promhttp.Handler()))

	// Start the server
	log.Println("Starting server on :8080")
	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
