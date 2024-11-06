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

func (wr *WalletRepo) Create(ctx context.Context, wallet *domain.Wallet) error {
	return nil
}

func (wr *WalletRepo) GetByID(ctx context.Context, walletID uuid.UUID) (*domain.Wallet, error) {
	var wallet domain.Wallet
	var balanceStr string
	query := `select * from wallets where id = $1`
	row := wr.dbPool.QueryRow(ctx, query, walletID)

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
