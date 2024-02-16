package services

type aggregated struct {
	PaymentId         int    `json:"payment_id"`
	ClientId          string `json:"client_id"`
	PaymentMethodType string `json:"payment_method_type"`
	PaymentMethodId   string `json:"payment_method_id"`
	Amount            int    `json:"amount"`
	Currency          string `json:"currency"`
}

func newPayment() *aggregated {
	return &aggregated{}
}
