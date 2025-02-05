package handlers

import (
	"fmt"
	"net/http"
)

func ArticlesHandler(w http.ResponseWriter, r *http.Request) {
	// vars := mux.Vars(r)
	fmt.Fprintf(w, "Articles!")
}
