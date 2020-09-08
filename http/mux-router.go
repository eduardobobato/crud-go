package router

import (
	"fmt"
	"net/http"

	"github.com/go-openapi/runtime/middleware"
	"github.com/gorilla/mux"
)

type muxRouter struct{}

var (
	muxDispatcher = mux.NewRouter()
)

// NewMuxRouter : Create a new Router
func NewMuxRouter() Router {
	return &muxRouter{}
}

func (*muxRouter) GET(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	muxDispatcher.HandleFunc(uri, f).Methods("GET")
}

func (*muxRouter) POST(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	muxDispatcher.HandleFunc(uri, f).Methods("POST")
}

func (*muxRouter) PUT(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	muxDispatcher.HandleFunc(uri, f).Methods("PUT")
}

func (*muxRouter) DELETE(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	muxDispatcher.HandleFunc(uri, f).Methods("DELTE")
}

func (*muxRouter) SWAGGER(uri string, yaml string) {
	opts := middleware.RedocOpts{SpecURL: yaml}
	sh := middleware.Redoc(opts, nil)
	muxDispatcher.Handle(yaml, http.FileServer(http.Dir("./")))
	muxDispatcher.Handle(uri, sh)
}

func (*muxRouter) SERVE(port string) {
	fmt.Printf("Server run in port " + port)
	http.ListenAndServe(port, muxDispatcher)
}
