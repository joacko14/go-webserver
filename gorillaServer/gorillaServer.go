package gorillaserver

import (
	"fmt"
	"net/http"
	"webapp/gorillaServer/handlers"

	"github.com/gorilla/mux"
)

const port = ":80"

func Start() {
	r := mux.NewRouter()
	r.HandleFunc("/", handlers.HomeHandler)
	r.HandleFunc("/product/{prodcutId}", handlers.ProductsHandler)
	r.HandleFunc("/articles/{author}", handlers.ArticlesHandler)
	http.Handle("/", r)
	http.ListenAndServe(port, nil)
	fmt.Println("Gorilla server started")
}
