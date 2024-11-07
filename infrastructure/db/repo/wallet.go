package pgsql

import (
	"context"
	"fmt"
	"wallet/internal/domain"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	// "github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type WalletRepo struct {
	dbPool *pgxpool.Pool
}

func NewWalletRepo(dbPool *pgxpool.Pool) *WalletRepo {
	return &WalletRepo{
		dbPool: dbPool,
	}
}

func (r *WalletRepo) Create(ctx context.Context, wallet *domain.Wallet) error {
	
	query := `
		insert into wallets (id, balance, created_at, updated_at)
		values ($1, $2, $3, $4)
	`
	_, err := r.dbPool.Exec(ctx, query, wallet.ID, wallet.Balance, wallet.CreatedAt, wallet.UpdatedAt)
	if err != nil {
		return fmt.Errorf("postgres wallet repo: create error: %w", err)
	}

	return nil
}

func (r *WalletRepo) GetByID(ctx context.Context, walletID uuid.UUID) (*domain.Wallet, error) {
	var wallet domain.Wallet
	var balanceStr string
	query := `select * from wallets where id = $1`
	row := r.dbPool.QueryRow(ctx, query, walletID)

	err := row.Scan(&wallet.ID, &balanceStr, &wallet.CreatedAt, &wallet.UpdatedAt)
	if err != nil {
		return nil, fmt.Errorf("postgres wallet repo: get by id error: %v", err)
	}

	wallet.Balance, err = decimal.NewFromString(balanceStr)
	if err != nil {
		return nil, fmt.Errorf("failed to convert balance to decimal: %v", err)
	}

	return &wallet, nil
} 	

func (r *WalletRepo) Update(ctx context.Context, wallet *domain.Wallet) error {
	query := `update wallets set balance = $1, updated_at = $2 where id = $3`

	_, err := r.dbPool.Exec(ctx, query, wallet.Balance, wallet.UpdatedAt, wallet.ID)
	if err != nil {
		return fmt.Errorf("postgres wallet repo: update error: %v", err)
	}
	return nil
}