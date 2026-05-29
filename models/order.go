package models

import (
    "time"
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
    gorm.Model
    ClientName   string
    ProjectType  ProjectType
    Status       Status
    DeliveryDate time.Time
}

type CreateOrderRequest struct {
    ClientName   string      `json:"client_name" binding:"required"`
    ProjectType  ProjectType `json:"project_type" binding:"required"`
    DeliveryDate time.Time   `json:"delivery_date" binding:"required"`
}

type UpdateOrderRequest struct {
    ClientName   string      `json:"client_name"`
    ProjectType  ProjectType `json:"project_type"`
    Status       Status      `json:"status"`
    DeliveryDate time.Time   `json:"delivery_date"`
}

