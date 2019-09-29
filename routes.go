package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Request struct {
	Image string `json:"image"`
}

func events(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		json.NewEncoder(w).Encode(err.Error())
		return
	}
	data, _ := ioutil.ReadAll(r.Body)
	var request Request
	request.Image = definition(data)
	json.NewEncoder(w).Encode(request)
}

func Start() {
	r := mux.NewRouter()
	r.HandleFunc("/recognize", events)
	err := http.ListenAndServe(":21000", r)
	log.Println(err)
}
