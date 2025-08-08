package main

import (
	"log"

	"github.com/AlexFJ498/middle-earth-leitmotifs-api/cmd/api/bootstrap"
)

func main() {
	if err := bootstrap.Run(); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
