package transaction1

import "time"

type Transaction1 struct {
	Id              int       `json:"id"`
	TransactionId   string    `json:"transaction_id"`
	MenuId          string    `json:"menu_id"`
	DateTransaction time.Time `json:"date_transaction"`
	Qty             int       `json:"qty"`
	TotalPrice      int       `json:"total_price"`
	CreatedAt       time.Time `json:"created_at"`
	CreatedBy       string    `json:"created_by"`
	CreatedOn       string    `json:"created_on"`
	ModifiedAt      time.Time `json:"modified_at"`
	ModifiedBy      string    `json:"modified_by"`
	ModifiedOn      string    `json:"modified_on"`
}
