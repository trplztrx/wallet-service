package usecase

import (
	"context"
	"fmt"
	"wallet/internal/repo"
	"wallet/internal/transport/dto"
)

type WalletOperationUsecase struct {
	walletRepo repo.WalletRepo
	transactionRepo repo.TransactionRepo
}

func NewWalletOperationUsecase(walletRepo repo.WalletRepo, transactionRepo repo.TransactionRepo) *WalletOperationUsecase {
	return &WalletOperationUsecase{
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
	// TODO: Хочется сделать абстрактные транзакции для отвязки от БД, но похоже не судьба мне это сделать  еххх, дедляйни((
}

