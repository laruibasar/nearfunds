package processor

import (
	"testing"
	"time"

	"github.com/google/uuid"

	"git.sr.ht/~laruibasar/nearfunds/internal/models"
)

func TestCreateOrder(t *testing.T) {
	p := New()

	tests := []struct {
		name           string
		order          models.Order
		expectedResult error
	}{
		{
			name: "Valid test order",
			order: models.Order{
				Account:     "ACC-1",
				Fund:        "FUND-A",
				Type:        "SUBSCRIPTION",
				Amount:      "1000.00",
				Correlation: uuid.New(),
				Date:        time.Now(),
			},
			expectedResult: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			err := p.CreateOrder(test.order)

			if err != test.expectedResult {
				t.Fatalf("Test \"%s\" failed with: %v", test.name, err)
			}
		})
	}
}
