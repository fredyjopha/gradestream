package main

import (
	"context"
	"log"
	"net/http"

	"github.com/fredyjopha/gradestream/internal/api"
	"github.com/fredyjopha/gradestream/internal/storage/postgres"
)

func main() {
	ctx := context.Background()
	pool, err := postgres.Connect(ctx, "postgres://postgres:devpass@localhost:5432/gradestream")
	if err != nil {
		log.Fatalf("connexion echoué: %v", err)
	}
	defer pool.Close()
	log.Println("GradeStream — ecoute sur :8080")
	log.Fatal(http.ListenAndServe(":8080", api.NewRouter()))
}
