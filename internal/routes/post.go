package routes

import "net/http"

type PostsHandler struct {
}

func (handler *PostsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Bonjour le monde"))
}
