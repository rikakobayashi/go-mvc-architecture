package presenter

import (
	"encoding/json"
	"net/http"
)

func Response(w http.ResponseWriter, body interface{}) {
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(body)
	if err != nil {
		Error(w, err)
		return
	}
}
