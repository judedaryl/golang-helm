package entities

import (
	"time"
)

type Product struct {
	ID        string `json:"id" gorm:"primaryKey"`
	Name      string `json:"name"`
	CreatedAt time.Time
}
