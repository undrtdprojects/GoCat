package transaction0

import (
	"GoCat/modules/transaction1"
	"time"
)

type Transaction0 struct {
	Id              string                      `json:"id"`
	UserId          int                         `json:"user_id"`
	PaymentId       int                         `json:"payment_id"`
	GrandTotalPrice int                         `json:"grand_total_price"`
	CreatedAt       time.Time                   `json:"created_at"`
	CreatedBy       string                      `json:"created_by"`
	CreatedOn       string                      `json:"created_on"`
	ModifiedAt      time.Time                   `json:"modified_at"`
	ModifiedBy      string                      `json:"modified_by"`
	ModifiedOn      string                      `json:"modified_on"`
	ListDetail      []transaction1.Transaction1 `json:"list_detail"`
}
