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
	ID          string `gorm:"primaryKey" validate:"required"` // Will contain a unique nano id because they are small and url friendly
	Name        string `json:"name" validate:"required"`
	Description string `json:"description"`

	// qrcode data
	DestinationURL string `json:"destination_url" validate:"required"`
	Active         bool
	QRCodeImage    []byte `gorm:"type:bytea" validate:"required"` //

	// time stamps
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`

	// the relation ship data
	UserID    string            `json:"userid" validate:"required"`
	User      User              `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE,OnUpdate:CASCADE;"`
	Analytics []QRCodeAnalytics `gorm:"foreignKey:QRCodeID"`
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
