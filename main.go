package main

import (
	"fmt"
<<<<<<< Updated upstream
	"github/joho/godotenv"
	"log"
	"os"
=======
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	"github.com/vamshi1997/go-web-app/helper"
>>>>>>> Stashed changes
)

func main() {
	buildAndRunCommnad := "go build && ./go-web-app"
<<<<<<< Updated upstream
	fmt.Printf("welcome to the simple server..., use this command to run application %v \n", buildAndRunCommnad)
=======
	fmt.Printf("Welcome to the Go server...\nUse this command to run application %v \n", buildAndRunCommnad)

	helper.Logger("Testing Logger")
>>>>>>> Stashed changes

	godotenv.Load(".env")

	portString := os.Getenv("PORT")
	if portString == "" {
		log.Fatal("Port is not found in environment")
	}
<<<<<<< Updated upstream
	fmt.Printf("Server listening to Port: %v", portString)
=======

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

	v1Router.Get("/status", statusCheck)

	router.Mount("/v1", v1Router)

	server := &http.Server{
		Handler: router,
		Addr:    ":" + portString,
	}

	fmt.Printf("Server listening to Port: %v \n", portString)
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
>>>>>>> Stashed changes
}
