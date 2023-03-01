package exchange

import (
	"fmt"
	"github.com/balance/api/utils/convert"
	js_error "github.com/balance/api/utils/error"
	"net/http"
)

func Conv(curr, balance string, wrt http.ResponseWriter) (string, string, error) {
	var js js_error.JsonError
	if curr == "EURO" {
		euro, err := convert.ConvertToEuro(balance)
		if err != nil {
			js.WriteJsError(wrt, fmt.Errorf("Failed convert rubles to euro:%w", err), http.StatusBadRequest)
			return "", "", err
		}

		curr = "EUR"
		balance = euro
		return curr, balance, nil
	} else if curr == "Dollars" {
		dollars, err := convert.ConvertToDollars(balance)
		if err != nil {
			js.WriteJsError(wrt, fmt.Errorf("Failed convert rubles to dollars:%w", err), http.StatusBadRequest)
			return "", "", err
		}

		curr = "Dollars"
		balance = dollars
		return curr, balance, nil
	} else if curr == "German" {
		german, err := convert.ConvertToGerman(balance)
		if err != nil {
			js.WriteJsError(wrt, fmt.Errorf("Failed convert rubles to german:%v", err), http.StatusBadRequest)
			return "", "", err
		}

		curr = "German"
		balance = german
		return curr, balance, nil
	} else {
		curr = "RUB"
		rub_balance := balance
		return curr, rub_balance, nil
	}

	return "", "", nil
}
