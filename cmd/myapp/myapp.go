package main

import (
	"context"
	"io"
	"log"
	"os"
	"os/signal"

	"github.com/set0xc3/crmGO/internal/app"
)

func run(ctx context.Context, w io.Writer, args []string) error {
	ctx, cancel := signal.NotifyContext(ctx, os.Interrupt)
	defer cancel()

	if err := app.Run(ctx); err != nil {
		return err
	}

	return nil
}

func main() {
	ctx := context.Background()

	if err := run(ctx, os.Stdout, os.Args); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}
