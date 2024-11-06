package dto

import (
	// "github.com/google/uuid"
	"github.com/shopspring/decimal"
)

// type WalletUpdateReq struct {
// 	WalletID uuid.UUID
// 	Balance decimal.Decimal 
// }

type WalletGetBalanceResponse struct {
	Balance decimal.Decimal `json:"balance"`
}


