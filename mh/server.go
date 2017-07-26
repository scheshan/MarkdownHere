package mh

import (
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

//IServer the server interface
type IServer interface {
	Start() error
}

type server struct {
	host string
	port int
}

func (ser *server) Start() error {
	addr := ser.host + ":" + strconv.Itoa(ser.port)
	return http.ListenAndServe(addr, buildRouter())
}

func buildRouter() *httprouter.Router {
	router := httprouter.New()

	router.GET("/", indexPage)

	return router
}

//NewHTTPServer create a new server
func NewHTTPServer(host string, port int) IServer {
	if host == "" {
		host = "localhost"
	}
	if port < 1 {
		port = 5000
	}

	return &server{
		host: host,
		port: port}
}
