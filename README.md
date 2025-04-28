# dev-irl-payment-service

```
aws secretsmanager create-secret \
  --name devirl/stripe \
  --description "Stripe credentials for Dev IRL Payment Service" \
  --secret-string '{"stripeSecretKey":"your_sk_test_key_here","stripeWebhookSecret":"your_whsec_key_here"}'
```
