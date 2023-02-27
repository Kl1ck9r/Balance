package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/balance/api/database/methods"
	js_error "github.com/balance/api/utils/error"
)

type GetUserBalance struct {
	User int64 `json:"id_user"`
	Currency string `json:"currency"` // you can choice currency,rubles or dollards
}

type BalanceUser struct {
	Ok       bool   `json:"ok"`       // true or false
	Currency string `json:"currency"` // only RUB
	Balance  string `json:"balance"`  // 200
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
	bl, cr, err := db.GetBalance(ctx, get.User)
	if err!=nil{
		js.WriteJsError(wrt, fmt.Errorf("Not Found user :%v", err), http.StatusBadRequest)
		return
	}

	fmt.Println("Balance and Currency: ",bl,cr)
	balance := &BalanceUser{
		Ok:       true,
		Currency: cr,
		Balance:  bl,
	}

	decode, err := json.MarshalIndent(&balance,"\n","\t")
	if err != nil {
		js.WriteJsError(wrt, fmt.Errorf("Marshal json:%v", err), http.StatusBadRequest)
		return
	}

	wrt.Write(decode)
}
