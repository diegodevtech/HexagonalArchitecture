package server

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/diegodevtech/hexagonal-architecture/adapters/web/handler"
	"github.com/diegodevtech/hexagonal-architecture/application"
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

type Webserver struct {
	Service application.ProductServiceInterface
}

func MakeNewWebserver() *Webserver {
	return &Webserver{}
}

func (w Webserver) Serve() {

	r := mux.NewRouter() //router
	n := negroni.New(
		negroni.NewLogger(),
	) //middleware de log

	handler.MakeProductHandlers(r, n, w.Service)
	http.Handle("/", r)


	server := &http.Server{
		ReadHeaderTimeout: 10 * time.Second,
		WriteTimeout:  10 * time.Second,
		Addr: ":9000",
		Handler: http.DefaultServeMux,
		ErrorLog: log.New(os.Stderr, "log: ", log.Lshortfile),
	}
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

}