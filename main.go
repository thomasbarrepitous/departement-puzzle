package main

import (
	"context"
	"departement/handlers"
	"departement/storage"
	"departement/utils"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

type Server struct {
	Server *http.Server
}

func NewServer(listenAddr string, ctx context.Context, storage *storage.Storage) *Server {
	r := NewRouter(ctx, storage).Router
	return &Server{
		Server: &http.Server{
			Addr:         listenAddr,
			Handler:      utils.Limit(r),
			WriteTimeout: 30 * time.Second,
			ReadTimeout:  30 * time.Second,
		},
	}
}

type Router struct {
	Router *mux.Router
}

func NewRouter(ctx context.Context, store *storage.Storage) *Router {
	r := mux.NewRouter()

	// TODO: Implement OAuth2 with Firebase ?
	// If so, refactor since it's probably not a "storage"
	// but rather a "service"
	// authStore := storage.NewFirebaseStorage(ctx)

	homeHandler := &handlers.HomeHandler{}
	notFoundHandler := &handlers.NotFoundHandler{}
	loginHandler := &handlers.LoginHandler{UserStore: store.Users, ProfileStore: store.Profiles}
	rankingHandler := &handlers.RankingHandler{RankingStore: store.Rankings}
	registerHandler := &handlers.RegisterHandler{UserStore: store.Users, ProfileStore: store.Profiles}
	userHandler := &handlers.UserHandler{UserStore: store.Users}
	profileHandler := &handlers.ProfileHandler{ProfileStore: store.Profiles, RankingStore: store.Rankings}
	gameHandler := &handlers.GameHandler{}
	playMenuHandler := &handlers.PlayMenuHandler{}

	// Non protected routes

	// Home
	r.HandleFunc("/", homeHandler.RenderHomePage)

	// Handle 404
	r.NotFoundHandler = http.HandlerFunc(notFoundHandler.RenderNotFoundPage)
	r.HandleFunc("/404", notFoundHandler.RenderNotFoundPage)

	// Login
	r.HandleFunc("/login", loginHandler.RenderLoginPage)
	r.HandleFunc("/api/auth/login", loginHandler.EmailLoginHandle).Methods("POST")
	r.HandleFunc("/api/auth/google", loginHandler.GoogleLoginHandle).Methods("POST")
	r.HandleFunc("/api/auth/google/callback", loginHandler.GoogleCallbackHandle)
	r.HandleFunc("/api/auth/logout", loginHandler.LogoutHandle)

	// Registration
	r.HandleFunc("/api/users", registerHandler.RegisterHandle).Methods("POST")
	r.HandleFunc("/register", registerHandler.RenderRegisterPage)

	// Play menu
	r.HandleFunc("/play", playMenuHandler.RenderPlayMenuPage)

	// Protected routes

	// Create subrouter for our protected routes
	protectedRouter := r.PathPrefix("/").Subrouter()
	protectedRouter.Use(utils.JWTVerifyMiddleware)

	// Create subrouter for our API routes
	apiRouter := r.PathPrefix("/api").Subrouter()
	apiRouter.Use(utils.JWTVerifyMiddleware)

	// Users
	apiRouter.HandleFunc("/users", userHandler.GetAllUsers).Methods("GET")
	apiRouter.HandleFunc("/users/{id}", userHandler.GetUserByID).Methods("GET")

	// Profile
	protectedRouter.HandleFunc("/profile", profileHandler.RenderProfilePage)

	// Rankings
	apiRouter.HandleFunc("/rankings", rankingHandler.GetAllRankings).Methods("GET")
	apiRouter.HandleFunc("/rankings", rankingHandler.CreateRanking).Methods("POST")

	// Load our assets (css, js, images, etc.)
	staticRoute := http.StripPrefix("/static/", http.FileServer(neuteredFileSystem{http.Dir("./static")}))
	r.PathPrefix("/static/").Handler(staticRoute)

	// Game related routes
	protectedRouter.HandleFunc("/departement", gameHandler.RenderGamePage)

	return &Router{r}
}

func main() {
	// Initialize env variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Initialize the postgres database
	store := storage.NewStorage(os.Getenv("ENV_TYPE"))

	// defer store.DB.Close()

	port := ":3000"
	log.Print("Listening on port ", port)
	s := NewServer(port, context.Background(), store)
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
