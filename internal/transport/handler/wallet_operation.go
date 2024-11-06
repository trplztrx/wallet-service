package handler

import (
	"net/http"
	// "wallet/internal/transport/dto"
	"wallet/internal/usecase"
)

type WalletOperationHandler struct {
	uc usecase.WalletOperationUsecase
}

func NewWalletOperationHandler(uc usecase.WalletOperationUsecase) *WalletOperationHandler {
	return &WalletOperationHandler{uc: uc}
}

func (h *WalletOperationHandler) Operation(w http.ResponseWriter, r *http.Request) {

}