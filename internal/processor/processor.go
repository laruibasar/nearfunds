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

// CreateOrder will have the logic to handle a new order.
func (p *processor) CreateOrder(order models.Order, clock time.Time) error {
	// 1. Check if the order exists.
	exist, err := p.db.FindOrderByCorrelation(order.Correlation)
	if err != nil {
		if err.Error() != "not found" {
			return err
		}
	}

	// 2. Order exists we validate request correlation and data.
	if exist != nil {
		return p.handleExistingOrder(order, *exist)
	}

	// 3. Order does not not exist, try to create.
	return p.createOrder(order, time.Now())
}

func (p *processor) handleExistingOrder(new, old models.Order) error {
	// Implement the validation logic.
	return errors.New("not implement")
}

func (p *processor) createOrder(order models.Order, clock time.Time) error {
	// Implement validation and store, including audit log.
	return errors.New("not implemented.")
}
