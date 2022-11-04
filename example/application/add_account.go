package application

import (
	"encoding/json"
	"net/http"
)

type AddAccRequest struct {
	UserID  int64
	Balance int64
}

func (a Application) AddAccount(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var req AddAccRequest
	err := decoder.Decode(&req)
	if err != nil {
		response500(w, err)
		return
	}
	u, err := a.users.Get(req.UserID)
	if err != nil {
		response500(w, err)
		return
	}
	acc, err := a.accounts.Add(*u, req.Balance)
	if err != nil {
		response500(w, err)
		return
	}
	response200(w, acc)
}
