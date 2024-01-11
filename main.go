package main

import (
	"departement/limit"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	// Create subrouter for our API routes
	apiRouter := r.PathPrefix("/api").Subrouter()
	apiRouter.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(map[string]bool{"ok": true})
	})
	apiRouter.HandleFunc("/ranking/{category}/{trois}/", rankingHandler)
	// Load our assets (css, js, images, etc.)
	staticRoute := http.StripPrefix("/static/", http.FileServer(neuteredFileSystem{http.Dir("./static")}))
	r.PathPrefix("/static/").Handler(staticRoute)
	// Load on root path our index.html
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./static/index.html")
	})

	log.Print("Listening on :3000")
	srv := &http.Server{
		Handler:      limit.Limit(r),
		Addr:         "127.0.0.1:3000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}

func rankingHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Category: %v, %v\n", vars["category"], vars["trois"])
}

// neuteredFileSystem is a custom implementation of http.FileSystem
// that disables directory listings.
// https://www.alexedwards.net/blog/disable-http-fileserver-directory-listings
type neuteredFileSystem struct {
	fs http.FileSystem
}

func (nfs neuteredFileSystem) Open(path string) (http.File, error) {
	f, err := nfs.fs.Open(path)
	if err != nil {
		return nil, err
	}

	s, err := f.Stat()
	if s.IsDir() {
		index := filepath.Join(path, "index.html")
		if _, err := nfs.fs.Open(index); err != nil {
			closeErr := f.Close()
			if closeErr != nil {
				return nil, closeErr
			}

			return nil, err
		}
	}

	return f, nil
}
