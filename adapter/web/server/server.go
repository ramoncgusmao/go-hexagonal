package server

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/ramoncgusmao/go-hexagonal/adapter/web/handler"
	"github.com/ramoncgusmao/go-hexagonal/application"
)

type Webserver struct {
	Service application.ProductServiceInterface
}

func MakeNewwebserver() *Webserver {
	return &Webserver{}
}

func (w Webserver) Serve() {

	router := mux.NewRouter()
	n := negroni.New(
		negroni.NewLogger(),
	)

	handler.MakeProductHandler(router, n, w.Service)
	http.Handle("/", router)

	server := &http.Server{
		ReadHeaderTimeout: 10 * time.Second,
		WriteTimeout:      10 * time.Second,
		Addr:              ":9000",
		Handler:           nil,
		ErrorLog:          log.New(os.Stderr, "log: ", log.Lshortfile),
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
