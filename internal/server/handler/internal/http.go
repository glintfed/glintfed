package internal

import (
	"encoding/json"
	"errors"
	"net/http"

	"glintfed/internal/lib/liberrs"
)

func WriteJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(v)
}

func WriteActivityJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/activity+json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(v)
}

func WriteJSONWithCORS(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	WriteJSON(w, status, v)
}

func WriteError(w http.ResponseWriter, err error) {
	var validationErr *ValidationError
	if errors.As(err, &validationErr) {
		WriteJSON(w, http.StatusBadRequest, validationErr)
		return
	}

	if errors.Is(err, liberrs.Todo) {
		http.Error(w, http.StatusText(http.StatusNotImplemented), http.StatusNotImplemented)
		return
	}

	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}
