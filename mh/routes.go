package mh

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func indexPage(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	res.Write([]byte("hello world"))
}
