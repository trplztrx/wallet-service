package handler

import (
	"encoding/json"
	"net/http"

	// "wallet/internal/transport/dto"
	"wallet/internal/transport/dto"
	"wallet/internal/usecase"

	"github.com/shopspring/decimal"
)

type WalletOperationHandler struct {
	uc usecase.WalletOperationUsecase
}

func NewWalletOperationHandler(uc usecase.WalletOperationUsecase) *WalletOperationHandler {
	return &WalletOperationHandler{uc: uc}
}

func (h *WalletOperationHandler) Operation(w http.ResponseWriter, r *http.Request) {
	var req dto.WalletOperationRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request format")
		return
	}

	if req.OperationType != "DEPOSIT" && req.OperationType != "WITHDRAW" {
		respondWithError(w, http.StatusBadRequest, "Invalid operation type. Must be 'DEPOSIT' or 'WITHDRAW'")
        return
	}

	if req.Amount.LessThanOrEqual(decimal.Zero) {
        respondWithError(w, http.StatusBadRequest, "Amount must be a positive value")
        return
    }

	err = h.uc.Operation(r.Context(), &req)
	if err != nil {
		switch err.Error() {
        case "wallet not found":
            respondWithError(w, http.StatusNotFound, "Wallet not found")
		default:
            respondWithError(w, http.StatusInternalServerError, err.Error())
        }
        return
	}
	
	w.WriteHeader(http.StatusOK)
}