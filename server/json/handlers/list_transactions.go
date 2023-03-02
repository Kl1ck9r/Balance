package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/balance/api/database/methods"
	js_error "github.com/balance/api/utils/error"
)

type getTransaction struct {
	Limit  int64 `json:"limit"`
	UserID int64 `json:"user_id"`
}

func GetTransaction(wrt http.ResponseWriter, req *http.Request) {
	var js js_error.JsonError
	var transaction getTransaction
	err := json.NewDecoder(req.Body).Decode(&transaction)
	if err != nil {
		js.WriteJsError(wrt, fmt.Errorf("decoding json:%v", err), http.StatusBadRequest)
		return
	}

	if transaction.UserID <= 0 {
		js.WriteJsError(wrt, fmt.Errorf("Wrong enter id,%v", err), http.StatusNotFound)
		return
	}

	var db methods.Postgres
	model, err := db.GetListTransaction(context.Background(), transaction.UserID, transaction.Limit)
	if err != nil {
		js.WriteJsError(wrt, fmt.Errorf("Error get list transactions:%v", err), http.StatusBadRequest)
		return
	}

	bytes, err := json.Marshal(&model)
	if err != nil {
		js.WriteJsError(wrt, fmt.Errorf("Error in mershaling struct:%v", err), http.StatusBadRequest)
		return
	}

	wrt.Write(bytes)
}
