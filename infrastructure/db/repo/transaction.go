package pgsql

import (
	"context"
	"fmt"
	"wallet/internal/domain"

	"github.com/jackc/pgx/v5/pgxpool"
	// "github.com/shopspring/decimal"
)

type TransactionRepo struct {
	dbPool *pgxpool.Pool
}

func NewTransactionRepo(dbPool *pgxpool.Pool) *TransactionRepo {
	return &TransactionRepo{
		dbPool: dbPool,
	}
}

func (r *TransactionRepo) Create(ctx context.Context, transaction *domain.Transaction) error {
	query := `insert into transactions (id, wallet_id, operation_type, amount, created_at) values ($1, $2, $3, $4, $5)`

	AmountStr := transaction.Amount.String() 
	_, err := r.dbPool.Exec(ctx, query, 
		transaction.ID, 
		transaction.WalletID,
		transaction.OperationType,
		AmountStr,
		transaction.CreatedAt,
	)
	if err != nil {
		return fmt.Errorf("postgres transaction repo: create error: %v", err)
	}

	return nil
}