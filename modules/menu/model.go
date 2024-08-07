package book

import "time"

type Book struct {
	Id         int       `json:"id"`
	Name       string    `json:"name"`
	Price      int       `json:"price"`
	CategoryId string    `json:"category_id"`
	CreatedAt  time.Time `json:"created_at"`
	CreatedBy  string    `json:"created_by"`
	ModifiedAt time.Time `json:"modified_at"`
	ModifiedBy string    `json:"modified_by"`
}

