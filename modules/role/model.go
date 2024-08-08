package role

import "time"

type Role struct {
	Id         int       `json:"id"`
	Name       string    `json:"name"`
	CreatedAt  time.Time `json:"created_at"`
	CreatedBy  string    `json:"created_by"`
	CreatedOn  string    `json:"created_on"`
	ModifiedAt time.Time `json:"modified_at"`
	ModifiedBy string    `json:"modified_by"`
	ModifiedOn string    `json:"modified_on"`
}
