package handlers

import (
	"fmt"
	"net/http"
)

func ProductsHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Products!")
}
