package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ProjectType int
type Status int

const (
	Pending    Status = iota
	InProgress Status = 1
	Completed  Status = 2
)

type Order struct {
	ID           string         `gorm:"primaryKey"                  json:"id"`
	ClientName   string         `gorm:"uniqueIndex:idx_order_dedup" json:"client_name"`
	ProjectType  ProjectType    `gorm:"uniqueIndex:idx_order_dedup" json:"project_type"`
	DeliveryDate time.Time      `gorm:"uniqueIndex:idx_order_dedup" json:"delivery_date"`
	Status       Status         `json:"status"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `gorm:"index"                       json:"-"`
}

func (o *Order) BeforeCreate(tx *gorm.DB) error {
	for {
		id := uuid.New().String()
		if err := tx.Where("id = ?", id).First(&Order{}).Error; err == gorm.ErrRecordNotFound {
			o.ID = id
			return nil
		}
	}
}

type CreateOrderRequest struct {
	ClientName   string      `json:"client_name"   binding:"required"`
	ProjectType  ProjectType `json:"project_type"  binding:"required"`
	DeliveryDate time.Time   `json:"delivery_date" binding:"required"`
}

type UpdateOrderRequest struct {
	ClientName   string      `json:"client_name"`
	ProjectType  ProjectType `json:"project_type"`
	Status       Status      `json:"status"`
	DeliveryDate time.Time   `json:"delivery_date"`
}
