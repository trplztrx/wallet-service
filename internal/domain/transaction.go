package domain

import (
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Transaction struct {
	ID            uuid.UUID
	WalletID      uuid.UUID
	OperationType string
	Amount        decimal.Decimal
	CreatedAt     time.Time
}
