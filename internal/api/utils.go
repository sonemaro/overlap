package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func GetJsonBody(r *http.Request, out interface{}) error {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}
	return json.Unmarshal(body, out)
}

func SendCreated(w http.ResponseWriter, emptyPayload bool) {
	resp := struct {
		Message string `json:"message"`
	}{
		Message: "resource created successfully",
	}

	w.WriteHeader(http.StatusCreated)
	if emptyPayload {
		w.Write(nil)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func SendBadRequest(w http.ResponseWriter) {
	resp := struct {
		Error string `json:"error"`
	}{
		Error: "bad request",
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(w).Encode(resp)
}

func SendJson(w http.ResponseWriter, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(payload)
}