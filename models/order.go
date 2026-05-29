package models

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ProjectType string
type Status int

const (
	Pending    Status = iota
	InProgress Status = 1
	Completed  Status = 2
)

var statusNames = map[Status]string{
	Pending:    "Pending",
	InProgress: "In Progress",
	Completed:  "Completed",
}

var statusValues = map[string]Status{
	"Pending":     Pending,
	"In Progress": InProgress,
	"Completed":   Completed,
}

func (s Status) MarshalJSON() ([]byte, error) {
	name, ok := statusNames[s]
	if !ok {
		return nil, fmt.Errorf("unknown status: %d", s)
	}
	return json.Marshal(name)
}

func (s *Status) UnmarshalJSON(data []byte) error {
	var name string
	if err := json.Unmarshal(data, &name); err != nil {
		return err
	}
	val, ok := statusValues[name]
	if !ok {
		return fmt.Errorf("unknown status: %q", name)
	}
	*s = val
	return nil
}

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

func (r CreateOrderRequest) Validate() error {
	if strings.TrimSpace(r.ClientName) == "" {
		return fmt.Errorf("client_name cannot be blank")
	}
	if strings.TrimSpace(string(r.ProjectType)) == "" {
		return fmt.Errorf("project_type cannot be blank")
	}
	if !r.DeliveryDate.After(time.Now()) {
		return fmt.Errorf("delivery_date must be in the future")
	}
	return nil
}

type UpdateOrderRequest struct {
	ClientName   string      `json:"client_name"`
	ProjectType  ProjectType `json:"project_type"`
	Status       Status      `json:"status"`
	DeliveryDate time.Time   `json:"delivery_date"`
}

func (r UpdateOrderRequest) Validate() error {
	if r.ClientName != "" && strings.TrimSpace(r.ClientName) == "" {
		return fmt.Errorf("client_name cannot be blank")
	}
	if r.ProjectType != "" && strings.TrimSpace(string(r.ProjectType)) == "" {
		return fmt.Errorf("project_type cannot be blank")
	}
	if !r.DeliveryDate.IsZero() && !r.DeliveryDate.After(time.Now()) {
		return fmt.Errorf("delivery_date must be in the future")
	}
	return nil
}
