package app

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"
	"wallet/config"
	pgsql "wallet/infrastructure/db/repo"
	"wallet/internal/transport/handler"
	"wallet/internal/usecase"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jackc/pgx/v5/pgxpool"
)

func Run(cfg *config.Config) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	connString := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", cfg.User, cfg.Password ,cfg.Host, cfg.Port, cfg.DBConfig.DatabaseName)
	pool, err := pgxpool.New(ctx, connString)
	if err != nil {
		log.Fatalf("can't connect to postgresql: %v", err.Error())
	}
	defer pool.Close()

	router := chi.NewRouter()
	router.Use(middleware.Logger)
	
	walletRepo := pgsql.NewWalletRepo(pool)
	transactionRepo := pgsql.NewTransactionRepo(pool)
	walletOperationRepo := pgsql.NewWalletOperationRepo(pool, walletRepo, transactionRepo)

	walletUsecase := usecase.NewWalletUsecase(walletRepo)
	walletOperationUsecase := usecase.NewWalletOperationUsecase(walletOperationRepo, walletRepo, transactionRepo)
	
	walletOperationHandler := handler.NewWalletOperationHandler(*walletOperationUsecase)
	walletHandler := handler.NewWalletHandler(*walletUsecase)

	router.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	router.Post("/api/v1/wallet", walletOperationHandler.Operation)
	router.Get("/api/v1/wallet/{wallet_id}", walletHandler.GetBalance)

	err = http.ListenAndServe(":8080", router)
	if err != nil {
		fmt.Println(err)
	}
}