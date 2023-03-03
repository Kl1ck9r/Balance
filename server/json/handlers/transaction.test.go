package handlers

import (
	"github.com/ant0ine/go-json-rest/rest"
	"github.com/ant0ine/go-json-rest/rest/test"
	"log"
	"testing"
)

func TestTransaction(t *testing.T) {
	api := rest.NewApi()
	api.Use(rest.DefaultDevStack...)

	request := &transaction{
		ToID:       1,
		FromID:     2,
		Amount:     "8300",
		Descrption: "Amount has been sent succusses",
	}

	router, err := rest.MakeRouter(
		rest.Post("/transaction", func(w rest.ResponseWriter, r *rest.Request) {
			w.WriteJson(request)
		}),
	)

	if err != nil {
		log.Fatal(err)
	}

	api.SetApp(router)
	recorded := test.RunRequest(t, api.MakeHandler(),
		test.MakeSimpleRequest("POST", "http://localhost:8080/transaction", nil))
	recorded.CodeIs(200)
	recorded.ContentTypeIsJson()
}
