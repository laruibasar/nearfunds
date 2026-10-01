// Package processor contains the business logic to handle requests.
package processor

import (
	"errors"
	"time"

	"github.com/laruibasar/nearfunds/internal/database"
	"github.com/laruibasar/nearfunds/internal/models"
)

type Processor interface {
	CreateOrder(order models.Order, clock time.Time) error
}

type processor struct {
	db database.Database
}

func New(db database.Database) *processor {
	return &processor{
		db: db,
	}
}

func (p *processor) CreateOrder(order models.Order, clock time.Time) error {
	return errors.New("not implemented")
}
