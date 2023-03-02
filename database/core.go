package database

import (
	"context"
	"fmt"
	"os"

	"github.com/balance/api/utils/zap"
	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
)

var lgzap, _ = zap.InitLogger()

func init() {
	if err := godotenv.Load("C:/Users/Ruslan/Desktop/BalanceAPI/.env"); err != nil {
		lgzap.Error(err.Error() + " :Not Found .env file")
	}
}

func ConnectDB() *pgx.Conn {
	ctx := context.Background()
	host := os.Getenv("DB_HOST")
	password := os.Getenv("DB_PASSWORD")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USERNAME")
	dbname := os.Getenv("DB_NAME")

	conn, err := pgx.Connect(ctx,
		fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
			host,
			port,
			user,
			password,
			dbname,
		),
	)

	if err != nil {
		lgzap.Error(err.Error() + " :Failed to connected db")
	}

	if err = conn.Ping(ctx); err != nil {
		lgzap.Error(err.Error() + " :Failed set connection to db")
	}

	lgzap.Info("Database success connected")

	return conn
}
