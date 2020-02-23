package api

import (
	"github.com/gorilla/mux"
	"github.com/sonemaro/overlap/internal/rectangle"
	"log"
	"net/http"
)


func SetupRouter() *mux.Router{
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", postRect).Methods("POST")
	router.HandleFunc("/", getRects).Methods("GET")
	router.HandleFunc("/unique", getUnique).Methods("GET")
	return router
}

type PostRectRequest struct {
	Main rectangle.Rectangle `json:"main"`
	Input []rectangle.Rectangle `json:"input"`
}

func postRect(w http.ResponseWriter, r *http.Request) {
	var j PostRectRequest
	err := GetJsonBody(r, &j)
	if err != nil {
		log.Print(err)
		SendBadRequest(w)
		return
	}
	overlaps := rectangle.GetOverlapInArray(&j.Main, j.Input)
	for _, r := range overlaps {
		rectangle.DB.Save(r)
	}
	log.Printf("input: %#v | main: %#v", j.Input, j.Main)
	log.Printf("overlaps: %v", overlaps)

	SendCreated(w, true)
}

func getRects(w http.ResponseWriter, r *http.Request) {
	SendJson(w, rectangle.DB.GetAll())
}

func getUnique(w http.ResponseWriter, r *http.Request) {
	rect, isNull := rectangle.DB.GetFirst()
	if isNull {
		SendJson(w, []struct{}{})
		return
	}
	SendJson(w, rect)
}
