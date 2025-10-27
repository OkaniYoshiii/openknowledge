package routes

import (
	"html/template"
	"log"
	"net/http"
)

type PostsHandler struct {
	*template.Template
}

func (handler *PostsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := handler.Template.Execute(w, nil); err != nil {
		log.Fatal(err)
	}
}
