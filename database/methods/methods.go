package methods

import (
	"context"
	"github.com/balance/api/database"
	"github.com/balance/api/utils/zap"
	"github.com/jackc/pgx/v5"
)

type Postgres struct {
	conn *pgx.Conn

	Repository interface {
		GetBalance(ctx context.Context, userID int64)
		TransferBalance(ctx context.Context)

		ChangeCurrency(ctx context.Context)
		DeleteUserBalance(ctx context.Context, userID int64)
	}
}

var lgzap, _ = zap.InitLogger()

func (db Postgres) GetBalance(ctx context.Context, userID int64) (string, string, error) {
	db.conn = database.ConnectDB()
	getBalance := db.conn.QueryRow(ctx, "SELECT user_balance,currency FROM Balance WHERE user_id=$1", userID)

	var balance, currency string

	if err := getBalance.Scan(&balance, &currency); err != nil {
		lgzap.Error(err.Error() + "Unable to scan in variable")
		return "", "", err
	}

	return balance, currency, nil
}

func (db Postgres)TransferBalance(ctx context.Context){
	
}

func (db Postgres) DeleteUserBalance(ctx context.Context, userID int64) {

}
