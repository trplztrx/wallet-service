package repo

import (
	"context"
	"wallet/internal/domain"
)

type WalletOperationRepo interface {
	Create(ctx context.Context, wallet *domain.Wallet, transaction *domain.Transaction) error
}