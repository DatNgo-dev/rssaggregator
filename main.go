package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("env")

	if err != nil {
		log.Fatal("Error: ", err)
	} 

	PORT := os.Getenv("PORT")
	if PORT == "" {
		log.Fatal("PORT is not found in the environment")
	}
	fmt.Println("PORT: ", PORT)
	
	router := chi.NewRouter()

	server := &http.Server{
		Handler: router,
		Addr: ":" + PORT,
	}

	log.Printf("Server starting on port %v", PORT)
	err = server.ListenAndServe()

	if err != nil {
		log.Fatal(err)
	}

	// router.Use(cors.Handler(cors.option{

	// }))
}