# wallet-service
## Сборка
1. Создание config.env файла по примеру из config.env.example в корне репозитория
2. docker-compose --env-file ./env-file build
3. docker-compose --env-file ./env-file up

## Создание кошелька (CLI)
docker exec wallet-app-1 /app/main create-wallet

## API
 - GET api/v1/wallets/{WALLET_UUID}
 - POST api/v1/wallet
  {
  walletId: UUID,
  operationType: DEPOSIT or WITHDRAW,
  amount: 1000
  }
