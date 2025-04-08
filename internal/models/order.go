package models

import (
	"time"
)

type Location struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Address   string  `json:"address"`
}

type Order struct {
	ID             uint       `json:"id" gorm:"primaryKey"`
	UserID         uint       `json:"user_id"`
	DriverID       *uint      `json:"driver_id"`
	PickupLocation Location   `json:"pickup_location" gorm:"embedded;embeddedPrefix:pickup_"`
	DestLocation   Location   `json:"destination_location" gorm:"embedded;embeddedPrefix:dest_"`
	Status         string     `json:"status" gorm:"default:'pending'"` // pending, accepted, ongoing, completed, cancelled
	Price          float64    `json:"price"`
	PaymentStatus  string     `json:"payment_status" gorm:"default:'unpaid'"` // unpaid, paid
	PaymentMethod  string     `json:"payment_method"`
	CreatedAt      time.Time  `json:"created_at"`
	UpdatedAt      time.Time  `json:"updated_at"`
	CompletedAt    *time.Time `json:"completed_at,omitempty"`
}
