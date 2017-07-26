package server

import (
	"net/http"
	"os"
)

type HttpHandler struct {
}

func (handler *HttpHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	dir, _ := os.Getwd()

	http.FileServer(http.Dir(dir)).ServeHTTP(w, r)

	//w.Write([]byte(r.URL.String()))
}
