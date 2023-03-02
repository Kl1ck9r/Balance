package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/balance/api/database/methods"
	js_error "github.com/balance/api/utils/error"
)

type DecsreaseBalance struct {
	UserID int64  `json:"user_id"`
	Amount string `json:"amount"`
}

type ResponseUserBalance struct {
	Balance     string `json:"balance"`
	Currency    string `json:"currency"`
	Description string `json:"description"`
}

func DescreaseBalance(wrt http.ResponseWriter, req *http.Request) {
	var js js_error.JsonError
	var descrease DecsreaseBalance
	err := json.NewDecoder(req.Body).Decode(&descrease)
	if err != nil {
		js.WriteJsError(wrt, fmt.Errorf("Unable decode struct:%v ", err), http.StatusBadRequest)
		return
	}

	if descrease.UserID <= 0 {
		js.WriteJsError(wrt, fmt.Errorf("User id cannot be negative :%v", err), http.StatusBadRequest)
		return
	}

	var db methods.Postgres
	err = db.DescreaseUserBalance(context.Background(), descrease.UserID, descrease.Amount)
	if err != nil {
		js.WriteJsError(wrt, fmt.Errorf("Failed to descreare user balance: %v ", err), http.StatusNotFound)
		return
	}

	balance, currency, err := db.GetBalance(context.Background(), descrease.UserID)
	if err != nil {
		js.WriteJsError(wrt, fmt.Errorf("User not found: %v ", err), http.StatusNotFound)
		return
	}

	strc_balance := strconv.FormatInt(balance, 36)
	reponse := &ResponseUserBalance{
		Balance:     strc_balance,
		Currency:    currency,
		Description: "Amount writter from account:" + descrease.Amount,
	}

	bytes, err := json.MarshalIndent(&reponse, "\n", "\t")
	if err != nil {
		js.WriteJsError(wrt, fmt.Errorf("Failed to get reponse: %v ", err), http.StatusBadRequest)
		return
	}

	wrt.Write(bytes)
}
