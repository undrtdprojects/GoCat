package common

import (
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

type DefaultFieldTable struct {
	CreatedAt  time.Time `json:"created_at" db:"created_at"`
	CreatedOn  string    `json:"created_on" db:"created_on"`
	ModifiedAt time.Time `json:"modified_at" db:"modified_at"`
	ModifiedOn string    `json:"modified_on" db:"modified_on"`
}

func (d *DefaultFieldTable) SetDefaultField(ctx *gin.Context) {
	hostname, _ := os.Hostname()

	d.CreatedAt = time.Now()
	d.CreatedOn = hostname
	d.ModifiedAt = time.Now()
	d.ModifiedOn = hostname
}
