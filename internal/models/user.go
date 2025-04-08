package models

import (
	"time"
)

type User struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name"`
	Email     string    `json:"email" gorm:"unique"`
	Password  string    `json:"-"` // Password tidak akan ditampilkan dalam JSON
	Phone     string    `json:"phone"`
	Role      string    `json:"role" gorm:"default:'user'"` // user atau driver
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
