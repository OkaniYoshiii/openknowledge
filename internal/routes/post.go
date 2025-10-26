package routes

import (
	"log"
	"net/http"
)

type PostsHandler struct {
	Config
}

func (handler *PostsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	tmpl := handler.BaseTemplate

	if err := tmpl.Execute(w, nil); err != nil {
		log.Fatal(err)
	}
}
