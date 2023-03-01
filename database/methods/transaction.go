package methods

import (
	"context"
	"fmt"
	"github.com/balance/api/database"
	"strconv"
)

func (p Postgres) TransactionBalance(ctx context.Context, toID, fromID int64, amount string) (string, string, error) {
	if toID <= 0 && fromID <= 0 {
		return "", "", fmt.Errorf("Wrong enter user id")
	}

	p.conn = database.ConnectDB()
	t_amount, _ := strconv.Atoi(amount)

	if toID <= 0 && t_amount <= 0 {
		lgzap.Fatal("ToID cannot less than zero,and amount")
	}

	if fromID <= 0 {
		lgzap.Fatal("FromID cannot less than zero")
	}

	err := p.DescreaseUserBalance(ctx, fromID, amount)
	if err != nil {
		lgzap.Error(err.Error() + "Failed to descreade user balance")
		return "", "", err
	}

	_, err = p.conn.Exec(ctx, "UPDATE Balance SET user_balance = user_balance + $1 WHERE user_id = $2", t_amount, toID)
	if err != nil {
		lgzap.Error(err.Error() + "transaction failed")
		return "", "", err
	}

	balanceToID, _, err := p.GetBalance(ctx, toID)
	conv_balanceToID := strconv.FormatInt(balanceToID, 36)
	if err != nil {
		lgzap.Error(err.Error() + "User to_id not found or balance less than zero")
		return "", "", err
	}

	balanceFromID, _, err := p.GetBalance(ctx, fromID)
	conv_balanceFromID := strconv.FormatInt(balanceFromID, 36)
	if err != nil {
		lgzap.Error(err.Error() + "User from_id not found or balance less than zero")
		return "", "", err
	}

	return conv_balanceToID, conv_balanceFromID, nil
}
