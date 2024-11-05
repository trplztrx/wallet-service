package repo

import (
	"context"
	"wallet/internal/domain"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type WalletRepo interface {
	Create(ctx context.Context, wallet *domain.Wallet) error
	Update(ctx context.Context, walletID, newBalance decimal.Decimal) error
	GetByID(ctx context.Context, walletID uuid.UUID) (*domain.Wallet, error)
}