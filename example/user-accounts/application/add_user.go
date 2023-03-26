package application

import (
	"encoding/json"
	"net/http"
)

type AddUserRequest struct {
	Name string
	Age  uint
}

func (a Application) AddUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		response405(w)
		return
	}
	decoder := json.NewDecoder(r.Body)
	var req AddUserRequest
	err := decoder.Decode(&req)
	if err != nil {
		response500(w, err)
		return
	}
	u, err := a.users.Add(req.Name, req.Age)
	if err != nil {
		response500(w, err)
		return
	}
	response200(w, u)
}

func response405(w http.ResponseWriter) {
	response(w, http.StatusMethodNotAllowed, "method not allowed")
}
