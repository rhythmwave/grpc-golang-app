package main

import (
	"go-bponline/m/cmd"
	"os"

	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	"github.com/lpernett/godotenv"
)

func main() {
	logger := log.NewLogfmtLogger(os.Stderr)

	// Load env file
	err := godotenv.Load()
	if err != nil {
		level.Error(logger).Log("msg", "Error loading .env file")
	}
	cmd.Execute()
}
