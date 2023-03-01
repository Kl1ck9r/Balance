package methods

import (
	"context"
	"github.com/balance/api/database"
	"strconv"
)

func (db Postgres) DescreaseUserBalance(ctx context.Context, userID int64, amount string) error {
	db.conn = database.ConnectDB()
	t_amount, _ := strconv.Atoi(amount)
	_, err := db.conn.Exec(ctx, "UPDATE Balance SET user_balance = user_balance - $1 WHERE user_id = $2", t_amount, userID)
	if err != nil {
		lgzap.Error(err.Error() + "Failed to descreade user balance")
		return err
	}

	return nil
}
