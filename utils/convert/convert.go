package convert

import (
	"github.com/balance/api/utils/zap"
	"strconv"
)

var lgzap, _ = zap.InitLogger()

func ConvertToEuro(balance string) (string, error) {
	conv, err := strconv.ParseFloat(balance, 64)
	if err != nil {
		lgzap.Error(err.Error() + "Unable to convert rubles to euro")
		return "", err
	}
	euro := (conv / 79.62)
	return strconv.FormatFloat(euro, 'f', 2, 64), nil
}

func ConvertToDollars(balance string) (string, error) {
	conv, err := strconv.ParseFloat(balance, 64)
	if err != nil {
		lgzap.Error(err.Error() + "Unable to convert rubles to dollars")
		return "", err
	}
	dollars := (conv / 75.43)
	return strconv.FormatFloat(dollars, 'f', 2, 64), nil
}

func ConvertToGerman(balance string) (string, error) {
	conv, err := strconv.ParseFloat(balance, 64)
	var german float64 = (conv / 31.07196)
	if err != nil {
		lgzap.Error(err.Error() + "Unable to convert rubles to dollars")
		return "", err
	}
	return strconv.FormatFloat(german, 'f', 2, 64), nil
}
