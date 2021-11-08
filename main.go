package main

import (
	"context"
	"github.com/guythatdrinkscoffee/CirculationApp/api/router"
	"github.com/guythatdrinkscoffee/CirculationApp/internal"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var ttlC internal.TTLCache

func init() {

	//Attempt to load the .env file
	if err := godotenv.Load(".env"); err != nil {
		log.Fatalln(err)
	}

	log.Println("Successfully loaded the .env file")

	//Init a new ttlC(cache)
	ttlC = internal.NewTLLCache()
}

func main() {
	//Define the Gin Router
	r := router.NewCirculationRouter(&ttlC)
	r.SetupRoutes()

	//Define a server in order to handle graceful shutdown
	srv := &http.Server{
		Addr:    ":8080",
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
		log.Fatalf("Forced server shutdown: ", err)
	}
}
