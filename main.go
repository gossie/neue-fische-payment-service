package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type paymentRequest struct {
	OrderId string `json:"orderId"`
	Item    string `json:"item"`
}

type paymentResponse struct {
	OrderId string  `json:"orderId"`
	Item    string  `json:"item"`
	Price   float64 `json:"price"`
}

func determinePrice(item string) float64 {
	if item == "PIZZA" {
		return 9.5
	}
	return 10.5
}

func pay(c *gin.Context) {
	var request paymentRequest

	if err := c.BindJSON(&request); err != nil {
		return
	}

	c.IndentedJSON(http.StatusOK, paymentResponse{
		OrderId: request.OrderId,
		Item:    request.Item,
		Price:   determinePrice(request.Item),
	})
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}

	router := gin.Default()
	router.POST("/api/payment", pay)

	router.Run(":" + port)
}
