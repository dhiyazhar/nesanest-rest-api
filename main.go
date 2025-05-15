package main

import (
	"net/http"
)

func setJsonResp(message []byte, httpCode int, res http.ResponseWriter) {
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(httpCode)
	res.Write(message)
}

type homeHandler struct{}

func (h *homeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	message := []byte(`{"message": "Golang REST API - Azhar 2025"}`)
	setJsonResp(message, http.StatusOK, w)
}

func main() {
	mux := http.NewServeMux()

	mux.Handle("/", &homeHandler{})

	http.ListenAndServe(":8080", mux)
}
