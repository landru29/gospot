// Package main is the main program entry.
package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	ctx := context.Background()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	rootCmd := rootCommand(ctx, quit)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
	}
}
