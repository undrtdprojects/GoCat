package transaction1

import "time"

type Transaction1 struct {
	Id              int       `json:"id"`
	TransactionId   string    `json:"transaction_id"`
	MenuId          int       `json:"menu_id"`
	DateTransaction time.Time `json:"date_transaction"`
	Qty             int       `json:"qty"`
	TotalPrice      int       `json:"total_price"`
	PaymentId       int       `json:"payment_id"`
	CreatedAt       time.Time `json:"created_at"`
	CreatedBy       string    `json:"created_by"`
	ModifiedAt      time.Time `json:"modified_at"`
	ModifiedBy      string    `json:"modified_by"`
}