package router

import (
	"github.com/balance/api/server/json/handlers"
	"github.com/balance/api/server/json/middleware"
	"github.com/gorilla/mux"
	"net/http"
)

var APIRouter =  mux.NewRouter()

func Init() {
	APIRouter.Use(middleware.JsonMiddleware)

	APIRouter.HandleFunc("/get", handlers.GetBalance).Methods(http.MethodGet)
	APIRouter.HandleFunc("/replenish/balance", handlers.ReplenishBalance).Methods(http.MethodPost)
	APIRouter.HandleFunc("/descrease", handlers.DescreaseBalance).Methods(http.MethodPost)
	APIRouter.HandleFunc("/transaction", handlers.Transaction).Methods(http.MethodPost)
	APIRouter.HandleFunc("/delete", handlers.DeleteBalance).Methods(http.MethodDelete)
	APIRouter.HandleFunc("/get/list/transactions",handlers.GetTransaction).Methods(http.MethodGet)
}
