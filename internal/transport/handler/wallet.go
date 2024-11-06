package handler

import (
	"encoding/json"
	"net/http"
	"wallet/internal/usecase"

	"github.com/google/uuid"
	// "github.com/shopspring/decimal"
)

type WalletHandler struct {
	uc usecase.WalletUsecase
}

func NewWalletHandler(uc usecase.WalletUsecase) *WalletHandler {
	return &WalletHandler{
		uc: uc,
	}
}

func (h *WalletHandler) GetBalance(w http.ResponseWriter, r *http.Request) {
	walletIDStr := r.URL.Query().Get("wallet_id")
	walletID, err := uuid.Parse(walletIDStr)
	if err != nil {
		// TODO: обертку HTTP ответа (*)
		return
	}

	response, err := h.uc.GetBalance(r.Context(), walletID)
	if err != nil {
		// TODO: (*)
		return
	}
	
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		// TODO: (*)
		return
	}
}

