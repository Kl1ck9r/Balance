package methods

import (
	"context"
	"fmt"

	"github.com/balance/api/database"
)

func (p Postgres) GetListTransaction(ctx context.Context, toID, limit int64) ([]ListTransaction, error) {

	if toID <= 0 {
		return []ListTransaction{}, fmt.Errorf("User id cannot is negative")
	}

	p.conn = database.ConnectDB()
	rows, err := p.conn.Query(ctx, "SELECT * FROM Transaction WHERE toID = $1 LIMIT $2", toID, limit)
	if err != nil {
		lgzap.Error(err.Error() + "Unable to get data from database")
		return []ListTransaction{}, err
	}

	var (
		model []ListTransaction
		ltr   ListTransaction
	)

	for rows.Next() {
		if err := rows.Scan(&ltr.ToID, &ltr.FromID, &ltr.Amount, &ltr.Description); err != nil {
			lgzap.Error(err.Error() + "Unable to scan field of struct")
			return []ListTransaction{}, err
		}

		model = append(model, ltr)
	}

	return model, nil
}
