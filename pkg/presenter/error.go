package presenter

import (
	"encoding/json"
	"net/http"
)

func Error(w http.ResponseWriter, e error) {
	w.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(w).Encode(e.Error())
}

func ErrorNotFound(w http.ResponseWriter, e error) {
	w.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(w).Encode(e.Error())
}
