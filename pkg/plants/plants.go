package plants

import (
	"time"
)

type Plant struct {
	ID          uint   `json:"id" gorm:"primarykey"`
	Name        string `json:"name"`
	Active      bool   `json:"active"`
	Description string `json:"description"`

	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at" gorm:"index"`
}
