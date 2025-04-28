package stripe

import (
	"github.com/stripe/stripe-go/v78"
	"github.com/stripe/stripe-go/v78/paymentintent"
)

func Init(secretKey string) {
	stripe.Key = secretKey
}

func CreatePaymentIntent(amount int64, currency, taskID, phase string) (*stripe.PaymentIntent, error) {
	params := &stripe.PaymentIntentParams{
		Amount:   stripe.Int64(amount),
		Currency: stripe.String(currency),
		AutomaticPaymentMethods: &stripe.PaymentIntentAutomaticPaymentMethodsParams{
			Enabled: stripe.Bool(true),
		},
		Metadata: map[string]string{
			"task_id": taskID,
			"payment_phase": phase,
		},
	}

	return paymentintent.New(params)
}
