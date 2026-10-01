// Package processor contains the business logic to handle requests.
package processor

import (
	"errors"
	"time"

	"github.com/laruibasar/nearfunds/internal/models"
)

type Processor interface {
	CreateOrder(order models.Order) error
}

type processor struct {
	// we will save the service to handle database interactions.
}

func New() *processor {
	return &processor{}
}

func (p *processor) CreateOrder(order models.Order, time time.Time) error {
	return errors.New("not implemented")
}
