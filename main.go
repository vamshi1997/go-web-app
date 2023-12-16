package main

import (
	"fmt"
	"github/joho/godotenv"
	"log"
	"os"
)

func main() {
	buildAndRunCommnad := "go build && ./go-web-app"
	fmt.Printf("welcome to the simple server..., use this command to run application %v \n", buildAndRunCommnad)

	godotenv.Load(".env")

	portString := os.Getenv("PORT")
	if portString == "" {
		log.Fatal("Port is not found in environment")
	}
	fmt.Printf("Server listening to Port: %v", portString)
}
