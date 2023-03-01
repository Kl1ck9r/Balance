package methods

import (
	"context"
	"fmt"

	"github.com/balance/api/database"
	"strconv"
)

func (db Postgres) GetBalance(ctx context.Context, userID int64) (int64, string, error) {
	db.conn = database.ConnectDB()
	getBalance := db.conn.QueryRow(ctx, "SELECT user_balance,currency FROM Balance WHERE user_id=$1", userID)

	var balance string
	var currency string

	if err := getBalance.Scan(&balance, &currency); err != nil {
		lgzap.Error(err.Error() + "Unable to scan in variable")
		return 0, "", err
	}

	checkBalance, _ := strconv.ParseInt(balance, 36, 64)
	if checkBalance < 0 && checkBalance == 0 {
		return 0, "You don't have enought money", fmt.Errorf("Balance less than zero")
	}

	return checkBalance, currency, nil
}
