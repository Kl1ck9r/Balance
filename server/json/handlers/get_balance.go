package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/balance/api/database/methods"
	"github.com/balance/api/exchange"
	js_error "github.com/balance/api/utils/error"
)

type GetUserBalance struct {
	UserID   int64  `json:"id_user"`
	Currency string `json:"currency"`
}

type BalanceUser struct {
	Ok       bool   `json:"ok"`       // true or false
	Currency string `json:"currency"` // we have only EURO,DOLLARS,and GERMAN currency
	Balance  string `json:"balance"`  // balance user(200$,3000₽,5000₴)
}

func Get(wrt http.ResponseWriter, req *http.Request) {
	var js js_error.JsonError
	var get GetUserBalance
	err := json.NewDecoder(req.Body).Decode(&get)
	if err != nil {
		js.WriteJsError(wrt, fmt.Errorf("decoding json:%v", err), http.StatusBadRequest)
		return
	}

	ctx := context.Background()
	var db methods.Postgres
	bl, cr, err := db.GetBalance(ctx, get.UserID)
	if err != nil {
		js.WriteJsError(wrt, fmt.Errorf("Not Found user :%w", err), http.StatusBadRequest)
		return
	}

	conv_balance := strconv.FormatInt(bl, 36)

	balance := &BalanceUser{
		Ok:       true,
		Currency: cr,
		Balance:  conv_balance,
	}

	newCurrency, newBalance, err := exchange.Conv(get.Currency, balance.Balance, wrt) // this method will convert currency if will want user
	if err != nil {
		js.WriteJsError(wrt, fmt.Errorf("Unable to convert in another currency :%w", err), http.StatusBadRequest)
		return
	}

	balance.Currency = newCurrency
	balance.Balance = newBalance

	decode, err := json.MarshalIndent(&balance, "\n", "\t")
	if err != nil {
		js.WriteJsError(wrt, fmt.Errorf("Marshal json:%v", err), http.StatusBadRequest)
		return
	}

	wrt.Write(decode)
}
