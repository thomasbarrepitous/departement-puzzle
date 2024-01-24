package main

import (
	"database/sql"
	"departement/db"
	"departement/handlers"
	"departement/storage"
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

type Router struct {
	Router *mux.Router
}

func NewRouter(db *sql.DB) *Router {
	store := storage.NewPostgresStorage(db)

	r := mux.NewRouter()

	// Non protected routes

	// Handle 404
	notFoundHandler := &handlers.NotFoundHandler{}
	r.HandleFunc("/404", notFoundHandler.RenderNotFoundPage)

	// Login
	loginHandler := &handlers.LoginHandler{Store: store}
	r.HandleFunc("/login", loginHandler.RenderLoginPage)
	r.HandleFunc("/api/auth/login", loginHandler.JWTLoginHandle).Methods("POST")
	r.HandleFunc("/api/auth/google", loginHandler.GoogleLoginHandle).Methods("POST")
	r.HandleFunc("/api/auth/google/callback", loginHandler.GoogleCallbackHandle)
	r.HandleFunc("/api/auth/logout", loginHandler.LogoutHandle)

	// Registration
	registerHandler := &handlers.RegisterHandler{Store: store}
	r.HandleFunc("/api/users", registerHandler.RegisterHandle).Methods("POST")
	r.HandleFunc("/register", registerHandler.RenderRegisterPage)

	// Protected routes

	// Create subrouter for our protected routes
	protectedRouter := r.PathPrefix("/").Subrouter()
	protectedRouter.Use(utils.JWTVerifyMiddleware)

	// Create subrouter for our API routes
	apiRouter := r.PathPrefix("/api").Subrouter()
	apiRouter.Use(utils.JWTVerifyMiddleware)

	// Users
	userHandler := &handlers.UserHandler{Store: store}
	apiRouter.HandleFunc("/users", userHandler.GetAllUsers).Methods("GET")
	apiRouter.HandleFunc("/users/{id}", userHandler.GetUserByID).Methods("GET")

	// Profile
	profileHandler := &handlers.ProfileHandler{Store: store}
	protectedRouter.HandleFunc("/profile", profileHandler.RenderProfilePage)

	// Rankings
	rankingHandler := &handlers.RankingHandler{Store: store}
	apiRouter.HandleFunc("/rankings", rankingHandler.GetAllRankings).Methods("GET")
	apiRouter.HandleFunc("/rankings", rankingHandler.CreateRanking).Methods("POST")

	// Load our assets (css, js, images, etc.)
	staticRoute := http.StripPrefix("/static/", http.FileServer(neuteredFileSystem{http.Dir("./static")}))
	r.PathPrefix("/static/").Handler(staticRoute)

	// Game related routes
	gameHandler := &handlers.GameHandler{}
	protectedRouter.HandleFunc("/", gameHandler.RenderGamePage)

	return &Router{r}
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

	// Initialize the router
	r := NewRouter(db)

	port := ":3000"
	log.Print("Listening on port ", port)
	s := NewServer(port, r.Router)
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
