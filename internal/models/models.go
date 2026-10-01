// Package models contains in app models, to later reflect on database, according with needs.
package models

import (
	"time"

	"github.com/google/uuid"
)

type Order struct {
	Account     string `json:"account_id"`
	Fund        string `json:"fund_id"`
	Type        string `json:"side"`
	Amount      string `json:"amount"`
	Units       string `json:"units"`
	Correlation uuid.UUID
	Date        time.Time
}
