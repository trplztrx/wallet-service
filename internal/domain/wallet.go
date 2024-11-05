package domain

import (
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Wallet struct {
	ID        uuid.UUID
	Balance   decimal.Decimal
	CreatedAt time.Time
	UpdatedAt time.Time
}
