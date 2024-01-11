package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

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

	v1Router.Get("/health", handlerHealth)
	v1Router.Get("/error", handlerErr)

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