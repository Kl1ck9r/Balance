package handlers

import (
	"github.com/ant0ine/go-json-rest/rest"
	"github.com/ant0ine/go-json-rest/rest/test"
	"log"
	"testing"
)

func TestDeleteBalance(t *testing.T) {
	api := rest.NewApi()
	api.Use(rest.DefaultDevStack...)
	router, err := rest.MakeRouter(
		rest.Delete("/delete", func(w rest.ResponseWriter, r *rest.Request) {
			w.WriteJson(map[string]int64{"user_id": 1})
		}),
	)

	if err != nil {
		log.Fatal(err)
	}
	
	api.SetApp(router)
	recorded := test.RunRequest(t, api.MakeHandler(),
		test.MakeSimpleRequest("DELETE", "http://localhost:8080/delete", nil))
	recorded.CodeIs(200)
	recorded.ContentTypeIsJson()
}
