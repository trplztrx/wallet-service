package usecase

import (
	"context"
	"fmt"
	"time"
	"wallet/internal/domain"
	"wallet/internal/repo"
	"wallet/internal/transport/dto"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type WalletUsecase struct {
	walletRepo repo.WalletRepo
}

func NewWalletUsecase(walletRepo repo.WalletRepo) *WalletUsecase{
	return &WalletUsecase{walletRepo: walletRepo}
}

func (u *WalletUsecase) CreateWallet(ctx context.Context) (*domain.Wallet, error) {
	wallet := &domain.Wallet{
		ID: uuid.New(),
		Balance: decimal.Zero,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err := u.walletRepo.Create(ctx, wallet)
	if err != nil {
		return nil, fmt.Errorf("wallet usecase: create error: %v", err)
	}

	return wallet, nil
}

func (u *WalletUsecase) GetBalance(ctx context.Context, walletID uuid.UUID) (decimal.Decimal, error) {
	wallet, err := u.walletRepo.GetByID(ctx, walletID)
	if err != nil {
		return nil, fmt.Errorf("wallet not found")
	}

	return wallet.Balance
}

