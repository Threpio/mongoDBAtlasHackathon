package user

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func (ac *AuthController) HLogin(w http.ResponseWriter, r *http.Request) {

	var authAttempt AuthAttempt
	defer r.Body.Close()

	body, err := ioutil.ReadAll(r.Body)
	err = json.Unmarshal(body, &authAttempt)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	sessionToken, err := ac.LoginUser(authAttempt)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, sessionToken)
	return
}

func (ac *AuthController) HRegister(w http.ResponseWriter, r *http.Request) {

	var newUserAttempt NewUserAttempt
	defer r.Body.Close()

	body, err := ioutil.ReadAll(r.Body)
	err = json.Unmarshal(body, &newUserAttempt)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	sessionToken, err := ac.RegisterUser(newUserAttempt)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, sessionToken)
	return
}