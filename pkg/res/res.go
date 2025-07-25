package res

import (
	"encoding/json"
	"net/http"
)

func Json(w http.ResponseWriter, data any, StatusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(StatusCode)
	json.NewEncoder(w).Encode(data)
}
