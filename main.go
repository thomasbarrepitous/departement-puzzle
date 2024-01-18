package main

import (
	"departement/db"
	"departement/handlers"
	"departement/utils"
	"log"
	"net/http"
	"path/filepath"
	"time"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

type Server struct {
	Server *http.Server
}

func NewServer(listenAddr string, r *mux.Router) *Server {
	return &Server{
		Server: &http.Server{
			Addr:         listenAddr,
			Handler:      utils.Limit(r),
			WriteTimeout: 15 * time.Second,
			ReadTimeout:  15 * time.Second,
		},
	}
}

func main() {
	// Initialize env variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Initialize the database
	db, err := db.ConnectDB(*db.NewDBConfig())
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	r := mux.NewRouter()

	// Create subrouter for our API routes
	apiRouter := r.PathPrefix("/api").Subrouter()

	// Users
	userHandler := &handlers.UserHandler{DB: db}
	apiRouter.HandleFunc("/users", userHandler.GetAllUsers).Methods("GET")
	apiRouter.HandleFunc("/users/{id}", userHandler.GetUserByID).Methods("GET")
	apiRouter.HandleFunc("/users", userHandler.CreateUser).Methods("POST")

	// Rankings
	rankingHandler := &handlers.RankingHandler{DB: db}
	apiRouter.HandleFunc("/rankings", rankingHandler.GetAllRankings).Methods("GET")
	apiRouter.HandleFunc("/rankings", rankingHandler.CreateRanking).Methods("POST")

	// Load our assets (css, js, images, etc.)
	staticRoute := http.StripPrefix("/static/", http.FileServer(neuteredFileSystem{http.Dir("./static")}))
	r.PathPrefix("/static/").Handler(staticRoute)

	// Load on root path our index.html
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./static/index.html")
	})

	port := ":3000"
	log.Print("Listening on port ", port)
	s := NewServer(port, r)
	log.Fatal(s.Server.ListenAndServe())
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
