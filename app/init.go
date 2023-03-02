package app

import (
	"flag"
	"log"
	"net/http"
	_ "github.com/balance/api/database"
	"time"

	"github.com/balance/api/server/json/router"
)

func Start() {
	router.Init()

	HTTP := flag.Bool("http", true, "run server http")
	key := flag.String("key", "private.key", "key PEM file")
	crt := flag.String("crt", "certificate.crt", "certificate PEM file")
	flag.Parse()

	api := &http.Server{
		Handler:      router.APIRouter,
		Addr:        "localhost:8080", 
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  time.Second * 60,
	}

	if *HTTP {
		log.Fatal(api.ListenAndServe())
	} else {
		log.Fatal(api.ListenAndServeTLS(*crt, *key))
	}

}
