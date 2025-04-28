package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	StripeSecretKey   string
	StripeWebhookSecret string
)

func LoadEnv() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	StripeSecretKey = os.Getenv("STRIPE_SECRET_KEY")
	StripeWebhookSecret = os.Getenv("STRIPE_WEBHOOK_SECRET")

	if StripeSecretKey == "" || StripeWebhookSecret == "" {
		log.Fatal("Missing Stripe keys in environment")
	}
}
