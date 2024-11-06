package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"wallet/config"
	pgsql "wallet/infrastructure/db/repo"
	"wallet/internal/app"
	"wallet/internal/usecase"

	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	cfg, err := config.ReadConfig()
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	createWalletCmd := flag.NewFlagSet("create-wallet", flag.ExitOnError)

	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "create-wallet":
			createWalletCmd.Parse(os.Args[2:])
			createWallet(cfg)
			return
		}
	}

	app.Run(cfg)
}

func createWallet(cfg *config.Config) {
	ctx := context.Background()

	connString := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", cfg.User, cfg.Password ,cfg.Host, cfg.Port, cfg.DBConfig.DatabaseName)
	pool, err := pgxpool.New(ctx, connString)
	if err != nil {
		log.Fatalf("can't connect to postgresql: %v", err.Error())
	}
	defer pool.Close()

	walletRepo := pgsql.NewWalletRepo(pool)
	walletUsecase := usecase.NewWalletUsecase(walletRepo)

	wallet, err := walletUsecase.CreateWallet(ctx)
	if err != nil {
		log.Fatalf("Failed to create wallet: %v", err)
	}
	fmt.Printf("Wallet created with ID: %s\n", wallet.ID)
}