package database

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Order database model data type.
type Order struct {
	gorm.Model // Allow to use across models to have ID, CreatedAt, ...

	ID          uint
	Account     string // Adapt to db models, will be an ID and can be used as entity struct in the model.
	Fund        string // Adapt to db models.
	Type        string
	Amount      uint64 // Adapt to real value.
	Units       uint64
	Correlation uuid.UUID
	Data        time.Time
}
