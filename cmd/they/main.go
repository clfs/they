package main

import (
	"context"
	"log"
	"os"

	"github.com/clfs/they/internal/engine"
)

func main() {
	engine := engine.New(os.Stdin, os.Stdout)

	ctx := context.Background()

	if err := engine.Run(ctx); err != nil {
		log.Fatal(err)
	}
}
