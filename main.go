package main

import (
	"departement/components"
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.Handle("/staticfiles/",
		http.StripPrefix("/staticfiles/", http.FileServer(http.Dir("./static"))))
	http.Handle("/jsfiles/",
		http.StripPrefix("/jsfiles/", http.FileServer(http.Dir("./js"))))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		components.Layout().Render(r.Context(), w)
	})

	fmt.Println("Listening on :3000")
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatal(err)
	}
}
