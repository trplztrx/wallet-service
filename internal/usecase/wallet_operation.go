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

type WalletOperationUsecase struct {
	walletOperationRepo repo.WalletOperationRepo
	walletRepo repo.WalletRepo
	transactionRepo repo.TransactionRepo
}

func NewWalletOperationUsecase(walletOperationRepo repo.WalletOperationRepo, walletRepo repo.WalletRepo, transactionRepo repo.TransactionRepo) *WalletOperationUsecase {
	return &WalletOperationUsecase{
		walletOperationRepo: walletOperationRepo,
		walletRepo: walletRepo,
		transactionRepo: transactionRepo,
	}
}

func (u *WalletOperationUsecase) Operation(ctx context.Context, operReq *dto.WalletOperationRequest) error {
	wallet, err := u.walletRepo.GetByID(ctx, operReq.WalletID)
	if err != nil {
		return fmt.Errorf("wallet not found")
	}

	switch operReq.OperationType {
	case "DEPOSIT":
		wallet.Balance = wallet.Balance.Add(operReq.Amount) // (opt)TODO: расширить до метода
	case "WITHDRAW":
		wallet.Balance = wallet.Balance.Sub(operReq.Amount) // (opt)TODO: расширить до метода
	}

	// (моя хотелка)TODO: Cделать абстрактные транзакции для отвязки от БД, но похоже не судьба мне это сделать  еххх, дедляйни((
	
	curTime := time.Now()
	wallet.UpdatedAt = curTime
	transaction := &domain.Transaction{
		ID:            uuid.New(),
		WalletID:      wallet.ID,
		OperationType: operReq.OperationType,
		Amount:        operReq.Amount,
		CreatedAt:     curTime,
	}
	err = u.walletOperationRepo.Create(ctx, wallet, transaction)
	if err != nil {
		return fmt.Errorf("failed process operation %v", err)
	}

	return nil
}

