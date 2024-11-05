package dto

import (
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type WalletOperationRequest struct {
	WalletID uuid.UUID
	OperationType string
	Amount decimal.Decimal
}