package server

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/diegodevtech/hexagonal-architecture/application"
	"github.com/urfave/negroni"
	"github.com/gorilla/mux"
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


	server := &http.Server{
		ReadHeaderTimeout: 10 * time.Second,
		WriteTimeout:  10 * time.Second,
		Addr: ":8080",
		Handler: http.DefaultServeMux,
		ErrorLog: log.New(os.Stderr, "log: ", log.Lshortfile),
	}
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

}