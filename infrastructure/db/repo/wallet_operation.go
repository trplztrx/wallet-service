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
			rbErr := tx.Rollback(ctx)
			if rbErr != nil {
				err = fmt.Errorf("postgres wallet operation repo: create error: %v", err.Error())
			}
		} else {
			err = tx.Commit(ctx)
		}
	}()

	// Не хватает да как будто бы в этом месте той самой абстракции транзакций ыхых

	query := `update wallets set balance = $1, updated_at = $2 where id = $3`

	_, err = tx.Exec(ctx, query, wallet.Balance, wallet.UpdatedAt, wallet.ID)
	if err != nil {
		return fmt.Errorf("postgres wallet repo: update error: %v", err)
	}

	query = `insert into transactions (id, wallet_id, operation_type, amount, created_at) values ($1, $2, $3, $4, $5)`

	AmountStr := transaction.Amount.String() 
	_, err = tx.Exec(ctx, query, 
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