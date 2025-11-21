package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID        string `gorm:"type:uuid;primaryKey"`
	Email     string `json:"email" gorm:"unique" validate:"required"`
	Name      string `json:"name" validate:"required"`
	Role      string `json:"role" validate:"required"`
	Password  string `json:"password" validate:"required"`
	CreatedAt time.Time
	QRCodes   []QRCode `gorm:"foreignKey:UserID"`
}

type QRCode struct {
	ID             string `gorm:"primaryKey"` // Will contain a unique nano id because they are small and url friendly
	Name           string `json:"name"`
	DestinationURL string `json:"destination_url"`
	Active         bool
	UserID         uuid.UUID
	User           User   `gorm:"foreignKey:UserID"`
	QRCodeImage    []byte `gorm:"type:bytea"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      gorm.DeletedAt    `gorm:"index"`
	Analytics      []QRCodeAnalytics `gorm:"foreignKey:QRCodeID"`
}

type QRCodeAnalytics struct {
	ID         uuid.UUID `gorm:"type:uuid;primaryKey"`
	QRCodeID   string
	QRCode     QRCode `gorm:"foreignKey:QRCodeID"`
	UserAgent  string
	DeviceType string
	Country    string
	City       string
	Referrer   string
	CreatedAt  time.Time
}
