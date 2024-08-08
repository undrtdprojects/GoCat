package menu

import "time"

type Menu struct {
	Id         string    `json:"id"`
	Name       string    `json:"name"`
	Price      int       `json:"price"`
	CategoryId string    `json:"category_id"`
	CreatedAt  time.Time `json:"created_at"`
	CreatedBy  string    `json:"created_by"`
	CreatedOn  string    `json:"created_on"`
	ModifiedAt time.Time `json:"modified_at"`
	ModifiedBy string    `json:"modified_by"`
	ModifiedOn string    `json:"modified_on"`
}
