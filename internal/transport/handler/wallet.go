package handler

import (
	"encoding/json"
	"net/http"
	"wallet/internal/usecase"

	"github.com/go-chi/chi/v5"
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
	walletIDStr := chi.URLParam(r, "wallet_id")
	walletID, err := uuid.Parse(walletIDStr)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid wallet ID")
		return
	}

	response, err := h.uc.GetBalance(r.Context(), walletID)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Failed to get balance")
		return
	}
	
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Failed to encode response")
		return
	}
}

