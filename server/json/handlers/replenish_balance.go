package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/balance/api/database/methods"
	js_error "github.com/balance/api/utils/error"
	"github.com/balance/api/utils/zap"
	"net/http"
	"strconv"
)

var lgzap, _ = zap.InitLogger()

type replenishBalance struct {
	UserID   int64  `json:"user_id"`
	Amount   string `json:"balance"`
	Currency string `json:"currency"`
}

type responseUser struct {
	Balance     string `json:"balance"`
	Currency    string `json:"currency"`
	Description string `json:"description"`
}

func ReplenishBalance(wrt http.ResponseWriter, req *http.Request) {
	var refillBalance replenishBalance
	var js js_error.JsonError
	err := json.NewDecoder(req.Body).Decode(&refillBalance)
	if err != nil {
		js.WriteJsError(wrt, fmt.Errorf("Unable decode:%w", err), http.StatusBadRequest)
		return
	}

	balance, err := strconv.Atoi(refillBalance.Amount)
	if err != nil {
		lgzap.Error(err.Error() + "Failed to convert string in number")
		return
	}

	if balance <= 0 {
		js.WriteJsError(wrt, fmt.Errorf("The replenishment amount cannot be less than zero :%w", err), http.StatusBadRequest)
		return
	}

	var db methods.Postgres
	err = db.ReplenishBalance(context.Background(), refillBalance.UserID, refillBalance.Amount, refillBalance.Currency)
	if err != nil {
		js.WriteJsError(wrt, fmt.Errorf("Coudn't refill user balance :%w", err), http.StatusBadRequest)
		return
	}

	get_balance, currency, err := db.GetBalance(context.Background(), refillBalance.UserID)
	conv_balance := strconv.FormatInt(get_balance, 36)

	response := &responseUser{
		Balance:     conv_balance,
		Currency:    currency,
		Description: "replenishment amount: " + refillBalance.Amount,
	}

	bytes, err := json.MarshalIndent(&response, "\n", "\t")
	if err != nil {
		js.WriteJsError(wrt, fmt.Errorf("Unable marshal struct :%w", err), http.StatusBadRequest)
		return
	}

	wrt.Write(bytes)
}
