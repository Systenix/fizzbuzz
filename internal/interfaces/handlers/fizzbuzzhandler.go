package handlers

import (
	"net/http"
	"strconv"

	"github.com/Systenix/fizzbuzz/internal/models"
	"github.com/Systenix/fizzbuzz/internal/services"
	"github.com/gin-gonic/gin"
)

type IFizzBuzzHandler interface {
	FizzBuzz(c *gin.Context)
	GetStatistics(c *gin.Context)
}

type FizzBuzzHandler struct {
	Service *services.FizzBuzzService
}

func NewFizzBuzzHandler(service *services.FizzBuzzService) *FizzBuzzHandler {
	return &FizzBuzzHandler{
		Service: service,
	}
}

func (h *FizzBuzzHandler) FizzBuzz(c *gin.Context) {
	// Extract query parameters
	int1, err := strconv.Atoi(c.Query("int1"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid int1 parameter"})
		return
	}
	int2, err := strconv.Atoi(c.Query("int2"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid int2 parameter"})
		return
	}
	limit, err := strconv.Atoi(c.Query("limit"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid limit parameter"})
		return
	}
	str1 := c.Query("str1")
	str2 := c.Query("str2")
	if str1 == "" || str2 == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "str1 and str2 parameters are required"})
		return
	}

	req := &models.FizzBuzzRequest{
		Int1:  int1,
		Int2:  int2,
		Limit: limit,
		Str1:  str1,
		Str2:  str2,
	}

	// Call service method
	res, err := h.Service.FizzBuzz(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"resp": models.FizzBuzzResponse{
		Result: res,
	}})
}

func (h *FizzBuzzHandler) GetStatistics(c *gin.Context) {

	// Call service method
	res, err := h.Service.GetStatistics(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"resp": res})
}
