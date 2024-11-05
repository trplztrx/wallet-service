package app

import (
	"context"
	"fmt"
	"log"
	"time"
	"wallet/config"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jackc/pgx/v5/pgxpool"
)

func Run(cfg *config.Config) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	connString := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", cfg.User, cfg.Password ,cfg.Host, cfg.Port, cfg.Db.Db)
	pool, err := pgxpool.New(ctx, connString)
	defer pool.Close()
	if err != nil {
		log.Fatalf("can't connect to postgresql: %v", err.Error())
	}

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	
	// r.Post("/api/v1/wallet", walletOperationHandler.Operation())
	// r.Get("api/v1/wallet/{id}", walletHandler.GetBalance())
}