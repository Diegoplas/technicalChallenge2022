package controller

import (
	"encoding/json"
	"net/http"
)

func HelloWize(w http.ResponseWriter, r *http.Request) {
	helloString := "Hello Wizeline"
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(helloString)
}
