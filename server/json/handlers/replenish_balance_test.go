package handlers

import (
	"github.com/ant0ine/go-json-rest/rest"
	"github.com/ant0ine/go-json-rest/rest/test"
	"log"
	"testing"
)

func TestReplenishBalance(t *testing.T) {
	api := rest.NewApi()
	api.Use(rest.DefaultDevStack...)

	request := &replenishBalance{
		UserID:   1,
		Amount:   "12905",
		Currency: "RUB",
	}

	router, err := rest.MakeRouter(
		rest.Post("/replenish/balance", func(w rest.ResponseWriter, r *rest.Request) {
			w.WriteJson(request)
		}),
	)

	if err != nil {
		log.Fatal(err)
	}

	api.SetApp(router)
	recorded := test.RunRequest(t, api.MakeHandler(),
		test.MakeSimpleRequest("POST", "http://localhost:8080/replenish/balance", nil))
	recorded.CodeIs(200)
	recorded.ContentTypeIsJson()
}
