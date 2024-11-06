package dto

import (
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type WalletOperationRequest struct {
	WalletID      uuid.UUID       `json:"wallet_id"`
	OperationType string          `json:"operation_type"`
	Amount        decimal.Decimal `json:"amount"`
}
