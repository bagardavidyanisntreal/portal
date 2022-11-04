package application

import (
	"fmt"
	"log"
	"net/http"
)

func response200(w http.ResponseWriter, msg any) {
	response(w, http.StatusOK, msg)
}

func response500(w http.ResponseWriter, err error) {
	response(w, http.StatusInternalServerError, err.Error())
}

func response(w http.ResponseWriter, code int, msg any) {
	w.WriteHeader(code)
	_, err := fmt.Fprint(w, msg)
	if err != nil {
		log.Printf("[RESPONSE]: %s", err)
	}
}
