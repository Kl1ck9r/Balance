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

type UserDelete struct {
	UserID int64 `json:"user_id"`
}

type ResponseUserDelete struct {
	DeletedUser int64  `json:"user_id"`
	Balance     string `json:"balance"`
	Description string `json:"description"`
}

func DeleteBalance(wrt http.ResponseWriter, req *http.Request) {
	var js js_error.JsonError
	var deleteUser UserDelete
	err := json.NewDecoder(req.Body).Decode(&deleteUser)
	if err != nil {
		js.WriteJsError(wrt, fmt.Errorf("decoding json:%v", err), http.StatusBadRequest)
		return
	}

	if deleteUser.UserID <= 0 {
		js.WriteJsError(wrt, fmt.Errorf("User id cannot be negative :%v", err), http.StatusBadRequest)
		return
	}

	var db methods.Postgres
	err = db.DeleteBalance(context.Background(), deleteUser.UserID)
	if err != nil {
		js.WriteJsError(wrt, fmt.Errorf("Failed delete user [DB]:%v", err), http.StatusBadRequest)
		return
	}

	get_balance, _, err := db.GetBalance(context.Background(), deleteUser.UserID)
	conv_balance := strconv.FormatInt(get_balance, 36)

	response := &ResponseUserDelete{
		DeletedUser: deleteUser.UserID,
		Balance:     conv_balance,
		Description: "User with this id:" +
			fmt.Sprint(deleteUser.UserID) + "Success deleted",
	}

	bytes, err := json.MarshalIndent(&response, "\n", "\t")
	if err != nil {
		js.WriteJsError(wrt, fmt.Errorf("failed marshal struct :%v", err), http.StatusBadRequest)
		return
	}

	wrt.Write(bytes)
}
