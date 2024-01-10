package main

import (
	"departement/limit"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/api/ranking", rankingHandler)
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./static")))

	log.Print("Listening on :3000")
	err := http.ListenAndServe(":3000", limit.Limit(r))
	if err != nil {
		log.Fatal(err)
	}
}

func rankingHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}
