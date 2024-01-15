package main

import (
	"departement/db_config"
	"departement/limit"
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"time"

	"github.com/gorilla/mux"
)

func getRanking(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "OK\n")
}

func postRanking(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "OK\n")
}

func getAPIHealth(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "OK\n")
}

func main() {
	// Initialize the database
	db, err := db_config.ConnectDB(*db_config.NewDBConfig())
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Test the database connection
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	r := mux.NewRouter()

	// Create subrouter for our API routes
	apiRouter := r.PathPrefix("/api").Subrouter()

	// Health check
	apiRouter.HandleFunc("/health", getAPIHealth).Methods("GET")

	// Ranking
	apiRouter.HandleFunc("/ranking/", getRanking).Methods("GET")
	apiRouter.HandleFunc("/ranking/", postRanking).Methods("POST")

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
		Addr:         ":3000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
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
