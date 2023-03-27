package application

import (
	"fmt"
	"log"
	"net/http"
)

func response200(w http.ResponseWriter, msg any) {
	response(w, http.StatusOK, msg)
}

func response400(w http.ResponseWriter, msg any) {
	response(w, http.StatusBadRequest, msg)
}

func response404(w http.ResponseWriter) {
	response(w, http.StatusNotFound, "not found")
}

func response(w http.ResponseWriter, code int, msg any) {
	w.WriteHeader(code)
	_, err := fmt.Fprint(w, msg)
	if err != nil {
		log.Printf("[response error]: %s", err)
	}
}
