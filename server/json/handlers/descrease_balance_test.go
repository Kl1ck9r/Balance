package handlers

import (
	"github.com/ant0ine/go-json-rest/rest"
	"github.com/ant0ine/go-json-rest/rest/test"
	"log"
	"testing"
)

func TestDescreaseBalance(t *testing.T) {
	api := rest.NewApi()
	api.Use(rest.DefaultDevStack...)
	
	request := &DecsreaseBalance{
		UserID: 2,
		Amount: "560",
	}

	router, err := rest.MakeRouter(
		rest.Post("/descrease", func(w rest.ResponseWriter, r *rest.Request) {
			w.WriteJson(request)
		}),
	)

	if err != nil {
		log.Fatal(err)
	}

	api.SetApp(router)
	recorded := test.RunRequest(t, api.MakeHandler(),
		test.MakeSimpleRequest("POST", "http://localhost:8080/descrease", nil))
	recorded.CodeIs(200)
	recorded.ContentTypeIsJson()
}
