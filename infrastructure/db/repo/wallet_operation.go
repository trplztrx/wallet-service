package pgsql

import (
	"context"
	"fmt"
	"wallet/internal/domain"
	"wallet/internal/repo"

	"github.com/jackc/pgx/v5/pgxpool"
)

type WalletOperationRepo struct {
	dbPool *pgxpool.Pool
	walletRepo repo.WalletRepo
	transactionRepo repo.TransactionRepo
}

func NewWalletOperationRepo(dbPool *pgxpool.Pool, walletRepo repo.WalletRepo, transactionRepo repo.TransactionRepo) *WalletOperationRepo {
	return &WalletOperationRepo{
		dbPool: dbPool,
		walletRepo: walletRepo,
		transactionRepo: transactionRepo,
	}
}

func (r *WalletOperationRepo) Create(ctx context.Context, wallet *domain.Wallet, transaction *domain.Transaction) error {
	tx, err := r.dbPool.Begin(ctx)
	if err != nil {
		return fmt.Errorf("postgres wallet operation repo: transaction begin error: %v", err)
	}
	defer func() {
		if err != nil {
			tx.Rollback(ctx)
		} else {
			err = tx.Commit(ctx)
		}
	}()

	err = r.walletRepo.Update(ctx, wallet)
	if err != nil {
		return err
	}
	err = r.transactionRepo.Create(ctx, transaction)
	if err != nil {
		return err
	}

	return nil
}