package main

import (
	"log"
	"stripe-payment-service/internal/api/handlers"
	"stripe-payment-service/internal/config"
	"stripe-payment-service/internal/stripe"

	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadEnv()
	stripe.Init(config.StripeSecretKey)

	r := gin.Default()

	r.POST("/payments/upfront", handlers.CreateUpfrontPayment)
	r.POST("/payments/time", handlers.CreateTimePayment)
	r.POST("/payments/reward", handlers.CreateRewardPayment)
	r.POST("/webhook", handlers.HandleWebhook)

	log.Println("Server running on :8080")
	r.Run(":8080")
}
