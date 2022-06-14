package models

type Transaction struct {
	Razorpay      RazorpayTransaction `json:"razorpay" bson:"razorpay"`
	PaymentStatus string              `json:"paymentStatus" bson:"paymentStatus"`
}

type RazorpayTransaction struct {
	Amount     float64       `json:"amount" bson:"amount"`
	CreatedAt  float64       `json:"created_at" bson:"createdAt"`
	Currency   string        `json:"currency" bson:"currency"`
	RazorpayID string        `json:"id" bson:"razorpayID"`
	OfferID    string        `json:"offer_id" bson:"offerID"`
	Receipt    string        `json:"receipt" bson:"receipt"`
	Events     []interface{} `json:"events" bson:"events"`
}
