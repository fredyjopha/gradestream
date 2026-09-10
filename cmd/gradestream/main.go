package main

import (
	"context"
	"fmt"
	"log"
	"github.com/fredyjopha/gradestream/internal/storage/postgres"
)

func main() {
	ctx := context.Background()
	pool, err := postgres.Connect(ctx, "postgres://postgres:devpass@localhost:5432/gradestream")
	if err != nil {
		log.Fatalf("connexion echoué: %v", err)
	}
	defer pool.Close()
	fmt.Println("GradeStream — connexion postgres OK")
}
