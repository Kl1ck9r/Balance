package handlers

import (
	"github.com/ant0ine/go-json-rest/rest"
	"github.com/ant0ine/go-json-rest/rest/test"
	"log"
	"testing"
)

func TestListTransaction(t *testing.T) {
	api := rest.NewApi()
	api.Use(rest.DefaultDevStack...)

	request := &getTransaction{
		Limit:  2,
		UserID: 1,
	}

	router, err := rest.MakeRouter(
		rest.Get("/replenish/balance", func(w rest.ResponseWriter, r *rest.Request) {
			w.WriteJson(request)
		}),
	)

	if err != nil {
		log.Fatal(err)
	}

	api.SetApp(router)
	recorded := test.RunRequest(t, api.MakeHandler(),
		test.MakeSimpleRequest("Get", "http://localhost:8080/get/list/transactions", nil))
	recorded.CodeIs(200)
	recorded.ContentTypeIsJson()
}
