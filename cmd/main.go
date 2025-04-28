package main

import (
	"log"
	"os"

	"github.com/AeddGynvael3110/sprint6/internal/server"
)

func main() {
	logger := log.New(os.Stdout, "MORSE_CONVERTER: ", log.LstdFlags)
	srv := server.NewServer(logger)

	logger.Println("Starting server on :8080")
	if err := srv.ListenAndServe(); err != nil {
		logger.Fatal(err)
	}
}
