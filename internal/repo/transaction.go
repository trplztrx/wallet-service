package repo

import (
	"context"
	"wallet/internal/domain"
)

type TransactionRepo interface {
	Create(ctx context.Context, transaction *domain.Transaction) error
}