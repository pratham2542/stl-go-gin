package internals

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

// ShutdownConfig defines how to shutdown the app
type ShutdownConfig struct {
	Timeout time.Duration // Grace period
}

func NewShutdownConfig(
	timout time.Duration,
) *ShutdownConfig {
	return &ShutdownConfig{
		Timeout: timout,
	}
}

// Shutdown waits for termination signals and gracefully shuts down the server.
func Shutdown(server *http.Server, cfg ShutdownConfig) {

	// quit channel will recieve SIGINT(ctrl + c) and SIGTERM (kill signal from os)
	// It will not listen to SIGKILL because SIGKILL is a forcefull shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	<-quit // quit chan will only unblock from here if we get any signal
	fmt.Println("Shutdown signal received")

	// ctx will allow background activities like server, DB operations , connections, logger etc to shutdown under timeout
	// otherwise the shutdown will put them into sleep.
	ctx, cancel := context.WithTimeout(context.Background(), cfg.Timeout)
	defer cancel()

	// wg will wait for all the activities to complete their cleanup
	var wg sync.WaitGroup

	// Graceful server shutdown
	go func() {
		if err := server.Shutdown(ctx); err != nil { // this is a internal http server shutdown method
			fmt.Println("Server shutdown error: " + err.Error())
		} else {
			fmt.Println("Server shutdown cleanly")
		}
	}()

	// Wait for cleanup to complete
	wg.Wait()
	fmt.Println("Cleanup complete")
}
