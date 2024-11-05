package usecase

import (
	"context"
	"fmt"
	"time"
	"wallet/internal/domain"
	"wallet/internal/repo"
	"wallet/internal/transport/dto"

	"github.com/google/uuid"
)

type TransactionUsecase struct {
	transactionRepo repo.TransactionRepo
}

func NewTransactionUsecase(transactionRepo repo.TransactionRepo) *TransactionUsecase {
	return &TransactionUsecase{transactionRepo: transactionRepo}
}

func (u *TransactionUsecase) CreateTransaction(ctx context.Context, transReq *dto.TransactionCreateRequest) (*domain.Transaction, error) {
	transaction := &domain.Transaction{
		ID:            uuid.New(),
		WalletID:      transReq.WalletID,
		OperationType: transReq.OperationType,
		Amount:        transReq.Amount,
		CreatedAt:     time.Now(),
	}

	err := u.transactionRepo.Create(ctx, transaction)
	if err != nil {
		return nil, fmt.Errorf("transaction usecase: create error: %v", err)
	}

	return transaction, nil
}
