package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"main/internal/handlers"
	"main/internal/repository"
	"main/internal/service"

	"github.com/joho/godotenv"
	"github.com/rs/cors"
)

func main() {
	mux := http.NewServeMux()

	err := godotenv.Load()
	if err != nil {
		log.Println("Note: .env file not found or error loading it")
	}

	repo := repository.NewDiskImgRepo()
	imgService := service.New(repo)
	imgHandler := handlers.New(imgService)

	mux.HandleFunc("/process_image_url", imgHandler.ProcessImageURL)
	mux.HandleFunc("/process_image_file", imgHandler.ProcessImageFile)

	handler := cors.AllowAll().Handler(mux)

	port := os.Getenv("PORT")
	if port == "" {
		port = "5001"
	}

	addr := fmt.Sprintf(":%s", port)
	log.Printf("Server is running on port %s", port)
	if err := http.ListenAndServe(addr, handler); err != nil {
		log.Fatalf("Error starting server: %s", err)
	}
}
