package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/stripe/stripe-go/v78/webhook"
	"stripe-payment-service/internal/stripe"
)

// PaymentRequest defines the payload for creating payments
type PaymentRequest struct {
	TaskID   string `json:"task_id" binding:"required"`
	Amount   int64  `json:"amount" binding:"required"`
	Currency string `json:"currency" binding:"required"`
}

// CreateUpfrontPayment handles the upfront initial payment for a task
func CreateUpfrontPayment(c *gin.Context) {
	var req PaymentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	intent, err := stripe.CreatePaymentIntent(req.Amount, req.Currency, req.TaskID, "upfront")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"client_secret": intent.ClientSecret,
		"payment_intent_id": intent.ID,
	})
}

// CreateTimePayment handles charging for time spent on a task
func CreateTimePayment(c *gin.Context) {
	var req PaymentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	intent, err := stripe.CreatePaymentIntent(req.Amount, req.Currency, req.TaskID, "time")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"client_secret": intent.ClientSecret,
		"payment_intent_id": intent.ID,
	})
}

// CreateRewardPayment handles charging the final reward after client approval
func CreateRewardPayment(c *gin.Context) {
	var req PaymentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	intent, err := stripe.CreatePaymentIntent(req.Amount, req.Currency, req.TaskID, "reward")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"client_secret": intent.ClientSecret,
		"payment_intent_id": intent.ID,
	})
}


// HandleWebhook processes Stripe webhook events
func HandleWebhook(c *gin.Context) {
	const MaxBodyBytes = int64(65536)
	c.Request.Body = http.MaxBytesReader(c.Writer, c.Request.Body, MaxBodyBytes)

	payload, err := c.GetRawData()
	if err != nil {
		c.String(http.StatusServiceUnavailable, "Error reading request body")
		return
	}

	endpointSecret := os.Getenv("STRIPE_WEBHOOK_SECRET")
	sigHeader := c.GetHeader("Stripe-Signature")

	event, err := webhook.ConstructEvent(payload, sigHeader, endpointSecret)
	if err != nil {
		c.String(http.StatusBadRequest, "Webhook signature verification failed")
		return
	}

	switch event.Type {
	case "payment_intent.succeeded":
		log.Println("Payment succeeded 🎉")
	case "payment_intent.payment_failed":
		log.Println("Payment failed 💥")
	default:
		log.Println("Unhandled event type:", event.Type)
	}

	c.String(http.StatusOK, "Received")
}
