package server

import (
	"net/http"
	"strconv"
)

type IHTTPServer interface {
	Start() error
}

type HTTPServer struct {
	Host string
	Port int
}

func (server *HTTPServer) Start() error {
	host := server.Host
	if host == "" {
		host = "0.0.0.0"
	}
	port := server.Port
	if port < 1 {
		port = 5000
	}

	addr := host + ":" + strconv.Itoa(port)

	return http.ListenAndServe(addr, &HttpHandler{})
}
