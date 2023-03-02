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

type transaction struct {
	ToID   int64  `json:"to_id"`
	FromID int64  `json:"from_id"`
	Amount string `json:"amount"`
	Descrption string `json:"description"`
}

type ResponseBalance struct {
	ToIDBalance   string `json:"to_id_balance"`
	FromIDBalance string `json:"from_id_balance"`
	Description   string `json:"description"`
}

func Transaction(wrt http.ResponseWriter, req *http.Request) {
	var js js_error.JsonError
	var transfer transaction
	err := json.NewDecoder(req.Body).Decode(&transfer)
	if err != nil {
		js.WriteJsError(wrt, fmt.Errorf("Failed decode: %v", err), http.StatusBadRequest)
		return
	}

	if transfer.ToID <= 0 {
		js.WriteJsError(wrt, fmt.Errorf("User id cannot be negative :%v", err), http.StatusBadRequest)
		return
	}

	if transfer.FromID <= 0 {
		js.WriteJsError(wrt, fmt.Errorf("User id cannot be negative :%v", err), http.StatusBadRequest)
		return
	}

	amountInt,_:=strconv.ParseInt(transfer.Amount,12,36)
	if amountInt <=0{
		js.WriteJsError(wrt, fmt.Errorf("User id cannot be negative :%v", err), http.StatusBadRequest)
		return
	}

	var db methods.Postgres
	toIDBalance, fromIDBalance, err := db.TransactionBalance(context.Background(), transfer.ToID, transfer.FromID, transfer.Amount)
	if err != nil {
		js.WriteJsError(wrt, fmt.Errorf("[DB ERROR]: %v", err), http.StatusBadRequest)
		return
	}

	bytes, err := json.Marshal(&ResponseBalance{
		ToIDBalance:   toIDBalance,
		FromIDBalance: fromIDBalance,
		Description:   "user balance after transaction: "+ toIDBalance,
	})

	if err != nil {
		js.WriteJsError(wrt, fmt.Errorf("Failed marshal struct: %v", err), http.StatusBadRequest)
		return
	}

	wrt.Write(bytes)
}
