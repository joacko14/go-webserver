package nativeServer

import (
	"fmt"
	"net/http"
)

const port = ":80"

func Start() {
	fmt.Println("Starting server on port", port)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var method = r.Method
		var name = "Anonymous"
		switch method {
		case http.MethodGet:
			name = r.URL.Query().Get("name")
		case http.MethodPost:
			name = r.FormValue("name")
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
		fmt.Fprintf(w, "Hello, %s!", name)
		fs := http.FileServer(http.Dir("static/"))
		http.Handle("/static/", http.StripPrefix("/static/", fs))
	})
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "OK")
	})
	http.ListenAndServe(port, nil)
}
