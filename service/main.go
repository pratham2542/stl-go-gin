package main

import (
	"fmt"
	"os"
	"test-go/config"
	"test-go/internals"
	"test-go/service/routes"
	"time"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Error parsing env!")
		os.Exit(1)
	}

	shutdownConfig := internals.NewShutdownConfig(time.Duration(5 * time.Second))

	appContext := &config.AppContext{}
	appContext = appContext.WithAppConfig()

	s := routes.NewServer(appContext).AddRoutes().Start()

	internals.Shutdown(s, *shutdownConfig)
}
