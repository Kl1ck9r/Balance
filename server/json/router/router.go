package router

import (
	"github.com/balance/api/server/json/handlers"
	"github.com/balance/api/server/json/middleware"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func InitRouters() {
	router := mux.NewRouter()
	router.Use(middleware.JsonMiddleware)

	router.HandleFunc("/get", handlers.Get).Methods(http.MethodGet)
	router.HandleFunc("/replenish/balance", handlers.ReplenishBalance).Methods(http.MethodPost)
	router.HandleFunc("/descrease", handlers.DescreaseBalance).Methods(http.MethodPost)
	router.HandleFunc("/transaction", handlers.Transaction).Methods(http.MethodPost)
	router.HandleFunc("/delete", handlers.DeleteBalance).Methods(http.MethodDelete)

	log.Fatal(http.ListenAndServe(":8080", router))
}
