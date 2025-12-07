package models

type UpdateQRCodeDTO struct { // DTO stands for Data Transfer Object
	ID             string `json:"id" validate:"required"`
	Name           string `json:"name" validate:"required"`
	Description    string `json:"description,omitempty"`
	DestinationURL string `json:"destination_url" validate:"required"`
	Active         bool   `json:"active"`
}

type CreateQRCodeDTO struct {
	ID             string `json:"id" validate:"required"`
	Name           string `json:"name" validate:"required"`
	Description    string `json:"description,omitempty"`
	DestinationURL string `json:"destination_url" validate:"required"`
	Active         bool   `json:"active"`
}
