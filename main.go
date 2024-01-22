package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/DatNgo-dev/rssaggregator/internal/database"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"

	_ "github.com/lib/pq"
)

type apiConfig struct {
	DB *database.Queries
}

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal(err)
	} 

	PORT := os.Getenv("PORT")
	if PORT == "" {
		log.Fatal("PORT is not found in the environment")
	}
	fmt.Println("PORT: ", PORT)

	dbUrl := os.Getenv("DB_URL")
	if dbUrl == "" {
		log.Fatal("DB_URL is not found in the environment")
	}
	fmt.Println("DB Url: ", dbUrl)

	conn, err := sql.Open("postgres", dbUrl)
	if err != nil {
		log.Fatal("Can't connect to database")
	}

	queries := database.New(conn)

	apiCfg := apiConfig{
		DB: queries,
	}
	
	router := chi.NewRouter()

	// must be define before we make other route like v1Router
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"https://*", "http://*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"*"},
		ExposedHeaders: []string{"Link"},
		AllowCredentials: false,
		MaxAge: 300,
	}))

	v1Router := chi.NewRouter()

	// Routes: 
	v1Router.Get("/health", handlerHealth)
	v1Router.Get("/error", handlerErr)

	// User Routes
	v1Router.Get("/users", apiCfg.middlewareAuth(apiCfg.handlerGetUser))
	v1Router.Post("/users", apiCfg.middlewareAuth(apiCfg.handlerCreateUser))

	// Feed Routes
	v1Router.Get("/feeds", apiCfg.middlewareAuth(apiCfg.handlerGetFeeds)) // get all feeds
	v1Router.Post("/feeds", apiCfg.middlewareAuth(apiCfg.handlerCreateFeed))

	// Feed Follows Routes
	v1Router.Get("/feed-follows", apiCfg.middlewareAuth(apiCfg.handlerGetAllFeedFollows))
	v1Router.Post("/feed-follows", apiCfg.middlewareAuth(apiCfg.handlerCreateFeedFollow))
	v1Router.Delete("/feed-follows/{feedFollowID}", apiCfg.middlewareAuth(apiCfg.handlerDeleteFeedFollow))

	router.Mount("/v1", v1Router)

	server := &http.Server{
		Handler: router,
		Addr: ":" + PORT,
	}

	log.Printf("Server starting on port %v", PORT)
	err = server.ListenAndServe()

	if err != nil {
		log.Fatal(err)
	}
}