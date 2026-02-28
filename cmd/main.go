package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Post("/eguro/{basta}", getNegra)

	http.ListenAndServe(":3000", r)
}

func getNegra(w http.ResponseWriter, r *http.Request) {
	basta := chi.URLParam(r, "basta")
	w.Write([]byte(fmt.Sprintf("Eguro, ahi va: %s", basta)))
	return
}
