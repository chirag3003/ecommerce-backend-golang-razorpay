package helpers

import (
	"github.com/chirag3003/ecommerce-golang-api/models"
)

func FillTransactionModel(data map[string]interface{}) (*models.Transaction, bool) {
	amount, ok := data["amount"].(float64)
	if !ok {
		return nil, false
	}
	createdAt, ok := data["created_at"].(float64)
	if !ok {
		return nil, false
	}
	currency, ok := data["currency"].(string)
	if !ok {
		return nil, false
	}
	razorpayID, ok := data["id"].(string)
	if !ok {
		return nil, false
	}
	offerID, ok := data["offer_id"].(string)
	if !ok {
		offerID = ""
	}
	receipt, ok := data["receipt"].(string)
	if !ok {
		return nil, false
	}

	transaction := &models.Transaction{
		PaymentStatus: "pending",
		Razorpay: models.RazorpayTransaction{
			Amount:     amount,
			CreatedAt:  createdAt,
			Currency:   currency,
			RazorpayID: razorpayID,
			OfferID:    offerID,
			Receipt:    receipt,
			Events:     []interface{}{},
		},
	}

	return transaction, true
}
