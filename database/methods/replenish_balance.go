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
		GetBalance(ctx context.Context, userID int64) (int64, string, error)
		ReplenishBalance(ctx context.Context, userID int64, balance, currency string) error
		DeleteBalance(ctx context.Context,userID int64)(error)

		TransactionBalance(ctx context.Context,toID,fromID int64,amount string)(string, string, error)
		DescreaseUserBalance(ctx context.Context, userID int64,amount string)(error)
	}
}

var lgzap, _ = zap.InitLogger()

func (db Postgres) ReplenishBalance(ctx context.Context, userID int64, balance, currency string) error {
	db.conn = database.ConnectDB()
	_, _, exists := db.GetBalance(ctx, userID)
	if exists != nil {
		sqlInsert := `INSERT INTO Balance(user_balance,user_id,currency)
			VALUES($1,$2,$3)`
		_, err := db.conn.Exec(ctx, sqlInsert, balance, userID, currency)
		if err != nil {
			lgzap.Error(err.Error() + "Unable to insert balance user")
			return err
		}
	} else {
		_, err := db.conn.Exec(ctx, "UPDATE Balance SET user_balance=$1  WHERE user_id=$2", balance, userID)
		if err != nil {
			lgzap.Error(err.Error() + "Unable to update balance user")
			return err
		}
	}

	return nil
}

