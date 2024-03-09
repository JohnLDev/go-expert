package main

import (
	"log/slog"

	"github.com/johnldev/requester/cmd/rootCmd"
)

func main() {
	slog.Info("App initialized")
	rootCmd.Execute()
}
