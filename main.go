package main

import (
	"context"
	"github.com/guythatdrinkscoffee/CirculationApp/api/router"
	"github.com/guythatdrinkscoffee/CirculationApp/config"
	"github.com/guythatdrinkscoffee/CirculationApp/internal"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var ttlC *internal.TTLCache
var c *config.Config

func init() {
	//Init a new ttlC(cache)
	ttlC = internal.NewTLLCache()
	c = config.GetConfig()
}

func main() {
	ginMode := c.GIN_MODE

	//Define the Gin Router
	r := router.NewCirculationRouter(ttlC, ginMode)
	r.SetupRoutes()

	//Define a server in order to handle graceful shutdown
	srv := &http.Server{
		Addr:    ":" + c.PORT,
		Handler: r.Router,
	}

	// Initializing the server in a goroutine so that
	// it won't block the graceful shutdown handling below
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	//Create a channel to handle graceful shutdown. Gin Example
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	msg := <-sigChan
	log.Printf("Shutting down server with %s", msg)

	//Clear the cache
	log.Println("Closing the cache")
	_ = ttlC.Cache.Close()

	ctx, term := context.WithTimeout(context.Background(), 5*time.Second)
	defer term()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Forced server shutdown: %s", err)
	}
}
