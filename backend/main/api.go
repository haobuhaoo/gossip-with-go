package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/haobuhaoo/gossip-with-go/internal/auth"
	"github.com/haobuhaoo/gossip-with-go/internal/comments"
	repo "github.com/haobuhaoo/gossip-with-go/internal/postgresql/sqlc"
	"github.com/haobuhaoo/gossip-with-go/internal/posts"
	"github.com/haobuhaoo/gossip-with-go/internal/topics"
	"github.com/haobuhaoo/gossip-with-go/internal/users"
	middleWare "github.com/haobuhaoo/gossip-with-go/middleware"
	"github.com/jackc/pgx/v5/pgxpool"
)

// application contains the configuration and database connection for the web server.
type application struct {
	config config
	db     *pgxpool.Pool
}

// config contains the server address string and database configurations.
type config struct {
	addr string
	db   dbConfig
}

// dbConfig contains the connection string for the PostgreSQL database.
type dbConfig struct {
	dsn string
}

// mount sets up the HTTP router, middleware, application routes.
// It returns a chi.Router that can be used by the HTTP server.
func (app *application) mount() http.Handler {
	r := chi.NewRouter()

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5173"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-type"},
		AllowCredentials: true,
		MaxAge:           300,
	}))
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	repo := repo.New(app.db)

	jwtSecret := os.Getenv("JWT_SECRET_KEY")
	authService := auth.NewService(repo)
	authHandler := auth.NewHandler(authService, jwtSecret)
	auth.Routes(r, authHandler)

	userService := users.NewService(repo)
	userHandler := users.NewHandler(userService)
	users.Routes(r, userHandler)

	r.Route("/api", func(r chi.Router) {
		r.Use(middleWare.JWTAuth(jwtSecret))

		r.Get("/me", authHandler.AuthenticateUser)

		topicService := topics.NewService(repo)
		topicHandler := topics.NewHandler(topicService)
		topics.Routes(r, topicHandler)

		postService := posts.NewService(repo)
		postHandler := posts.NewHandler(postService)
		posts.Routes(r, postHandler)

		commentService := comments.NewService(repo, app.db)
		commentHandler := comments.NewHandler(commentService)
		comments.Routes(r, commentHandler)
	})

	return r
}

// run starts the HTTP server with the given handler.
// It sets read, write, and idle timeouts and blocks until the server stops or an error occurs.
func (app *application) run(h http.Handler) error {
	svr := &http.Server{
		Addr:         app.config.addr,
		Handler:      h,
		WriteTimeout: time.Second * 30,
		ReadTimeout:  time.Second * 10,
		IdleTimeout:  time.Minute,
	}

	log.Printf("Starting server at Port %s", svr.Addr)

	return svr.ListenAndServe()
}
