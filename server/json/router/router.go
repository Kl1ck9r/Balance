package router

import (
	"github.com/gorilla/mux"
	"github.com/balance/api/server/json/handlers"
	"net/http"
	"log"
)

func InitRouters(){
	router := mux.NewRouter()

	router.HandleFunc("/get",handlers.Get).Methods(http.MethodGet)

	log.Fatal(http.ListenAndServe(":8080", router))
}
