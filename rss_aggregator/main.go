package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/0xOthmane/golang_projects/rss_aggregator/internal/database"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"

	_ "github.com/lib/pq"
)

type apiConfig struct {
	DB *database.Queries
}

func main() {
	godotenv.Load() // Load .env file
	portString := os.Getenv("PORT")
	dbURL := os.Getenv("DB_URL")
	if portString == "" {
		log.Fatal("PORT not found")
	}
	if dbURL == "" {
		log.Fatal("DB_URL not found")
	}
	conn, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal("Unable to connect to the db.")
	}
	db := database.New(conn)

	apiCfg := apiConfig{
		DB: db,
	}
	router := chi.NewRouter()
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))
	v1Router := chi.NewRouter()
	// v1Router.HandleFunc("/healthz", handlerReadiness)
	v1Router.Get("/healthz", handlerReadiness)
	v1Router.Get("/error", handlerError)
	v1Router.Post("/users", apiCfg.handlerCreateUser)
	v1Router.Get("/users", apiCfg.middlewareAuth(apiCfg.handlerGetUserByAPIKey))
	v1Router.Post("/feeds", apiCfg.middlewareAuth(apiCfg.handlerCreateFeed))
	v1Router.Get("/feeds", apiCfg.handlerGetFeeds)
	v1Router.Post("/feed_follow",apiCfg.middlewareAuth(apiCfg.handlerCreateFeedFollow))
	router.Mount("/v1", v1Router)
	srv := &http.Server{
		Handler: router,
		Addr:    ":" + portString,
	}
	log.Printf("Server listening on port %v", portString)
	errSrv := srv.ListenAndServe()
	if errSrv != nil {
		log.Fatal(err)
	}
	fmt.Println("PORT:", portString)
}
